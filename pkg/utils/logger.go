package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var once sync.Once
var Logger zerolog.Logger

func ConfigBasicLogger() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimestampFieldName = time.RFC3339Nano

		logLevel, err := zerolog.ParseLevel(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = zerolog.InfoLevel
		}

		var output io.Writer = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		buildInfo, _ := debug.ReadBuildInfo()
		Logger = zerolog.New(output).Level(zerolog.Level(logLevel)).With().Timestamp().Str("go_version", buildInfo.GoVersion).Logger()
	})
}
