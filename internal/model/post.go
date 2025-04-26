package model

type PostKind int

//go:generate go tool enumer -type PostKind -trimprefix PostKind -json
const (
	PostKindUndefined PostKind = iota
	PostKindKudo
	PostKindFine
)

type Post struct {
	Kind        PostKind `json:"kind"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}
