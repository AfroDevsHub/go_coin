package warehouse

import "github.com/google/uuid"

type loginhistory struct {
	id       uuid.UUID
	login_id uuid.UUID
}

func (l *loginhistory) String() string {
	l.Dict()
	return "User ID: " + l.login_id.String()
}

func (l *loginhistory) Dict() map[string]interface{} {
	return map[string]interface{}{
		"login": l.login_id,
	}
}

type LoginHistory struct{}

func (*LoginHistory) Create_login_history() *loginhistory {
	return &loginhistory{
		id:       uuid.New(),
		login_id: uuid.New(),
	}

}
