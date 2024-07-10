package dataset

import "log/slog"

type Service struct {
	log        *slog.Logger
	repository Repsitory 
}

func NewService(log *slog.Logger, repo Repsitory) *Service {
	return nil
}
