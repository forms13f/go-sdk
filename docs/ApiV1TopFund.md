# ApiV1TopFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cik** | Pointer to **string** | The Central Index Key (CIK) of the fund. | [optional] 
**Name** | Pointer to **string** | The name of the fund. | [optional] 
**PeriodOfReport** | Pointer to **string** | The period of report date (end of quarter). | [optional] 
**Pnl** | Pointer to **NullableFloat64** | The portfolio PnL percentage for the given quarter. | [optional] 

## Methods

### NewApiV1TopFund

`func NewApiV1TopFund() *ApiV1TopFund`

NewApiV1TopFund instantiates a new ApiV1TopFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1TopFundWithDefaults

`func NewApiV1TopFundWithDefaults() *ApiV1TopFund`

NewApiV1TopFundWithDefaults instantiates a new ApiV1TopFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCik

`func (o *ApiV1TopFund) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *ApiV1TopFund) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *ApiV1TopFund) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *ApiV1TopFund) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetName

`func (o *ApiV1TopFund) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1TopFund) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1TopFund) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiV1TopFund) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPeriodOfReport

`func (o *ApiV1TopFund) GetPeriodOfReport() string`

GetPeriodOfReport returns the PeriodOfReport field if non-nil, zero value otherwise.

### GetPeriodOfReportOk

`func (o *ApiV1TopFund) GetPeriodOfReportOk() (*string, bool)`

GetPeriodOfReportOk returns a tuple with the PeriodOfReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfReport

`func (o *ApiV1TopFund) SetPeriodOfReport(v string)`

SetPeriodOfReport sets PeriodOfReport field to given value.

### HasPeriodOfReport

`func (o *ApiV1TopFund) HasPeriodOfReport() bool`

HasPeriodOfReport returns a boolean if a field has been set.

### GetPnl

`func (o *ApiV1TopFund) GetPnl() float64`

GetPnl returns the Pnl field if non-nil, zero value otherwise.

### GetPnlOk

`func (o *ApiV1TopFund) GetPnlOk() (*float64, bool)`

GetPnlOk returns a tuple with the Pnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnl

`func (o *ApiV1TopFund) SetPnl(v float64)`

SetPnl sets Pnl field to given value.

### HasPnl

`func (o *ApiV1TopFund) HasPnl() bool`

HasPnl returns a boolean if a field has been set.

### SetPnlNil

`func (o *ApiV1TopFund) SetPnlNil(b bool)`

 SetPnlNil sets the value for Pnl to be an explicit nil

### UnsetPnl
`func (o *ApiV1TopFund) UnsetPnl()`

UnsetPnl ensures that no value is present for Pnl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


