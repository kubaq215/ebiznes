package repository

import (
    "echo-app/internal/model"
    "gorm.io/gorm"
)

type ProduktRepository struct {
    DB *gorm.DB
}

func NewProduktRepository(db *gorm.DB) *ProduktRepository {
    return &ProduktRepository{
        DB: db,
    }
}

// Create dodaje nowy produkt do bazy danych
func (r *ProduktRepository) Create(produkt *model.Produkt) error {
    return r.DB.Create(produkt).Error
}

// GetAll zwraca wszystkie produkty z bazy danych
func (r *ProduktRepository) GetAll() ([]model.Produkt, error) {
    var produkty []model.Produkt
    result := r.DB.Find(&produkty)
    return produkty, result.Error
}

// GetByID zwraca produkt o podanym ID
func (r *ProduktRepository) GetByID(id uint) (*model.Produkt, error) {
    var produkt model.Produkt
    result := r.DB.First(&produkt, id)
    return &produkt, result.Error
}

// Update aktualizuje produkt w bazie danych
func (r *ProduktRepository) Update(produkt *model.Produkt) error {
    return r.DB.Save(produkt).Error
}

// Delete usuwa produkt o podanym ID z bazy danych
func (r *ProduktRepository) Delete(id uint) error {
    return r.DB.Delete(&model.Produkt{}, id).Error
}

