package pkg

type LogData struct {
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
	PlayerId int
	Data     LogData
}

type LogResponse struct {
}
