/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetFundamentalsResponseData struct for GetFundamentalsResponseData
type GetFundamentalsResponseData struct {
	CurrencyName *string `json:"currencyName,omitempty"`
	Id *string `json:"id,omitempty"`
	Points *[]GetFundamentalsResponseDataPoints `json:"points,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewGetFundamentalsResponseData instantiates a new GetFundamentalsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFundamentalsResponseData() *GetFundamentalsResponseData {
	this := GetFundamentalsResponseData{}
	return &this
}

// NewGetFundamentalsResponseDataWithDefaults instantiates a new GetFundamentalsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFundamentalsResponseDataWithDefaults() *GetFundamentalsResponseData {
	this := GetFundamentalsResponseData{}
	return &this
}

// GetCurrencyName returns the CurrencyName field value if set, zero value otherwise.
func (o *GetFundamentalsResponseData) GetCurrencyName() string {
	if o == nil || o.CurrencyName == nil {
		var ret string
		return ret
	}
	return *o.CurrencyName
}

// GetCurrencyNameOk returns a tuple with the CurrencyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseData) GetCurrencyNameOk() (*string, bool) {
	if o == nil || o.CurrencyName == nil {
		return nil, false
	}
	return o.CurrencyName, true
}

// HasCurrencyName returns a boolean if a field has been set.
func (o *GetFundamentalsResponseData) HasCurrencyName() bool {
	if o != nil && o.CurrencyName != nil {
		return true
	}

	return false
}

// SetCurrencyName gets a reference to the given string and assigns it to the CurrencyName field.
func (o *GetFundamentalsResponseData) SetCurrencyName(v string) {
	o.CurrencyName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetFundamentalsResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetFundamentalsResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetFundamentalsResponseData) SetId(v string) {
	o.Id = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *GetFundamentalsResponseData) GetPoints() []GetFundamentalsResponseDataPoints {
	if o == nil || o.Points == nil {
		var ret []GetFundamentalsResponseDataPoints
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseData) GetPointsOk() (*[]GetFundamentalsResponseDataPoints, bool) {
	if o == nil || o.Points == nil {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *GetFundamentalsResponseData) HasPoints() bool {
	if o != nil && o.Points != nil {
		return true
	}

	return false
}

// SetPoints gets a reference to the given []GetFundamentalsResponseDataPoints and assigns it to the Points field.
func (o *GetFundamentalsResponseData) SetPoints(v []GetFundamentalsResponseDataPoints) {
	o.Points = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetFundamentalsResponseData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFundamentalsResponseData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetFundamentalsResponseData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetFundamentalsResponseData) SetTitle(v string) {
	o.Title = &v
}

func (o GetFundamentalsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyName != nil {
		toSerialize["currencyName"] = o.CurrencyName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Points != nil {
		toSerialize["points"] = o.Points
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableGetFundamentalsResponseData struct {
	value *GetFundamentalsResponseData
	isSet bool
}

func (v NullableGetFundamentalsResponseData) Get() *GetFundamentalsResponseData {
	return v.value
}

func (v *NullableGetFundamentalsResponseData) Set(val *GetFundamentalsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFundamentalsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFundamentalsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFundamentalsResponseData(val *GetFundamentalsResponseData) *NullableGetFundamentalsResponseData {
	return &NullableGetFundamentalsResponseData{value: val, isSet: true}
}

func (v NullableGetFundamentalsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFundamentalsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


