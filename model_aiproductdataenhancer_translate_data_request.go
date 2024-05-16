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

// checks if the AiproductdataenhancerTranslateDataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerTranslateDataRequest{}

// AiproductdataenhancerTranslateDataRequest struct for AiproductdataenhancerTranslateDataRequest
type AiproductdataenhancerTranslateDataRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	LanguageCode *AiproductdataenhancerLanguageCode `json:"languageCode,omitempty"`
	DataToTranslate []AiproductdataenhancerDataToTranslate `json:"dataToTranslate,omitempty"`
}

// NewAiproductdataenhancerTranslateDataRequest instantiates a new AiproductdataenhancerTranslateDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerTranslateDataRequest() *AiproductdataenhancerTranslateDataRequest {
	this := AiproductdataenhancerTranslateDataRequest{}
	var languageCode AiproductdataenhancerLanguageCode = AIPRODUCTDATAENHANCERLANGUAGECODE_UNKNOWN
	this.LanguageCode = &languageCode
	return &this
}

// NewAiproductdataenhancerTranslateDataRequestWithDefaults instantiates a new AiproductdataenhancerTranslateDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerTranslateDataRequestWithDefaults() *AiproductdataenhancerTranslateDataRequest {
	this := AiproductdataenhancerTranslateDataRequest{}
	var languageCode AiproductdataenhancerLanguageCode = AIPRODUCTDATAENHANCERLANGUAGECODE_UNKNOWN
	this.LanguageCode = &languageCode
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AiproductdataenhancerTranslateDataRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataRequest) GetLanguageCode() AiproductdataenhancerLanguageCode {
	if o == nil || IsNil(o.LanguageCode) {
		var ret AiproductdataenhancerLanguageCode
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataRequest) GetLanguageCodeOk() (*AiproductdataenhancerLanguageCode, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataRequest) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given AiproductdataenhancerLanguageCode and assigns it to the LanguageCode field.
func (o *AiproductdataenhancerTranslateDataRequest) SetLanguageCode(v AiproductdataenhancerLanguageCode) {
	o.LanguageCode = &v
}

// GetDataToTranslate returns the DataToTranslate field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataRequest) GetDataToTranslate() []AiproductdataenhancerDataToTranslate {
	if o == nil || IsNil(o.DataToTranslate) {
		var ret []AiproductdataenhancerDataToTranslate
		return ret
	}
	return o.DataToTranslate
}

// GetDataToTranslateOk returns a tuple with the DataToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataRequest) GetDataToTranslateOk() ([]AiproductdataenhancerDataToTranslate, bool) {
	if o == nil || IsNil(o.DataToTranslate) {
		return nil, false
	}
	return o.DataToTranslate, true
}

// HasDataToTranslate returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataRequest) HasDataToTranslate() bool {
	if o != nil && !IsNil(o.DataToTranslate) {
		return true
	}

	return false
}

// SetDataToTranslate gets a reference to the given []AiproductdataenhancerDataToTranslate and assigns it to the DataToTranslate field.
func (o *AiproductdataenhancerTranslateDataRequest) SetDataToTranslate(v []AiproductdataenhancerDataToTranslate) {
	o.DataToTranslate = v
}

func (o AiproductdataenhancerTranslateDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerTranslateDataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if !IsNil(o.DataToTranslate) {
		toSerialize["dataToTranslate"] = o.DataToTranslate
	}
	return toSerialize, nil
}

type NullableAiproductdataenhancerTranslateDataRequest struct {
	value *AiproductdataenhancerTranslateDataRequest
	isSet bool
}

func (v NullableAiproductdataenhancerTranslateDataRequest) Get() *AiproductdataenhancerTranslateDataRequest {
	return v.value
}

func (v *NullableAiproductdataenhancerTranslateDataRequest) Set(val *AiproductdataenhancerTranslateDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerTranslateDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerTranslateDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerTranslateDataRequest(val *AiproductdataenhancerTranslateDataRequest) *NullableAiproductdataenhancerTranslateDataRequest {
	return &NullableAiproductdataenhancerTranslateDataRequest{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerTranslateDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerTranslateDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

