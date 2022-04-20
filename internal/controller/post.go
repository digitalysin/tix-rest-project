package controller

import (
	"tix-rest-project/internal/service"
	"tix-rest-project/internal/shared/dto"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	Post struct {
		dig.In
		Service service.Holder
	}
)

// All godoc
// @Tags posts-controller
// @Summary Get All Posts
// @Description Put all mandatory parameter
// @Accept json
// @Produce json
// @Success 200 {object} dto.FindPostResponse
// @Router /posts [get]
func (impl *Post) Find(c echo.Context) error {
	var (
		req = dto.FindPostRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := impl.Service.Post.Find(c.Request().Context(), &req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

// All godoc
// @Tags posts-controller
// @Summary Get post by id
// @Description Put all mandatory parameter
// @Accept json
// @Produce json
// @Param id path string true "post id"
// @Success 200 {object} dto.FindPostByIdResponse
// @Router /posts/{id} [get]
func (impl *Post) FindById(c echo.Context) error {
	var (
		req = dto.FindPostByIdRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := impl.Service.Post.FindById(c.Request().Context(), &req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

// All godoc
// @Tags posts-controller
// @Summary Create a new post
// @Description Put all mandatory parameter
// @Param CreatePostRequest body dto.CreatePostRequest true "CreatePostRequest"
// @Accept json
// @Produce json
// @Success 200 {object} dto.CreatePostResponse
// @Router /posts [post]
func (impl *Post) Create(c echo.Context) error {
	var (
		req = dto.CreatePostRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := impl.Service.Post.Create(c.Request().Context(), &req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

// All godoc
// @Tags posts-controller
// @Summary Update existing post
// @Description Put all mandatory parameter
// @Param UpdatePostRequest body dto.UpdatePostRequest true "UpdatePostRequest"
// @Param id path string true "post id"
// @Accept json
// @Produce json
// @Success 200 {object} dto.UpdatePostRequest
// @Router /posts/{id} [patch]
func (impl *Post) Update(c echo.Context) error {
	var (
		req = dto.UpdatePostRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := impl.Service.Post.Update(c.Request().Context(), &req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}

// All godoc
// @Tags posts-controller
// @Summary Delete existing post
// @Description Put all mandatory parameter
// @Param id path string true "post id"
// @Accept json
// @Produce json
// @Success 200 {object} dto.DeletePostRequest
// @Router /posts/{id} [delete]
func (impl *Post) Delete(c echo.Context) error {
	var (
		req = dto.DeletePostRequest{}
	)

	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := impl.Service.Post.Delete(c.Request().Context(), &req)

	if err != nil {
		return err
	}

	return c.JSON(200, res)
}
