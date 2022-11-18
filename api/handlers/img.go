package handlers

import (
	"fmt"
	"io"

	"github.com/rjahon/img-client/api/models"
	"github.com/rjahon/img-client/genproto/img_service"

	"github.com/gin-gonic/gin"
)

// Save Img godoc
// @ID create_img
// @Router /img [POST]
// @Summary Create Img
// @Description Save an image
// @Tags Img
// @Accept mpfd
// @Produce json
// @Param file formData file true "img"
// @Success 201 {object} models.Response{data=img_service.Img} "Image Saved Successfully"
// @Response 400 {object} models.Response{data=string} "Invalid Argument"
// @Failure 500 {object} models.Response{data=string} "Server Error"
func (h *Handler) CreateImg(c *gin.Context) {
	var form Form

	if err := c.ShouldBind(&form); err != nil {
		h.handleResponse(c, models.BadRequest, err.Error())
		return
	}

	file, err := form.File.Open()
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}
	defer file.Close()

	imgData, err := io.ReadAll(file)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}

	img, err := h.services.ImgService().Create(c, &img_service.CreateRequest{
		Title: form.File.Filename,
		Body:  imgData,
	})
	if err != nil {
		h.handleResponse(c, models.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, models.Created, img)
}

// Get Img By ID godoc
// @ID get_img_by_id
// @Router /img/{id} [GET]
// @Summary Get Img By ID
// @Description Get image by ID
// @Tags Img
// @Accept json
// @Produce mpfd
// @Param id path string true "id"
// @Success 200 {object} string "Image Retrieved Successfully"
// @Response 400 {object} models.Response{data=string} "Invalid Argument"
// @Failure 500 {object} models.Response{data=string} "Server Error"
func (h *Handler) GetImgByID(c *gin.Context) {
	res, err := h.services.ImgService().Get(c, &img_service.ImgPrimaryKey{
		Id: c.Param("id"),
	})
	if err != nil {
		h.handleResponse(c, models.GRPCError, err.Error())
		return
	}

	file := fmt.Sprintf("./out/%s", res.Title)
	err = h.Btoi(c, res.Body, file)
	if err != nil {
		h.handleResponse(c, models.InternalServerError, err.Error())
		return
	}

	c.FileAttachment(file, res.Title)
	h.handleResponse(c, models.OK, res)
}

// Get Imgs godoc
// @ID get_imgs
// @Router /img [GET]
// @Summary List Imgs
// @Description List all images
// @Tags Img
// @Accept json
// @Produce json
// @Param img query models.GetImagesModel false "filters"
// @Success 200 {object} models.Response{data=img_service.GetListResponse} "Images Retrieved Successfully"
// @Failure 500 {object} models.Response{data=string} "Server Error"
func (h *Handler) GetImgs(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", h.cfg.DefaultLimit)
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", h.cfg.DefaultOffset)
	if err != nil {
		return
	}

	list, err := h.services.ImgService().GetList(c, &img_service.GetListRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handleResponse(c, models.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, models.OK, list)
}
