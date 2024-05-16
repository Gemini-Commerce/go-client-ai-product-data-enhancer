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

// checks if the AiproductdataenhancerTranslateDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerTranslateDataResponse{}

// AiproductdataenhancerTranslateDataResponse struct for AiproductdataenhancerTranslateDataResponse
type AiproductdataenhancerTranslateDataResponse struct {
	DataTranslated *map[string]string `json:"dataTranslated,omitempty"`
	ConfidenceRate *float32 `json:"confidenceRate,omitempty"`
	CompletionRate *float32 `json:"completionRate,omitempty"`
}

// NewAiproductdataenhancerTranslateDataResponse instantiates a new AiproductdataenhancerTranslateDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerTranslateDataResponse() *AiproductdataenhancerTranslateDataResponse {
	this := AiproductdataenhancerTranslateDataResponse{}
	return &this
}

// NewAiproductdataenhancerTranslateDataResponseWithDefaults instantiates a new AiproductdataenhancerTranslateDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerTranslateDataResponseWithDefaults() *AiproductdataenhancerTranslateDataResponse {
	this := AiproductdataenhancerTranslateDataResponse{}
	return &this
}

// GetDataTranslated returns the DataTranslated field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataResponse) GetDataTranslated() map[string]string {
	if o == nil || IsNil(o.DataTranslated) {
		var ret map[string]string
		return ret
	}
	return *o.DataTranslated
}

// GetDataTranslatedOk returns a tuple with the DataTranslated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataResponse) GetDataTranslatedOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DataTranslated) {
		return nil, false
	}
	return o.DataTranslated, true
}

// HasDataTranslated returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataResponse) HasDataTranslated() bool {
	if o != nil && !IsNil(o.DataTranslated) {
		return true
	}

	return false
}

// SetDataTranslated gets a reference to the given map[string]string and assigns it to the DataTranslated field.
func (o *AiproductdataenhancerTranslateDataResponse) SetDataTranslated(v map[string]string) {
	o.DataTranslated = &v
}

// GetConfidenceRate returns the ConfidenceRate field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataResponse) GetConfidenceRate() float32 {
	if o == nil || IsNil(o.ConfidenceRate) {
		var ret float32
		return ret
	}
	return *o.ConfidenceRate
}

// GetConfidenceRateOk returns a tuple with the ConfidenceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataResponse) GetConfidenceRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ConfidenceRate) {
		return nil, false
	}
	return o.ConfidenceRate, true
}

// HasConfidenceRate returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataResponse) HasConfidenceRate() bool {
	if o != nil && !IsNil(o.ConfidenceRate) {
		return true
	}

	return false
}

// SetConfidenceRate gets a reference to the given float32 and assigns it to the ConfidenceRate field.
func (o *AiproductdataenhancerTranslateDataResponse) SetConfidenceRate(v float32) {
	o.ConfidenceRate = &v
}

// GetCompletionRate returns the CompletionRate field value if set, zero value otherwise.
func (o *AiproductdataenhancerTranslateDataResponse) GetCompletionRate() float32 {
	if o == nil || IsNil(o.CompletionRate) {
		var ret float32
		return ret
	}
	return *o.CompletionRate
}

// GetCompletionRateOk returns a tuple with the CompletionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTranslateDataResponse) GetCompletionRateOk() (*float32, bool) {
	if o == nil || IsNil(o.CompletionRate) {
		return nil, false
	}
	return o.CompletionRate, true
}

// HasCompletionRate returns a boolean if a field has been set.
func (o *AiproductdataenhancerTranslateDataResponse) HasCompletionRate() bool {
	if o != nil && !IsNil(o.CompletionRate) {
		return true
	}

	return false
}

// SetCompletionRate gets a reference to the given float32 and assigns it to the CompletionRate field.
func (o *AiproductdataenhancerTranslateDataResponse) SetCompletionRate(v float32) {
	o.CompletionRate = &v
}

func (o AiproductdataenhancerTranslateDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerTranslateDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataTranslated) {
		toSerialize["dataTranslated"] = o.DataTranslated
	}
	if !IsNil(o.ConfidenceRate) {
		toSerialize["confidenceRate"] = o.ConfidenceRate
	}
	if !IsNil(o.CompletionRate) {
		toSerialize["completionRate"] = o.CompletionRate
	}
	return toSerialize, nil
}

type NullableAiproductdataenhancerTranslateDataResponse struct {
	value *AiproductdataenhancerTranslateDataResponse
	isSet bool
}

func (v NullableAiproductdataenhancerTranslateDataResponse) Get() *AiproductdataenhancerTranslateDataResponse {
	return v.value
}

func (v *NullableAiproductdataenhancerTranslateDataResponse) Set(val *AiproductdataenhancerTranslateDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerTranslateDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerTranslateDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerTranslateDataResponse(val *AiproductdataenhancerTranslateDataResponse) *NullableAiproductdataenhancerTranslateDataResponse {
	return &NullableAiproductdataenhancerTranslateDataResponse{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerTranslateDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerTranslateDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


