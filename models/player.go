package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	Name 	   string
	Nickname   string
	ID         primitive.ObjectID `bson:"_id,omitempty"`
}
/*
pentakill int
bicrave int
deco-ping int
firstblood int
firstkill int
*/
