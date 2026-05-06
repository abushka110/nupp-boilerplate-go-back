package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type OrganizationController struct {
	organizationService *app.OrganizationService
}

func NewOrganizationController(os *app.OrganizationService) OrganizationController {
	return OrganizationController{
		organizationService: os,
	}
}

// Create organization
func (c OrganizationController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		type CreateOrgRequest struct {
			Name        string  `json:"name" validate:"required,gte=1,max=100"`
			Description string  `json:"description" validate:"max=500"`
			City        string  `json:"city" validate:"required,gte=1,max=100"`
			Address     string  `json:"address" validate:"required,gte=1,max=255"`
			Lat         float64 `json:"lat" validate:"required"`
			Lon         float64 `json:"lon" validate:"required"`
		}

		req := CreateOrgRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("OrganizationController.Create: %s", err)
			BadRequest(w, err)
			return
		}

		v := validator.New()
		if err := v.Struct(req); err != nil {
			log.Printf("OrganizationController.Create: %s", err)
			BadRequest(w, err)
			return
		}

		org := domain.Organization{
			UserId:      int64(user.Id),
			Name:        req.Name,
			Description: req.Description,
			City:        req.City,
			Address:     req.Address,
			Lat:         req.Lat,
			Lon:         req.Lon,
		}

		err := c.organizationService.CreateOrganization(&org)
		if err != nil {
			log.Printf("OrganizationController.Create: %s", err)
			InternalServerError(w, err)
			return
		}

		Created(w, resources.NewOrganizationDto(&org))
	}
}

// Get single organization
func (c OrganizationController) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("OrganizationController.Get: %s", err)
			BadRequest(w, err)
			return
		}

		org, err := c.organizationService.GetOrganization(id, int64(user.Id))
		if err != nil {
			log.Printf("OrganizationController.Get: %s", err)
			InternalServerError(w, err)
			return
		}

		if org == nil {
			NotFound(w, err)
			return
		}

		Success(w, resources.NewOrganizationDto(org))
	}
}

// Get all organizations with pagination
func (c OrganizationController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		page := uint64(1)
		countPerPage := uint64(10)

		if p := r.URL.Query().Get("page"); p != "" {
			if pageNum, err := strconv.ParseUint(p, 10, 64); err == nil {
				page = pageNum
			}
		}

		if c := r.URL.Query().Get("count"); c != "" {
			if count, err := strconv.ParseUint(c, 10, 64); err == nil {
				countPerPage = count
			}
		}

		pagination := &domain.Pagination{
			Page:         page,
			CountPerPage: countPerPage,
		}

		organizations, err := c.organizationService.GetAllOrganizations(int64(user.Id), pagination)
		if err != nil {
			log.Printf("OrganizationController.GetAll: %s", err)
			InternalServerError(w, err)
			return
		}

		total, err := c.organizationService.GetOrganizationCount(int64(user.Id))
		if err != nil {
			log.Printf("OrganizationController.GetAll: %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.NewOrganizationsDtoList(organizations, total))
	}
}

// Update organization
func (c OrganizationController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("OrganizationController.Update: %s", err)
			BadRequest(w, err)
			return
		}

		type UpdateOrgRequest struct {
			Name        string  `json:"name" validate:"required,gte=1,max=100"`
			Description string  `json:"description" validate:"max=500"`
			City        string  `json:"city" validate:"required,gte=1,max=100"`
			Address     string  `json:"address" validate:"required,gte=1,max=255"`
			Lat         float64 `json:"lat" validate:"required"`
			Lon         float64 `json:"lon" validate:"required"`
		}

		req := UpdateOrgRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("OrganizationController.Update: %s", err)
			BadRequest(w, err)
			return
		}

		v := validator.New()
		if err := v.Struct(req); err != nil {
			log.Printf("OrganizationController.Update: %s", err)
			BadRequest(w, err)
			return
		}

		org := domain.Organization{
			Id:          id,
			UserId:      int64(user.Id),
			Name:        req.Name,
			Description: req.Description,
			City:        req.City,
			Address:     req.Address,
			Lat:         req.Lat,
			Lon:         req.Lon,
		}

		err = c.organizationService.UpdateOrganization(&org)
		if err != nil {
			log.Printf("OrganizationController.Update: %s", err)
			InternalServerError(w, err)
			return
		}

		Success(w, resources.NewOrganizationDto(&org))
	}
}

// Delete organization
func (c OrganizationController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("OrganizationController.Delete: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.organizationService.DeleteOrganization(id, int64(user.Id))
		if err != nil {
			log.Printf("OrganizationController.Delete: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
