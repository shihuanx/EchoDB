package routers

import (
	"github.com/gin-gonic/gin"
	"memoryDataBase/controller"
)

func SetUpStudentRouter(studentController *controller.StudentController) *gin.Engine {
	r := gin.Default()
	// 创建一个学生组
	studentGroup := r.Group("/student")

	studentGroup.POST("", studentController.AddStudent)
	studentGroup.GET("/:id", studentController.GetStudent)
	studentGroup.PUT("", studentController.UpdateStudent)
	studentGroup.DELETE("/:id", studentController.DeleteStudent)
	studentGroup.POST("/add-course", studentController.AddCourse)
	studentGroup.GET("/get-all-course", studentController.GetAllCourse)
	studentGroup.POST("/choose-course", studentController.ChooseCourse)

	r.GET("/JoinRaftCluster", studentController.JoinRaftCluster)

	r.GET("/LeaderHandleCommand", studentController.LeaderHandleCommand)

	r.GET("/GetLeaderPortAddress", studentController.GetLeaderPortAddress)

	r.GET("/DeleteFatalPeer", studentController.DeleteFatalPeer)

	r.GET("/loadCourseRemains", studentController.LoadCourseRemains)

	return r

}
