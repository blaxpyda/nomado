package data

import "thugcorp.io/nomado/models"

type PropertyStorage interface {
	GetTopProperties()([]models.Property, error)
	GetRandomProperties()([]models.Property, error)
	// GetPropertyByID(id string) (models.Property, error)
	// SearchPropertiesByName(name string) ([]models.Property, error)
	// GetPropertiesByLocation(location string) ([]models.Property, error)
	// GetPropertiesByPriceRange(minPrice, maxPrice float64) ([]models.Property, error)
	// GetPropertiesByType(propertyType string) ([]models.Property, error)
}