package domain

import "time"

/*User refere-se a abstração de um usuário com acesso ao sistema*/
type User struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	NickName      string    `json:"nick_name"`
	NumberOfPosts uint64    `json:"post_amount"`
	CreationDate  time.Time `json:"created_in"`
	LastPost      time.Time `json:"last_post_date"`
	LastActivity  time.Time `json:"last_login"`
	Avatar        string    `json:"avatar"`
	Role          string    `json:"role"`
	Mask          string    `json:"mask"`
	Signature     string    `json:"signature"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	Password      string    `json:"pwd"`
}
