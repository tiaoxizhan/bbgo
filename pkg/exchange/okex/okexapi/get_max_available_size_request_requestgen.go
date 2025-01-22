// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/v5/account/max-avail-size -type GetMaxAvailableSizeRequest -responseDataType []MaxAvailableResponse -rateLimiter 1+20/2s"; DO NOT EDIT.

package okexapi

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"net/url"
	"reflect"
	"regexp"
)

var GetMaxAvailableSizeRequestLimiter = rate.NewLimiter(10, 1)

func (g *GetMaxAvailableSizeRequest) InstrumentID(instrumentID string) *GetMaxAvailableSizeRequest {
	g.instrumentID = instrumentID
	return g
}

func (g *GetMaxAvailableSizeRequest) Currency(currency string) *GetMaxAvailableSizeRequest {
	g.currency = &currency
	return g
}

func (g *GetMaxAvailableSizeRequest) TdMode(tdMode TradeMode) *GetMaxAvailableSizeRequest {
	g.tdMode = tdMode
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetMaxAvailableSizeRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetMaxAvailableSizeRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check instrumentID field -> json key instId
	instrumentID := g.instrumentID

	// assign parameter of instrumentID
	params["instId"] = instrumentID
	// check currency field -> json key ccy
	if g.currency != nil {
		currency := *g.currency

		// assign parameter of currency
		params["ccy"] = currency
	} else {
	}
	// check tdMode field -> json key tdMode
	tdMode := g.tdMode

	// TEMPLATE check-valid-values
	switch tdMode {
	case TradeModeCash, TradeModeIsolated, TradeModeCross:
		params["tdMode"] = tdMode

	default:
		return nil, fmt.Errorf("tdMode value %v is invalid", tdMode)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of tdMode
	params["tdMode"] = tdMode

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetMaxAvailableSizeRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetMaxAvailableSizeRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetMaxAvailableSizeRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetMaxAvailableSizeRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetMaxAvailableSizeRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetMaxAvailableSizeRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetMaxAvailableSizeRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

// GetPath returns the request path of the API
func (g *GetMaxAvailableSizeRequest) GetPath() string {
	return "/api/v5/account/max-avail-size"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetMaxAvailableSizeRequest) Do(ctx context.Context) ([]MaxAvailableResponse, error) {
	if err := GetMaxAvailableSizeRequestLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	var apiURL string

	apiURL = g.GetPath()

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse

	type responseUnmarshaler interface {
		Unmarshal(data []byte) error
	}

	if unmarshaler, ok := interface{}(&apiResponse).(responseUnmarshaler); ok {
		if err := unmarshaler.Unmarshal(response.Body); err != nil {
			return nil, err
		}
	} else {
		// The line below checks the content type, however, some API server might not send the correct content type header,
		// Hence, this is commented for backward compatibility
		// response.IsJSON()
		if err := response.DecodeJSON(&apiResponse); err != nil {
			return nil, err
		}
	}

	type responseValidator interface {
		Validate() error
	}

	if validator, ok := interface{}(&apiResponse).(responseValidator); ok {
		if err := validator.Validate(); err != nil {
			return nil, err
		}
	}
	var data []MaxAvailableResponse
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
