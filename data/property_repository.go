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
			FROM properties
			ORDER BY price DESC
			LIMIT $1`
	return r.getProperties(query)
}

func (r *PropertyRepository) GetRandomProperties() ([]models.Property, error) {
	query := ` SELECT id, name, description, transaction_type, price, status, agent_id, location, property_type, area, bedrooms, bathrooms, created_at, updated_at
			FROM properties
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

var (
	ErrPropertyNotFound = errors.New("Property not found")
)
