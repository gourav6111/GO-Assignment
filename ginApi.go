package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var students []Student

func main() {

	students = append(students, Student{ID: "1", Name: "abc"})
	students = append(students, Student{ID: "2", Name: "def"})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		for _, item := range students {
			c.String(http.StatusOK, "\n Hello %s", item)
		}
	})

	r.POST("/", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")

		students = append(students, Student{ID: id, Name: name})
		for _, item := range students {
			c.String(http.StatusOK, "\n Hello %s", item)
		}
	})

	r.PUT("/", func(c *gin.Context) {
		id := c.Query("id")
		name := c.Query("name")

		for index, item := range students {
			if item.ID == id {
				students = append(students[:index], students[index+1:]...)
				var Student Student
				Student.ID = id
				Student.Name = name
				students = append(students, Student)

			}
		}
		for _, item := range students {
			c.String(http.StatusOK, "\n Hello %s", item)
		}
	})

	r.DELETE("/", func(c *gin.Context) {
		id := c.Query("id")
		for index, item := range students {
			if item.ID == id {
				students = append(students[:index], students[index+1:]...)
			}
		}
		for _, item := range students {
			c.String(http.StatusOK, "\n Hello %s", item)
		}
	})
	r.Run()
}
