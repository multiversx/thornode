/*
Thornode API

Thornode REST API.

Contact: devs@thorchain.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// TradeUnitApiService TradeUnitApi service
type TradeUnitApiService service

type ApiTradeUnitRequest struct {
	ctx context.Context
	ApiService *TradeUnitApiService
	asset string
	height *int64
}

// optional block height, defaults to current tip
func (r ApiTradeUnitRequest) Height(height int64) ApiTradeUnitRequest {
	r.height = &height
	return r
}

func (r ApiTradeUnitRequest) Execute() (*TradeUnitResponse, *http.Response, error) {
	return r.ApiService.TradeUnitExecute(r)
}

/*
TradeUnit Method for TradeUnit

Returns the total units and depth of a trade asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param asset
 @return ApiTradeUnitRequest
*/
func (a *TradeUnitApiService) TradeUnit(ctx context.Context, asset string) ApiTradeUnitRequest {
	return ApiTradeUnitRequest{
		ApiService: a,
		ctx: ctx,
		asset: asset,
	}
}

// Execute executes the request
//  @return TradeUnitResponse
func (a *TradeUnitApiService) TradeUnitExecute(r ApiTradeUnitRequest) (*TradeUnitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TradeUnitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TradeUnitApiService.TradeUnit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/thorchain/trade/unit/{asset}"
	localVarPath = strings.Replace(localVarPath, "{"+"asset"+"}", url.PathEscape(parameterToString(r.asset, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.height != nil {
		localVarQueryParams.Add("height", parameterToString(*r.height, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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