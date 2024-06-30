package errors

import "errors"

var (
	ErrNoMenuType = errors.New("Меню с таким наименованием не существует")
)
