package websocket

import (
	"backend/entity"
	"backend/model"
	"backend/repository"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Hub            *Hub
	User           *entity.User
	Conn           *websocket.Conn
	Send           chan *model.SendMessage
	ChatRepository *repository.ChatRepository
	DB             *gorm.DB
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		readObj := new(model.ReadMessage)
		err = json.Unmarshal(message, readObj)
		if err != nil {
			log.Println(err.Error())
			return
		}

		modelMessage := &model.BroadcastMessage{
			Message: readObj.Message,
			Sender: &model.SenderMessage{
				Name: c.User.Name,
				ID:   c.User.ID,
			},
			Receiver: readObj.Receiver,
		}

		c.Hub.Broadcast <- modelMessage

	}
}

func (c *Client) WritePump() {

	for {
		select {
		case message := <-c.Send:

			marshal, err := json.Marshal(model.WriteMessage{
				Sender:   message.Sender,
				Receiver: message.Receiver,
				Message:  string(message.Message),
				Time:     time.Now().Format("2006-1-2 15:4:5"),
				Type:     message.Type,
			})

			if err != nil {
				log.Println(err.Error())
			}

			err = c.DB.Transaction(func(tx *gorm.DB) error {
				if message.Type == "private" && message.Sender.ID == c.User.ID {
					chat := &entity.Chat{
						Message:    message.Message,
						SenderID:   message.Sender.ID,
						ReceiverID: message.Receiver.ID,
					}

					err := c.ChatRepository.Create(tx, chat)
					if err != nil {
						return err
					}
				}

				return c.Conn.WriteMessage(websocket.TextMessage, marshal)
			})

			if err != nil {
				log.Println(err.Error())
			}

		}
	}
}
