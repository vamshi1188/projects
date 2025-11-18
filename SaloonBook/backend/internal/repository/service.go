package repository

import (
	"context"
	"saloonbook/internal/models"
)

type ServiceRepository struct {
	services []models.Service
}

func NewServiceRepository() *ServiceRepository {
	return &ServiceRepository{
		services: []models.Service{
			{ID: 1, Name: "Haircut", DurationMinutes: 30, Price: 20, Description: "Professional haircut service"},
			{ID: 2, Name: "Beard Trim", DurationMinutes: 20, Price: 15, Description: "Expert beard trimming and styling"},
			{ID: 3, Name: "Hair Color", DurationMinutes: 60, Price: 50, Description: "Premium hair coloring service"},
			{ID: 4, Name: "Massage", DurationMinutes: 45, Price: 35, Description: "Relaxing head and shoulder massage"},
			{ID: 5, Name: "Face Wash", DurationMinutes: 25, Price: 18, Description: "Deep cleansing facial treatment"},
		},
	}
}

func (r *ServiceRepository) List(ctx context.Context) ([]models.Service, error) {
	return r.services, nil
}

func (r *ServiceRepository) GetByID(ctx context.Context, id int) (*models.Service, error) {
	for _, s := range r.services {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, nil
}
