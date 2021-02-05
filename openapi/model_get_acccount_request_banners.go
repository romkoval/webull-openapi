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

// GetAcccountRequestBanners struct for GetAcccountRequestBanners
type GetAcccountRequestBanners struct {
	ImgUrl *string `json:"imgUrl,omitempty"`
	Link *string `json:"link,omitempty"`
	SubTitle *string `json:"subTitle,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGetAcccountRequestBanners instantiates a new GetAcccountRequestBanners object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAcccountRequestBanners() *GetAcccountRequestBanners {
	this := GetAcccountRequestBanners{}
	return &this
}

// NewGetAcccountRequestBannersWithDefaults instantiates a new GetAcccountRequestBanners object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAcccountRequestBannersWithDefaults() *GetAcccountRequestBanners {
	this := GetAcccountRequestBanners{}
	return &this
}

// GetImgUrl returns the ImgUrl field value if set, zero value otherwise.
func (o *GetAcccountRequestBanners) GetImgUrl() string {
	if o == nil || o.ImgUrl == nil {
		var ret string
		return ret
	}
	return *o.ImgUrl
}

// GetImgUrlOk returns a tuple with the ImgUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestBanners) GetImgUrlOk() (*string, bool) {
	if o == nil || o.ImgUrl == nil {
		return nil, false
	}
	return o.ImgUrl, true
}

// HasImgUrl returns a boolean if a field has been set.
func (o *GetAcccountRequestBanners) HasImgUrl() bool {
	if o != nil && o.ImgUrl != nil {
		return true
	}

	return false
}

// SetImgUrl gets a reference to the given string and assigns it to the ImgUrl field.
func (o *GetAcccountRequestBanners) SetImgUrl(v string) {
	o.ImgUrl = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *GetAcccountRequestBanners) GetLink() string {
	if o == nil || o.Link == nil {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestBanners) GetLinkOk() (*string, bool) {
	if o == nil || o.Link == nil {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *GetAcccountRequestBanners) HasLink() bool {
	if o != nil && o.Link != nil {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *GetAcccountRequestBanners) SetLink(v string) {
	o.Link = &v
}

// GetSubTitle returns the SubTitle field value if set, zero value otherwise.
func (o *GetAcccountRequestBanners) GetSubTitle() string {
	if o == nil || o.SubTitle == nil {
		var ret string
		return ret
	}
	return *o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestBanners) GetSubTitleOk() (*string, bool) {
	if o == nil || o.SubTitle == nil {
		return nil, false
	}
	return o.SubTitle, true
}

// HasSubTitle returns a boolean if a field has been set.
func (o *GetAcccountRequestBanners) HasSubTitle() bool {
	if o != nil && o.SubTitle != nil {
		return true
	}

	return false
}

// SetSubTitle gets a reference to the given string and assigns it to the SubTitle field.
func (o *GetAcccountRequestBanners) SetSubTitle(v string) {
	o.SubTitle = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetAcccountRequestBanners) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestBanners) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetAcccountRequestBanners) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetAcccountRequestBanners) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAcccountRequestBanners) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAcccountRequestBanners) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAcccountRequestBanners) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAcccountRequestBanners) SetType(v string) {
	o.Type = &v
}

func (o GetAcccountRequestBanners) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImgUrl != nil {
		toSerialize["imgUrl"] = o.ImgUrl
	}
	if o.Link != nil {
		toSerialize["link"] = o.Link
	}
	if o.SubTitle != nil {
		toSerialize["subTitle"] = o.SubTitle
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetAcccountRequestBanners struct {
	value *GetAcccountRequestBanners
	isSet bool
}

func (v NullableGetAcccountRequestBanners) Get() *GetAcccountRequestBanners {
	return v.value
}

func (v *NullableGetAcccountRequestBanners) Set(val *GetAcccountRequestBanners) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAcccountRequestBanners) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAcccountRequestBanners) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAcccountRequestBanners(val *GetAcccountRequestBanners) *NullableGetAcccountRequestBanners {
	return &NullableGetAcccountRequestBanners{value: val, isSet: true}
}

func (v NullableGetAcccountRequestBanners) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAcccountRequestBanners) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


