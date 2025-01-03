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

// checks if the AiproductdataenhancerGenerateProductDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerGenerateProductDataResponse{}

// AiproductdataenhancerGenerateProductDataResponse struct for AiproductdataenhancerGenerateProductDataResponse
type AiproductdataenhancerGenerateProductDataResponse struct {
	ProductDataGenerated *map[string]string `json:"productDataGenerated,omitempty"`
	CompletionRate       *float32           `json:"completionRate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiproductdataenhancerGenerateProductDataResponse AiproductdataenhancerGenerateProductDataResponse

// NewAiproductdataenhancerGenerateProductDataResponse instantiates a new AiproductdataenhancerGenerateProductDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerGenerateProductDataResponse() *AiproductdataenhancerGenerateProductDataResponse {
	this := AiproductdataenhancerGenerateProductDataResponse{}
	return &this
}

// NewAiproductdataenhancerGenerateProductDataResponseWithDefaults instantiates a new AiproductdataenhancerGenerateProductDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerGenerateProductDataResponseWithDefaults() *AiproductdataenhancerGenerateProductDataResponse {
	this := AiproductdataenhancerGenerateProductDataResponse{}
	return &this
}

// GetProductDataGenerated returns the ProductDataGenerated field value if set, zero value otherwise.
func (o *AiproductdataenhancerGenerateProductDataResponse) GetProductDataGenerated() map[string]string {
	if o == nil || IsNil(o.ProductDataGenerated) {
		var ret map[string]string
		return ret
	}
	return *o.ProductDataGenerated
}

// GetProductDataGeneratedOk returns a tuple with the ProductDataGenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerGenerateProductDataResponse) GetProductDataGeneratedOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ProductDataGenerated) {
		return nil, false
	}
	return o.ProductDataGenerated, true
}

// HasProductDataGenerated returns a boolean if a field has been set.
func (o *AiproductdataenhancerGenerateProductDataResponse) HasProductDataGenerated() bool {
	if o != nil && !IsNil(o.ProductDataGenerated) {
		return true
	}

	return false
}

// SetProductDataGenerated gets a reference to the given map[string]string and assigns it to the ProductDataGenerated field.
func (o *AiproductdataenhancerGenerateProductDataResponse) SetProductDataGenerated(v map[string]string) {
	o.ProductDataGenerated = &v
}

// GetCompletionRate returns the CompletionRate field value if set, zero value otherwise.
func (o *AiproductdataenhancerGenerateProductDataResponse) GetCompletionRate() float32 {
	if o == nil || IsNil(o.CompletionRate) {
		var ret float32
		return ret
	}
	return *o.CompletionRate
}

// GetCompletionRateOk returns a tuple with the CompletionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerGenerateProductDataResponse) GetCompletionRateOk() (*float32, bool) {
	if o == nil || IsNil(o.CompletionRate) {
		return nil, false
	}
	return o.CompletionRate, true
}

// HasCompletionRate returns a boolean if a field has been set.
func (o *AiproductdataenhancerGenerateProductDataResponse) HasCompletionRate() bool {
	if o != nil && !IsNil(o.CompletionRate) {
		return true
	}

	return false
}

// SetCompletionRate gets a reference to the given float32 and assigns it to the CompletionRate field.
func (o *AiproductdataenhancerGenerateProductDataResponse) SetCompletionRate(v float32) {
	o.CompletionRate = &v
}

func (o AiproductdataenhancerGenerateProductDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerGenerateProductDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductDataGenerated) {
		toSerialize["productDataGenerated"] = o.ProductDataGenerated
	}
	if !IsNil(o.CompletionRate) {
		toSerialize["completionRate"] = o.CompletionRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiproductdataenhancerGenerateProductDataResponse) UnmarshalJSON(data []byte) (err error) {
	varAiproductdataenhancerGenerateProductDataResponse := _AiproductdataenhancerGenerateProductDataResponse{}

	err = json.Unmarshal(data, &varAiproductdataenhancerGenerateProductDataResponse)

	if err != nil {
		return err
	}

	*o = AiproductdataenhancerGenerateProductDataResponse(varAiproductdataenhancerGenerateProductDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "productDataGenerated")
		delete(additionalProperties, "completionRate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AiproductdataenhancerGenerateProductDataResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AiproductdataenhancerGenerateProductDataResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAiproductdataenhancerGenerateProductDataResponse struct {
	value *AiproductdataenhancerGenerateProductDataResponse
	isSet bool
}

func (v NullableAiproductdataenhancerGenerateProductDataResponse) Get() *AiproductdataenhancerGenerateProductDataResponse {
	return v.value
}

func (v *NullableAiproductdataenhancerGenerateProductDataResponse) Set(val *AiproductdataenhancerGenerateProductDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerGenerateProductDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerGenerateProductDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerGenerateProductDataResponse(val *AiproductdataenhancerGenerateProductDataResponse) *NullableAiproductdataenhancerGenerateProductDataResponse {
	return &NullableAiproductdataenhancerGenerateProductDataResponse{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerGenerateProductDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerGenerateProductDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
