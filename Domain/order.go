package domain

import "time"

type Order struct {
	ID             string    `json:"id" bson:"_id"`
	UserId         string    `json:"user_id" bson:"user_id"`
	EntryPrice     float32   `json:"entry_price" bson:"entry_price"`
	ExitPrice      float32   `json:"exit_price" bson:"exit_price"`
	Quanitity      float32   `json:"quanitity" bson:"quanitity"`
	EntryTimestamp time.Time `json:"entry_timestamp" bson:"entry_timestamp"`
	ExitTimestamp  time.Time `json:"exit_timestamp" bson:"exit_timestamp"`
}
