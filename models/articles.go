package models

import "time"

type Artigo struct {
	ID        int       `json:"id"`
	Descricao string    `json:"descricao" validate:"required, min=10"`
	CreatedAt time.Time `json:"created_at"`
}
