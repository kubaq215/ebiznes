package service

import (
    "echo-app/internal/model"
    "echo-app/internal/repository"
)

// ProduktService definiuje interfejs dla serwisu produktów.
type ProduktService struct {
    repo *repository.ProduktRepository
}

// NewProduktService tworzy nową instancję ProduktService.
func NewProduktService(repo *repository.ProduktRepository) *ProduktService {
    return &ProduktService{
        repo: repo,
    }
}

// CreateProdukt tworzy nowy produkt.
func (s *ProduktService) CreateProdukt(produkt *model.Produkt) error {
    return s.repo.Create(produkt)
}

// GetProdukty zwraca wszystkie produkty.
func (s *ProduktService) GetProdukty() ([]model.Produkt, error) {
    return s.repo.GetAll()
}

// GetProdukt zwraca produkt na podstawie jego ID.
func (s *ProduktService) GetProdukt(id uint) (*model.Produkt, error) {
    return s.repo.GetByID(id)
}

// UpdateProdukt aktualizuje istniejący produkt.
func (s *ProduktService) UpdateProdukt(produkt *model.Produkt) error {
    return s.repo.Update(produkt)
}

// DeleteProdukt usuwa produkt na podstawie jego ID.
func (s *ProduktService) DeleteProdukt(id uint) error {
    return s.repo.Delete(id)
}

