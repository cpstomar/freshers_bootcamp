package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hello/Models"
	"net/http"
)

//GetUsers ... Get all users
func GetMarks(c *gin.Context) {
	var marks []Models.Marks
	err := Models.GetAllMarks_(&marks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marks)
	}
}

//CreateUser ... Create User
func CreateMark(c *gin.Context) {
	var mark Models.Marks
	c.BindJSON(&mark)
	err := Models.CreateMark_(&mark)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, mark)
	}
}

//GetUsers ... Get all users
func GetStudents(c *gin.Context) {
	var students []Models.Student
	err := Models.GetAllStudents_(&students)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

//CreateUser ... Create User
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent_(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//GetUserByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetStudentByID_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//UpdateUser ... Update the user information
func UpdateStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID_(&student, id)
	if err != nil {
		c.JSON(http.StatusNotFound, student)
	}
	c.BindJSON(&student)
	err = Models.UpdateStudent_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

//DeleteUser ... Delete the user
func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent_(&student, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}