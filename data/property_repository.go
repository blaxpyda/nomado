package data

import (
	"database/sql"
	"errors"
	"thugcorp.io/nomado/logger"
	"thugcorp.io/nomado/models"
)

type PropertyRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewPropertyRepository(db *sql.DB, log *logger.Logger) (*PropertyRepository, error) {
	return &PropertyRepository{
		db:     db,
		logger: log,
	}, nil
}

const defaultLimit = 3

func (r *PropertyRepository) GetTopProperties() ([]models.Property, error) {
	query := ` SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			ORDER BY price DESC
			LIMIT $1`
	return r.getProperties(query)
}

func (r *PropertyRepository) GetRandomProperties() ([]models.Property, error) {
	query := ` SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			ORDER BY RANDOM() DESC
			LIMIT $1`
	return r.getProperties(query)
}

func (r *PropertyRepository) getProperties(query string) ([]models.Property, error) {
	rows, err := r.db.Query(query, defaultLimit)
	if err != nil {
		r.logger.Error("Failed to query properties", err)
		return nil, err
	}
	defer rows.Close()

	var properties []models.Property
	for rows.Next() {
		var property models.Property
		if err := rows.Scan(&property.ID, &property.Name, &property.Description, &property.TransactionType, &property.Price, &property.Status, &property.AgentID, &property.Location, &property.PropertyType, &property.Area, &property.Bedrooms, &property.Bathrooms, &property.CreatedAt, &property.UpdatedAt); err != nil {
			r.logger.Error("Failed to scan property row", err)
			return nil, err
		}
		properties = append(properties, property)
	}
	return properties, nil
}

func (r *PropertyRepository) GetPropertyByID(id string) (models.Property, error) {
	var property models.Property
	query := `SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&property.ID, &property.Name, &property.Description, &property.TransactionType, &property.Price, &property.Status, &property.AgentID, &property.Location, &property.PropertyType, &property.Area, &property.Bedrooms, &property.Bathrooms, &property.CreatedAt, &property.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Property{}, ErrPropertyNotFound
		}
		r.logger.Error("Failed to query property by ID", err)
		return models.Property{}, err
	}
	return property, nil
}

func (r *PropertyRepository) SearchPropertiesByName(name string) ([]models.Property, error) {
	pattern := "%" + name + "%"
	query := `SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			WHERE name ILIKE $1`
	return r.getPropertiesWithParam(query, pattern)
}

func (r *PropertyRepository) GetPropertiesByLocation(location string) ([]models.Property, error) {
	pattern := "%" + location + "%"
	query := `SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			WHERE location ILIKE $1`
	return r.getPropertiesWithParam(query, pattern)
}


func (r *PropertyRepository) GetPropertiesByPriceRange(minPrice, maxPrice float64) ([]models.Property, error) {
	if minPrice < 0 || maxPrice < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if minPrice > maxPrice {
		return nil, errors.New("minPrice cannot be greater than maxPrice")
	}
	if minPrice == 0 && maxPrice == 0 {
		return nil, errors.New("at least one of minPrice or maxPrice must be specified")
	}
	query := `SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			WHERE price BETWEEN $1 AND $2`
	return r.getPropertiesWithParam(query, minPrice, maxPrice)
}

func (r *PropertyRepository) GetPropertiesByType(propertyType string) ([]models.Property, error) {
	if propertyType == "" {
		return nil, errors.New("property type cannot be empty")
	}
	pattern := "%" + propertyType + "%"
	query := `SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM property
			WHERE property_type ILIKE $1`
	return r.getPropertiesWithParam(query, pattern)
}

func (r *PropertyRepository) getPropertiesWithParam(query string, args ...interface{}) ([]models.Property, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		r.logger.Error("Failed to query properties with parameters", err)
		return nil, err
	}
	defer rows.Close()

	var properties []models.Property
	for rows.Next() {
		var property models.Property
		if err := rows.Scan(&property.ID, &property.Name, &property.Description, &property.TransactionType, &property.Price, &property.Status, &property.AgentID, &property.Location, &property.PropertyType, &property.Area, &property.Bedrooms, &property.Bathrooms, &property.CreatedAt, &property.UpdatedAt); err != nil {
			r.logger.Error("Failed to scan property row with parameters", err)
			return nil, err
		}
		properties = append(properties, property)
	}
	return properties, nil
}

var (
	ErrPropertyNotFound = errors.New("Property not found")
)
