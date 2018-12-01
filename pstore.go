package main

const (
	ErrPlayerNotFound = PlayerStoreErr("player not found")
)

type PlayerStoreErr string

func (e PlayerStoreErr) Error() string {
	return string(e)
}

type PlayerStore struct {
	Collection map[string]Player
}

func (p PlayerStore) Get(id string) (Player, error) {
	player, ok := p.Collection[id]

	if !ok {
		return player, ErrPlayerNotFound
	}

	return player, nil
}
