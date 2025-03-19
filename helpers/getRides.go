package helpers

var Vehicles = []map[string]any{
	{
		"type": "Sedan",
		"name": "Toyota",
		"totalSeats": 5,
		"price": 10,
	},
	{
		"type": "Hatchback",
		"name": "Nissian",
		"totalSeats": 8,
		"price": 25,
	},
	{
		"type": "SUV",
		"name": "Mercedes",
		"totalSeats": 10,
		"price": 30,
	},
}

func GetVehicle(passengers int) []map[string]any {
	var selectedVehicles []map[string]any

	for _, item := range Vehicles {
		if item["totalSeats"].(int) >= passengers {
			selectedVehicles = append(selectedVehicles, item)
		}
	}

	return selectedVehicles
}