package coffee

import (
	"database/sql"
	"fmt"

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

	rows, err := s.DB.Query("select id, variety, bitterness, description from coffee")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var c Coffee
		err = rows.Scan(&c.ID, &c.Variety, &c.Bitterness, &c.Description)
		if err != nil {
			return nil, err
		}
		result = append(result, &c)
	}
	return result, nil
}

func (s *Service) Get(ID int64) (*Coffee, error) {

	var c Coffee
	stmt, err := s.DB.Prepare("select id, variety, bitterness, description from coffee where coffee = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&c.ID, &c.Variety, &c.Bitterness, &c.Description)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *Service) Store(c *Coffee) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into coffee (variety, bitterness, description) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Variety, c.Bitterness, c.Description)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(c *Coffee) error {
	if c.ID == 0 {
		return fmt.Errorf("ID cant be iqual 0")
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update coffee set variety = ?, bitterness = ?, description = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Variety, c.Bitterness, c.Description, c.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("ID cant be iqual 0")
	}
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("delete from coffee where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
