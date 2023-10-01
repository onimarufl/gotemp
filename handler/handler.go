package handler

import (
	"gotemp/model"
	"gotemp/repository"
	"gotemp/service"
	"net/http"
	"strconv"

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

	if req.Firstname == "" {
		c.JSON(http.StatusBadRequest, "Request Firstname")
		return
	}

	resp, err := h.service.PostService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h Handler) HandlerInsertData(c *gin.Context) {
	req := repository.InsertDataRequest{}
	c.Bind(&req)

	if err := h.service.InsertData(req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "success")
}

func (h *Handler) HandlerInquiryDataByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	userResp, err := h.service.InquiryDataByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, &userResp)
}

func (h *Handler) HandlerInquiryAllData(c *gin.Context) {

	userResp, err := h.service.InquiryAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, &userResp)
}

func (h *Handler) HandlerUpdateDataByID(c *gin.Context) {
	req := repository.UpdateDataRequest{}
	c.Bind(&req)

	userResp, err := h.service.UpdateDataByID(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, &userResp)
}

func (h *Handler) HandlerDeleteDataByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	err = h.service.DeleteDataByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, "success")
}
