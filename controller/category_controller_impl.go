package controller

import (
	"course-golang-restful-api/helper"
	"course-golang-restful-api/model"
	"course-golang-restful-api/model/request"
	"course-golang-restful-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	categoryCreateRequest := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(httpRequest, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(httpRequest.Context(), categoryCreateRequest)
	response := model.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(httpRequest, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(httpRequest.Context(), categoryUpdateRequest)
	response := model.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(httpRequest.Context(), id)
	response := model.GeneralResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(httpRequest.Context(), id)
	response := model.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(httpRequest.Context())
	response := model.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(writer, response)
}
