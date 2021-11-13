package api

import "time"

type PersonDetails struct {
	Guid    string `json:"guid" bson:"guid"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	//Age     int    `json:"age" bson:"age"`
	Company string `json:"company" bson:"company"`
	//CreatedAt time.Time `json:"created_at" bson:"created_at"`
	//UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	IsActive   bool      `json:"isActive" bson:"isActive"`
	Picuture   string    `json:"picture" bson:"picture"`
	Age        float64   `json:"age" bson:"age"`
	Gender     string    `json:"gender" bson:"gender"`
	Email      string    `json:"email" bson:"email"`
	Phone      string    `json:"phone" bson:"phone"`
	Registered time.Time `json:"registered" bson:"registered"`
}

var PersonList []PersonDetails
