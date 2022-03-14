package coffee

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Coffee, error)
	Get(ID int64) (*Coffee, error)
	Store(c *Coffee) error
	Update(c *Coffee) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Coffee, error) {
	var result []*Coffee
	return result, nil
}

func (s *Service) Get(ID int64) (*Coffee, error) {
	var result *Coffee
	return result, nil
}

func (s *Service) Store(c *Coffee) error {
	return nil
}

func (s *Service) Update(c *Coffee) error {
	return nil
}

func (s *Service) Remove(ID int64) error {
	return nil
}
