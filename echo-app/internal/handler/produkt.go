package handler

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "gorm.io/gorm"

    "echo-app/internal/model"
)

type ProduktHandler struct {
    DB *gorm.DB
}

func NewProduktHandler(db *gorm.DB) *ProduktHandler {
    return &ProduktHandler{DB: db}
}

// CreateProdukt dodaje nowy produkt do bazy danych
func (h *ProduktHandler) CreateProdukt(c echo.Context) error {
    produkt := new(model.Produkt)
    if err := c.Bind(produkt); err != nil {
        return c.JSON(http.StatusBadRequest, "Niepoprawne dane")
    }

    if result := h.DB.Create(&produkt); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error.Error())
    }

    return c.JSON(http.StatusCreated, produkt)
}

// GetProdukty zwraca listę wszystkich produktów
func (h *ProduktHandler) GetProdukty(c echo.Context) error {
    var produkty []model.Produkt
    if result := h.DB.Find(&produkty); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error.Error())
    }

    return c.JSON(http.StatusOK, produkty)
}

// UpdateProdukt aktualizuje produkt o podanym ID
func (h *ProduktHandler) UpdateProdukt(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Niepoprawne ID")
    }

    produkt := new(model.Produkt)
    if err := h.DB.First(&produkt, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Produkt nie znaleziony")
    }

    if err := c.Bind(produkt); err != nil {
        return c.JSON(http.StatusBadRequest, "Niepoprawne dane")
    }

    h.DB.Save(&produkt)

    return c.JSON(http.StatusOK, produkt)
}

// DeleteProdukt usuwa produkt o podanym ID
func (h *ProduktHandler) DeleteProdukt(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Niepoprawne ID")
    }

    if result := h.DB.Delete(&model.Produkt{}, id); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error.Error())
    }

    return c.JSON(http.StatusOK, "Produkt usunięty")
}
