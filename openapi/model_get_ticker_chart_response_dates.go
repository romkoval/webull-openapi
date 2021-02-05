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

// GetTickerChartResponseDates struct for GetTickerChartResponseDates
type GetTickerChartResponseDates struct {
	End *string `json:"end,omitempty"`
	Start *string `json:"start,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGetTickerChartResponseDates instantiates a new GetTickerChartResponseDates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTickerChartResponseDates() *GetTickerChartResponseDates {
	this := GetTickerChartResponseDates{}
	return &this
}

// NewGetTickerChartResponseDatesWithDefaults instantiates a new GetTickerChartResponseDates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTickerChartResponseDatesWithDefaults() *GetTickerChartResponseDates {
	this := GetTickerChartResponseDates{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *GetTickerChartResponseDates) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponseDates) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *GetTickerChartResponseDates) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *GetTickerChartResponseDates) SetEnd(v string) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetTickerChartResponseDates) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponseDates) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetTickerChartResponseDates) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *GetTickerChartResponseDates) SetStart(v string) {
	o.Start = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetTickerChartResponseDates) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTickerChartResponseDates) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetTickerChartResponseDates) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetTickerChartResponseDates) SetType(v string) {
	o.Type = &v
}

func (o GetTickerChartResponseDates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetTickerChartResponseDates struct {
	value *GetTickerChartResponseDates
	isSet bool
}

func (v NullableGetTickerChartResponseDates) Get() *GetTickerChartResponseDates {
	return v.value
}

func (v *NullableGetTickerChartResponseDates) Set(val *GetTickerChartResponseDates) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTickerChartResponseDates) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTickerChartResponseDates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTickerChartResponseDates(val *GetTickerChartResponseDates) *NullableGetTickerChartResponseDates {
	return &NullableGetTickerChartResponseDates{value: val, isSet: true}
}

func (v NullableGetTickerChartResponseDates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTickerChartResponseDates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


