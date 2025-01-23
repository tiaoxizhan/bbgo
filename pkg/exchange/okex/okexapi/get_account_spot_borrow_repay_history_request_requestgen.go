// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -url /api/v5/account/spot-borrow-repay-history -type GetAccountSpotBorrowRepayHistoryRequest -responseDataType []MarginHistoryEntry"; DO NOT EDIT.

package okexapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetAccountSpotBorrowRepayHistoryRequest) EventType(eventType MarginEventType) *GetAccountSpotBorrowRepayHistoryRequest {
	g.eventType = &eventType
	return g
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) Currency(currency string) *GetAccountSpotBorrowRepayHistoryRequest {
	g.currency = &currency
	return g
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) After(after time.Time) *GetAccountSpotBorrowRepayHistoryRequest {
	g.after = &after
	return g
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) Before(before time.Time) *GetAccountSpotBorrowRepayHistoryRequest {
	g.before = &before
	return g
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) Limit(limit uint64) *GetAccountSpotBorrowRepayHistoryRequest {
	g.limit = &limit
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check eventType field -> json key type
	if g.eventType != nil {
		eventType := *g.eventType

		// TEMPLATE check-valid-values
		switch eventType {
		case MarginEventTypeAutoBorrow, MarginEventTypeAutoRepay, MarginEventTypeManualBorrow, MarginEventTypeManualRepay:
			params["type"] = eventType

		default:
			return nil, fmt.Errorf("type value %v is invalid", eventType)

		}
		// END TEMPLATE check-valid-values

		// assign parameter of eventType
		params["type"] = eventType
	} else {
	}
	// check currency field -> json key ccy
	if g.currency != nil {
		currency := *g.currency

		// assign parameter of currency
		params["ccy"] = currency
	} else {
	}
	// check after field -> json key after
	if g.after != nil {
		after := *g.after

		// assign parameter of after
		// convert time.Time to milliseconds time stamp
		params["after"] = strconv.FormatInt(after.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check before field -> json key before
	if g.before != nil {
		before := *g.before

		// assign parameter of before
		// convert time.Time to milliseconds time stamp
		params["before"] = strconv.FormatInt(before.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetParametersQuery() (url.Values, error) {
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
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetAccountSpotBorrowRepayHistoryRequest) GetSlugsMap() (map[string]string, error) {
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
func (g *GetAccountSpotBorrowRepayHistoryRequest) GetPath() string {
	return "/api/v5/account/spot-borrow-repay-history"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetAccountSpotBorrowRepayHistoryRequest) Do(ctx context.Context) ([]MarginHistoryEntry, error) {

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
	var data []MarginHistoryEntry
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return data, nil
}
