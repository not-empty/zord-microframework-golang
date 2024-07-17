package handlers

import (
	requestContext "go-skeleton/internal/application/context"
	"go-skeleton/internal/application/domain/dummy"
	"go-skeleton/internal/application/providers/filters"
	"go-skeleton/internal/application/providers/pagination"
	_ "go-skeleton/internal/application/services"
	dummyCreate "go-skeleton/internal/application/services/dummy/CREATE"
	dummyDelete "go-skeleton/internal/application/services/dummy/DELETE"
	dummyEdit "go-skeleton/internal/application/services/dummy/EDIT"
	dummyGet "go-skeleton/internal/application/services/dummy/GET"
	dummyList "go-skeleton/internal/application/services/dummy/LIST"
	dummyRepository "go-skeleton/internal/repositories/dummy"
	"go-skeleton/pkg/idCreator"
	"go-skeleton/pkg/logger"
	"go-skeleton/pkg/registry"
	"go-skeleton/pkg/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DummyHandlers struct {
	DummyRepository *dummyRepository.DummyRepository

	logger    *logger.Logger
	idCreator *idCreator.IdCreator
	validator *validator.Validator
}

func NewDummyHandlers(reg *registry.Registry) *DummyHandlers {
	return &DummyHandlers{
		DummyRepository: reg.Inject("dummyRepository").(*dummyRepository.DummyRepository),
		logger:          reg.Inject("logger").(*logger.Logger),
		idCreator:       reg.Inject("idCreator").(*idCreator.IdCreator),
		validator:       reg.Inject("validator").(*validator.Validator),
	}
}

// HandleGetDummy Get Dummy
// @Summary      Get a Dummy
// @Tags         Dummy
// @Accept       json
// @Produce      json
// @Param        dummy_id path string true "Dummy ID"
// @Param        Tenant header string true "tenant name"
// @Success      200  {object}  dummyGet.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /dummy/{dummy_id} [get]
func (hs *DummyHandlers) HandleGetDummy(context echo.Context) error {
	s := dummyGet.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyGet.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(422, errors)
	}
	tenant := context.Get("tenant").(string)
	request := dummyGet.NewRequest(data)
	ctx := requestContext.NewPrepareContext(tenant)
	ctx.SetContext(request.Domain)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleCreateDummy Create Dummy
// @Summary      Create Dummy
// @Tags         Dummy
// @Accept       json
// @Produce      json
// @Param        request body dummyCreate.Data true "body model"
// @Param        Tenant header string true "tenant name"
// @Success      200  {object}  dummyCreate.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /dummy [post]
func (hs *DummyHandlers) HandleCreateDummy(context echo.Context) error {
	s := dummyCreate.NewService(hs.logger, hs.DummyRepository, hs.idCreator)
	data := new(dummyCreate.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	tenant := context.Get("tenant").(string)
	request := dummyCreate.NewRequest(data, hs.validator)
	ctx := requestContext.NewPrepareContext(tenant)
	ctx.SetContext(request.Domain)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusCreated, response)
}

// HandleEditDummy Edit Dummy
// @Summary      Edit Dummy
// @Tags         Dummy
// @Accept       json
// @Produce      json
// @Param        dummy_id path string true "Dummy ID"
// @Param        request body dummyEdit.Data true "body model"
// @Param        Tenant header string true "tenant name"
// @Success      200  {object}  dummyEdit.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /dummy/{dummy_id} [put]
func (hs *DummyHandlers) HandleEditDummy(context echo.Context) error {
	s := dummyEdit.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyEdit.Data)

	id := context.Param("id")
	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}

	tenant := context.Get("tenant").(string)
	request := dummyEdit.NewRequest(id, data, hs.validator)
	ctx := requestContext.NewPrepareContext(tenant)
	ctx.SetContext(request.Domain)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleListDummy List Dummy
// @Summary      List Dummy
// @Tags         Dummy
// @Accept       json
// @Produce      json
// @Param        page  query   int  true  "valid int"
// @Param        name  query   string  false  "value example: eql|lik,value"
// @Param        email  query   string  false  "value example: lik,value"
// @Param        Tenant header string true "tenant name"
// @Success      200  {object}  dummyList.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /dummy [get]
func (hs *DummyHandlers) HandleListDummy(context echo.Context) error {
	s := dummyList.NewService(
		hs.logger,
		hs.DummyRepository,
		pagination.NewPaginationProvider[dummy.Dummy](hs.DummyRepository),
	)

	data := new(dummyList.Data)
	bindErr := echo.QueryParamsBinder(context).
		Int("page", &data.Page).
		String("name", &data.Name).
		String("email", &data.Email).
		BindErrors()

	if bindErr != nil {
		return context.JSON(http.StatusBadRequest, bindErr)
	}

	f := filters.NewFilters()

	tenant := context.Get("tenant").(string)
	request := dummyList.NewRequest(data, f)
	ctx := requestContext.NewPrepareContext(tenant)
	ctx.SetContext(request.Domain)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}

// HandleDeleteDummy Delete Dummy
// @Summary      Delete Dummy
// @Tags         Dummy
// @Accept       json
// @Produce      json
// @Param        dummy_id path string true "Dummy ID"
// @Param        Tenant header string true "tenant name"
// @Success      200  {object}  dummyDelete.Response
// @Failure      400  {object}  services.Error
// @Failure      404  {object}  services.Error
// @Failure      500  {object}  services.Error
// @Router       /dummy/{dummy_id} [delete]
func (hs *DummyHandlers) HandleDeleteDummy(context echo.Context) error {
	s := dummyDelete.NewService(hs.logger, hs.DummyRepository)
	data := new(dummyDelete.Data)

	if errors := context.Bind(data); errors != nil {
		return context.JSON(http.StatusBadRequest, errors)
	}
	tenant := context.Get("tenant").(string)
	request := dummyDelete.NewRequest(data)
	ctx := requestContext.NewPrepareContext(tenant)
	ctx.SetContext(request.Domain)
	s.Execute(
		request,
	)

	response, err := s.GetResponse()
	if err != nil {
		return context.JSON(err.Status, err)
	}
	return context.JSON(http.StatusOK, response)
}
