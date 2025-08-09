package model

import "time"

type TodoDocItem struct {
	ID       int       `json:"id"`
	Created  time.Time `json:"created"`
	Name     string    `json:"name"`
	Comment  string    `json:"comment"`
	Priority int       `json:"priority"`
	Labels   []string  `json:"labels"`
}

type CpltDocItem struct {
	ID        int       `json:"id"`
	Created   time.Time `json:"created"`
	Completed time.Time `json:"completed"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	Labels    []string  `json:"labels"`
}

func Todo2CpltDocItem(item *TodoDocItem) *CpltDocItem {
	return &CpltDocItem{
		ID:        item.ID,
		Created:   item.Created,
		Completed: time.Time{},
		Name:      item.Name,
		Comment:   item.Comment,
		Labels:    item.Labels,
	}
}
