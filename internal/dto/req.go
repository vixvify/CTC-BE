package dto

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Teamname string `json:"teamname"`
	School   string `json:"school"`
	Call_1   string `json:"call_1"`
	Call_2   string `json:"call_2"`
	Relation string `json:"relation"`
	Name_1   string `json:"name_1"`
	Name_2   string `json:"name_2"`
	Name_3   string `json:"name_3"`
	Name_4   string `json:"name_4"`
	Verified string `json:"verified"`
	Video    string `json:"video"`
	Quiz_1   string `json:"quiz_1"`
	Quiz_2   string `json:"quiz_2"`
	Quiz_3   string `json:"quiz_3"`
	Quiz_4   string `json:"quiz_4"`
	Quiz_5   string `json:"quiz_5"`
}
