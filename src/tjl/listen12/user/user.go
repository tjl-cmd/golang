package user

type User struct {
	Username  string
	Sex       string
	age       int
	AvatarUrl string
}

func NewUser(username string, sex string, age int, avatar string) *User {
	/*user := &User{
		Username:  username,
		Sex:       sex,
		Age:       age,
		AvatarUrl: avatar,
	}*/
	user := new(User)
	user.Username = username
	user.Sex = sex
	user.age = age
	user.AvatarUrl = avatar
	return user
}
