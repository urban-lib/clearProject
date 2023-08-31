package entity

import "time"

type User struct {
	ID        int       `db:"id" goqu:"skipinsert,skipupdate"`
	Login     string    `db:"login"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"createdAt" goqu:"skipupdate"`
	LastLogin time.Time `db:"lastLogin"`
	IsActive  bool      `db:"isActive"`
}

func (this User) ToMap() map[string]interface{} {
	//TODO: use reflect to make map from model
	return nil
}
