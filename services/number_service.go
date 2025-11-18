package services

import (
	"errors"
	"numberniceic/models"
	"numberniceic/repository"
)

type NumberService interface {
	GetAllNumbers() ([]models.Number, error)
	GetNumberByPairNumber(pairNumber string) (models.Number, error)
}

type numberService struct {
	repo repository.NumberRepository
}

func NewNumberService(repo repository.NumberRepository) NumberService {
	return &numberService{repo: repo}
}

func (s *numberService) GetNumberByPairNumber(pairNumber string) (models.Number, error) {

	if len(pairNumber) != 2 {
		return models.Number{}, errors.New("pairnumber must be 2 characters")
	}

	number, err := s.repo.GetByPairNumber(pairNumber)
	if err != nil {
		return models.Number{}, err
	}

	return number, nil
}

func (s *numberService) GetAllNumbers() ([]models.Number, error) {

	numbers, err := s.repo.GetAllNumbers()
	if err != nil {
		return nil, err
	}

	return numbers, nil
}
