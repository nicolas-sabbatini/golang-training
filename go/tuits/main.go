package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tuit struct {
	ID      int    `json:"id,string,omitempty"`
	User    string `json:"user"`
	Content string `json:"content"`
}

type ResError struct {
	Error string `json:"error"`
}

var tuits = []Tuit{
	{ID: 0, User: "Nico", Content: "Test"},
	{ID: 1, User: "GO", Content: "Hello from GO !"},
}

func getTuits(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tuits)
}

func getTuitById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		var err = ResError{Error: "Invalid ID"}
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	for _, tuit := range tuits {
		if tuit.ID == id {
			c.IndentedJSON(http.StatusOK, tuit)
			return
		}
	}
}

func postTuit(c *gin.Context) {
	var newTuit Tuit
	if err := c.BindJSON(&newTuit); err != nil {
		fmt.Println(err)
		return
	}
	if newTuit.User == "" || newTuit.Content == "" {
		var err = ResError{Error: "User and Content are required"}
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	newTuit.ID = len(tuits)
	tuits = append(tuits, newTuit)
	c.IndentedJSON(http.StatusCreated, newTuit)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/tuits", getTuits)
	r.GET("/tuits/:id", getTuitById)
	r.POST("/tuits", postTuit)

	r.Run()
}
