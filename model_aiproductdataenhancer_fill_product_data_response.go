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

// checks if the AiproductdataenhancerFillProductDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerFillProductDataResponse{}

// AiproductdataenhancerFillProductDataResponse struct for AiproductdataenhancerFillProductDataResponse
type AiproductdataenhancerFillProductDataResponse struct {
	JobId                *string `json:"jobId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiproductdataenhancerFillProductDataResponse AiproductdataenhancerFillProductDataResponse

// NewAiproductdataenhancerFillProductDataResponse instantiates a new AiproductdataenhancerFillProductDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerFillProductDataResponse() *AiproductdataenhancerFillProductDataResponse {
	this := AiproductdataenhancerFillProductDataResponse{}
	return &this
}

// NewAiproductdataenhancerFillProductDataResponseWithDefaults instantiates a new AiproductdataenhancerFillProductDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerFillProductDataResponseWithDefaults() *AiproductdataenhancerFillProductDataResponse {
	this := AiproductdataenhancerFillProductDataResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *AiproductdataenhancerFillProductDataResponse) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerFillProductDataResponse) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *AiproductdataenhancerFillProductDataResponse) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *AiproductdataenhancerFillProductDataResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o AiproductdataenhancerFillProductDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerFillProductDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiproductdataenhancerFillProductDataResponse) UnmarshalJSON(data []byte) (err error) {
	varAiproductdataenhancerFillProductDataResponse := _AiproductdataenhancerFillProductDataResponse{}

	err = json.Unmarshal(data, &varAiproductdataenhancerFillProductDataResponse)

	if err != nil {
		return err
	}

	*o = AiproductdataenhancerFillProductDataResponse(varAiproductdataenhancerFillProductDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "jobId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AiproductdataenhancerFillProductDataResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AiproductdataenhancerFillProductDataResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAiproductdataenhancerFillProductDataResponse struct {
	value *AiproductdataenhancerFillProductDataResponse
	isSet bool
}

func (v NullableAiproductdataenhancerFillProductDataResponse) Get() *AiproductdataenhancerFillProductDataResponse {
	return v.value
}

func (v *NullableAiproductdataenhancerFillProductDataResponse) Set(val *AiproductdataenhancerFillProductDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerFillProductDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerFillProductDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerFillProductDataResponse(val *AiproductdataenhancerFillProductDataResponse) *NullableAiproductdataenhancerFillProductDataResponse {
	return &NullableAiproductdataenhancerFillProductDataResponse{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerFillProductDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerFillProductDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
