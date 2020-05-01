package domain

import "time"

/*User refere-se a abstração de um usuário com acesso ao sistema*/
type User struct {
	id            uint64
	name          string
	nickName      string
	numberOfPosts uint64
	creationDate  time.Time
	lastPost      time.Time
	lastActivity  time.Time
	avatar        string
	role          string
	mask          string
	signature     string
	email         string
	status        string
	password      string
}
