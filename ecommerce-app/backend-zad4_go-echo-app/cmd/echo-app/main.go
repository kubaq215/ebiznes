package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "echo-app/internal/handler"
    "echo-app/internal/model"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "net/http"
)

func main() {
    // Initialization of Echo
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // CORS middleware setup to allow requests from the React frontend
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"}, // URL of the React app and all origins
        AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
        AllowHeaders: []string{"Content-Type", "Authorization"},
    }))

    // Database initialization
    db, err := gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
    if err != nil {
        e.Logger.Fatal("Error connecting to the database: ", err)
    }

    // Automatic migrations for models
    db.AutoMigrate(&model.Produkt{}, &model.Koszyk{}, &model.Kategoria{}) // Assuming models are already defined

    // Initialization of handlers (controllers)
    produktHandler := handler.NewProduktHandler(db)
    koszykHandler := handler.NewKoszykHandler(db)

    // Routing for root path
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "200 OK")
    })

    // Routing for products
    e.GET("/produkty", produktHandler.GetProdukty)
    e.POST("/produkty", produktHandler.CreateProdukt)
    e.PUT("/produkty/:id", produktHandler.UpdateProdukt)
    e.DELETE("/produkty/:id", produktHandler.DeleteProdukt)

    // Routing for cart
    e.GET("/koszyk/:id", koszykHandler.GetKoszyk)
    e.POST("/newkoszyk", koszykHandler.CreateKoszyk)
    e.POST("/koszyk/:id/:produkt_id", koszykHandler.AddItem)
    e.DELETE("/koszyk/:id/:produkt_id", koszykHandler.RemoveItem)
    e.DELETE("/koszyk/:id", koszykHandler.RemoveKoszyk)

    // Routing for payment
    e.POST("/pay", handler.PaymentHandler)  // Using PaymentHandler from the handler package

    // Log message that the server is running
    e.Logger.Info("Server is running on port :8080")
    // Start the server
    e.Logger.Fatal(e.Start(":8080"))
}
