package model

import (
    "gorm.io/gorm"
)

type Koszyk struct {
    gorm.Model
    Produkty []Produkt `gorm:"many2many:koszyk_produkty;"`
}

