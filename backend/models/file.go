package models

import "github.com/99designs/gqlgen/graphql"

type File struct {
	Path string `json:"path"`
}

type UploadFile struct {
	Content graphql.Upload `json:"content"`
}
