package controller

import (
	"blog/application/domain"
	"blog/application/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService ports.PostService
}

func NewPostController(apiGroup *gin.RouterGroup, postService ports.PostService) {
	controller := &PostController{
		postService: postService,
	}

	apiGroup.GET("/posts/:id", controller.findbyAuthorId)
	apiGroup.POST("/posts", controller.create)
	apiGroup.PATCH("/posts/:id", controller.update)
}

func (c *PostController) findbyAuthorId(ctx *gin.Context) {
	id := ctx.Param("id")
	posts, err := c.postService.FindByAuthorId(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) create(ctx *gin.Context) {
	var post domain.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postService.Create(&post); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post created successfully"})
}

func (c *PostController) update(ctx *gin.Context) {
	id := ctx.Param("id")
	var request struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.postService.Update(id, request.Title, request.Content); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
