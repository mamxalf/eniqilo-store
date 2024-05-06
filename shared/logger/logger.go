package logger

import (
	"github.com/mamxalf/eniqilo-store/config"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes the logger
func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	// format the output as needed here

	log.Logger = log.Output(output)
	log.Trace().Msg("Zerolog initialized.")
}

// FatalError logs fatal with message.
func FatalError(err error, message string) {
	log.Fatal().Err(err).Msg(message)
}

// ErrorWithStack logs and error and its stack trace with custom formatting.
func ErrorWithStack(err error) {
	log.Error().Msgf("%+v", errors.WithStack(err))
}

// ErrorInterfaceWithMessage logs and error and show interface data usually parameters/req data.
func ErrorInterfaceWithMessage(err error, message string, key string, args interface{}) {
	log.Error().Err(err).Interface(key, args).Msg(message)
}

// ErrorWithMessage logs and error with message.
func ErrorWithMessage(err error, message string) {
	log.Error().Err(err).Msg(message)
}

// Info just logs info with message.
func Info(message string) {
	log.Info().Msg(message)
}

// WarningWithMessage logs and warning with message.
func WarningWithMessage(err error, message string) {
	log.Warn().Err(err).Msg(message)
}

// SetLogLevel sets the desired log level specified in env var.
func SetLogLevel(config *config.Config) {
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		level = zerolog.TraceLevel
		log.Trace().Str("loglevel", level.String()).Msg("Environment has no log level set up, using default.")
	} else {
		log.Trace().Str("loglevel", level.String()).Msg("Desired log level detected.")
	}
	zerolog.SetGlobalLevel(level)
}
