package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) sampleRoutes(router gin.IRoutes) {
	router.GET("/person/:person_id/info", s.userHandler.GetPersonInfoByPersonID())
	router.POST("/person/create", s.userHandler.CreatePerson())
}
