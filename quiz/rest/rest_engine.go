/*
 * =============================================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2022
 * The source code for this program is not published or otherwise divested of its trade secrets,
 * irrespective of what has been deposited with the U.S. Copyright Office.
 * =============================================================================================
 */

package rest

import (
	"net/http"
	middlewares "quiz/quiz/middlewares"
	definition "quiz/quiz/services/definition"

	"github.com/gin-gonic/gin"
)

func New(service definition.Quiz) *gin.Engine {
	engine := gin.New()
	publicRoutes := engine.Group("")
	quizRoutes := engine.Group("")
	adminRoutes := engine.Group("/admin")

	engine.Use(
		gin.Recovery(),
	)

	quizRoutes.Use(middlewares.JWTAuthCustomer())
	adminRoutes.Use(middlewares.JWTAuth())

	// root path to return OK response code
	publicRoutes.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	NewQuizController(service, quizRoutes)
	NewAdminController(service, adminRoutes)
	NewRegisterController(service, publicRoutes)

	return engine
}
