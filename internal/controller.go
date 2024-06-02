package internal

import (
	"context"
	"nba-players-statistics/pkg"
)

type StatisticsController struct {
	service StatisticsService
}

func NewStatisticsController(service StatisticsService) *StatisticsController {
	return &StatisticsController{
		service: service,
	}
}

func (s *StatisticsController) Health(ctx context.Context, _ pkg.HealthRequest) (pkg.HealthResponse, error, int) {
	return pkg.HealthResponse{
		Healthy: true,
		Message: "I'm healthy!",
	}, nil, 200
}
