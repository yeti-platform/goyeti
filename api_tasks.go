/*
FastAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goyeti

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TasksAPIService TasksAPI service
type TasksAPIService service

type ApiDeleteExportApiV2TasksExportExportNameDeleteRequest struct {
	ctx        context.Context
	ApiService *TasksAPIService
	exportName string
}

func (r ApiDeleteExportApiV2TasksExportExportNameDeleteRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.DeleteExportApiV2TasksExportExportNameDeleteExecute(r)
}

/*
DeleteExportApiV2TasksExportExportNameDelete Delete Export

Deletes an ExportTask.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param exportName
	@return ApiDeleteExportApiV2TasksExportExportNameDeleteRequest
*/
func (a *TasksAPIService) DeleteExportApiV2TasksExportExportNameDelete(ctx context.Context, exportName string) ApiDeleteExportApiV2TasksExportExportNameDeleteRequest {
	return ApiDeleteExportApiV2TasksExportExportNameDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		exportName: exportName,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *TasksAPIService) DeleteExportApiV2TasksExportExportNameDeleteExecute(r ApiDeleteExportApiV2TasksExportExportNameDeleteRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.DeleteExportApiV2TasksExportExportNameDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/export/{export_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"export_name"+"}", url.PathEscape(parameterValueToString(r.exportName, "exportName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiExportContentApiV2TasksExportExportIdContentGetRequest struct {
	ctx        context.Context
	ApiService *TasksAPIService
	exportId   string
}

func (r ApiExportContentApiV2TasksExportExportIdContentGetRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ExportContentApiV2TasksExportExportIdContentGetExecute(r)
}

/*
ExportContentApiV2TasksExportExportIdContentGet Export Content

Downloads the latest contents of a given ExportTask.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param exportId
	@return ApiExportContentApiV2TasksExportExportIdContentGetRequest
*/
func (a *TasksAPIService) ExportContentApiV2TasksExportExportIdContentGet(ctx context.Context, exportId string) ApiExportContentApiV2TasksExportExportIdContentGetRequest {
	return ApiExportContentApiV2TasksExportExportIdContentGetRequest{
		ApiService: a,
		ctx:        ctx,
		exportId:   exportId,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *TasksAPIService) ExportContentApiV2TasksExportExportIdContentGetExecute(r ApiExportContentApiV2TasksExportExportIdContentGetRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.ExportContentApiV2TasksExportExportIdContentGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/export/{export_id}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"export_id"+"}", url.PathEscape(parameterValueToString(r.exportId, "exportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiNewExportApiV2TasksExportNewPostRequest struct {
	ctx              context.Context
	ApiService       *TasksAPIService
	newExportRequest *NewExportRequest
}

func (r ApiNewExportApiV2TasksExportNewPostRequest) NewExportRequest(newExportRequest NewExportRequest) ApiNewExportApiV2TasksExportNewPostRequest {
	r.newExportRequest = &newExportRequest
	return r
}

func (r ApiNewExportApiV2TasksExportNewPostRequest) Execute() (*ExportTaskOutput, *http.Response, error) {
	return r.ApiService.NewExportApiV2TasksExportNewPostExecute(r)
}

/*
NewExportApiV2TasksExportNewPost New Export

Creates a new ExportTask in the database.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiNewExportApiV2TasksExportNewPostRequest
*/
func (a *TasksAPIService) NewExportApiV2TasksExportNewPost(ctx context.Context) ApiNewExportApiV2TasksExportNewPostRequest {
	return ApiNewExportApiV2TasksExportNewPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ExportTaskOutput
func (a *TasksAPIService) NewExportApiV2TasksExportNewPostExecute(r ApiNewExportApiV2TasksExportNewPostRequest) (*ExportTaskOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportTaskOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.NewExportApiV2TasksExportNewPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/export/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.newExportRequest == nil {
		return localVarReturnValue, nil, reportError("newExportRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.newExportRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPatchExportApiV2TasksExportExportNamePatchRequest struct {
	ctx                context.Context
	ApiService         *TasksAPIService
	exportName         string
	patchExportRequest *PatchExportRequest
}

func (r ApiPatchExportApiV2TasksExportExportNamePatchRequest) PatchExportRequest(patchExportRequest PatchExportRequest) ApiPatchExportApiV2TasksExportExportNamePatchRequest {
	r.patchExportRequest = &patchExportRequest
	return r
}

func (r ApiPatchExportApiV2TasksExportExportNamePatchRequest) Execute() (*ExportTaskOutput, *http.Response, error) {
	return r.ApiService.PatchExportApiV2TasksExportExportNamePatchExecute(r)
}

/*
PatchExportApiV2TasksExportExportNamePatch Patch Export

Pathes an existing ExportTask in the database.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param exportName
	@return ApiPatchExportApiV2TasksExportExportNamePatchRequest
*/
func (a *TasksAPIService) PatchExportApiV2TasksExportExportNamePatch(ctx context.Context, exportName string) ApiPatchExportApiV2TasksExportExportNamePatchRequest {
	return ApiPatchExportApiV2TasksExportExportNamePatchRequest{
		ApiService: a,
		ctx:        ctx,
		exportName: exportName,
	}
}

// Execute executes the request
//
//	@return ExportTaskOutput
func (a *TasksAPIService) PatchExportApiV2TasksExportExportNamePatchExecute(r ApiPatchExportApiV2TasksExportExportNamePatchRequest) (*ExportTaskOutput, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExportTaskOutput
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.PatchExportApiV2TasksExportExportNamePatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/export/{export_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"export_name"+"}", url.PathEscape(parameterValueToString(r.exportName, "exportName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchExportRequest == nil {
		return localVarReturnValue, nil, reportError("patchExportRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchExportRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRunApiV2TasksTaskNameRunPostRequest struct {
	ctx        context.Context
	ApiService *TasksAPIService
	taskName   interface{}
	taskParams *TaskParams
}

func (r ApiRunApiV2TasksTaskNameRunPostRequest) TaskParams(taskParams TaskParams) ApiRunApiV2TasksTaskNameRunPostRequest {
	r.taskParams = &taskParams
	return r
}

func (r ApiRunApiV2TasksTaskNameRunPostRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RunApiV2TasksTaskNameRunPostExecute(r)
}

/*
RunApiV2TasksTaskNameRunPost Run

Runs a task asynchronously.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taskName
	@return ApiRunApiV2TasksTaskNameRunPostRequest
*/
func (a *TasksAPIService) RunApiV2TasksTaskNameRunPost(ctx context.Context, taskName interface{}) ApiRunApiV2TasksTaskNameRunPostRequest {
	return ApiRunApiV2TasksTaskNameRunPostRequest{
		ApiService: a,
		ctx:        ctx,
		taskName:   taskName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *TasksAPIService) RunApiV2TasksTaskNameRunPostExecute(r ApiRunApiV2TasksTaskNameRunPostRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.RunApiV2TasksTaskNameRunPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/{task_name}/run"
	localVarPath = strings.Replace(localVarPath, "{"+"task_name"+"}", url.PathEscape(parameterValueToString(r.taskName, "taskName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.taskParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSearchApiV2TasksSearchPostRequest struct {
	ctx               context.Context
	ApiService        *TasksAPIService
	taskSearchRequest *TaskSearchRequest
}

func (r ApiSearchApiV2TasksSearchPostRequest) TaskSearchRequest(taskSearchRequest TaskSearchRequest) ApiSearchApiV2TasksSearchPostRequest {
	r.taskSearchRequest = &taskSearchRequest
	return r
}

func (r ApiSearchApiV2TasksSearchPostRequest) Execute() (*TaskSearchResponse, *http.Response, error) {
	return r.ApiService.SearchApiV2TasksSearchPostExecute(r)
}

/*
SearchApiV2TasksSearchPost Search

Searches for tasks.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchApiV2TasksSearchPostRequest
*/
func (a *TasksAPIService) SearchApiV2TasksSearchPost(ctx context.Context) ApiSearchApiV2TasksSearchPostRequest {
	return ApiSearchApiV2TasksSearchPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TaskSearchResponse
func (a *TasksAPIService) SearchApiV2TasksSearchPostExecute(r ApiSearchApiV2TasksSearchPostRequest) (*TaskSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TaskSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.SearchApiV2TasksSearchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taskSearchRequest == nil {
		return localVarReturnValue, nil, reportError("taskSearchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.taskSearchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiToggleApiV2TasksTaskNameTogglePostRequest struct {
	ctx        context.Context
	ApiService *TasksAPIService
	taskName   interface{}
}

func (r ApiToggleApiV2TasksTaskNameTogglePostRequest) Execute() (*ResponseToggleApiV2TasksTaskNameTogglePost, *http.Response, error) {
	return r.ApiService.ToggleApiV2TasksTaskNameTogglePostExecute(r)
}

/*
ToggleApiV2TasksTaskNameTogglePost Toggle

Toggles the enabled status on a task.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taskName
	@return ApiToggleApiV2TasksTaskNameTogglePostRequest
*/
func (a *TasksAPIService) ToggleApiV2TasksTaskNameTogglePost(ctx context.Context, taskName interface{}) ApiToggleApiV2TasksTaskNameTogglePostRequest {
	return ApiToggleApiV2TasksTaskNameTogglePostRequest{
		ApiService: a,
		ctx:        ctx,
		taskName:   taskName,
	}
}

// Execute executes the request
//
//	@return ResponseToggleApiV2TasksTaskNameTogglePost
func (a *TasksAPIService) ToggleApiV2TasksTaskNameTogglePostExecute(r ApiToggleApiV2TasksTaskNameTogglePostRequest) (*ResponseToggleApiV2TasksTaskNameTogglePost, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResponseToggleApiV2TasksTaskNameTogglePost
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TasksAPIService.ToggleApiV2TasksTaskNameTogglePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/tasks/{task_name}/toggle"
	localVarPath = strings.Replace(localVarPath, "{"+"task_name"+"}", url.PathEscape(parameterValueToString(r.taskName, "taskName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
