package pkg

type PlayerStats struct {
	Points        uint8   `json:"points"`
	Rebounds      uint8   `json:"rebounds"`
	Assists       uint8   `json:"assists"`
	Steals        uint8   `json:"steals"`
	Blocks        uint8   `json:"blocks"`
	Turnovers     uint8   `json:"turnovers"`
	Fouls         uint8   `json:"fouls"`
	MinutesPlayed float32 `json:"minutesPlayed"`
}

type LogRequest struct {
	GameNumber uint16      `json:"game_number"`
	PlayerId   int         `json:"player_id"`
	TeamId     int         `json:"team_id"`
	PlayerData PlayerStats `json:"player_data"`
}

type LogResponse struct {
}
