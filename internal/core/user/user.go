// create user struct type
package user

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Uid      int    `json:"uid"`
}
