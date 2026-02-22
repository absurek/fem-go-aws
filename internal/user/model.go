package user

type UserData struct {
	Username string `json:"username" dynamodbav:"username"`
	Password string `json:"password" dynamodbav:"password"`
}

type User struct {
	Username     string `json:"username" dynamodbav:"username"`
	PasswordHash string `json:"password_hash" dynamodbav:"password_hash"`
}
