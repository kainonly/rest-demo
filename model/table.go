package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Table struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	ShopId          primitive.ObjectID `bson:"shop_id" json:"shop_id"`
	AreaId          primitive.ObjectID `bson:"area_id" json:"area_id"`
	Sn              string             `bson:"sn" json:"sn"`
	Alias           string             `bson:"alias" json:"alias"`
	Seats           int64              `bson:"seats" json:"seats"`
	MinimumSpending float64            `bson:"minimum_spending" json:"minimum_spending"`
	Runtime         int64              `bson:"runtime" json:"runtime"`
	Sort            int64              `bson:"sort" json:"sort"`
	CreateTime      time.Time          `bson:"create_time" json:"create_time"`
	UpdateTime      time.Time          `bson:"update_time" json:"update_time"`
}
