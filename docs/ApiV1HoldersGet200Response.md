# ApiV1HoldersGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cusip** | Pointer to **string** | The CUSIP of the securities. | [optional] 
**PeriodOfReport** | Pointer to **string** | The period of report date. | [optional] 
**Ciks** | Pointer to **[]string** | An array of CIKs. | [optional] 

## Methods

### NewApiV1HoldersGet200Response

`func NewApiV1HoldersGet200Response() *ApiV1HoldersGet200Response`

NewApiV1HoldersGet200Response instantiates a new ApiV1HoldersGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1HoldersGet200ResponseWithDefaults

`func NewApiV1HoldersGet200ResponseWithDefaults() *ApiV1HoldersGet200Response`

NewApiV1HoldersGet200ResponseWithDefaults instantiates a new ApiV1HoldersGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCusip

`func (o *ApiV1HoldersGet200Response) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ApiV1HoldersGet200Response) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ApiV1HoldersGet200Response) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ApiV1HoldersGet200Response) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetPeriodOfReport

`func (o *ApiV1HoldersGet200Response) GetPeriodOfReport() string`

GetPeriodOfReport returns the PeriodOfReport field if non-nil, zero value otherwise.

### GetPeriodOfReportOk

`func (o *ApiV1HoldersGet200Response) GetPeriodOfReportOk() (*string, bool)`

GetPeriodOfReportOk returns a tuple with the PeriodOfReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfReport

`func (o *ApiV1HoldersGet200Response) SetPeriodOfReport(v string)`

SetPeriodOfReport sets PeriodOfReport field to given value.

### HasPeriodOfReport

`func (o *ApiV1HoldersGet200Response) HasPeriodOfReport() bool`

HasPeriodOfReport returns a boolean if a field has been set.

### GetCiks

`func (o *ApiV1HoldersGet200Response) GetCiks() []string`

GetCiks returns the Ciks field if non-nil, zero value otherwise.

### GetCiksOk

`func (o *ApiV1HoldersGet200Response) GetCiksOk() (*[]string, bool)`

GetCiksOk returns a tuple with the Ciks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiks

`func (o *ApiV1HoldersGet200Response) SetCiks(v []string)`

SetCiks sets Ciks field to given value.

### HasCiks

`func (o *ApiV1HoldersGet200Response) HasCiks() bool`

HasCiks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


