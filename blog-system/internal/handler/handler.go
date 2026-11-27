package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/internal/handler/middleware"
	"project/internal/service"
)

type BlogHandler struct {
	router  *gin.Engine
	service *service.Service
}

func NewHandler(router *gin.Engine, s *service.Service) *BlogHandler {
	return &BlogHandler{
		router:  router,
		service: s,
	}
}

func (h *BlogHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "The service is up and running"})
}

func (h *BlogHandler) InitRoutes() {
	h.router.Use(
		middleware.CORS(),
		middleware.Recovery(),
	)

	{
		h.router.GET("/health", h.HealthCheck)
		h.router.POST("/users/add", h.SignUp)
		h.router.POST("/users/login", h.SignIn)
	}

	v1 := h.router.Group("/v1")
	v1.Use(
		middleware.Authenticate(),
	)

	postGroup := v1.Group("/posts")
	{
		postGroup.GET("", h.GetPosts)
		postGroup.GET("/page/:page-id", h.GetPosts)
		postGroup.GET("/:id", h.GetPostByID)
		postGroup.POST("/add", h.AddPost)
		postGroup.DELETE("/delete/:id", h.DeletePost)
		postGroup.PUT("/update", h.EditPost)
	}

	userGroup := v1.Group("/users")
	{
		userGroup.GET("", h.GetUsers)
		userGroup.GET("/page/:page-id", h.GetUsers)
		userGroup.GET("/:id", h.GetUserByID)
		userGroup.GET("/:id/drafts", h.GetDrafts)
		userGroup.GET("/:id/drafts/page/:page-id", h.GetDrafts)
		userGroup.DELETE("/delete", h.DeleteUser)
		userGroup.PUT("/update", h.EditUser)
		userGroup.GET("/:id/posts", h.GetPostsByAuthor)
		userGroup.GET("/:id/posts/page/:page-id", h.GetPostsByAuthor)
		userGroup.GET("/search/:username/page/:page-id", h.GetUserByUsername)
		userGroup.GET("/search/:username", h.GetUserByUsername)
	}

	subscribeGroup := userGroup.Group("/subscribers/:user-id")
	{
		subscribeGroup.GET("", h.ViewSubscribers)
		subscribeGroup.GET("/page/:page-id", h.ViewSubscribers)
		subscribeGroup.POST("/add", h.Subscribe)
		subscribeGroup.DELETE("/remove", h.Unsubscribe)
	}

	commentGroup := v1.Group("/posts/:id/comments")
	{
		commentGroup.GET("", h.GetComments)
		commentGroup.GET("/page/:page-id", h.GetComments)
		commentGroup.GET("/recent", h.GetRecentPostComments)
		commentGroup.GET("/recent/page/:page-id", h.GetRecentPostComments)
		commentGroup.GET("/popular", h.GetPopularPostComments)
		commentGroup.GET("/popular/page/:page-id", h.GetRecentPostComments)
		commentGroup.GET("/parent/:comment-id", h.GetParentComment)
		commentGroup.GET("/children/:comment-id", h.GetChildrenComments)
		commentGroup.GET("/children/:comment-id/page/:page-id", h.GetChildrenComments)
		commentGroup.POST("/add", h.AddComment)
		commentGroup.PUT("/update", h.UpdateComment)
		commentGroup.DELETE("/delete/:comment-id", h.DeleteComment)
	}

	postLikesGroup := postGroup.Group("/likes/:id")
	{
		postLikesGroup.GET("", h.ListPostLikes)
		postLikesGroup.GET("/page/page-id", h.ListPostLikes)
		postLikesGroup.POST("/add", h.LikePost)
		postLikesGroup.DELETE("/remove", h.UnlikePost)
	}

	commentLikesGroup := commentGroup.Group("/likes/:comment-id")
	{
		commentLikesGroup.GET("", h.ListCommentLikes)
		commentLikesGroup.POST("/add", h.LikeComment)
		commentLikesGroup.DELETE("/remove", h.UnlikeComment)
	}

	profileGroup := v1.Group("/profiles")
	{
		profileGroup.GET("/:id", h.GetProfile)
		profileGroup.POST("/add", h.AddProfile)
		profileGroup.DELETE("/delete", h.DeleteProfile)
		profileGroup.PUT("/update", h.UpdateProfile)
	}

}

// cd blog-system/cmd/app

//func (h *BlogHandler) InitRoutes() {
//
//	// User Routes
//	{
//		h.mux.HandleFunc("/users", h.GetUsers)
//		h.mux.HandleFunc("/users/{id}", h.GetUserByID)
//		h.mux.HandleFunc("/users/add", h.AddUser)
//		h.mux.HandleFunc("/users/edit", h.EditUser)
//		h.mux.HandleFunc("/users/search", h.SearchUsername)
//	}
//
//	// Post Routes
//	{
//		h.mux.HandleFunc("/posts", h.GetPosts)
//		h.mux.HandleFunc("/posts/{id}", h.GetPostByID)
//		//h.mux.HandleFunc("/posts/search", h.GetPostsByTitle/Contents)
//		h.mux.HandleFunc("/posts/user/{id}", h.GetPostsByAuthor)
//		h.mux.HandleFunc("/posts/add", h.AddPost)
//		h.mux.HandleFunc("/posts/edit", h.EditPost)
//		h.mux.HandleFunc("/posts/{id}/delete", h.DeletePost)
//	}
//
//}
