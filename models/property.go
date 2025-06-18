package models

type Property struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	TransactionType string `json:"transaction_type"`
	Price       float64 `json:"price"`
	Status      string `json:"status"`
	AgentID     string `json:"agent_id"`
	Location    string `json:"location"`
	PropertyType string `json:"property_type"`
	Area		float64 `json:"area"`
	Bedrooms    int     `json:"bedrooms"`
	Bathrooms   int     `json:"bathrooms"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}