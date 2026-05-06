package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type OrganizationService struct {
	repo database.OrganizationRepository
}

func NewOrganizationService(repo database.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// CreateOrganization creates a new organization
func (s *OrganizationService) CreateOrganization(org *domain.Organization) error {
	return s.repo.Create(org)
}

// GetOrganization retrieves an organization by ID
func (s *OrganizationService) GetOrganization(id, userId int64) (*domain.Organization, error) {
	org, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	if org == nil || org.UserId != userId {
		return nil, nil
	}

	return org, nil
}

// GetAllOrganizations retrieves all organizations for a user with pagination
func (s *OrganizationService) GetAllOrganizations(userId int64, pagination *domain.Pagination) ([]*domain.Organization, error) {
	return s.repo.GetAllByUserId(userId, pagination)
}

// UpdateOrganization updates an existing organization
func (s *OrganizationService) UpdateOrganization(org *domain.Organization) error {
	return s.repo.Update(org)
}

// DeleteOrganization deletes an organization
func (s *OrganizationService) DeleteOrganization(id, userId int64) error {
	return s.repo.Delete(id, userId)
}

// GetOrganizationCount returns the total count of organizations for a user
func (s *OrganizationService) GetOrganizationCount(userId int64) (int64, error) {
	return s.repo.GetCountByUserId(userId)
}
