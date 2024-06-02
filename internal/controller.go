package internal

import (
	"context"
	"log"

	"nba-players-statistics/pkg"
)

type StatisticsController struct {
	logger  *log.Logger
	service StatisticsService
}

func NewStatisticsController(service StatisticsService, logger *log.Logger) *StatisticsController {
	return &StatisticsController{
		logger:  logger,
		service: service,
	}
}

func (s *StatisticsController) Log(ctx context.Context, req pkg.LogRequest) (pkg.LogResponse, error, int) {
	err := s.service.Log(ctx, req)
	if err != nil {
		return pkg.LogResponse{}, err, 500
	}

	return pkg.LogResponse{}, nil, 201
}

func (s *StatisticsController) GetStats(
	ctx context.Context,
	req pkg.GetStatisticsRequest) (pkg.GetStatisticsResponse, error, int) {

	err := s.service.GetStatistics(ctx, req)
	if err != nil {
		return pkg.GetStatisticsResponse{}, err, 500
	}

	//TODO: Implement and return real values
	return pkg.GetStatisticsResponse{}, nil, 200
}

func (s *StatisticsController) Health(ctx context.Context, _ pkg.HealthRequest) (pkg.HealthResponse, error, int) {
	return pkg.HealthResponse{
		Healthy: true,
		Message: "I'm healthy!",
	}, nil, 200
}
