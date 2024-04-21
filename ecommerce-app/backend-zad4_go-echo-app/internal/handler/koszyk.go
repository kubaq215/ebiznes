package handler

import (
    "net/http"
    "strconv"

    "github.com/labstack/echo/v4"
    "gorm.io/gorm"

    "echo-app/internal/model"
)

type KoszykHandler struct {
    DB *gorm.DB
}

func NewKoszykHandler(db *gorm.DB) *KoszykHandler {
    return &KoszykHandler{DB: db}
}

// CreateKoszyk tworzy nowy, pusty koszyk i zwraca jego ID.
func (h *KoszykHandler) CreateKoszyk(c echo.Context) error {
    // Tworzenie nowego koszyka
    koszyk := &model.Koszyk{}

    // Zapisanie nowego koszyka w bazie danych
    if err := h.DB.Create(koszyk).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Nie można utworzyć koszyka")
    }

    // Zwrócenie ID nowo utworzonego koszyka
    return c.JSON(http.StatusCreated, koszyk)
}

// RemoveKoszyk usuwa cały koszyk wraz z zawartością.
func (h *KoszykHandler) RemoveKoszyk(c echo.Context) error {
    koszykID, _ := strconv.Atoi(c.Param("id"))

    // Usunięcie koszyka (GORM automatycznie usunie wszystkie powiązane produkty dzięki kaskadowemu usuwaniu, jeśli zostało to odpowiednio skonfigurowane w modelu)
    if result := h.DB.Delete(&model.Koszyk{}, koszykID); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Nie można usunąć koszyka")
    }

    return c.JSON(http.StatusOK, "Koszyk usunięty")
}

// AddItem dodaje produkt do koszyka
func (h *KoszykHandler) AddItem(c echo.Context) error {
    koszykID, _ := strconv.Atoi(c.Param("id"))
    produktID, _ := strconv.Atoi(c.Param("produkt_id"))

    koszyk := &model.Koszyk{}
    if err := h.DB.Preload("Produkty").First(koszyk, koszykID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Koszyk nie znaleziony")
    }

    produkt := &model.Produkt{}
    if err := h.DB.First(produkt, produktID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Produkt nie znaleziony")
    }

    koszyk.Produkty = append(koszyk.Produkty, *produkt)

    if err := h.DB.Save(koszyk).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Nie można dodać produktu do koszyka")
    }

    return c.JSON(http.StatusOK, "Produkt dodany do koszyka")
}

// GetKoszyk zwraca zawartość koszyka
func (h *KoszykHandler) GetKoszyk(c echo.Context) error {
    koszykID, _ := strconv.Atoi(c.Param("id"))

    koszyk := &model.Koszyk{}
    if err := h.DB.Preload("Produkty").First(koszyk, koszykID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Koszyk nie znaleziony")
    }

    return c.JSON(http.StatusOK, koszyk.Produkty)
}

// RemoveItem usuwa produkt z koszyka
func (h *KoszykHandler) RemoveItem(c echo.Context) error {
    koszykID, _ := strconv.Atoi(c.Param("id"))
    produktID, _ := strconv.Atoi(c.Param("produkt_id"))

    // Wyszukiwanie koszyka
    koszyk := &model.Koszyk{}
    if err := h.DB.First(koszyk, koszykID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Koszyk nie znaleziony")
    }

    // Wyszukiwanie produktu, który ma zostać usunięty
    produkt := &model.Produkt{}
    if err := h.DB.First(produkt, produktID).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Produkt nie znaleziony")
    }

    // Usuwanie produktu z relacji
    if err := h.DB.Model(&koszyk).Association("Produkty").Delete(produkt); err != nil {
        return c.JSON(http.StatusInternalServerError, "Nie można usunąć produktu z koszyka")
    }

    return c.JSON(http.StatusOK, "Produkt usunięty z koszyka")
}

