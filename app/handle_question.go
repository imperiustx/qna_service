package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
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

func (a *Application) handleQuestionUpdate(w http.ResponseWriter, r *http.Request) {
	var question map[string]string
	questionID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	filter := bson.M{"_id": questionID}
	update := bson.D{
		{"$set", bson.D{{"contain", question["contain"]}, {"updatedAt", time.Now().UTC().Unix()}}},
	}

	err = a.db.Update(questionsCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated question successfully")
}

func (a *Application) handleQuestionVoteUp(w http.ResponseWriter, r *http.Request) {
	questionID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	filter := bson.M{"_id": questionID}
	update := bson.D{
		{"$inc", bson.D{{"votes.up", 1}}},
	}

	err := a.db.Update(questionsCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated question vote up successfully")
}

func (a *Application) handleQuestionVoteDown(w http.ResponseWriter, r *http.Request) {
	questionID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	filter := bson.M{"_id": questionID}
	update := bson.D{
		{"$inc", bson.D{{"votes.down", 1}}},
	}

	err := a.db.Update(questionsCollection, filter, update)
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, "updated question vote down successfully")
}

func (a *Application) handleGetAllQuestions(w http.ResponseWriter, r *http.Request) {

	data, err := a.db.FindAll(questionsCollection, bson.M{})
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}

func (a *Application) handleGetAllOpenQuestions(w http.ResponseWriter, r *http.Request) {

	data, err := a.db.FindAll(questionsCollection, bson.M{"openQuestion": true})
	if err != nil {
		a.RenderError(w, 400, err)
		return
	}

	a.RenderJSON(w, http.StatusOK, data)
}
