package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"chat-api/internal/hub"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct {
	Hub *hub.Hub
}

func NewChatHandler(h *hub.Hub) *ChatHandler {
	return &ChatHandler{Hub: h}
}

func (h *ChatHandler) ServeWS(c *gin.Context) {
	userID := c.GetUint("user_id")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}
	client := &hub.Client{
		UserID: userID,
		Send:   make(chan []byte),
		Hub:    h.Hub,
	}
	h.Hub.Register(client)

	go writePump(client, conn)
	go readPump(client, conn)
}

func readPump(client *hub.Client, conn *websocket.Conn) {
	defer func() {
		client.Hub.Unregister(client)
		conn.Close()
	}()

	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			break
		}

		chatMsg := hub.ChatMessage{
			UserID:  client.UserID,
			Content: string(rawMessage),
		}

		encoded, err := json.Marshal(chatMsg)
		if err != nil {
			log.Printf("failed to encode message: %v", err)
			continue
		}

		client.Hub.Broadcast(encoded)
	}
}

func writePump(client *hub.Client, conn *websocket.Conn) {
	defer conn.Close()
	for message := range client.Send {
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
}
