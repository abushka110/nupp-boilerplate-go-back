package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type CreateOrganizationRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=100"`
	Description string  `json:"description" validate:"max=500"`
	City        string  `json:"city" validate:"required,gte=1,max=100"`
	Address     string  `json:"address" validate:"required,gte=1,max=255"`
	Lat         float64 `json:"lat" validate:"required"`
	Lon         float64 `json:"lon" validate:"required"`
}

type UpdateOrganizationRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=100"`
	Description string  `json:"description" validate:"max=500"`
	City        string  `json:"city" validate:"required,gte=1,max=100"`
	Address     string  `json:"address" validate:"required,gte=1,max=255"`
	Lat         float64 `json:"lat" validate:"required"`
	Lon         float64 `json:"lon" validate:"required"`
}

func (r CreateOrganizationRequest) ToDomainModel(userId int64) (interface{}, error) {
	return domain.Organization{
		UserId:      userId,
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Address:     r.Address,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}

func (r UpdateOrganizationRequest) ToDomainModel(id, userId int64) (interface{}, error) {
	return domain.Organization{
		Id:          id,
		UserId:      userId,
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Address:     r.Address,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}
