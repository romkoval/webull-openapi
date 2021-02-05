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

// GetAcccountRequestAccountMembers struct for GetAcccountRequestAccountMembers
type GetAcccountRequestAccountMembers struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewGetAcccountRequestAccountMembers instantiates a new GetAcccountRequestAccountMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAcccountRequestAccountMembers() *GetAcccountRequestAccountMembers {
	this := GetAcccountRequestAccountMembers{}
	return &this
}

// NewGetAcccountRequestAccountMembersWithDefaults instantiates a new GetAcccountRequestAccountMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAcccountRequestAccountMembersWithDefaults() *GetAcccountRequestAccountMembers {
	this := GetAcccountRequestAccountMembers{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *GetAcccountRequestAccountMembers) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestAccountMembers) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *GetAcccountRequestAccountMembers) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *GetAcccountRequestAccountMembers) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetAcccountRequestAccountMembers) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestAccountMembers) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetAcccountRequestAccountMembers) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetAcccountRequestAccountMembers) SetValue(v string) {
	o.Value = &v
}

func (o GetAcccountRequestAccountMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableGetAcccountRequestAccountMembers struct {
	value *GetAcccountRequestAccountMembers
	isSet bool
}

func (v NullableGetAcccountRequestAccountMembers) Get() *GetAcccountRequestAccountMembers {
	return v.value
}

func (v *NullableGetAcccountRequestAccountMembers) Set(val *GetAcccountRequestAccountMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAcccountRequestAccountMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAcccountRequestAccountMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAcccountRequestAccountMembers(val *GetAcccountRequestAccountMembers) *NullableGetAcccountRequestAccountMembers {
	return &NullableGetAcccountRequestAccountMembers{value: val, isSet: true}
}

func (v NullableGetAcccountRequestAccountMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAcccountRequestAccountMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


