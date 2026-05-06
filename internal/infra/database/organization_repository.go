package database

import (
	"fmt"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const OrganizationsTableName = "organizations"

type organization struct {
	Id          int64      `db:"id,omitempty"`
	UserId      int64      `db:"user_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	City        string     `db:"city"`
	Address     string     `db:"address"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date,omitempty"`
	UpdatedDate time.Time  `db:"updated_date,omitempty"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type OrganizationRepository interface {
	Create(org *domain.Organization) error
	GetById(id int64) (*domain.Organization, error)
	GetAllByUserId(userId int64, pagination *domain.Pagination) ([]*domain.Organization, error)
	Update(org *domain.Organization) error
	Delete(id, userId int64) error
	GetCountByUserId(userId int64) (int64, error)
}

type organizationRepository struct {
	coll db.Collection
}

func NewOrganizationRepository(dbSession db.Session) OrganizationRepository {
	return organizationRepository{
		coll: dbSession.Collection(OrganizationsTableName),
	}
}

// Create adds a new organization to the database
func (r organizationRepository) Create(org *domain.Organization) error {
	model := r.mapDomainToModel(*org)
	err := r.coll.InsertReturning(&model)
	if err != nil {
		return err
	}
	org.Id = model.Id
	org.CreatedDate = model.CreatedDate
	org.UpdatedDate = model.UpdatedDate
	return nil
}

// GetById retrieves an organization by ID
func (r organizationRepository) GetById(id int64) (*domain.Organization, error) {
	var model organization
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&model)
	if err != nil {
		if err == db.ErrNoMoreRows {
			return nil, nil
		}
		return nil, err
	}
	org := r.mapModelToDomain(model)
	return &org, nil
}

// GetAllByUserId retrieves all organizations for a specific user with pagination
func (r organizationRepository) GetAllByUserId(userId int64, pagination *domain.Pagination) ([]*domain.Organization, error) {
	var models []organization
	offset := (pagination.Page - 1) * pagination.CountPerPage

	err := r.coll.Find(db.Cond{"user_id": userId, "deleted_date": nil}).
		OrderBy("-created_date").
		Limit(int(pagination.CountPerPage)).
		Offset(int(offset)).
		All(&models)

	if err != nil {
		return nil, err
	}

	organizations := make([]*domain.Organization, len(models))
	for i, model := range models {
		org := r.mapModelToDomain(model)
		organizations[i] = &org
	}

	return organizations, nil
}

// Update modifies an existing organization
func (r organizationRepository) Update(org *domain.Organization) error {
	org.UpdatedDate = time.Now()
	model := r.mapDomainToModel(*org)

	err := r.coll.Find(db.Cond{"id": org.Id, "user_id": org.UserId, "deleted_date": nil}).Update(&model)
	if err != nil {
		return err
	}

	return nil
}

// Delete performs soft delete of an organization
func (r organizationRepository) Delete(id, userId int64) error {
	now := time.Now()
	err := r.coll.Find(db.Cond{"id": id, "user_id": userId, "deleted_date": nil}).
		Update(map[string]interface{}{
			"deleted_date": &now,
			"updated_date": now,
		})

	if err != nil {
		if err == db.ErrNoMoreRows {
			return fmt.Errorf("organization not found")
		}
		return err
	}

	return nil
}

// GetCountByUserId returns the total count of organizations for a user
func (r organizationRepository) GetCountByUserId(userId int64) (int64, error) {
	count, err := r.coll.Find(db.Cond{"user_id": userId, "deleted_date": nil}).Count()
	return int64(count), err
}

func (r organizationRepository) mapDomainToModel(d domain.Organization) organization {
	return organization{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r organizationRepository) mapModelToDomain(m organization) domain.Organization {
	return domain.Organization{
		Id:          m.Id,
		UserId:      m.UserId,
		Name:        m.Name,
		Description: m.Description,
		City:        m.City,
		Address:     m.Address,
		Lat:         m.Lat,
		Lon:         m.Lon,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}
