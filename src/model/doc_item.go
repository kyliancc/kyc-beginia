package model

import "time"

type DocItem struct {
	ID        int       `json:"id"`
	Created   time.Time `json:"created"`
	Completed time.Time `json:"completed"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	Priority  int       `json:"priority"`
	Labels    []string  `json:"labels"`
	Done      bool      `json:"done"`
}
