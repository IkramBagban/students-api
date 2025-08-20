package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *gorm.DB

func main() {
	fmt.Println("Hello, World!")

	var err error

	dsn := "host=localhost user=postgres password=pass123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database") // this will stop the application
	}
	db.AutoMigrate(&Student{}) // this will create the students table in the database

	r := gin.Default() // this will create a new gin router, meaning we can define our API routes

	// routes
	r.GET("/students", getStudents)
	r.POST("/student", createStudent)

	r.Run(":8000")
}

func getStudents(c *gin.Context) {
	var students []Student
	db.Find(&students)
	c.JSON(http.StatusOK, students)
}

func createStudent(c *gin.Context) {
	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}
	db.Create(&student)
	c.JSON(http.StatusCreated, student)

}
