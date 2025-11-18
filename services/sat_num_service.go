package services

import (
	"numberniceic/models"
	"numberniceic/repository"
)

// SatNumService defines the interface for the sat_num service.
type SatNumService interface {
	GetAllSatNums() ([]models.SatNum, error)
}

type satNumService struct {
	repo repository.SatNumRepository
}

// NewSatNumService creates a new SatNumService.
func NewSatNumService(repo repository.SatNumRepository) SatNumService {
	return &satNumService{repo: repo}
}

// GetAllSatNums retrieves all sat_num records.
func (s *satNumService) GetAllSatNums() ([]models.SatNum, error) {
	// ในอนาคต เราสามารถเพิ่ม business logic ตรงนี้ได้
	// แต่ตอนนี้ เราจะแค่ส่งต่อไปให้ repository
	satNums, err := s.repo.GetAllSatNums()
	if err != nil {
		return nil, err // ส่ง error กลับขึ้นไป
	}
	return satNums, nil
}
