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

// checks if the ApiV1Filer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1Filer{}

// ApiV1Filer struct for ApiV1Filer
type ApiV1Filer struct {
	// The Central Index Key (CIK) of the filer.
	Cik *string `json:"cik,omitempty"`
	// An array of company names associated with the CIK.
	CompanyNames []string `json:"company_names,omitempty"`
}

// NewApiV1Filer instantiates a new ApiV1Filer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1Filer() *ApiV1Filer {
	this := ApiV1Filer{}
	return &this
}

// NewApiV1FilerWithDefaults instantiates a new ApiV1Filer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1FilerWithDefaults() *ApiV1Filer {
	this := ApiV1Filer{}
	return &this
}

// GetCik returns the Cik field value if set, zero value otherwise.
func (o *ApiV1Filer) GetCik() string {
	if o == nil || IsNil(o.Cik) {
		var ret string
		return ret
	}
	return *o.Cik
}

// GetCikOk returns a tuple with the Cik field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1Filer) GetCikOk() (*string, bool) {
	if o == nil || IsNil(o.Cik) {
		return nil, false
	}
	return o.Cik, true
}

// HasCik returns a boolean if a field has been set.
func (o *ApiV1Filer) HasCik() bool {
	if o != nil && !IsNil(o.Cik) {
		return true
	}

	return false
}

// SetCik gets a reference to the given string and assigns it to the Cik field.
func (o *ApiV1Filer) SetCik(v string) {
	o.Cik = &v
}

// GetCompanyNames returns the CompanyNames field value if set, zero value otherwise.
func (o *ApiV1Filer) GetCompanyNames() []string {
	if o == nil || IsNil(o.CompanyNames) {
		var ret []string
		return ret
	}
	return o.CompanyNames
}

// GetCompanyNamesOk returns a tuple with the CompanyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1Filer) GetCompanyNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CompanyNames) {
		return nil, false
	}
	return o.CompanyNames, true
}

// HasCompanyNames returns a boolean if a field has been set.
func (o *ApiV1Filer) HasCompanyNames() bool {
	if o != nil && !IsNil(o.CompanyNames) {
		return true
	}

	return false
}

// SetCompanyNames gets a reference to the given []string and assigns it to the CompanyNames field.
func (o *ApiV1Filer) SetCompanyNames(v []string) {
	o.CompanyNames = v
}

func (o ApiV1Filer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1Filer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cik) {
		toSerialize["cik"] = o.Cik
	}
	if !IsNil(o.CompanyNames) {
		toSerialize["company_names"] = o.CompanyNames
	}
	return toSerialize, nil
}

type NullableApiV1Filer struct {
	value *ApiV1Filer
	isSet bool
}

func (v NullableApiV1Filer) Get() *ApiV1Filer {
	return v.value
}

func (v *NullableApiV1Filer) Set(val *ApiV1Filer) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1Filer) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1Filer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1Filer(val *ApiV1Filer) *NullableApiV1Filer {
	return &NullableApiV1Filer{value: val, isSet: true}
}

func (v NullableApiV1Filer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1Filer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


