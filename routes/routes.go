package routes

import (
	"belajar-gin/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/buku", controllers.Read)
}