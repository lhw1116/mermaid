package model

import "time"

type Articles struct {
	ID         int
	Title      string
	Category   string
	Tag        string
	Author     string
	Context    string
	Html       string
	CreateTime time.Time
	UpdateTime time.Time
}

func (a *Articles) TableName() string {
	return "feelingliu_article"
}
