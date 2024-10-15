// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Result -url /v5/account/wallet-balance -type GetWalletBalancesRequest -responseDataType .WalletBalancesResponse -rateLimiter 1+45/1s"; DO NOT EDIT.

package bybitapi

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"net/url"
	"reflect"
	"regexp"
)

var GetWalletBalancesRequestLimiter = rate.NewLimiter(45.00000045, 1)

func (g *GetWalletBalancesRequest) AccountType(accountType AccountType) *GetWalletBalancesRequest {
	g.accountType = accountType
	return g
}

func (g *GetWalletBalancesRequest) Coin(coin string) *GetWalletBalancesRequest {
	g.coin = &coin
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetWalletBalancesRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check accountType field -> json key accountType
	accountType := g.accountType

	// TEMPLATE check-valid-values
	switch accountType {
	case "UNIFIED":
		params["accountType"] = accountType

	default:
		return nil, fmt.Errorf("accountType value %v is invalid", accountType)

	}
	// END TEMPLATE check-valid-values

	// assign parameter of accountType
	params["accountType"] = accountType
	// check coin field -> json key coin
	if g.coin != nil {
		coin := *g.coin

		// assign parameter of coin
		params["coin"] = coin
	} else {
	}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetWalletBalancesRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetWalletBalancesRequest) GetParametersQuery() (url.Values, error) {
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
func (g *GetWalletBalancesRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetWalletBalancesRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetWalletBalancesRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetWalletBalancesRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetWalletBalancesRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetWalletBalancesRequest) GetSlugsMap() (map[string]string, error) {
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
func (g *GetWalletBalancesRequest) GetPath() string {
	return "/v5/account/wallet-balance"
}

// Do generates the request object and send the request object to the API endpoint
func (g *GetWalletBalancesRequest) Do(ctx context.Context) (*WalletBalancesResponse, error) {
	if err := GetWalletBalancesRequestLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	// no body params
	var params interface{}
	query, err := g.GetQueryParameters()
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
	var data WalletBalancesResponse
	if err := json.Unmarshal(apiResponse.Result, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
