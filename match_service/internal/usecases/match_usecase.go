package usecases

import (
	"EPL/match_service/internal/entity"
	"EPL/match_service/internal/interfaces/repository"
	"time"
)

type CreateMatch struct {
	Repository repository.MatchCRUDRepository
}
type DeleteMatch struct {
	Repository repository.MatchCRUDRepository
}
type FindMatchByID struct {
	Repository repository.MatchCRUDRepository
}
type FindAllMatches struct {
	Repository repository.MatchCRUDRepository
}
type UpdateMatch struct {
	Repository repository.MatchCRUDRepository
}

func (uc *CreateMatch) Execute(match entity.NewMatch) (entity.NewMatch, error) {
	match.CreatedAt = time.Now()
	match.UpdatedAt = time.Now()
	return uc.Repository.Insert(match)
}
func (uc *DeleteMatch) Execute(id int) error {
	return uc.Repository.Delete(id)
}
func (uc *FindMatchByID) Execute(id int) (entity.NewMatch, error) {
	return uc.Repository.FindByID(id)
}
func (uc *FindAllMatches) Execute() ([]entity.NewMatch, error) {
	return uc.Repository.GetAll()
}
func (uc *UpdateMatch) Execute(match entity.NewMatch) (entity.NewMatch, error) {
	match.UpdatedAt = time.Now()
	return uc.Repository.Update(match)
}
