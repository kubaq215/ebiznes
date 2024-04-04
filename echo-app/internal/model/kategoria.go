package model

import (
    "gorm.io/gorm"
)

type Kategoria struct {
    gorm.Model
    Nazwa    string
    Produkty []Produkt `gorm:"foreignKey:KategoriaID"`
}

