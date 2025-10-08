package server

import (
	"fmt"
	"net/http"
	"rock/config"
	"rock/tracker"
)

func Serve() error {
	http.HandleFunc("/start", handleStart)
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
