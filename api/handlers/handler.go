package handlers

import (
	"bytes"
	"image"
	"image/png"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/rjahon/img-client/api/models"
	"github.com/rjahon/img-client/config"
	"github.com/rjahon/img-client/grpc/client"
	"github.com/rjahon/img-client/logger"

	"github.com/gin-gonic/gin"
)

type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status models.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, models.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) ParseQueryParam(c *gin.Context, key string, defaultValue string) (int32, error) {
	valueStr := c.DefaultQuery(key, defaultValue)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return 0, err
	}

	return int32(value), nil
}

func (h *Handler) Btoi(c *gin.Context, data []byte, location string) (err error) {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return err
	}

	file, err := os.Create(location)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return err
	}

	if err := png.Encode(file, img); err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return err
	}

	return nil
}
