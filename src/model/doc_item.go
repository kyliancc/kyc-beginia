package model

import "time"

type DocItem struct {
	ID        int
	Created   time.Time
	Completed time.Time
	Name      string
	Comment   string
	Priority  int
	Labels    []string
	Done      bool
}
