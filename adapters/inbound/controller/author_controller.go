package controller

import (
	"blog/adapters/inbound/dtos"
	"blog/application/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorService ports.AuthorService
}

func NewAuthorController(apiGroup *gin.RouterGroup, authorService ports.AuthorService) {
	controller := &AuthorController{
		authorService: authorService,
	}

	apiGroup.POST("/authors/register", controller.register)
}

func (c *AuthorController) register(ctx *gin.Context) {
	var request dtos.RegisterAuthorRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, token, err := c.authorService.Register(request.ToAuthor(), request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resource := dtos.ToRegisterAuthorResponse(author, token)
	ctx.JSON(http.StatusOK, resource)
}
