package handlers

import (
	"net/http"
)

type PriceRequestBody struct {
	ServiceType string `json:"serviceType" validate:"required"`
	PickupDate string `json:"pickupDate" validate:"required"`
	PickupTime string `json:"pickupTime" validate:"required"`
	PickupLocation string `json:"pickupLocation" validate:"required"`
	DropoffLocation string `json:"dropoffLocation" validate:"required"`
	Milage float32 `json:"mileage" validate:"required"`
	Passengers int `json:"passengers" validate:"required"`
	Luggage int `json:"luggage" validate:"required"`
}

type CarPriceResponse struct {
	CarType string `json:"carType"`
	TotalSeats int `json:"totalSeats"`
	Price int `json:"price"`
}

func Rides(w http.ResponseWriter, r *http.Request) {
	var requestBody PriceRequestBody

	if err := ReadRequestBody(r, &requestBody); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPrice := calculatePrice(requestBody.Milage)

	responseBody := CarPriceResponse{
		CarType: "Sedan",
		TotalSeats: 4,
		Price: int(totalPrice),
	}
	
	WriteJSONResponse(w, http.StatusAccepted, responseBody)
}

func calculatePrice(mileage float32) float32 {
	var totalPrice float32 = 60.00

	if mileage > 2 {
		totalPrice += (mileage - 2) * 4

	}
	return totalPrice
}