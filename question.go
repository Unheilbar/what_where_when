package what_where_when

type Question struct {
	Question   string   `json:"question"`
	Answer     string   `json:"correct_answer"`
	Difficulty string   `json:"difficulty"`
	Category   string   `json:"category"`
	Options    []string `json:"incorrect_answers"`
}

type TriviaApiResponse struct {
	Response_code uint `json:"response_code"`
	Results       []Question
}
