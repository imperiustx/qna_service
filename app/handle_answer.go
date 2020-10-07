package app

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *Application) handleAnswerCreate(w http.ResponseWriter, r *http.Request) {
	var answer Answer

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	answer.ID = primitive.NewObjectID()
	answer.QuestionID = iToObjectID(answer.QuestionID)
	answer.UserID = iToObjectID(answer.UserID)
	data, err := a.db.Create(answersCollection, answer)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}
