package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"rock/config"
	"rock/tracker"
)

type httpError struct {
	Msg  string
	Code int
}

type requestData struct {
	Ticket  string `json:"ticket"`
	Comment string `json:"comment"`
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
	bodyRd, err := getBodyData(req)

	if handleHttpError(w, err) {
		return
	}

	comment := bodyRd.Comment
	ticket := bodyRd.Ticket

	if ticket == "" || comment == "" {
		formRd, err := getFormData(req)

		if handleHttpError(w, err) {
			return
		}

		if ticket == "" {
			ticket = formRd.Ticket
		}

		if comment == "" {
			comment = formRd.Comment
		}
	}

	if ticket == "" {
		if handleHttpError(w, &httpError{"Incomplete data", http.StatusBadGateway}) {
			return
		}
	}

	fmt.Printf("START: %s\n", ticket)
	tracker.Start(ticket, comment)
}

func handleStop(w http.ResponseWriter, req *http.Request) {
	bodyRd, err := getBodyData(req)

	if handleHttpError(w, err) {
		return
	}

	comment := bodyRd.Comment

	if comment == "" {
		formRd, err := getFormData(req)

		if handleHttpError(w, err) {
			return
		}

		comment = formRd.Comment
	}

	fmt.Printf("STOP\n")
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
			fmt.Println(err.Error())
			return true
		}
	}
	return false
}

func getFormData(req *http.Request) (requestData, error){
	if err := req.ParseForm(); err != nil {
		return requestData{}, &httpError{"Failed to parse form data", http.StatusBadRequest}
	}

	ticket := req.FormValue("ticket")
	comment := req.FormValue("comment")

	rd := requestData{
		Ticket:  ticket,
		Comment: comment,
	}

	return rd, nil
}

func getBodyData(req *http.Request) (requestData, error) {
	defer req.Body.Close()
	rd := requestData{}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return rd, err
	}

	if len(data) == 0 {
		return requestData{}, nil
	}

	err = json.Unmarshal(data, &rd)
	if err != nil {
		return rd, err
	}

	return rd, nil
}
