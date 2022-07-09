package repository

import (
	"github.com/bagasfathoni/go-sqlx/config"
	"github.com/bagasfathoni/go-sqlx/model"
	"github.com/jmoiron/sqlx"
)

type studentRepository struct {
	repo *sqlx.DB
}

type StudentRepository interface {
	GetAll() ([]model.Student, error)
}

func (s *studentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student
	err := s.repo.Select(&students, config.SELECT_ALL_STUDENTS)
	return students, err
}

func InitStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepository{repo: db}
}
