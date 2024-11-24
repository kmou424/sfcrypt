package common

import (
	"github.com/kmou424/sfcrypt/app/buildinfo"
	"log/slog"
	"os"
)

func init() {
	if !buildinfo.Debug {
		return
	}
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
}
