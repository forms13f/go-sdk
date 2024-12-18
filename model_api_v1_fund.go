/*
SEC form 13F API

API for SEC form filings such as 13F.

API version: 1.0.0
Contact: forms13f@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package forms13f

import (
	"encoding/json"
)

// checks if the ApiV1Fund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1Fund{}

// ApiV1Fund struct for ApiV1Fund
type ApiV1Fund struct {
	// The name of the company.
	Name *string `json:"name,omitempty"`
	// The CIK of the company.
	CIK *string `json:"CIK,omitempty"`
}

// NewApiV1Fund instantiates a new ApiV1Fund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1Fund() *ApiV1Fund {
	this := ApiV1Fund{}
	return &this
}

// NewApiV1FundWithDefaults instantiates a new ApiV1Fund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1FundWithDefaults() *ApiV1Fund {
	this := ApiV1Fund{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiV1Fund) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1Fund) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiV1Fund) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiV1Fund) SetName(v string) {
	o.Name = &v
}

// GetCIK returns the CIK field value if set, zero value otherwise.
func (o *ApiV1Fund) GetCIK() string {
	if o == nil || IsNil(o.CIK) {
		var ret string
		return ret
	}
	return *o.CIK
}

// GetCIKOk returns a tuple with the CIK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1Fund) GetCIKOk() (*string, bool) {
	if o == nil || IsNil(o.CIK) {
		return nil, false
	}
	return o.CIK, true
}

// HasCIK returns a boolean if a field has been set.
func (o *ApiV1Fund) HasCIK() bool {
	if o != nil && !IsNil(o.CIK) {
		return true
	}

	return false
}

// SetCIK gets a reference to the given string and assigns it to the CIK field.
func (o *ApiV1Fund) SetCIK(v string) {
	o.CIK = &v
}

func (o ApiV1Fund) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1Fund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CIK) {
		toSerialize["CIK"] = o.CIK
	}
	return toSerialize, nil
}

type NullableApiV1Fund struct {
	value *ApiV1Fund
	isSet bool
}

func (v NullableApiV1Fund) Get() *ApiV1Fund {
	return v.value
}

func (v *NullableApiV1Fund) Set(val *ApiV1Fund) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1Fund) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1Fund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1Fund(val *ApiV1Fund) *NullableApiV1Fund {
	return &NullableApiV1Fund{value: val, isSet: true}
}

func (v NullableApiV1Fund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1Fund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


