package model

type IDParam struct {
	ID int `uri:"id" binding:"required"`
}
