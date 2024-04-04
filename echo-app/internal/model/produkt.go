package model

import (
    "gorm.io/gorm"
)

type Produkt struct {
    gorm.Model
    Nazwa       string
    Cena        float64
    Opis        string
    KategoriaID uint
    Kategoria   Kategoria `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

