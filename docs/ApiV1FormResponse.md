# ApiV1FormResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPnl** | Pointer to **NullableFloat64** | Estimated total PnL percentage across all holdings in the form since the period of report. | [optional] 
**Entries** | Pointer to [**[]ApiV1FormEntry**](ApiV1FormEntry.md) | The holdings reported in the form. | [optional] 

## Methods

### NewApiV1FormResponse

`func NewApiV1FormResponse() *ApiV1FormResponse`

NewApiV1FormResponse instantiates a new ApiV1FormResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1FormResponseWithDefaults

`func NewApiV1FormResponseWithDefaults() *ApiV1FormResponse`

NewApiV1FormResponseWithDefaults instantiates a new ApiV1FormResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPnl

`func (o *ApiV1FormResponse) GetTotalPnl() float64`

GetTotalPnl returns the TotalPnl field if non-nil, zero value otherwise.

### GetTotalPnlOk

`func (o *ApiV1FormResponse) GetTotalPnlOk() (*float64, bool)`

GetTotalPnlOk returns a tuple with the TotalPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPnl

`func (o *ApiV1FormResponse) SetTotalPnl(v float64)`

SetTotalPnl sets TotalPnl field to given value.

### HasTotalPnl

`func (o *ApiV1FormResponse) HasTotalPnl() bool`

HasTotalPnl returns a boolean if a field has been set.

### SetTotalPnlNil

`func (o *ApiV1FormResponse) SetTotalPnlNil(b bool)`

 SetTotalPnlNil sets the value for TotalPnl to be an explicit nil

### UnsetTotalPnl
`func (o *ApiV1FormResponse) UnsetTotalPnl()`

UnsetTotalPnl ensures that no value is present for TotalPnl, not even an explicit nil
### GetEntries

`func (o *ApiV1FormResponse) GetEntries() []ApiV1FormEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ApiV1FormResponse) GetEntriesOk() (*[]ApiV1FormEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ApiV1FormResponse) SetEntries(v []ApiV1FormEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ApiV1FormResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


