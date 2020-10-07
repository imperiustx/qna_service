package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Application) handleAdminCreate(w http.ResponseWriter, r *http.Request) {
	var admin Admin

	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}
	admin.ID = primitive.NewObjectID()

	data, err := a.db.Create(adminsCollection, admin)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}

func (a *Application) handleAdminUpdate(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	adminID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": adminID}
	update := bson.D{
		{"$set", bson.D{{"password", admin.Password}}},
	}

	err = a.db.Update(adminsCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated admin successfully")
}
