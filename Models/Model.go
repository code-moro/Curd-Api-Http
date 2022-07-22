package Models

//Create Struct
type Movie struct {
	ID     int `json:"_id" bson:"_id"`
	Year   int           `json:"year,omitempty" bson:"year,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Dricetor string          `json:"director" bson:"director,omitempty"`
	Rating   float32          `json:"rating,omitempty" bson:"rating,omitempty"`
}

