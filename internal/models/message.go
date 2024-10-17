package models

type Message struct {
	Content string `json:"content"`
	Step    int    `json:"step"`
}
