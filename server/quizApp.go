package main

import (
	"log"

	"github.com/rancaka/quiz-app/server/model"
)

func main() {
	quiz, err := model.GetQuiz(1)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", quiz)

	question, err := model.GetQuestion(1)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v", question)
}
