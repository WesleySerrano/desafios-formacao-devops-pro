package main

type PlayerRepository struct {
	players []Player
}

func (playerRepository *PlayerRepository) save(p Player) Player {
	if p.Id <= 0 {
		p.Id = uint64(len(playerRepository.players) + 1)
	}
	playerRepository.players = append(playerRepository.players, p)

	return playerRepository.players[len(playerRepository.players)-1]
}

func (playerRepository PlayerRepository) findAll() []Player {
	return playerRepository.players
}

func (playerRepository PlayerRepository) findById(id uint64) *Player {

	for _, player := range playerRepository.players {
		if player.Id == id {
			return &player
		}
	}

	return nil
}
