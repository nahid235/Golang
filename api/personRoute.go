package api

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.POST("PersonDetails/add", AddPersonDetails())
	r.POST("PersonDetails/update", UpdatePersonDetails())
	r.GET("PersonDetails/view", ViewPersonDetails())
	r.POST("PersonDetails/delete", DeletePersonDetails())
}
