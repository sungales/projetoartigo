package main

type Artigo struct { 
	ID int `json:"id"`
	Descricao string `json:"descricao"`
	CreatedAt string `json:"created_at"`
}