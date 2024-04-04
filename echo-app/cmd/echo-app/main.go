package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "echo-app/internal/handler"
    "echo-app/internal/model"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func main() {
    // Inicjalizacja Echo
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Inicjalizacja bazy danych
    db, err := gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
    if err != nil {
        e.Logger.Fatal("Error connecting to the database: ", err)
    }

    // Automatyczne migracje dla modeli
    db.AutoMigrate(&model.Produkt{}, &model.Koszyk{}, &model.Kategoria{}) // Zakładając, że modele są już zdefiniowane

    // Inicjalizacja handlerów (kontrolerów)
    produktHandler := handler.NewProduktHandler(db)
    koszykHandler := handler.NewKoszykHandler(db)

    // Routing dla produktów
    e.GET("/produkty", produktHandler.GetProdukty)
    e.POST("/produkty", produktHandler.CreateProdukt)
    e.PUT("/produkty/:id", produktHandler.UpdateProdukt)
    e.DELETE("/produkty/:id", produktHandler.DeleteProdukt)

    // Routing dla koszyka
    e.GET("/koszyk/:id", koszykHandler.GetKoszyk)
    e.POST("/newkoszyk", koszykHandler.CreateKoszyk)
    e.POST("/koszyk/:id/:produkt_id", koszykHandler.AddItem)
    e.DELETE("/koszyk/:id/:produkt_id", koszykHandler.RemoveItem)
    e.DELETE("/koszyk/:id", koszykHandler.RemoveKoszyk)

    // Write a message that the server is running
    e.Logger.Info("Server is running on port :8080")
    // Uruchomienie serwera
    e.Logger.Fatal(e.Start(":8080"))
}

