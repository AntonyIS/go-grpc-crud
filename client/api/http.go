package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	m "github.com/AntonyIS/go-grpc-crud/server/domain"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type movieHandler interface {
	CreateMovie(*gin.Context)
	GetMovie(*gin.Context)
	GetMovies(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
}

type handler struct{}

func gRPCClient() *pb.MovieServiceClient {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewMovieServiceClient(conn)

	return &client
}

func NewHandler() movieHandler {
	return handler{}
}

func (handler) CreateMovie(ctx *gin.Context) {
	movie := m.Movie{}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := *gRPCClient()

	res, err := client.CreateMovie(ctx, &pb.MovieRequest{
		Name:        movie.Name,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Image:       movie.Image,
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	movie.ID = res.Id
	ctx.JSON(http.StatusOK, movie)
}

func (handler) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	client := *gRPCClient()

	res, err := client.GetMovie(ctx, &pb.MovieID{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errors": "Movie not found",
		})
		return
	}

	movie := make(map[string]string)
	movie["id"] = res.GetId()
	movie["name"] = res.GetName()
	movie["description"] = res.GetDescription()
	movie["release_date"] = res.GetReleaseDate()
	movie["image"] = res.GetImage()
	ctx.JSON(http.StatusOK, movie)
}

func (handler) GetMovies(ctx *gin.Context) {
	client := *gRPCClient()
	res, err := client.GetMovies(ctx, &pb.EmptyRequest{})

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	movies := res.GetMovies()
	if len(movies) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Movies": []m.Movie{},
		})
		return
	}
	ctx.JSON(http.StatusOK, res.GetMovies())

}

func (handler) UpdateMovie(ctx *gin.Context) {
	movie := m.Movie{}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := *gRPCClient()

	res, err := client.UpdateMovie(ctx, &pb.MovieRequest{
		Id:          movie.ID,
		Name:        movie.Name,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Image:       movie.Image,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	movie.ID = res.Id
	ctx.JSON(http.StatusOK, movie)
}

func (handler) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	client := *gRPCClient()
	res, err := client.DeleteMovie(ctx, &pb.MovieID{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": res.GetMessage(),
	})
}
