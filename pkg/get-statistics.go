package pkg

type GetStatisticsRequest struct {
}

type PlayerSeasonAvg struct {
}

type TeamSeasonAvg struct {
}

type GetStatisticsResponse struct {
	PlayerSeasonAvg PlayerSeasonAvg `json:"playerSeasonAvg"`
	TeamSeasonAvg   TeamSeasonAvg   `json:"teamSeasonAvg"`
}
