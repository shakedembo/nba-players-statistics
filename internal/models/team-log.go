package models

import "github.com/uptrace/bun"

type TeamLog struct {
	bun.BaseModel `bun:"table:teams_logs,alias:tl"`

	TeamId     int   `bun:",pk,notnull"`
	GameNumber uint8 `bun:",pk,notnull"`
	SeasonId   int   `bun:",pk,notnull"`
	Team       *Team `bun:"rel:belongs-to,join:team_id=id"`
	Game       *Game `bun:"rel:belongs-to,join:game_number=number,join:season_id=season_id"`
	Points     uint8
	Rebounds   uint8
	Assists    uint8
	Steals     uint8
	Blocks     uint8
	Turnovers  uint8
	Fouls      uint8
}
