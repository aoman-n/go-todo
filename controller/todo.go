package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var Todos = []Todo{
	{Id: 1, Title: "number. 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: 2, Title: "number. 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{Id: 3, Title: "number. 3", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func IndexGET(c *gin.Context) {
	// res, err := json.Marshal(Todos)
	// if err != nil {
	// 	log.Fatalln("err: ", err)
	// 	c.String(http.StatusInternalServerError, "Server Error")
	// }
	// c.String(http.StatusOK, string(res))
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todos":  Todos,
	})
}

func PostsPost(c *gin.Context) {
	title := c.PostForm("title")
	id := Todos[len(Todos)-1].Id + 1
	todo := Todo{
		Id:        id,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	Todos = append(Todos, todo)
	log.Printf("create new todo. %v \n", todo)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"todo":   todo,
	})
}
