package server

import (
	"errors"
	"fmt"
	"net/http"
	"rock/config"
	"rock/tracker"
)

type httpError struct {
	Msg  string
	Code int
}

func (e *httpError) Error() string {
	return fmt.Sprintf("Error [%d]: %s", e.Code, e.Msg)
}

func Serve() error {
	http.HandleFunc("/start", handleStart)
	http.HandleFunc("/stop", handleStop)
	err := http.ListenAndServe(addr(), nil)
	if err != nil {
		return err
	}
	return nil
}

func addr() string {
	port := config.ServerPort()
	return fmt.Sprintf(":%d", port)
}

func handleStart(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	ticket := req.FormValue("ticket")
	if ticket == "" {
		http.Error(w, "Missing ticket", http.StatusBadRequest)
		return
	}
	comment := req.FormValue("comment")
	fmt.Printf("START: %s\n", ticket)
	tracker.Start(ticket, comment)
}

func handleStop(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	comment := req.FormValue("comment")
	fmt.Printf("STOP")
	tracker.Stop(comment)
}

func handleHttpError(w http.ResponseWriter, err error) bool {
	var httpError *httpError
	if err != nil {
		if errors.As(err, &httpError) {
			http.Error(w, httpError.Msg, httpError.Code)
			return true
		} else {
			http.Error(w, "Someting unexpected happened", http.StatusInternalServerError)
			return true
		}
	}
	return false
}
