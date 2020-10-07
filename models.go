package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Admin is the model of administrator in this service
type Admin struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
}

// User is the model of user in this service
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64              `json:"updatedAt" bson:"updatedAt"`
}

// Question is the model of question in this service
type Question struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	Contain      string             `json:"contain" bson:"contain"`
	Votes        Vote               `json:"votes" bson:"votes"`
	CreatedAt    int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64              `json:"updatedAt" bson:"updatedAt"`
	OpenQuestion bool               `json:"openQuestion" bson:"openQuestion"`
	Tags         []string           `json:"tags" bson:"tags"`
}

// Answer is the model of answer in this service
type Answer struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
	QuestionID     primitive.ObjectID `json:"question_id" bson:"question_id"`
	Contain        string             `json:"contain" bson:"contain"`
	Votes          Vote               `json:"votes" bson:"votes"`
	CreatedAt      int64              `json:"createdAt" bson:"createdAt"`
	UpdatedAt      int64              `json:"updatedAt" bson:"updatedAt"`
	ApprovedAnswer bool               `json:"approvedAnswer" bson:"approvedAnswer"`
}

// Vote is the model of vote in this service
type Vote struct {
	Up   int16 `json:"up" bson:"up"`
	Down int16 `json:"down" bson:"down"`
}
