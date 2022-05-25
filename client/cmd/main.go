package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMovieServiceClient(conn)

	router := gin.Default()

	router.POST("/api/v1/movies", func(ctx *gin.Context) {

	})

	router.GET("/api/v1/movies/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := client.GetMovie(ctx, &pb.MovieID{Id: id})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Error: %s", err),
			})
			return
		}
		movie := make(map[string]string)
		movie["id"] = res.GetId()
		movie["name"] = res.GetName()
		movie["description"] = res.GetDescription()
		movie["release_date"] = res.GetReleaseDate()
		movie["image"] = res.GetReleaseDate()
		ctx.JSON(http.StatusOK, movie)
	})

	router.Run((":5000"))

}
