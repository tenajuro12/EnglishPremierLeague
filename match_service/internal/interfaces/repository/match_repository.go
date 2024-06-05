package repository

import (
	"EPL/match_service/internal/entity"
)

type MatchCRUDRepository interface {
	Insert(match entity.NewMatch) (entity.NewMatch, error)
	FindByID(id int) (entity.NewMatch, error)
	Update(match entity.NewMatch) (entity.NewMatch, error)
	Delete(id int) error
	GetAll() ([]entity.NewMatch, error)
}
