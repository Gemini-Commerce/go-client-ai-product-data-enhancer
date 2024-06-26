/*
aiproductdataenhancer/service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aiproductdataenhancer

import (
	"encoding/json"
)

// checks if the AiproductdataenhancerProductInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerProductInformation{}

// AiproductdataenhancerProductInformation struct for AiproductdataenhancerProductInformation
type AiproductdataenhancerProductInformation struct {
	Brand *string `json:"brand,omitempty"`
	Code *string `json:"code,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewAiproductdataenhancerProductInformation instantiates a new AiproductdataenhancerProductInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerProductInformation() *AiproductdataenhancerProductInformation {
	this := AiproductdataenhancerProductInformation{}
	return &this
}

// NewAiproductdataenhancerProductInformationWithDefaults instantiates a new AiproductdataenhancerProductInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerProductInformationWithDefaults() *AiproductdataenhancerProductInformation {
	this := AiproductdataenhancerProductInformation{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *AiproductdataenhancerProductInformation) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerProductInformation) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *AiproductdataenhancerProductInformation) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *AiproductdataenhancerProductInformation) SetBrand(v string) {
	o.Brand = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AiproductdataenhancerProductInformation) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerProductInformation) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AiproductdataenhancerProductInformation) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AiproductdataenhancerProductInformation) SetCode(v string) {
	o.Code = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiproductdataenhancerProductInformation) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerProductInformation) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiproductdataenhancerProductInformation) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiproductdataenhancerProductInformation) SetTitle(v string) {
	o.Title = &v
}

func (o AiproductdataenhancerProductInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerProductInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAiproductdataenhancerProductInformation struct {
	value *AiproductdataenhancerProductInformation
	isSet bool
}

func (v NullableAiproductdataenhancerProductInformation) Get() *AiproductdataenhancerProductInformation {
	return v.value
}

func (v *NullableAiproductdataenhancerProductInformation) Set(val *AiproductdataenhancerProductInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerProductInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerProductInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerProductInformation(val *AiproductdataenhancerProductInformation) *NullableAiproductdataenhancerProductInformation {
	return &NullableAiproductdataenhancerProductInformation{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerProductInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerProductInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


