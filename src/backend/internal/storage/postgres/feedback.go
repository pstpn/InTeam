package mypostgres

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/storage/postgres"
)

type FeedbackStorage struct {
	*postgres.Postgres
}

func NewFeedbackStorage(db *postgres.Postgres) storage.IFeedbackStorage {
	return &FeedbackStorage{db}
}

func (r *FeedbackStorage) Create(ctx context.Context, feedback *model.Feedback) error {
	query := r.Builder.
		Insert(feedbackTable).
		Columns(
			racketIDField,
			userIDField,
			feedbackField,
			ratingField,
			dateField).
		Values(
			feedback.RacketID,
			feedback.UserID,
			feedback.Text,
			feedback.Rating,
			feedback.Date).
		Suffix("returning feedback")

	sql, ars, err := query.ToSql()
	if err != nil {
		return err
	}

	row := r.Pool.QueryRow(ctx, sql, ars...)
	err = row.Scan(
		&feedback.Text,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *FeedbackStorage) Update(ctx context.Context, feedback *model.Feedback) error {
	query := r.Builder.
		Update(feedbackTable).
		Set(feedbackField, feedback.Text).
		Set(ratingField, feedback.Rating).
		Set(dateField, feedback.Date).
		Where(squirrel.And{
			squirrel.Eq{userIDField: feedback.UserID},
			squirrel.Eq{racketIDField: feedback.RacketID}})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *FeedbackStorage) Remove(ctx context.Context, req *dto.RemoveFeedbackReq) error {
	query := r.Builder.
		Delete(feedbackTable).
		Where(squirrel.And{
			squirrel.Eq{userIDField: req.UserID},
			squirrel.Eq{racketIDField: req.RacketID}})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *FeedbackStorage) GetFeedback(ctx context.Context, req *dto.GetFeedbackReq) (*model.Feedback, error) {
	query := r.Builder.
		Select("*").
		From(feedbackTable).
		Where(squirrel.And{
			squirrel.Eq{userIDField: req.UserID},
			squirrel.Eq{racketIDField: req.RacketID}})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.Pool.QueryRow(ctx, sql, args...)

	return r.rowToFeedbackModel(row)
}

func (r *FeedbackStorage) GetFeedbacksByUserID(ctx context.Context, id int) ([]*model.Feedback, error) {
	query := r.Builder.
		Select("*").
		From(feedbackTable).
		Where(squirrel.Eq{userIDField: id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*model.Feedback

	for rows.Next() {
		feedback, err := r.rowToFeedbackModel(rows)
		if err != nil {
			return nil, err
		}

		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

func (r *FeedbackStorage) GetFeedbacksByRacketID(ctx context.Context, id int) ([]*model.FeedbackWithUsername, error) {
	query := r.Builder.
		Select(
			nameField,
			surnameField,
			racketIDField,
			userIDField,
			feedbackField,
			ratingField,
			dateField,
		).
		From(feedbackTable).
		Join(on(feedbackTable, userTable, userIDField, idField)).
		Where(squirrel.Eq{racketIDField: id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var feedbacks []*model.FeedbackWithUsername

	for rows.Next() {
		feedback, err := r.rowToFeedbackWithUsernameModel(rows)
		if err != nil {
			return nil, err
		}

		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

func (r *FeedbackStorage) rowToFeedbackModel(row pgx.Row) (*model.Feedback, error) {
	var feedback model.Feedback

	err := row.Scan(
		&feedback.RacketID,
		&feedback.UserID,
		&feedback.Text,
		&feedback.Rating,
		&feedback.Date,
	)
	if err != nil {
		return nil, err
	}

	return &feedback, nil
}

func (r *FeedbackStorage) rowToFeedbackWithUsernameModel(row pgx.Row) (*model.FeedbackWithUsername, error) {
	var feedback model.FeedbackWithUsername
	var name, surname string

	err := row.Scan(
		&name,
		&surname,
		&feedback.RacketID,
		&feedback.UserID,
		&feedback.Text,
		&feedback.Rating,
		&feedback.Date,
	)
	if err != nil {
		return nil, err
	}

	feedback.Username = name + " " + surname

	return &feedback, nil
}
