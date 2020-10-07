package app

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Application) handleQuestionCreate(w http.ResponseWriter, r *http.Request) {
	var question Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	question.ID = primitive.NewObjectID()
	question.UserID = iToObjectID(question.UserID)
	data, err := a.db.Create(questionsCollection, question)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}
