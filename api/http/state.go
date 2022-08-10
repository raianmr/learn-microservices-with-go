package http

import "gorm.io/gorm"

// TODO refactor `logic`
type State struct {
	DB  *gorm.DB
	Msg string
}
