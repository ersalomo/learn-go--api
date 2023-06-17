package model

import "time"

type Book struct {
	ID        int
	Title     string `json:"id" binding:"string,required"`
	Publisher string `json:"publisher" binding:"string,required"`
	Year      int    `json:"Year" binding:"number,required"`
	Category  string `json:"category" binding:"string,required"`
	Qty       int    `json:"qty" binding:"number,required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
