package controller

import (
	"net/http"
	"strconv"

	"enigmacamp.com/blog-apps/model"
	"enigmacamp.com/blog-apps/shared/common"
	"enigmacamp.com/blog-apps/usecase"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	blogUC usecase.BlogUseCase
	rg     *gin.RouterGroup
}

func (b *BlogController) createHandler(ctx *gin.Context) {
	var payload model.Blog
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	task, err := b.blogUC.CreateNewBlog(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, task, "Created")
}

func (b *BlogController) listHandler(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	size, _ := strconv.Atoi(ctx.Query("size"))
	tasks, paging, err := b.blogUC.FindAllBlog(page, size)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range tasks {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, paging, "Ok")
}

func (t *BlogController) Route() {
	t.rg.POST("/blogs", t.createHandler)
	t.rg.GET("/blogs", t.listHandler)
}

func NewBlogController(blogUC usecase.BlogUseCase, rg *gin.RouterGroup) *BlogController {
	return &BlogController{
		blogUC: blogUC,
		rg:     rg,
	}
}
