package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var Todos = map[int]Todo{
	1: {Id: 1, Title: "number. 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	2: {Id: 2, Title: "number. 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	3: {Id: 3, Title: "number. 3", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}
var LastId = 3

func TodosIndexEndpoint(c *gin.Context) {
	resTodos := make([]Todo, len(Todos))
	i := 0
	for _, todo := range Todos {
		resTodos[i] = todo
		i++
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todos":  resTodos,
	})
}

func TodosShowEndpoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "ng",
		})
	}
	resTodo, ok := Todos[id]
	if ok == false {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ng",
			"message": "todo not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todo":   resTodo,
	})
}

func TodosCreateEndpoint(c *gin.Context) {
	title := c.PostForm("title")
	id := LastId + 1
	LastId = id
	todo := Todo{
		Id:        id,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	Todos[LastId+1] = todo
	log.Printf("create new todo. %v \n", todo)
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"todo":   todo,
	})
}

func TodosDeleteEndpoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "ng",
		})
	}

	_, ok := Todos[id]
	if ok == false {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "ng",
			"message": "todo not found",
		})
	}

	delete(Todos, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "deleted.",
	})
}
