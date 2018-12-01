package main

const (
	ErrPlayerNotFound      = PlayerStoreErr("player not found")
	ErrPlayerAlreadyExists = PlayerStoreErr("player already exists")
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

func (p *PlayerStore) Add(player Player) (Player, error) {
	plyr, err := p.Get(player.Id)

	switch err {
	case ErrPlayerNotFound:
		p.Collection[player.Id] = player
		return player, nil
	case nil:
		return plyr, ErrPlayerAlreadyExists
	default:
		return plyr, err
	}
}

func (p *PlayerStore) Update(player Player) (Player, error) {
	plyr, err := p.Get(player.Id)

	switch err {
	case ErrPlayerNotFound:
		return plyr, ErrPlayerNotFound
	case nil:
		p.Collection[player.Id] = player
		return player, nil
	default:
		return plyr, err
	}
}
