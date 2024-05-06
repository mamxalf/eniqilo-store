package infra

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/mamxalf/eniqilo-store/config"

	_ "github.com/lib/pq" // add this
	"github.com/rs/zerolog/log"
)

type PostgresConn struct {
	PG *sqlx.DB
}

func ProvidePostgresConn(config *config.Config) *PostgresConn {
	return &PostgresConn{
		PG: CreatePostgresDBConnection(config),
	}
}

func CreatePostgresDBConnection(config *config.Config) *sqlx.DB {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?%s",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
		config.DbParams)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("host", config.DbHost).
			Str("port", config.DbPort).
			Str("dbName", config.DbName).
			Msg("Failed connecting to Postgres database")
	} else {
		log.
			Info().
			Str("host", config.DbHost).
			Str("port", config.DbPort).
			Str("dbName", config.DbName).
			Msg("Connected to Postgres database")
	}

	return db
}
