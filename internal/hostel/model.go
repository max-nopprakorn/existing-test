package hostel

// Hostel stores an information of the hostel
type Hostel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Detail      string  `json:"detai"`
	Geolocation Map     `json:"geolocation"`
}

// Map stores the geolocation
type Map struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
