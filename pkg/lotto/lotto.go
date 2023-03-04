package lotto

import (
	"github.com/gookit/slog"
	"github.com/yawn77/sphelper"
)

// URLs
const ()

// regex
const ()

// errors
const ()

type Error string

func (e Error) Error() string { return string(e) }

func Play(yearOnly bool) {
	creds, err := sphelper.GetCredentials()
	if err != nil {
		slog.Error("failed to get credentials: %v", err)
		return
	}
	slog.Infof("play lotto for %s", creds.Username)
	slog.Info("played lotto successfully")
}
