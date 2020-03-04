package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/rancaka/quicy/server/model"
)

type ResponseData struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/api/getAllQuiz", getQuizzes)
	http.HandleFunc("/api/getQuiz", getQuiz)
	http.HandleFunc("/api/getQuestion", getQuestion)
	http.ListenAndServe(":8080", nil)
}

func getQuizzes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	responseData := &ResponseData{}

	quizzes, err := model.GetQuizzes()
	if err != nil {
		responseData.Error = err.Error()
		encoder.Encode(responseData)
		return
	}

	responseData.Data = map[string]interface{}{
		"list": quizzes,
	}

	encoder.Encode(responseData)
}

func getQuiz(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	responseData := &ResponseData{}

	query := r.URL.Query()
	IDValues, ok := query["id"]
	if !ok {
		responseData.Error = errors.New("missing parameter id").Error()
		encoder.Encode(responseData)
		return
	}

	if len(IDValues) < 1 {
		responseData.Error = errors.New("missing value for id").Error()
		encoder.Encode(responseData)
		return
	}

	quizID, err := strconv.Atoi(IDValues[0])
	if err != nil {
		responseData.Error = err.Error()
		encoder.Encode(responseData)
		return
	}

	quiz, err := model.GetQuiz(quizID)
	if err != nil {
		responseData.Error = err.Error()
		encoder.Encode(responseData)
		return
	}

	responseData.Data = quiz
	encoder.Encode(responseData)
}

func getQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	responseData := &ResponseData{}

	query := r.URL.Query()
	IDValues, ok := query["id"]
	if !ok {
		responseData.Error = errors.New("missing parameter id").Error()
		encoder.Encode(responseData)
		return
	}

	if len(IDValues) < 1 {
		responseData.Error = errors.New("missing value for id").Error()
		encoder.Encode(responseData)
		return
	}

	questionID, err := strconv.Atoi(IDValues[0])
	if err != nil {
		responseData.Error = err.Error()
		encoder.Encode(responseData)
		return
	}

	question, err := model.GetQuestion(questionID)
	if err != nil {
		responseData.Error = err.Error()
		encoder.Encode(responseData)
		return
	}

	responseData.Data = question
	encoder.Encode(responseData)
}
