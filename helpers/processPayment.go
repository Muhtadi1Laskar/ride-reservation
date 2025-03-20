package helpers

import (
	"fmt"
)


func ProcessPayment(price float32, originalPrice float32) (string, error) {
	if price == originalPrice {
		return "Payment successful", nil
	}
	return "", fmt.Errorf("payment processing failed")
}