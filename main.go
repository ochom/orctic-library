package database

import "github.com/ochom/orctic-database/repository"

// New ...
func New(dbPlatform repository.Platform) (repository.Repo, error) {
	return repository.New(dbPlatform)
}
