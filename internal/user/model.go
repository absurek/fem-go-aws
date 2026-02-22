package user

type RegisterUser struct {
	Username string `json:"username" dynamodbav:"username"`
	Password string `json:"password" dynamodbav:"password"`
}
