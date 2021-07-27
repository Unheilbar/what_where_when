package service

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/unheilbar/what_where_when"
)

type QuestionsService struct {
	lastTokenUpdate time.Time
	token           string
}

func CreateQuestionsService() *QuestionsService {
	return &QuestionsService{
		lastTokenUpdate: time.Now(),
		token:           getToken(),
	}
}

var myClient = &http.Client{Timeout: 5 * time.Second}

func (q *QuestionsService) GetQuestionsList() ([]what_where_when.Question, error) {
	if time.Since(q.lastTokenUpdate) > 3600*5*time.Second {
		q.lastTokenUpdate = time.Now()
		q.token = getToken()
	}

	var resultapi what_where_when.TriviaApiResponse

	result := make([]what_where_when.Question, 10)

	url := "https://opentdb.com/api.php?amount=10&token=" + q.token

	res, err := myClient.Get(url)

	if err != nil {
		logrus.Errorf("Error on address %v, %v", url, err.Error())
		return result, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resultapi)

	if err != nil {
		logrus.Errorf("Error on address %v, %v", url, err.Error())
		return result, err
	}

	result = resultapi.Results

	return result, nil
}

func getToken() string {
	return "9aad7bdef0a4f4da105e31b87f4d3a6f0b5db6c14eca39c0a123b3137c549b6a"
}
