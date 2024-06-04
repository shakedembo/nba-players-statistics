package internal

import (
	"context"
	"errors"
	"log"
	"nba-players-statistics/internal/utils"

	"nba-players-statistics/pkg"
)

type StatisticsService interface {
	Log(ctx context.Context, request pkg.LogRequest) error
	GetStatistics(ctx context.Context, request pkg.GetStatisticsRequest) error
}

type SimpleStatisticsService struct {
	logger    *log.Logger
	dao       StatisticsDao
	validator utils.Validator[pkg.PlayerStats]
}

func (s *SimpleStatisticsService) Log(ctx context.Context, request pkg.LogRequest) error {
	context.WithValue(ctx, utils.PlayerId, request.PlayerId)

	if !s.validator.IsValid(request.PlayerData) {
		errText := "the data that was provided is not valid, didn't log it in the system"
		s.logger.Printf(errText)
		return errors.New(errText)
	}

	return s.dao.Log(ctx, request)
}

func (s *SimpleStatisticsService) GetStatistics(ctx context.Context, request pkg.GetStatisticsRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewSimpleStatisticsService(dao StatisticsDao, logger *log.Logger) *SimpleStatisticsService {
	return &SimpleStatisticsService{
		logger: logger,
		dao:    dao,
	}
}
