package service

import (
	"context"
	"fmt"
	"time"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
	"backend/pkg/common"
	"backend/pkg/logger"
)

type IFeedbackService interface {
	CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error)
	RemoveFeedback(ctx context.Context, req *dto.RemoveFeedbackReq) error
	UpdateFeedback(ctx context.Context, req *dto.UpdateFeedbackReq) error
	GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error)
	GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.FeedbackWithUsername, error)
}

type FeedbackService struct {
	l    logger.Interface
	repo storage.IFeedbackStorage
}

func NewFeedbackService(l logger.Interface, repo storage.IFeedbackStorage) *FeedbackService {
	return &FeedbackService{
		l:    l,
		repo: repo,
	}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error) {
	var feedback model.Feedback
	common.Copy(&feedback, req)

	feedback.Date = time.Now()

	err := s.repo.Create(ctx, &feedback)
	if err != nil {
		s.l.Errorf("create feedback fail, error %s", err.Error())
		return nil, fmt.Errorf("create feedback fail, error %w", err)
	}

	return &feedback, nil
}

func (s *FeedbackService) RemoveFeedback(ctx context.Context, req *dto.RemoveFeedbackReq) error {
	_, err := s.repo.GetFeedback(ctx,
		&dto.GetFeedbackReq{
			RacketID: req.RacketID,
			UserID:   req.UserID,
		},
	)
	if err != nil {
		s.l.Errorf("get feedback fail, error %s", err.Error())
		return fmt.Errorf("get feedback fail, error %w", err)
	}

	err = s.repo.Remove(ctx, req)
	if err != nil {
		s.l.Errorf("remove feedback fail, error %s", err.Error())
		return fmt.Errorf("remove feedback fail, error %w", err)
	}

	return nil
}

func (s *FeedbackService) UpdateFeedback(ctx context.Context, req *dto.UpdateFeedbackReq) error {
	feedback, err := s.repo.GetFeedback(ctx,
		&dto.GetFeedbackReq{
			RacketID: req.RacketID,
			UserID:   req.UserID,
		},
	)
	if err != nil {
		s.l.Errorf("get feedback fail, error %s", err.Error())
		return fmt.Errorf("get feedback fail, error %w", err)
	}

	common.Copy(&feedback, req)

	err = s.repo.Update(ctx, feedback)
	if err != nil {
		s.l.Errorf("update feedback fail, error %s", err.Error())
		return fmt.Errorf("update feedback fail, error %w", err)
	}

	return nil
}

func (s *FeedbackService) GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.FeedbackWithUsername, error) {
	feedbacks, err := s.repo.GetFeedbacksByRacketID(ctx, racketID)
	if err != nil {
		s.l.Errorf("get feedback by racket id fail, error %s", err.Error())
		return nil, fmt.Errorf("get feedback by racket id fail, error %w", err)
	}

	return feedbacks, nil
}

func (s *FeedbackService) GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error) {
	feedbacks, err := s.repo.GetFeedbacksByUserID(ctx, userID)
	if err != nil {
		s.l.Errorf("get feedback by user id fail, error %s", err.Error())
		return nil, fmt.Errorf("get feedback by user id fail, error %w", err)
	}

	return feedbacks, nil
}
