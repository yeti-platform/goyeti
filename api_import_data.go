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
	"os"
)

// ImportDataAPIService ImportDataAPI service
type ImportDataAPIService service

type ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest struct {
	ctx          context.Context
	ApiService   *ImportDataAPIService
	mispFileJson *os.File
}

func (r ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest) MispFileJson(mispFileJson *os.File) ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest {
	r.mispFileJson = mispFileJson
	return r
}

func (r ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.ImportMispJsonApiV2ImportDataImportMispJsonPostExecute(r)
}

/*
ImportMispJsonApiV2ImportDataImportMispJsonPost Import Misp Json

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest
*/
func (a *ImportDataAPIService) ImportMispJsonApiV2ImportDataImportMispJsonPost(ctx context.Context) ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest {
	return ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return interface{}
func (a *ImportDataAPIService) ImportMispJsonApiV2ImportDataImportMispJsonPostExecute(r ApiImportMispJsonApiV2ImportDataImportMispJsonPostRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportDataAPIService.ImportMispJsonApiV2ImportDataImportMispJsonPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/import_data/import_misp_json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mispFileJson == nil {
		return localVarReturnValue, nil, reportError("mispFileJson is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var mispFileJsonLocalVarFormFileName string
	var mispFileJsonLocalVarFileName string
	var mispFileJsonLocalVarFileBytes []byte

	mispFileJsonLocalVarFormFileName = "misp_file_json"
	mispFileJsonLocalVarFile := r.mispFileJson

	if mispFileJsonLocalVarFile != nil {
		fbs, _ := io.ReadAll(mispFileJsonLocalVarFile)

		mispFileJsonLocalVarFileBytes = fbs
		mispFileJsonLocalVarFileName = mispFileJsonLocalVarFile.Name()
		mispFileJsonLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: mispFileJsonLocalVarFileBytes, fileName: mispFileJsonLocalVarFileName, formFileName: mispFileJsonLocalVarFormFileName})
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
