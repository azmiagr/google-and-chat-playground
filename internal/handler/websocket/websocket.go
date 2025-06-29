package websocket

import (
	"google-login/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketHandker struct {
	hub *Hub
}

func NewWebSocketHandler(hub *Hub) *WebSocketHandker {
	return &WebSocketHandker{hub}
}

func (h *WebSocketHandker) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to upgrade connection", err)
		return
	}

	client := &Client{
		ClientID: uuid.New(),
		Conn:     conn,
		Send:     make(chan []byte),
	}

	h.hub.Register <- client

	go client.readPump(h.hub)
	go client.writePump()
}
