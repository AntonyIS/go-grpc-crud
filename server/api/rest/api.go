package rest

import (
	"fmt"
	"net/http"

	srv "github.com/AntonyIS/go-grpc-crud/server/api"
	"github.com/AntonyIS/go-grpc-crud/server/domain"
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

func NewHandler() movieHandler {
	return &handler{}
}

func (handler) CreateMovie(ctx *gin.Context) {
	movie := domain.Movie{}
	srv := *srv.Service()

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := srv.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (handler) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	srv := *srv.Service()
	res, err := srv.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (handler) GetMovies(ctx *gin.Context) {
	srv := *srv.Service()
	movies, err := srv.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (handler) UpdateMovie(ctx *gin.Context) {
	srv := *srv.Service()
	movie := domain.Movie{}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := srv.UpdateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (handler) DeleteMovie(ctx *gin.Context) {
	srv := *srv.Service()
	id := ctx.Param("id")

	err := srv.DeleteMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Movie deleted successfuly",
	})
}
