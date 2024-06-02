package internal

import (
	"context"
	"log"

	"nba-players-statistics/pkg"
)

type StatisticsService interface {
	Log(ctx context.Context, request pkg.LogRequest) error
	GetStatistics(ctx context.Context, request pkg.GetStatisticsRequest) error
}

type SimpleStatisticsService struct {
	logger *log.Logger
	dao    StatisticsDao
}

func (s *SimpleStatisticsService) Log(ctx context.Context, request pkg.LogRequest) error {
	//TODO implement me
	panic("implement me")
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
