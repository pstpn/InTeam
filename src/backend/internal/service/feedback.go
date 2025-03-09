package service

import (
	"backend/pkg/common"
	"backend/pkg/logger"
	"context"
	"fmt"
	"time"

	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/storage"
)

type IFeedbackService interface {
	CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error)
	RemoveFeedback(ctx context.Context, req *dto.RemoveFeedbackReq) error
	UpdateFeedback(ctx context.Context, req *dto.UpdateFeedbackReq) error
	GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error)
	GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.Feedback, error)
}

type FeedbackService struct {
	logger logger.Interface
	repo   storage.IFeedbackStorage
}

func NewFeedbackService(logger logger.Interface, repo storage.IFeedbackStorage) *FeedbackService {
	return &FeedbackService{
		logger: logger,
		repo:   repo,
	}
}

func (s *FeedbackService) CreateFeedback(ctx context.Context, req *dto.CreateFeedbackReq) (*model.Feedback, error) {
	var feedback model.Feedback
	common.Copy(&feedback, req)

	feedback.Date = time.Now()

	err := s.repo.Create(ctx, &feedback)
	if err != nil {
		s.logger.Errorf("create feedback fail, error %s", err.Error())
		return nil, fmt.Errorf("create feedback fail, error %s", err)
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
		s.logger.Errorf("get feedback fail, error %s", err.Error())
		return fmt.Errorf("get feedback fail, error %s", err)
	}

	err = s.repo.Remove(ctx, req)
	if err != nil {
		s.logger.Errorf("remove feedback fail, error %s", err.Error())
		return fmt.Errorf("remove feedback fail, error %s", err)
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
		s.logger.Errorf("get feedback fail, error %s", err.Error())
		return fmt.Errorf("get feedback fail, error %s", err)
	}

	common.Copy(&feedback, req)

	err = s.repo.Update(ctx, feedback)
	if err != nil {
		s.logger.Errorf("update feedback fail, error %s", err.Error())
		return fmt.Errorf("update feedback fail, error %s", err)
	}

	return nil
}

func (s *FeedbackService) GetFeedbacksByRacketID(ctx context.Context, racketID int) ([]*model.Feedback, error) {
	feedbacks, err := s.repo.GetFeedbacksByRacketID(ctx, racketID)
	if err != nil {
		s.logger.Errorf("get feedback by racket id fail, error %s", err.Error())
		return nil, fmt.Errorf("get feedback by racket id fail, error %s", err)
	}

	return feedbacks, nil
}

func (s *FeedbackService) GetFeedbacksByUserID(ctx context.Context, userID int) ([]*model.Feedback, error) {
	feedbacks, err := s.repo.GetFeedbacksByUserID(ctx, userID)
	if err != nil {
		s.logger.Errorf("get feedback by user id fail, error %s", err.Error())
		return nil, fmt.Errorf("get feedback by user id fail, error %s", err)
	}

	return feedbacks, nil
}
