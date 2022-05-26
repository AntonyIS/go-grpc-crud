package rest

import (
	"github.com/gin-gonic/gin"
)

type movieHandler interface {
	CreateMovie(*gin.Context)
	GetMovie(*gin.Context)
	GetMovies(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
}

type handler struct {
}

func NewHandler()
