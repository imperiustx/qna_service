package app

import (
	"fmt"
	"net/http"
	"time"
)

func (a *Application) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	data := "health"
	a.RenderJSON(w, http.StatusOK, data)
}

func (a *Application) handleResourceNotFound(w http.ResponseWriter, r *http.Request) {
	a.RenderError(w, http.StatusConflict, "Error Resource Not Found")
}

func (a *Application) handleWelcome(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()
	date := now.Format("2006-01-02T15:04:05Z07:00")

	data := fmt.Sprintf("welcome to qna_backend at %s", date)
	a.RenderJSON(w, http.StatusOK, data)

}
