# ApiV1Filer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | Pointer to **string** | The Central Index Key (CIK) of the filer. | [optional] 
**CompanyNames** | Pointer to **[]string** | An array of company names associated with the CIK. | [optional] 

## Methods

### NewApiV1Filer

`func NewApiV1Filer() *ApiV1Filer`

NewApiV1Filer instantiates a new ApiV1Filer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1FilerWithDefaults

`func NewApiV1FilerWithDefaults() *ApiV1Filer`

NewApiV1FilerWithDefaults instantiates a new ApiV1Filer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *ApiV1Filer) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *ApiV1Filer) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *ApiV1Filer) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *ApiV1Filer) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetCompanyNames

`func (o *ApiV1Filer) GetCompanyNames() []string`

GetCompanyNames returns the CompanyNames field if non-nil, zero value otherwise.

### GetCompanyNamesOk

`func (o *ApiV1Filer) GetCompanyNamesOk() (*[]string, bool)`

GetCompanyNamesOk returns a tuple with the CompanyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNames

`func (o *ApiV1Filer) SetCompanyNames(v []string)`

SetCompanyNames sets CompanyNames field to given value.

### HasCompanyNames

`func (o *ApiV1Filer) HasCompanyNames() bool`

HasCompanyNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


