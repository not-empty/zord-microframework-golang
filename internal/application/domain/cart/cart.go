package cart

import "time"

type Cart struct {
	Id         string
	ModifiedAt time.Time
	UserId     string
}
