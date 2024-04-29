package users

import "time"

type Users struct {
	Id        int64
	Name      string
	CreatedOn time.Time
}
