package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Application) handleUserCreate(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	user.ID = primitive.NewObjectID()
	data, err := a.db.Create(usersCollection, user)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}

func (a *Application) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	var user User
	userID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": userID}
	update := bson.D{
		{"$set", bson.D{{"password", user.Password}}},
	}

	err = a.db.Update(usersCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated user successfully")
}
