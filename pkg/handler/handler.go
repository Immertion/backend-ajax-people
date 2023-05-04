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
	router.POST("/sign-out", h.signOut)
	router.POST("/test", h.test)

	apiPublic := router.Group("/api")
	{
		users := apiPublic.Group("/users", h.userIdentify)
		{
			users.GET("/", h.getAllUsers)
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)

			activation := users.Group("/activation")
			{
				activation.POST("/check", h.checkActivationUser)
			}

			users.POST("/select", h.selectUsers)
		}

		coincidence := apiPublic.Group("/coincidence")
		{
			coincidence.POST("/", h.coincidenceSend)
			coincidence.PUT("/:id", h.coincidenceAccept)
		}

		registerData := apiPublic.Group("/register-data")
		{
			registerData.GET("/faculties", h.getAllFaculties)
			registerData.GET("/interests", h.getAllInterests)
			registerData.GET("/user-statuses", h.getAllStatuses)
			registerData.GET("/education-levels", h.getAllEdLevels)
			registerData.GET("/schools", h.getAllSchools)
		}

		posts := apiPublic.Group("/posts")
		{
			posts.GET("/:id", h.getPostById)
			posts.GET("/", h.getAllPosts)
			posts.POST("/", h.createPost)
			posts.PUT("/:id", h.updatePost)
			posts.DELETE("/:id", h.deletePost)
		}

		tags := apiPublic.Group("/tags")
		{
			tags.GET("/:id", h.getTagById)
			tags.GET("/", h.getAllTags)
			tags.POST("/", h.createTag)
			tags.DELETE("/:id", h.deleteTag)
		}
	}

	apiPrivate := router.Group("/api-private", h.userIdentifyAdmin)
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
