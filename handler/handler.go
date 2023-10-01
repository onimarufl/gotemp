package handler

import (
	"gotemp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) Handler {
	return Handler{service: service}
}

func (h Handler) HandlerGetMessage(c *gin.Context) {
	message := c.Query("message")

	resp := h.service.GetMessage(message)

	c.JSON(http.StatusOK, resp)
}
