package handler

import (
    "github.com/labstack/echo/v4"
    "net/http"
    "fmt"
)

// PaymentRequest struct to capture incoming payment data
type PaymentRequest struct {
    Amount string `json:"amount"`
	CardNumber string `json:"cardNumber"`
    CartID string   `json:"cart_id"`
	Cvv string `json:"cvv"`
	ExpiryDate string `json:"expiryDate"`
}

// PaymentHandler handles the payment processing
func PaymentHandler(c echo.Context) error {
    var paymentRequest PaymentRequest
    if err := c.Bind(&paymentRequest); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid payment data"})
    }

    // Here you would typically call a payment API or service
    // Simulate a payment success
    fmt.Printf("Received payment for cart ID %d of amount %.2f\n", paymentRequest.CartID, paymentRequest.Amount)

    return c.JSON(http.StatusOK, echo.Map{"status": "payment successful"})
}
