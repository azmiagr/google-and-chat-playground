package rest

import (
	"fmt"
	"google-login/internal/handler/websocket"
	"google-login/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router    *gin.Engine
	service   *service.Service
	websocket *websocket.WebSocketHandker
}

func NewRest(service *service.Service, wsHandler *websocket.WebSocketHandker) *Rest {
	return &Rest{
		router:    gin.Default(),
		service:   service,
		websocket: wsHandler,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")

	auth := routerGroup.Group("/auth")
	auth.GET("/google", r.GoogleLogin)
	auth.GET("mangujo/callback/google", r.GoogleCallback)

	routerGroup.GET("/ws", r.websocket.HandleWebSocket)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
