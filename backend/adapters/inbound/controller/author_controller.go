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

	apiGroup.POST("/authors/login", controller.login)
	apiGroup.POST("/authors/register", controller.register)
}

func (c *AuthorController) login(ctx *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userAgent := ctx.Request.UserAgent()
	ipAddress := ctx.Request.RemoteAddr

	authorLogin, err := c.authorService.Login(request.Email, request.Password, userAgent, ipAddress)
	if err != nil {
		switch err.Error() {
		case "invalid credentials":
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		case "account is disabled":
			ctx.JSON(http.StatusForbidden, gin.H{"error": "account is disabled"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	ctx.JSON(http.StatusOK, authorLogin)
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
