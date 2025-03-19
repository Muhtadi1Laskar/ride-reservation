package handlers

import (
	"net/http"
	"ride-reservation/helpers"
	"ride-reservation/services"
)

type PriceRequestBody struct {
	ServiceType     string  `json:"serviceType" validate:"required"`
	PickupDate      string  `json:"pickupDate" validate:"required"`
	PickupTime      string  `json:"pickupTime" validate:"required"`
	PickupLocation  string  `json:"pickupLocation" validate:"required"`
	DropoffLocation string  `json:"dropoffLocation" validate:"required"`
	Milage          float32 `json:"mileage" validate:"required"`
	Passengers      int     `json:"passengers" validate:"required"`
	Luggage         int     `json:"luggage" validate:"required"`
}

type CarType struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Seats int     `json:"seats"`
	Price float32 `json:"price"`
}

type CarPriceResponse struct {
	Vehicle []CarType `json:"vehicle"`
	Mileage float32   `json:"mileage"`
}

func Rides(w http.ResponseWriter, r *http.Request) {
	var requestBody PriceRequestBody

	if err := ReadRequestBody(r, &requestBody); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	mileage := services.GetMileage(requestBody.PickupLocation, requestBody.DropoffLocation)
	vehicles := helpers.GetVehicle(requestBody.Passengers)

	var carTypes []CarType
	for _, vehicle := range vehicles {
		carTypes = append(carTypes, CarType{
			Type:  vehicle["type"].(string),
			Name:  vehicle["name"].(string),
			Seats: vehicle["totalSeats"].(int),
			Price: calculatePrice(mileage, vehicle["price"].(int)),
		})
	}

	responseBody := CarPriceResponse{
		Vehicle: carTypes,
		Mileage: mileage,
	}

	WriteJSONResponse(w, http.StatusAccepted, responseBody)
}

func calculatePrice(mileage float32, vehicleFare int) float32 {
	var totalPrice float32 = 60.00

	if mileage > 2 {
		totalPrice += (mileage - 2) * 4

	}

	return totalPrice + float32(vehicleFare)
}
