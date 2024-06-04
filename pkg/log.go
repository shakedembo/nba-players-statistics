package pkg

type PlayerStats struct {
	Points        uint8
	Rebounds      uint8
	Assists       uint8
	Steals        uint8
	Blocks        uint8
	Turnovers     uint8
	Fouls         uint8
	MinutesPlayed float32
}

type LogRequest struct {
	GameId     int
	PlayerId   int
	PlayerData PlayerStats
}

type LogResponse struct {
}
