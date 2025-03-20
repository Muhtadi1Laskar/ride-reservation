package handlers

import (
	"fmt"
	"net/http"
	"ride-reservation/helpers"
)

type PaymentRequestBody struct {
	PaymentID string  `json:"paymentId" validate:"required"`
	Amount    float32 `json:"amount" validate:"required"`
	Email     string  `json:"email" validate:"required"`
}

type FinalRequestBody struct {
	BasicInfo   PriceRequestBody   `json:"basicInfo" validate:"required"`
	Vehicle     CarType            `json:"vehicle" validate:"required"`
	PaymentInfo PaymentRequestBody `json:"paymentInfo" validate:"required"`
}

type RideDetails struct {
	ServiceType     string `json:"serviceType"`
	PickupDate      string `json:"pickupDate"`
	PickupTime      string `json:"pickupTime"`
	PickupLocation  string `json:"pickupLocation"`
	DropoffLocation string `json:"dropoffLocation"`
	Passengers      int    `json:"passengers"`
	Luggage         int    `json:"luggage"`
}

type VehicleInformation struct {
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Seats   int     `json:"seats"`
	Luggage int     `json:"luggage"`
	Price   float32 `json:"price"`
}

type Confirmation struct {
	Email           string `json:"email"`
	BookingRef      string `json:"bookingRef"`
	PaymentStatus   string `json:"paymentStatus"`
	EstimatedTravel string `json:"estimatedTravel,omitempty"`
}

type SuccessResponseBody struct {
	Message            string             `json:"message"`
	Instructions       string             `json:"instructions"`
	RideDetails        RideDetails        `json:"rideDetails"`
	VehicleInformation VehicleInformation `json:"vehicleInformation"`
	Confirmation       Confirmation       `json:"confirmation"`
}

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody FinalRequestBody

	if err := ReadRequestBody(r, &requestBody); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	checkPaymentStatus, err := helpers.ProcessPayment(requestBody.PaymentInfo.Amount, requestBody.Vehicle.Price)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	instruction := fmt.Sprintf("Check you email %s for the ride resertaion", requestBody.PaymentInfo.Email)

	response := SuccessResponseBody{
		Message:      checkPaymentStatus,
		Instructions: instruction,
		RideDetails: RideDetails{
			ServiceType: requestBody.BasicInfo.ServiceType,
			PickupDate: requestBody.BasicInfo.PickupDate,
			PickupTime: requestBody.BasicInfo.PickupTime,
			PickupLocation: requestBody.BasicInfo.PickupLocation,
			DropoffLocation: requestBody.BasicInfo.DropoffLocation,
			Passengers: requestBody.BasicInfo.Passengers,
			Luggage: requestBody.BasicInfo.Luggage,
		},
		VehicleInformation: VehicleInformation{
			Type: requestBody.Vehicle.Type,
			Name: requestBody.Vehicle.Name,
			Seats: requestBody.Vehicle.Seats,
			Luggage: requestBody.Vehicle.Luggage,
			Price: requestBody.Vehicle.Price,
		},
		Confirmation: Confirmation{
			Email: requestBody.PaymentInfo.Email,
			BookingRef: "Empty",
			PaymentStatus: "Completed",
			EstimatedTravel: "Empty",
		},
	}

	WriteJSONResponse(w, http.StatusAccepted, response)
}
