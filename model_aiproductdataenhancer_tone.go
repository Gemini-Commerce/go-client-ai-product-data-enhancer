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

// checks if the AiproductdataenhancerTone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiproductdataenhancerTone{}

// AiproductdataenhancerTone struct for AiproductdataenhancerTone
type AiproductdataenhancerTone struct {
	ToneType             *AiproductdataenhancerToneType `json:"toneType,omitempty"`
	CustomTone           *string                        `json:"customTone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AiproductdataenhancerTone AiproductdataenhancerTone

// NewAiproductdataenhancerTone instantiates a new AiproductdataenhancerTone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiproductdataenhancerTone() *AiproductdataenhancerTone {
	this := AiproductdataenhancerTone{}
	var toneType AiproductdataenhancerToneType = AIPRODUCTDATAENHANCERTONETYPE_UNKNOWN
	this.ToneType = &toneType
	return &this
}

// NewAiproductdataenhancerToneWithDefaults instantiates a new AiproductdataenhancerTone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiproductdataenhancerToneWithDefaults() *AiproductdataenhancerTone {
	this := AiproductdataenhancerTone{}
	var toneType AiproductdataenhancerToneType = AIPRODUCTDATAENHANCERTONETYPE_UNKNOWN
	this.ToneType = &toneType
	return &this
}

// GetToneType returns the ToneType field value if set, zero value otherwise.
func (o *AiproductdataenhancerTone) GetToneType() AiproductdataenhancerToneType {
	if o == nil || IsNil(o.ToneType) {
		var ret AiproductdataenhancerToneType
		return ret
	}
	return *o.ToneType
}

// GetToneTypeOk returns a tuple with the ToneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTone) GetToneTypeOk() (*AiproductdataenhancerToneType, bool) {
	if o == nil || IsNil(o.ToneType) {
		return nil, false
	}
	return o.ToneType, true
}

// HasToneType returns a boolean if a field has been set.
func (o *AiproductdataenhancerTone) HasToneType() bool {
	if o != nil && !IsNil(o.ToneType) {
		return true
	}

	return false
}

// SetToneType gets a reference to the given AiproductdataenhancerToneType and assigns it to the ToneType field.
func (o *AiproductdataenhancerTone) SetToneType(v AiproductdataenhancerToneType) {
	o.ToneType = &v
}

// GetCustomTone returns the CustomTone field value if set, zero value otherwise.
func (o *AiproductdataenhancerTone) GetCustomTone() string {
	if o == nil || IsNil(o.CustomTone) {
		var ret string
		return ret
	}
	return *o.CustomTone
}

// GetCustomToneOk returns a tuple with the CustomTone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiproductdataenhancerTone) GetCustomToneOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTone) {
		return nil, false
	}
	return o.CustomTone, true
}

// HasCustomTone returns a boolean if a field has been set.
func (o *AiproductdataenhancerTone) HasCustomTone() bool {
	if o != nil && !IsNil(o.CustomTone) {
		return true
	}

	return false
}

// SetCustomTone gets a reference to the given string and assigns it to the CustomTone field.
func (o *AiproductdataenhancerTone) SetCustomTone(v string) {
	o.CustomTone = &v
}

func (o AiproductdataenhancerTone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiproductdataenhancerTone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToneType) {
		toSerialize["toneType"] = o.ToneType
	}
	if !IsNil(o.CustomTone) {
		toSerialize["customTone"] = o.CustomTone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiproductdataenhancerTone) UnmarshalJSON(data []byte) (err error) {
	varAiproductdataenhancerTone := _AiproductdataenhancerTone{}

	err = json.Unmarshal(data, &varAiproductdataenhancerTone)

	if err != nil {
		return err
	}

	*o = AiproductdataenhancerTone(varAiproductdataenhancerTone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "toneType")
		delete(additionalProperties, "customTone")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *AiproductdataenhancerTone) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *AiproductdataenhancerTone) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableAiproductdataenhancerTone struct {
	value *AiproductdataenhancerTone
	isSet bool
}

func (v NullableAiproductdataenhancerTone) Get() *AiproductdataenhancerTone {
	return v.value
}

func (v *NullableAiproductdataenhancerTone) Set(val *AiproductdataenhancerTone) {
	v.value = val
	v.isSet = true
}

func (v NullableAiproductdataenhancerTone) IsSet() bool {
	return v.isSet
}

func (v *NullableAiproductdataenhancerTone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiproductdataenhancerTone(val *AiproductdataenhancerTone) *NullableAiproductdataenhancerTone {
	return &NullableAiproductdataenhancerTone{value: val, isSet: true}
}

func (v NullableAiproductdataenhancerTone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiproductdataenhancerTone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
