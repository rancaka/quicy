package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type WeightRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Quiz struct {
	QuizID          int         `json:"quizID"`
	WeightRange     WeightRange `json:"weightRange"`
	FirstQuestionID int         `json:"firstQuestionID"`
}

func GetQuiz(quizID int) (*Quiz, error) {

	b, err := ioutil.ReadFile("model/quizzes.json")
	if err != nil {
		log.Printf("error when ioutil.ReadFile getQuiz: %v\n", err)
		return nil, err
	}

	quizzes := []Quiz{}
	err = json.Unmarshal(b, &quizzes)
	if err != nil {
		log.Printf("error when json.Unmarshal getQuiz: %v\n", err)
		return nil, err
	}

	for _, quiz := range quizzes {
		if quiz.QuizID == quizID {
			return &quiz, nil
		}
	}

	return nil, errors.New("quiz not found")
}
