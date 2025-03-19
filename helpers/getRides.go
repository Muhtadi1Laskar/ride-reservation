package helpers

var Vehicles = []map[string]any{
	{
		"type": "Sedan",
		"name": "Toyota",
		"totalSeats": 5,
		"luggageCapacity": 4,
		"price": 10,
	},
	{
		"type": "Hatchback",
		"name": "Nissian",
		"totalSeats": 8,
		"luggageCapacity": 6,
		"price": 25,
	},
	{
		"type": "SUV",
		"name": "Mercedes",
		"totalSeats": 10,
		"luggageCapacity": 8,
		"price": 30,
	},
}

func GetVehicle(passengers, luggage int) []map[string]any {
	var selectedVehicles []map[string]any

	for _, item := range Vehicles {
		if item["totalSeats"].(int) >= passengers && item["luggageCapacity"].(int) >= luggage {
			selectedVehicles = append(selectedVehicles, item)
		}
	}

	return selectedVehicles
}