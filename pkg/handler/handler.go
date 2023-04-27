package handler

import (
	"backend_ajax-people/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/sign-up", h.signUp)
	router.POST("/sign-in", h.signIn)
	router.POST("/test", h.test)

	apiPublic := router.Group("/api")
	{
		users := apiPublic.Group("/users")
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)

			activation := users.Group("/activation")
			{
				activation.POST("/check", h.checkActivationUser)
			}
		}

		registerData := api.Group("/register-data")

		faculty := apiPublic.Group("/faculty")
		{
			registerData.GET("/faculties", h.getAllFaculties)
			registerData.GET("/interests", h.getAllInterests)
			registerData.GET("/user-statuses", h.getAllStatuses)
			registerData.GET("/education-levels", h.getAllEdLevels)
			registerData.GET("/schools", h.getAllSchools)
		}
	}

	apiPrivate := router.Group("/apiP")
	{
		users := apiPrivate.Group("/users")
		{
			users.POST("/", h.createUser)
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.DELETE("/:id", h.deleteUser)
			users.PUT("/:id", h.updateUser)
		}
	}

	return router
}
