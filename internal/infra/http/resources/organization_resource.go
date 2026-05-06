package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type OrganizationDto struct {
	Id          int64      `json:"id"`
	UserId      int64      `json:"userId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	Lat         float64    `json:"lat"`
	Lon         float64    `json:"lon"`
	CreatedDate time.Time  `json:"createdDate"`
	UpdatedDate time.Time  `json:"updatedDate"`
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
}

type OrganizationsDto struct {
	Items []OrganizationDto `json:"items"`
	Total int64             `json:"total"`
}

func NewOrganizationDto(org *domain.Organization) OrganizationDto {
	return OrganizationDto{
		Id:          org.Id,
		UserId:      org.UserId,
		Name:        org.Name,
		Description: org.Description,
		City:        org.City,
		Address:     org.Address,
		Lat:         org.Lat,
		Lon:         org.Lon,
		CreatedDate: org.CreatedDate,
		UpdatedDate: org.UpdatedDate,
		DeletedDate: org.DeletedDate,
	}
}

func NewOrganizationsDtoList(orgs []*domain.Organization, total int64) OrganizationsDto {
	items := make([]OrganizationDto, len(orgs))
	for i, org := range orgs {
		items[i] = NewOrganizationDto(org)
	}
	return OrganizationsDto{
		Items: items,
		Total: total,
	}
}
