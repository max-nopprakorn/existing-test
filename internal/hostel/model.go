package hostel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Hostel stores an information of the hostel
type Hostel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Price       float32            `json:"price"`
	Detail      string             `json:"detai"`
	IsAvailable bool               `json:"available"`
	Geolocation Map                `json:"geolocation"`
}

// Map stores the geolocation
type Map struct {
	Latitude   float64 `json:"latitude"`
	Longtitude float64 `json:"longtitude"`
}
