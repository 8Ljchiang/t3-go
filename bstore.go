package main

const (
	ErrBoardNotFound      = PlayerStoreErr("board not found")
	ErrBoardAlreadyExists = PlayerStoreErr("board already exists")
)

type BoardStoreErr string

func (e BoardStoreErr) Error() string {
	return string(e)
}

type Store interface {
	Get(id string) (Board, error)
}

type BoardStore struct {
	Collection map[string]Board
}

func (b BoardStore) Get(id string) (Board, error) {
	board, ok := b.Collection[id]

	if !ok {
		return board, ErrBoardNotFound
	}

	return board, nil
}

func (b *BoardStore) Add(board Board) (Board, error) {
	brd, err := b.Get(board.Id)

	switch err {
	case ErrBoardNotFound:
		b.Collection[board.Id] = board
		return board, nil
	case nil:
		return brd, ErrBoardAlreadyExists
	default:
		return brd, err
	}
}

func (b *BoardStore) Update(board Board) (Board, error) {
	brd, err := b.Get(board.Id)

	switch err {
	case ErrBoardNotFound:
		return brd, ErrBoardNotFound
	case nil:
		b.Collection[board.Id] = board
		return board, nil
	default:
		return brd, err
	}
}

func (b *BoardStore) Remove(boardId string) {
	delete(b.Collection, boardId)
}
