package main

import (
	api "github.com/AntonyIS/go-grpc-crud/client/api"
	"github.com/gin-gonic/gin"
)

func main() {
	handler := api.NewHandler()
	r := gin.Default()

	r.POST("/api/v1/movies", handler.CreateMovie)
	r.GET("/api/v1/movies", handler.GetMovies)
	r.GET("/api/v1/movies/:id", handler.GetMovie)
	r.PUT("/api/v1/movies", handler.UpdateMovie)
	r.DELETE("/api/v1/movies/:id", handler.DeleteMovie)
	r.Run((":5000"))

}
