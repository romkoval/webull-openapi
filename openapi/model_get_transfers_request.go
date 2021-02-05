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

// GetTransfersRequest struct for GetTransfersRequest
type GetTransfersRequest struct {
	PageSize *float32 `json:"pageSize,omitempty"`
	LastRecordId *string `json:"lastRecordId,omitempty"`
}

// NewGetTransfersRequest instantiates a new GetTransfersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransfersRequest() *GetTransfersRequest {
	this := GetTransfersRequest{}
	var pageSize float32 = 256
	this.PageSize = &pageSize
	var lastRecordId string = "0"
	this.LastRecordId = &lastRecordId
	return &this
}

// NewGetTransfersRequestWithDefaults instantiates a new GetTransfersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransfersRequestWithDefaults() *GetTransfersRequest {
	this := GetTransfersRequest{}
	var pageSize float32 = 256
	this.PageSize = &pageSize
	var lastRecordId string = "0"
	this.LastRecordId = &lastRecordId
	return &this
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *GetTransfersRequest) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfersRequest) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *GetTransfersRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *GetTransfersRequest) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetLastRecordId returns the LastRecordId field value if set, zero value otherwise.
func (o *GetTransfersRequest) GetLastRecordId() string {
	if o == nil || o.LastRecordId == nil {
		var ret string
		return ret
	}
	return *o.LastRecordId
}

// GetLastRecordIdOk returns a tuple with the LastRecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransfersRequest) GetLastRecordIdOk() (*string, bool) {
	if o == nil || o.LastRecordId == nil {
		return nil, false
	}
	return o.LastRecordId, true
}

// HasLastRecordId returns a boolean if a field has been set.
func (o *GetTransfersRequest) HasLastRecordId() bool {
	if o != nil && o.LastRecordId != nil {
		return true
	}

	return false
}

// SetLastRecordId gets a reference to the given string and assigns it to the LastRecordId field.
func (o *GetTransfersRequest) SetLastRecordId(v string) {
	o.LastRecordId = &v
}

func (o GetTransfersRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	if o.LastRecordId != nil {
		toSerialize["lastRecordId"] = o.LastRecordId
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransfersRequest struct {
	value *GetTransfersRequest
	isSet bool
}

func (v NullableGetTransfersRequest) Get() *GetTransfersRequest {
	return v.value
}

func (v *NullableGetTransfersRequest) Set(val *GetTransfersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransfersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransfersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransfersRequest(val *GetTransfersRequest) *NullableGetTransfersRequest {
	return &NullableGetTransfersRequest{value: val, isSet: true}
}

func (v NullableGetTransfersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransfersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


