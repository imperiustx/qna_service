package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func env(name, alt string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return alt
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func llog(level, message string, args ...interface{}) {
	mBs, _ := json.Marshal(fmt.Sprintf(message, args...))
	fmt.Printf(`{"time":"%s","level":"%s","message":`+string(mBs)+"}\n",
		time.Now().UTC().Format(time.RFC3339), level)
}

//RenderJSON renders response data in json format
func (a *Application) RenderJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	d := map[string]interface{}{
		"success": true,
		"data":    data,
		"error":   "",
	}
	check(json.NewEncoder(w).Encode(d))
}

//RenderError renders response error in json format
func (a *Application) RenderError(w http.ResponseWriter, code int, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	d := map[string]interface{}{
		"success": false,
		"data":    nil,
		"error":   err,
	}
	check(json.NewEncoder(w).Encode(d))
}
