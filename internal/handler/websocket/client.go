package websocket

import (
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID uuid.UUID
	Conn     *websocket.Conn
	Send     chan []byte
	UserID   int
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.Register <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}
		log.Printf("recv from %s: %s", c.ClientID, message)

		hub.Broadcast <- message
	}
}

func (c *Client) writePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("write error:", err)
			break
		}
	}
}
