package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pokpak101/sa-64-example/controller"
	"github.com/pokpak101/sa-64-example/entity"
	"github.com/pokpak101/sa-64-example/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// User Routes
			protected.GET("/users", controller.ListUsers)
			protected.GET("/user/:id", controller.GetUser)
			protected.PATCH("/users", controller.UpdateUser)
			protected.DELETE("/users/:id", controller.DeleteUser)

			////

			protected.GET("/assesses", controller.ListAssess)

			protected.GET("/assess/:id", controller.GetAssess)

			protected.POST("/assesses", controller.CreateAssess)

			protected.PATCH("/assesses", controller.UpdateAssess)

			protected.DELETE("/assesses/:id", controller.DeleteAssess)

			///

			protected.GET("/cases", controller.ListCase)

			protected.GET("/case/:id", controller.GetCase)

			protected.POST("/cases", controller.CreateCase)

			protected.PATCH("/cases", controller.UpdateCase)

			protected.DELETE("/cases/:id", controller.DeleteCase)

			///

			protected.GET("/states", controller.ListState)

			protected.GET("/state/:id", controller.GetState)

			protected.POST("/states", controller.CreateState)

			protected.PATCH("/states", controller.UpdateState)

			protected.DELETE("/states/:id", controller.DeleteState)

			///

			protected.GET("/symptoms", controller.ListSymptom)

			protected.GET("/symptom/:id", controller.GetSymptom)

			protected.POST("/symptoms", controller.CreateSymptom)

			protected.PATCH("/symptoms", controller.UpdateSymptom)

			protected.DELETE("/symptoms/:id", controller.DeleteSymptom)

			///

			protected.GET("/assessment_sheets", controller.ListAssessmentSheet)

			protected.GET("/assessment_sheet/:id", controller.GetAssessmentSheet)

			protected.POST("/assessment_sheets", controller.CreateAssessmentSheet)

			protected.PATCH("/assessment_sheets", controller.UpdateAssessmentSheet)

			protected.DELETE("/assessment_sheets/:id", controller.DeleteAssessmentSheet)

		}
	}

	// User Routes
	r.POST("/users", controller.CreateUser)

	// Authentication Routes
	r.POST("/login", controller.Login)

	r.POST("/createassess", controller.CreateAssessmentSheet)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
