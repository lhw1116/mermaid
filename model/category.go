package model

import "time"

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (a *Category) TableName() string {
	return "feelingliu_category"
}
