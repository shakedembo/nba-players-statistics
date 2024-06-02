package internal

type StatisticsService interface {
	HealthHandler() string
}

type SimpleStatisticsService struct {
}

func NewSimpleStatisticsService() *SimpleStatisticsService {
	return &SimpleStatisticsService{}
}

func (s *SimpleStatisticsService) HealthHandler() string {
	return "HEALTHY"
}
