package utils

import "nba-players-statistics/pkg"

type Validator[T any] interface {
	IsValid(T) bool
}

type PlayerLogValidator struct {
}

func (p PlayerLogValidator) IsValid(data pkg.PlayerStats) bool {
	return data.Fouls <= 6 &&
		data.MinutesPlayed <= 48 &&
		data.MinutesPlayed >= 0
}
