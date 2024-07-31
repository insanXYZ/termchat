package websocket

import (
	"backend/entity"
	"backend/model"
	"backend/repository"
	"gorm.io/gorm"
)

type Hub struct {
	Clients        map[string]*Client
	Broadcast      chan *model.BroadcastMessage
	Register       chan *Client
	Cache          map[string]*entity.User
	Unregister     chan *Client
	DB             *gorm.DB
	UserRepository *repository.UserRepository
}

func NewHub(db *gorm.DB, repo *repository.UserRepository) *Hub {
	return &Hub{
		Broadcast:      make(chan *model.BroadcastMessage),
		Register:       make(chan *Client),
		Unregister:     make(chan *Client),
		Clients:        make(map[string]*Client),
		Cache:          make(map[string]*entity.User),
		DB:             db,
		UserRepository: repo,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.User.ID] = client
		case client := <-h.Unregister:
			if _, ok := h.Clients[client.User.ID]; ok {
				delete(h.Clients, client.User.ID)
			}
		case message := <-h.Broadcast:

			sendObj := &model.SendMessage{
				Sender:  message.Sender,
				Message: message.Message,
			}

			if message.Receiver == "global" {
				sendObj.Type = "global"

				for _, client := range h.Clients {
					client.Send <- sendObj
				}

			} else {

				if _, ok := h.Cache[message.Receiver]; !ok {

					cacheUsers := &entity.User{ID: message.Receiver}

					err := h.UserRepository.Take(h.DB, cacheUsers)
					if err != nil {
						return
					}

					h.Cache[message.Receiver] = cacheUsers
				}

				user := h.Cache[message.Receiver]

				sendObj.Type = "private"
				sendObj.Receiver = &model.ReceiverMessage{
					Name: user.Name,
					ID:   user.ID,
				}

				if receiver, okR := h.Clients[message.Receiver]; okR {
					receiver.Send <- sendObj
				}
				if sender, okS := h.Clients[message.Sender.ID]; okS {
					sender.Send <- sendObj
				}

			}

		}
	}
}
