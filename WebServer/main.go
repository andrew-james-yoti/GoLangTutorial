package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"webserver.com/tutorial/joke"
)

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", JokeHandler)
		api.POST("/jokes/like/:jokeID", LikeJoke)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	c.JSON(http.StatusOK, joke.GetJokes())
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		joke.SetJokeLike(jokeid)

		c.JSON(http.StatusOK, joke.GetJokes())
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
