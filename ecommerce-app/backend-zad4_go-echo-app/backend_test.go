package handler_test

import (
    "bytes"
    // "encoding/json"
    "net/http"
    "net/http/httptest"
    "strconv"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "echo-app/internal/handler"
    "echo-app/internal/model"
)

var produktJSON = `{"Nazwa":"Test Produkt","Opis":"Test Opis","Cena":10.0,"KategoriaID":1}`

func setupTestDB() *gorm.DB {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&model.Produkt{}, &model.Koszyk{}, &model.Kategoria{})
    return db
}

func setupEcho() *echo.Echo {
    e := echo.New()
    return e
}

func TestCreateProdukt(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    req := httptest.NewRequest(http.MethodPost, "/produkty", bytes.NewBufferString(produktJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.CreateProdukt(c))
    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestCreateProduktInvalidData(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    invalidJSON := `{"Nazwa":"", "Opis":"Test Opis", "Cena":10.0, "KategoriaID":1}`
    req := httptest.NewRequest(http.MethodPost, "/produkty", bytes.NewBufferString(invalidJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.CreateProdukt(c))
    assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetProduktyNoProducts(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    req := httptest.NewRequest(http.MethodGet, "/produkty", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.GetProdukty(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateProduktValidIDAndData(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)

    updateJSON := `{"Nazwa":"Updated Produkt","Opis":"Updated Opis","Cena":15.0,"KategoriaID":1}`
    req := httptest.NewRequest(http.MethodPut, "/produkty/"+strconv.Itoa(int(produkt.ID)), bytes.NewBufferString(updateJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.UpdateProdukt(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateProduktInvalidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    updateJSON := `{"Nazwa":"Updated Produkt","Opis":"Updated Opis","Cena":15.0,"KategoriaID":1}`
    req := httptest.NewRequest(http.MethodPut, "/produkty/999", bytes.NewBufferString(updateJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("999")

    assert.NoError(t, handler.UpdateProdukt(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteProduktValidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)

    req := httptest.NewRequest(http.MethodDelete, "/produkty/"+strconv.Itoa(int(produkt.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.DeleteProdukt(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteProduktInvalidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewProduktHandler(db)

    req := httptest.NewRequest(http.MethodDelete, "/produkty/999", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("999")

    assert.NoError(t, handler.DeleteProdukt(c))
    assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestCreateKoszyk(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    req := httptest.NewRequest(http.MethodPost, "/newkoszyk", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.CreateKoszyk(c))
    assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestRemoveKoszykValidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)

    req := httptest.NewRequest(http.MethodDelete, "/koszyk/"+strconv.Itoa(int(koszyk.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)))

    assert.NoError(t, handler.RemoveKoszyk(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRemoveKoszykInvalidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    req := httptest.NewRequest(http.MethodDelete, "/koszyk/999", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("999")

    assert.NoError(t, handler.RemoveKoszyk(c))
    assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestAddItemValidCartAndProductIDs(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)
    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)

    req := httptest.NewRequest(http.MethodPost, "/koszyk/"+strconv.Itoa(int(koszyk.ID))+"/"+strconv.Itoa(int(produkt.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)), strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.AddItem(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestAddItemInvalidCartID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)

    req := httptest.NewRequest(http.MethodPost, "/koszyk/999/"+strconv.Itoa(int(produkt.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues("999", strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.AddItem(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestAddItemInvalidProductID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)

    req := httptest.NewRequest(http.MethodPost, "/koszyk/"+strconv.Itoa(int(koszyk.ID))+"/999", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)), "999")

    assert.NoError(t, handler.AddItem(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestGetKoszykValidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)

    req := httptest.NewRequest(http.MethodGet, "/koszyk/"+strconv.Itoa(int(koszyk.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)))

    assert.NoError(t, handler.GetKoszyk(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetKoszykInvalidID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    req := httptest.NewRequest(http.MethodGet, "/koszyk/999", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id")
    c.SetParamValues("999")

    assert.NoError(t, handler.GetKoszyk(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestRemoveItemValidIDs(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)
    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)
    db.Model(&koszyk).Association("Produkty").Append(&produkt)

    req := httptest.NewRequest(http.MethodDelete, "/koszyk/"+strconv.Itoa(int(koszyk.ID))+"/"+strconv.Itoa(int(produkt.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)), strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.RemoveItem(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestRemoveItemInvalidCartID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    produkt := model.Produkt{Nazwa: "Test Produkt", Opis: "Test Opis", Cena: 10.0, KategoriaID: 1}
    db.Create(&produkt)

    req := httptest.NewRequest(http.MethodDelete, "/koszyk/999/"+strconv.Itoa(int(produkt.ID)), nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues("999", strconv.Itoa(int(produkt.ID)))

    assert.NoError(t, handler.RemoveItem(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestRemoveItemInvalidProductID(t *testing.T) {
    db := setupTestDB()
    e := setupEcho()
    handler := handler.NewKoszykHandler(db)

    koszyk := model.Koszyk{}
    db.Create(&koszyk)

    req := httptest.NewRequest(http.MethodDelete, "/koszyk/"+strconv.Itoa(int(koszyk.ID))+"/999", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)
    c.SetParamNames("id", "produkt_id")
    c.SetParamValues(strconv.Itoa(int(koszyk.ID)), "999")

    assert.NoError(t, handler.RemoveItem(c))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestPaymentHandlerValidPaymentData(t *testing.T) {
    e := setupEcho()
    req := httptest.NewRequest(http.MethodPost, "/pay", bytes.NewBufferString(`{"amount":"10.00","cardNumber":"4111111111111111","cvv":"123","expiryDate":"12/24","cart_id":"1"}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.PaymentHandler(c))
    assert.Equal(t, http.StatusOK, rec.Code)
}

func TestPaymentHandlerInvalidPaymentData(t *testing.T) {
    e := setupEcho()
    req := httptest.NewRequest(http.MethodPost, "/pay", bytes.NewBufferString(`{"amount":"","cardNumber":"4111111111111111","cvv":"123","expiryDate":"12/24","cart_id":"1"}`))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    assert.NoError(t, handler.PaymentHandler(c))
    assert.Equal(t, http.StatusBadRequest, rec.Code)
}

// func TestPaymentHandlerServiceUnavailable(t *testing.T) {
//     e := setupEcho()
//     req := httptest.NewRequest(http.MethodPost, "/pay", bytes.NewBufferString(`{"amount":"10.00","cardNumber":"4111111111111111","cvv":"123","expiryDate":"12/24","cart_id":"1"}`))
//     req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//     rec := httptest.NewRecorder()
//     c := e.NewContext(req, rec)

//     // Simulate service unavailable by changing the handler function
//     originalHandler := handler.PaymentHandler
//     handler.PaymentHandler = func(c echo.Context) error {
//         return c.JSON(http.StatusServiceUnavailable, "Payment service unavailable")
//     }
//     defer func() { handler.PaymentHandler = originalHandler }()

//     assert.NoError(t, handler.PaymentHandler(c))
//     assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
// }
