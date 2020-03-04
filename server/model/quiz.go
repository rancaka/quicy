package model

import (
	"errors"
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
	CurrentQuestion int         `json:"currentQuestionID"`
}

func GetQuizzes() ([]*Quiz, error) {

	quizzes := []*Quiz{}
	err := ReadJSON("model/quizzes.json", &quizzes)
	if err != nil {
		log.Printf("error on ReadJSON: %v", err)
		return nil, err
	}

	return quizzes, nil
}

func GetQuiz(quizID int) (*Quiz, error) {

	quizzes, err := GetQuizzes()
	if err != nil {
		log.Printf("error on getQuizzes: %v", err)
		return nil, err
	}

	for _, quiz := range quizzes {
		if quiz.QuizID == quizID {
			return quiz, nil
		}
	}

	return nil, errors.New("quiz not found")
}
