package infra

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mamxalf/eniqilo-store/config"
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
		"postgres://%s:%s@%s:%s/%s?%s",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
		config.DbParams)
	// Create a configuration object
	db, err := sqlx.ConnectContext(context.Background(), "pgx", connStr)
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
