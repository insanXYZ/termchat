package converter

import (
	"backend/entity"
	"backend/model"
	"time"
)

func ChatToWriteMessage(chat *entity.Chat) *model.WriteMessage {
	return &model.WriteMessage{
		Sender: &model.SenderMessage{
			Name: chat.Sender.Name,
			ID:   chat.Sender.ID,
		},
		Receiver: &model.ReceiverMessage{
			Name: chat.Receiver.Name,
			ID:   chat.Receiver.ID,
		},
		Message: chat.Message,
		Time:    chat.CreatedAt.Format(time.DateTime),
		Type:    "private",
	}
}
