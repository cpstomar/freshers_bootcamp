package Routes

import (
	"github.com/gin-gonic/gin"
	"hello/Controllers"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("")
	{
		grp1.GET("student", Controllers.GetStudents)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudent)
		grp1.DELETE("student/:id", Controllers.DeleteStudent)

		grp1.GET("mark", Controllers.GetMarks)
		grp1.POST("mark", Controllers.CreateMark)
	}
	return r
}
