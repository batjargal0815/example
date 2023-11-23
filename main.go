package user

type User struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetUserInfo(user *User) string {
	userInfo := user.UserName + " " + user.FirstName + " " + user.LastName
	return userInfo
}
