/*
SEC form 13F API

API for SEC form filings such as 13F.

API version: 1.0.0
Contact: forms13f@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forms13f

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiApiV1FilerGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	cik *string
}

// The Central Index Key (CIK) of the filer.
func (r ApiApiV1FilerGetRequest) Cik(cik string) ApiApiV1FilerGetRequest {
	r.cik = &cik
	return r
}

func (r ApiApiV1FilerGetRequest) Execute() (*ApiV1Filer, *http.Response, error) {
	return r.ApiService.ApiV1FilerGetExecute(r)
}

/*
ApiV1FilerGet Retrieve a filer by CIK.

Returns a filer with the specified CIK (Central Index Key) and associated company names.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FilerGetRequest
*/
func (a *DefaultAPIService) ApiV1FilerGet(ctx context.Context) ApiApiV1FilerGetRequest {
	return ApiApiV1FilerGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiV1Filer
func (a *DefaultAPIService) ApiV1FilerGetExecute(r ApiApiV1FilerGetRequest) (*ApiV1Filer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiV1Filer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FilerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/filer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cik == nil {
		return localVarReturnValue, nil, reportError("cik is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "", "")
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

type ApiApiV1FilersGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	offset *int32
	limit *int32
}

// Skip previous offset CIKs
func (r ApiApiV1FilersGetRequest) Offset(offset int32) ApiApiV1FilersGetRequest {
	r.offset = &offset
	return r
}

// Return max limit CIKs
func (r ApiApiV1FilersGetRequest) Limit(limit int32) ApiApiV1FilersGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1FilersGetRequest) Execute() ([]ApiV1Filer, *http.Response, error) {
	return r.ApiService.ApiV1FilersGetExecute(r)
}

/*
ApiV1FilersGet Retrieve unique filers.

Returns a list of unique filers with their CIK and associated company names.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FilersGetRequest
*/
func (a *DefaultAPIService) ApiV1FilersGet(ctx context.Context) ApiApiV1FilersGetRequest {
	return ApiApiV1FilersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiV1Filer
func (a *DefaultAPIService) ApiV1FilersGetExecute(r ApiApiV1FilersGetRequest) ([]ApiV1Filer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiV1Filer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FilersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/filers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offset != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	} else {
        var defaultValue int32 = 0
        r.offset = &defaultValue
	}
	if r.limit != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
        var defaultValue int32 = 100
        r.limit = &defaultValue
	}
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

type ApiApiV1FilingsGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	from *string
	to *string
	offset *int32
	limit *int32
}

// All filings returned will be on or after this filing date.
func (r ApiApiV1FilingsGetRequest) From(from string) ApiApiV1FilingsGetRequest {
	r.from = &from
	return r
}

// All filings returned will be on or before this filing date.
func (r ApiApiV1FilingsGetRequest) To(to string) ApiApiV1FilingsGetRequest {
	r.to = &to
	return r
}

// Skip the first offset filings.
func (r ApiApiV1FilingsGetRequest) Offset(offset int32) ApiApiV1FilingsGetRequest {
	r.offset = &offset
	return r
}

// Return at most limit filings.
func (r ApiApiV1FilingsGetRequest) Limit(limit int32) ApiApiV1FilingsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1FilingsGetRequest) Execute() ([]ApiV1Form, *http.Response, error) {
	return r.ApiService.ApiV1FilingsGetExecute(r)
}

/*
ApiV1FilingsGet Retrieve 13F filings for all filers in the time range.

Get the 13F filings for all filers in the specified time range. Uses a form filing date as timestamp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FilingsGetRequest
*/
func (a *DefaultAPIService) ApiV1FilingsGet(ctx context.Context) ApiApiV1FilingsGetRequest {
	return ApiApiV1FilingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiV1Form
func (a *DefaultAPIService) ApiV1FilingsGetExecute(r ApiApiV1FilingsGetRequest) ([]ApiV1Form, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiV1Form
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FilingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/filings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "", "")
	if r.offset != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	} else {
        var defaultValue int32 = 0
        r.offset = &defaultValue
	}
	if r.limit != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
        var defaultValue int32 = 100
        r.limit = &defaultValue
	}
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

type ApiApiV1FormGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	accessionNumber *string
	cik *string
	offset *int32
	limit *int32
}

// The accession number of the form entry.
func (r ApiApiV1FormGetRequest) AccessionNumber(accessionNumber string) ApiApiV1FormGetRequest {
	r.accessionNumber = &accessionNumber
	return r
}

// The Central Index Key (CIK) of the form entry.
func (r ApiApiV1FormGetRequest) Cik(cik string) ApiApiV1FormGetRequest {
	r.cik = &cik
	return r
}

// The offset for pagination.
func (r ApiApiV1FormGetRequest) Offset(offset int32) ApiApiV1FormGetRequest {
	r.offset = &offset
	return r
}

// The limit for pagination.
func (r ApiApiV1FormGetRequest) Limit(limit int32) ApiApiV1FormGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1FormGetRequest) Execute() ([]ApiV1FormEntry, *http.Response, error) {
	return r.ApiService.ApiV1FormGetExecute(r)
}

