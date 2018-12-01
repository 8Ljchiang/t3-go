package main

const (
	ErrBoardNotFound      = PlayerStoreErr("board not found")
	ErrBoardAlreadyExists = PlayerStoreErr("board already exists")
)

type BoardStoreErr string

func (e BoardStoreErr) Error() string {
	return string(e)
}

type BoardStore struct {
	Collection map[string]Board
}
