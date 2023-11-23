package step_handlers

import (
	"skat_bot/internal/service"
	"skat_bot/internal/service/auth"
	log "skat_bot/pkg/logger"
)

type StepHandler struct {
	S    service.Service
	log  *log.Logger
	Auth auth.Auth
}

func New(s service.Service, logger *log.Logger, auth auth.Auth) StepHandler {
	return StepHandler{
		s,
		logger,
		auth,
	}
}
