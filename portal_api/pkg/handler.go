package pkg

import (
	"encoding/json"
	"github.com/divoc/kernel_library/services"
	"github.com/divoc/portal-api/config"
	"github.com/divoc/portal-api/swagger_gen/restapi/operations"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GenericResponse struct {
	statusCode int
}

type GenericJsonResponse struct {
	body interface{}
}

func (o *GenericJsonResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	bytes, err := json.Marshal(o.body)
	if err != nil {
		rw.WriteHeader(500)
		rw.Write([]byte("JSON Marshalling error"))
	}
	rw.WriteHeader(200)
	rw.Write(bytes)
}

func (o *GenericResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses
	rw.WriteHeader(o.statusCode)
}

func NewGenericStatusOk() middleware.Responder {
	return &GenericResponse{statusCode: 200}
}

func NewGenericJSONResponse(body interface{}) middleware.Responder {
	return &GenericJsonResponse{body: body}
}

func NewGenericServerError() middleware.Responder {
	return &GenericResponse{statusCode: 500}
}
func SetupHandlers(api *operations.DivocPortalAPIAPI) {
	api.CreateMedicineHandler = operations.CreateMedicineHandlerFunc(createMedicineHandler)
	api.CreateProgramHandler = operations.CreateProgramHandlerFunc(createProgramHandler)
	api.PostFacilitiesHandler = operations.PostFacilitiesHandlerFunc(postFacilitiesHandler)
	api.PostVaccinatorsHandler = operations.PostVaccinatorsHandlerFunc(postVaccinatorsHandler)
	api.GetFacilitiesHandler = operations.GetFacilitiesHandlerFunc(getFacilitiesHandler)
	api.GetVaccinatorsHandler = operations.GetVaccinatorsHandlerFunc(getVaccinatorsHandler)
	api.GetMedicinesHandler = operations.GetMedicinesHandlerFunc(getMedicinesHandler)
	api.GetProgramsHandler = operations.GetProgramsHandlerFunc(getProgramsHandler)
	api.PostEnrollmentsHandler = operations.PostEnrollmentsHandlerFunc(postEnrollmentsHandler)
	api.CreateFacilityUsersHandler = operations.CreateFacilityUsersHandlerFunc(createFacilityUserHandler)
	api.GetFacilityUsersHandler = operations.GetFacilityUsersHandlerFunc(getFacilityUserHandler)
	api.GetFacilityGroupsHandler = operations.GetFacilityGroupsHandlerFunc(getFacilityGroupHandler)
	api.GetEnrollmentsHandler = operations.GetEnrollmentsHandlerFunc(getEnrollmentsHandler)
	api.UpdateFacilitiesHandler = operations.UpdateFacilitiesHandlerFunc(updateFacilitiesHandler)
}

func getEnrollmentsHandler(params operations.GetEnrollmentsParams, principal interface{}) middleware.Responder {
	return services.GetEntityType("Enrollment")
}

func getProgramsHandler(params operations.GetProgramsParams, principal interface{}) middleware.Responder {
	return services.GetEntityType("Program")
}

func getMedicinesHandler(params operations.GetMedicinesParams, principal interface{}) middleware.Responder {
	return services.GetEntityType("Medicine")
}

func getVaccinatorsHandler(params operations.GetVaccinatorsParams, principal interface{}) middleware.Responder {
	return services.GetEntityType("Vaccinator")
}

func getFacilitiesHandler(params operations.GetFacilitiesParams, principal interface{}) middleware.Responder {
	return services.GetEntityType("Facility")
}

func createMedicineHandler(params operations.CreateMedicineParams, principal interface{}) middleware.Responder {
	log.Infof("Create medicine %+v", params.Body)
	objectId := "Medicine"
	requestBody, err := json.Marshal(params.Body)
	if err != nil {
		return operations.NewCreateMedicineBadRequest()
	}
	requestMap := make(map[string]interface{})
	err = json.Unmarshal(requestBody, &requestMap)
	if err != nil {
		log.Info(err)
		return NewGenericServerError()
	}
	return services.MakeRegistryCreateRequest(requestMap, objectId)
}

func createProgramHandler(params operations.CreateProgramParams, principal interface{}) middleware.Responder {
	log.Infof("Create Program %+v", params.Body)
	objectId := "Program"
	requestBody, err := json.Marshal(params.Body)
	if err != nil {
		return operations.NewCreateProgramBadRequest()
	}
	requestMap := make(map[string]interface{})
	err = json.Unmarshal(requestBody, &requestMap)
	if err != nil {
		log.Info(err)
		return NewGenericServerError()
	}
	return services.MakeRegistryCreateRequest(requestMap, objectId)
}

func postEnrollmentsHandler(params operations.PostEnrollmentsParams, principal interface{}) middleware.Responder {
	data := NewScanner(params.File)
	defer params.File.Close()
	for data.Scan() {
		createEnrollment(&data)
		log.Info(data.Text("mobile"), data.Text("name"))
	}
	return operations.NewPostEnrollmentsOK()
}

func postFacilitiesHandler(params operations.PostFacilitiesParams, principal interface{}) middleware.Responder {
	data := NewScanner(params.File)
	defer params.File.Close()
	for data.Scan() {
		createFacility(&data, params.HTTPRequest.Header.Get("Authorization"))
		log.Info(data.Text("serialNum"), data.Text("facilityName"))
	}
	return operations.NewPostFacilitiesOK()
}

func postVaccinatorsHandler(params operations.PostVaccinatorsParams, principal interface{}) middleware.Responder {
	data := NewScanner(params.File)
	defer params.File.Close()
	for data.Scan() {
		createVaccinator(&data)
		log.Info("Created ", data.Text("serialNum"), data.Text("facilityName"))
	}
	return operations.NewPostFacilitiesOK()
}

func registryUrl(operationId string) string {
	url := config.Config.Registry.Url + "/" + operationId
	return url
}

func createFacilityUserHandler(params operations.CreateFacilityUsersParams, principal interface{}) middleware.Responder {
	err := CreateFacilityUser(params.Body, params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		log.Error(err)
		return operations.NewCreateFacilityUsersBadRequest()
	}
	return operations.NewCreateFacilityUsersOK()
}

func getFacilityUserHandler(params operations.GetFacilityUsersParams, principal interface{}) middleware.Responder {
	users, err := GetFacilityUsers(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		log.Error(err)
		return operations.NewCreateFacilityUsersBadRequest()
	}
	return &operations.GetFacilityUsersOK{Payload: users}
}

func getFacilityGroupHandler(params operations.GetFacilityGroupsParams, principal interface{}) middleware.Responder {
	groups, err := GetFacilityGroups(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		log.Error(err)
		return operations.NewGetFacilityGroupsBadRequest()
	}
	return &operations.GetFacilityGroupsOK{Payload: groups}
}

func updateFacilitiesHandler(params operations.UpdateFacilitiesParams, principal interface{}) middleware.Responder {
	for _, updateRequest := range params.Body {
		requestBody, err := json.Marshal(updateRequest)
		if err != nil {
			return operations.NewUpdateFacilitiesBadRequest()
		}
		requestMap := make(map[string]interface{})
		err = json.Unmarshal(requestBody, &requestMap)
		resp, err := services.UpdateRegistry("Facility", requestMap)
		if err != nil {
			log.Error(err)
		} else {
			log.Print(resp)
		}
	}
	return operations.NewUpdateFacilitiesOK()
}
