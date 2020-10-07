package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Application) handleAnswerUpdate(w http.ResponseWriter, r *http.Request) {
	var answer map[string]string
	answerID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	filter := bson.M{"_id": answerID}
	update := bson.D{
		{"$set", bson.D{{"contain", answer["contain"]}, {"updatedAt", time.Now().UTC().Unix()}}},
	}

	err = a.db.Update(answersCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated answer successfully")
}

func (a *Application) handleAnswerVoteUp(w http.ResponseWriter, r *http.Request) {
	answerID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	filter := bson.M{"_id": answerID}
	update := bson.D{
		{"$inc", bson.D{{"votes.up", 1}}},
	}

	err := a.db.Update(answersCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated answer vote up successfully")
}

func (a *Application) handleAnswerVoteDown(w http.ResponseWriter, r *http.Request) {
	answerID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	filter := bson.M{"_id": answerID}
	update := bson.D{
		{"$inc", bson.D{{"votes.down", 1}}},
	}

	err := a.db.Update(answersCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated answer vote down successfully")
}