/*
ApiV1FormGet Get SEC Form 13F.

Retrieve a content of form 13F by accession number and CIK.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FormGetRequest
*/
func (a *DefaultAPIService) ApiV1FormGet(ctx context.Context) ApiApiV1FormGetRequest {
	return ApiApiV1FormGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiV1FormEntry
func (a *DefaultAPIService) ApiV1FormGetExecute(r ApiApiV1FormGetRequest) ([]ApiV1FormEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiV1FormEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FormGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/form"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accessionNumber == nil {
		return localVarReturnValue, nil, reportError("accessionNumber is required and must be specified")
	}
	if r.cik == nil {
		return localVarReturnValue, nil, reportError("cik is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "accession_number", r.accessionNumber, "", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "", "")
	if r.offset != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	} else {
        var defaultValue int32 = 0
        r.offset = &defaultValue
	}
	if r.limit != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
        var defaultValue int32 = 100
        r.limit = &defaultValue
	}
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

type ApiApiV1FormsGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	cik *string
	from *string
	to *string
	offset *int32
	limit *int32
}

// The Central Index Key (CIK) of the filer.
func (r ApiApiV1FormsGetRequest) Cik(cik string) ApiApiV1FormsGetRequest {
	r.cik = &cik
	return r
}

// All forms returned will be on or after this period of report date.
func (r ApiApiV1FormsGetRequest) From(from string) ApiApiV1FormsGetRequest {
	r.from = &from
	return r
}

// All forms returned will be on or before this period of report date.
func (r ApiApiV1FormsGetRequest) To(to string) ApiApiV1FormsGetRequest {
	r.to = &to
	return r
}

// Skip the first offset forms.
func (r ApiApiV1FormsGetRequest) Offset(offset int32) ApiApiV1FormsGetRequest {
	r.offset = &offset
	return r
}

// Return at most limit forms.
func (r ApiApiV1FormsGetRequest) Limit(limit int32) ApiApiV1FormsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1FormsGetRequest) Execute() ([]ApiV1Form, *http.Response, error) {
	return r.ApiService.ApiV1FormsGetExecute(r)
}

/*
ApiV1FormsGet Retrieve SEC forms 13F for a filer.

This API returns SEC 13F forms between dates for a filer with CIK. Uses a period of report as timestamp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FormsGetRequest
*/
func (a *DefaultAPIService) ApiV1FormsGet(ctx context.Context) ApiApiV1FormsGetRequest {
	return ApiApiV1FormsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiV1Form
func (a *DefaultAPIService) ApiV1FormsGetExecute(r ApiApiV1FormsGetRequest) ([]ApiV1Form, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiV1Form
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FormsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/forms"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cik == nil {
		return localVarReturnValue, nil, reportError("cik is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "cik", r.cik, "", "")
	if r.from != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "", "")
	} else {
        var defaultValue string = "2010-01-01"
        r.from = &defaultValue
	}
	if r.to != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "", "")
	} else {
        var defaultValue string = "2030-01-01"
        r.to = &defaultValue
	}
	if r.offset != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	} else {
        var defaultValue int32 = 0
        r.offset = &defaultValue
	}
	if r.limit != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
        var defaultValue int32 = 100
        r.limit = &defaultValue
	}
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

type ApiApiV1FundsGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	name *string
	offset *int32
	limit *int32
}

// Filter companies by by a substring of their name
func (r ApiApiV1FundsGetRequest) Name(name string) ApiApiV1FundsGetRequest {
	r.name = &name
	return r
}

// Skip previous offset companies
func (r ApiApiV1FundsGetRequest) Offset(offset int32) ApiApiV1FundsGetRequest {
	r.offset = &offset
	return r
}

// Return max limit companies
func (r ApiApiV1FundsGetRequest) Limit(limit int32) ApiApiV1FundsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiApiV1FundsGetRequest) Execute() ([]ApiV1Fund, *http.Response, error) {
	return r.ApiService.ApiV1FundsGetExecute(r)
}

/*
ApiV1FundsGet Retrieve unique companies.

Returns a list of unique companies with their names and associated CIKs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1FundsGetRequest
*/
func (a *DefaultAPIService) ApiV1FundsGet(ctx context.Context) ApiApiV1FundsGetRequest {
	return ApiApiV1FundsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiV1Fund
func (a *DefaultAPIService) ApiV1FundsGetExecute(r ApiApiV1FundsGetRequest) ([]ApiV1Fund, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiV1Fund
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1FundsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/funds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "", "")
	}
	if r.offset != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "", "")
	} else {
        var defaultValue int32 = 0
        r.offset = &defaultValue
	}
	if r.limit != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
        var defaultValue int32 = 100
        r.limit = &defaultValue
	}
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

type ApiApiV1TickersGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	cusips *[]string
	tickers *[]string
}

// Array of CUSIPs to match
func (r ApiApiV1TickersGetRequest) Cusips(cusips []string) ApiApiV1TickersGetRequest {
	r.cusips = &cusips
	return r
}

// Array of tickers to match
func (r ApiApiV1TickersGetRequest) Tickers(tickers []string) ApiApiV1TickersGetRequest {
	r.tickers = &tickers
	return r
}

func (r ApiApiV1TickersGetRequest) Execute() ([]TickerInfo, *http.Response, error) {
	return r.ApiService.ApiV1TickersGetExecute(r)
}

/*
ApiV1TickersGet Returns cusip, ticker, and company name for provided matching cusips or tickers parameters

Either `cusips` or `tickers` parameter must be provided.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiV1TickersGetRequest
*/
func (a *DefaultAPIService) ApiV1TickersGet(ctx context.Context) ApiApiV1TickersGetRequest {
	return ApiApiV1TickersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TickerInfo
func (a *DefaultAPIService) ApiV1TickersGetExecute(r ApiApiV1TickersGetRequest) ([]TickerInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TickerInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.ApiV1TickersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/tickers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.cusips != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "cusips", r.cusips, "", "csv")
	}
	if r.tickers != nil {
        parameterAddToHeaderOrQuery(localVarQueryParams, "tickers", r.tickers, "", "csv")
	}
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
