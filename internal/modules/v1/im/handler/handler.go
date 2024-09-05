package handler

import (
	"Meow-backend/internal/modules/v1/im/service"
)

type IMHandler struct {
	imService service.IMService
}

func NewIMHandler(imService service.IMService) *IMHandler {
	return &IMHandler{
		imService: imService,
	}
}
