package config

type DbConfig struct {
	User     string
	Password string
}

const (
	PsqlUser     = "postgres" //os.Getenv("PSQL_USER")
	PsqlPassword = "postgres" //os.Getenv("PSQL_PASS")
)

func NewDbConfig() *DbConfig {
	return &DbConfig{
		User:     PsqlUser,
		Password: PsqlPassword,
	}
}
