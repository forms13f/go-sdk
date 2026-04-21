# ApiV1Fund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the company. | [optional] 
**CIK** | Pointer to **string** | The CIK of the company. | [optional] 

## Methods

### NewApiV1Fund

`func NewApiV1Fund() *ApiV1Fund`

NewApiV1Fund instantiates a new ApiV1Fund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1FundWithDefaults

`func NewApiV1FundWithDefaults() *ApiV1Fund`

NewApiV1FundWithDefaults instantiates a new ApiV1Fund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiV1Fund) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1Fund) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1Fund) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1Fund) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCIK

`func (o *ApiV1Fund) GetCIK() string`

GetCIK returns the CIK field if non-nil, zero value otherwise.

### GetCIKOk

`func (o *ApiV1Fund) GetCIKOk() (*string, bool)`

GetCIKOk returns a tuple with the CIK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIK

`func (o *ApiV1Fund) SetCIK(v string)`

SetCIK sets CIK field to given value.

### HasCIK

`func (o *ApiV1Fund) HasCIK() bool`

HasCIK returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


