package models

type Artigo struct { 
	ID int `json:"id"`
	Descricao string `json:"descricao" validate:"required, min=10"`
	CreatedAt string `json:"created_at" validate:"required"`
}