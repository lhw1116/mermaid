package model

import "time"

type Tag struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (a *Tag) TableName() string {
	return "feelingliu_tag"
}
