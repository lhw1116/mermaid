package model

import "time"

type Articles struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Category   string    `json:"category"`
	Tag        string    `json:"tag"`
	Author     string    `json:"author"`
	Context    string    `json:"context"`
	Html       string    `json:"html"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (a *Articles) TableName() string {
	return "feelingliu_article"
}
