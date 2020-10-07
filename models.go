package main

// Admin is the model of administrator in this service
type Admin struct {
	Username  string `json:"userName" bson:"userName"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}

// User is the model of user in this service
type User struct {
	Username  string `json:"userName" bson:"userName"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64  `json:"updatedAt" bson:"updatedAt"`
}

// Question is the model of question in this service
type Question struct {
	Contain      string   `json:"contain" bson:"contain"`
	Votes        Vote     `json:"votes" bson:"votes"`
	CreatedAt    int64    `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64    `json:"updatedAt" bson:"updatedAt"`
	OpenQuestion bool     `json:"openQuestion" bson:"openQuestion"`
	Tags         []string `json:"tags" bson:"tags"`
}

// Answer is the model of answer in this service
type Answer struct {
	Contain        string `json:"contain" bson:"contain"`
	Votes          Vote   `json:"votes" bson:"votes"`
	CreatedAt      int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt      int64  `json:"updatedAt" bson:"updatedAt"`
	ApprovedAnswer bool   `json:"approvedAnswer" bson:"approvedAnswer"`
}

// Vote is the model of vote in this service
type Vote struct {
	Up   int16 `json:"up" bson:"up"`
	Down int16 `json:"down" bson:"down"`
}
