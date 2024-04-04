package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TodoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status boolean            `json:"status,omitempty"`
}

