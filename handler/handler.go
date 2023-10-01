package handler

import (
	"gotemp/model"
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

func (h Handler) HandlerPostService(c *gin.Context) {
	req := model.Request{}
	c.Bind(&req)

	resp, err := h.service.PostService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, resp)
}
