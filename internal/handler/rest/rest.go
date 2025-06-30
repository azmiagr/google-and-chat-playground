package rest

import (
	"fmt"
	"google-login/internal/handler/websocket"
	"google-login/internal/service"
	"google-login/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	websocket  *websocket.WebSocketHandker
	middleware middleware.Interface
}

func NewRest(service *service.Service, wsHandler *websocket.WebSocketHandker, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		websocket:  wsHandler,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")

	auth := routerGroup.Group("/auth")
	auth.GET("/google", r.GoogleLogin)
	auth.GET("mangujo/callback/google", r.GoogleCallback)

	user := routerGroup.Group("/users")
	user.Use(r.middleware.AuthenticateUser)
	user.GET("/get-messages/:convoID", r.GetMessagesByConversationID)
	user.GET("/get-conversations", r.GetUserConversations)
	user.POST("/messages", r.SendMessage)
	user.POST("/create-conversation", r.CreateConversation)

	routerGroup.GET("/ws", r.websocket.HandleWebSocket)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
