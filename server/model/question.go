package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type QuestionType string

const singleAnswer QuestionType = "singleAnswer"

type Question struct {
	QuizID         int          `json:"quizID"`
	QuestionID     int          `json:"questionID"`
	NextQuestionID int          `json:"nextQuestionID"`
	QuestionType   QuestionType `json:"questionType"`
	Weight         int          `json:"weight"`
	Duration       int          `json:"duration"`
	IsPublished    bool         `json:"isPublished"`
	IsDone         bool         `json:"isDone"`
	Detail         Detail       `json:"detail"`
}

type DetailContent struct {
	ImageURL string `json:"imageURL"`
	Text     string `json:"text"`
	IsAnswer bool   `json:"isAnswer"`
}

type Detail struct {
	Questions []DetailContent `json:"questions"`
	Options   []DetailContent `json:"options"`
}

func GetQuestion(questionID int) (*Question, error) {

	b, err := ioutil.ReadFile("model/questions.json")
	if err != nil {
		log.Printf("error when ioutil.ReadFile getQuestion: %v; questionID: %v\n", err, questionID)
		return nil, err
	}

	questions := []*Question{}
	err = json.Unmarshal(b, &questions)
	if err != nil {
		log.Printf("error when json.Unmarshal getQuestion: %v; questionID: %v\n", err, questionID)
		return nil, err
	}

	for _, question := range questions {
		if question.QuestionID == questionID {
			if !question.IsPublished {
				return nil, errors.New("question not published")
			}
			return question, nil
		}
	}

	return nil, errors.New("question not found")
}
