# ApiV1FormEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionNumber** | Pointer to **string** | The accession number of the form entry. | [optional] 
**Cik** | Pointer to **string** | The Central Index Key (CIK) of the form entry. | [optional] 
**NameOfIssuer** | Pointer to **string** | The name of the issuer. | [optional] 
**TitleOfClass** | Pointer to **string** | The title of the class of securities. | [optional] 
**Cusip** | Pointer to **string** | The CUSIP number of the securities. | [optional] 
**Ticker** | Pointer to **string** | The ticker of the securities. | [optional] 
**Value** | Pointer to **NullableInt64** | The value of the securities. | [optional] 
**SshPrnamt** | Pointer to **NullableInt64** | The number of shares or principal amount. | [optional] 
**SshPrnamtType** | Pointer to **NullableString** | The type of shares or principal amount. | [optional] 
**InvestmentDiscretion** | Pointer to **NullableString** | The investment discretion. | [optional] 
**VotingAuthoritySole** | Pointer to **NullableInt64** | The sole voting authority. | [optional] 
**VotingAuthorityShared** | Pointer to **NullableInt64** | The shared voting authority. | [optional] 
**VotingAuthorityNone** | Pointer to **NullableInt64** | The no voting authority. | [optional] 

## Methods

### NewApiV1FormEntry

`func NewApiV1FormEntry() *ApiV1FormEntry`

NewApiV1FormEntry instantiates a new ApiV1FormEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1FormEntryWithDefaults

`func NewApiV1FormEntryWithDefaults() *ApiV1FormEntry`

NewApiV1FormEntryWithDefaults instantiates a new ApiV1FormEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionNumber

`func (o *ApiV1FormEntry) GetAccessionNumber() string`

GetAccessionNumber returns the AccessionNumber field if non-nil, zero value otherwise.

### GetAccessionNumberOk

`func (o *ApiV1FormEntry) GetAccessionNumberOk() (*string, bool)`

GetAccessionNumberOk returns a tuple with the AccessionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionNumber

`func (o *ApiV1FormEntry) SetAccessionNumber(v string)`

SetAccessionNumber sets AccessionNumber field to given value.

### HasAccessionNumber

`func (o *ApiV1FormEntry) HasAccessionNumber() bool`

HasAccessionNumber returns a boolean if a field has been set.

### GetCik

`func (o *ApiV1FormEntry) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *ApiV1FormEntry) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *ApiV1FormEntry) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *ApiV1FormEntry) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetNameOfIssuer

`func (o *ApiV1FormEntry) GetNameOfIssuer() string`

GetNameOfIssuer returns the NameOfIssuer field if non-nil, zero value otherwise.

### GetNameOfIssuerOk

`func (o *ApiV1FormEntry) GetNameOfIssuerOk() (*string, bool)`

GetNameOfIssuerOk returns a tuple with the NameOfIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOfIssuer

`func (o *ApiV1FormEntry) SetNameOfIssuer(v string)`

SetNameOfIssuer sets NameOfIssuer field to given value.

### HasNameOfIssuer

`func (o *ApiV1FormEntry) HasNameOfIssuer() bool`

HasNameOfIssuer returns a boolean if a field has been set.

### GetTitleOfClass

`func (o *ApiV1FormEntry) GetTitleOfClass() string`

GetTitleOfClass returns the TitleOfClass field if non-nil, zero value otherwise.

### GetTitleOfClassOk

`func (o *ApiV1FormEntry) GetTitleOfClassOk() (*string, bool)`

GetTitleOfClassOk returns a tuple with the TitleOfClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleOfClass

`func (o *ApiV1FormEntry) SetTitleOfClass(v string)`

SetTitleOfClass sets TitleOfClass field to given value.

### HasTitleOfClass

`func (o *ApiV1FormEntry) HasTitleOfClass() bool`

HasTitleOfClass returns a boolean if a field has been set.

### GetCusip

`func (o *ApiV1FormEntry) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ApiV1FormEntry) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ApiV1FormEntry) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ApiV1FormEntry) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetTicker

`func (o *ApiV1FormEntry) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *ApiV1FormEntry) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *ApiV1FormEntry) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *ApiV1FormEntry) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetValue

`func (o *ApiV1FormEntry) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiV1FormEntry) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiV1FormEntry) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiV1FormEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ApiV1FormEntry) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ApiV1FormEntry) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSshPrnamt

`func (o *ApiV1FormEntry) GetSshPrnamt() int64`

GetSshPrnamt returns the SshPrnamt field if non-nil, zero value otherwise.

### GetSshPrnamtOk

`func (o *ApiV1FormEntry) GetSshPrnamtOk() (*int64, bool)`

GetSshPrnamtOk returns a tuple with the SshPrnamt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrnamt

`func (o *ApiV1FormEntry) SetSshPrnamt(v int64)`

SetSshPrnamt sets SshPrnamt field to given value.

### HasSshPrnamt

`func (o *ApiV1FormEntry) HasSshPrnamt() bool`

HasSshPrnamt returns a boolean if a field has been set.

### SetSshPrnamtNil

`func (o *ApiV1FormEntry) SetSshPrnamtNil(b bool)`

 SetSshPrnamtNil sets the value for SshPrnamt to be an explicit nil

### UnsetSshPrnamt
`func (o *ApiV1FormEntry) UnsetSshPrnamt()`

UnsetSshPrnamt ensures that no value is present for SshPrnamt, not even an explicit nil
### GetSshPrnamtType

`func (o *ApiV1FormEntry) GetSshPrnamtType() string`

GetSshPrnamtType returns the SshPrnamtType field if non-nil, zero value otherwise.

### GetSshPrnamtTypeOk

`func (o *ApiV1FormEntry) GetSshPrnamtTypeOk() (*string, bool)`

GetSshPrnamtTypeOk returns a tuple with the SshPrnamtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPrnamtType

`func (o *ApiV1FormEntry) SetSshPrnamtType(v string)`

SetSshPrnamtType sets SshPrnamtType field to given value.

### HasSshPrnamtType

`func (o *ApiV1FormEntry) HasSshPrnamtType() bool`

HasSshPrnamtType returns a boolean if a field has been set.

### SetSshPrnamtTypeNil

`func (o *ApiV1FormEntry) SetSshPrnamtTypeNil(b bool)`

 SetSshPrnamtTypeNil sets the value for SshPrnamtType to be an explicit nil

### UnsetSshPrnamtType
`func (o *ApiV1FormEntry) UnsetSshPrnamtType()`

UnsetSshPrnamtType ensures that no value is present for SshPrnamtType, not even an explicit nil
### GetInvestmentDiscretion

`func (o *ApiV1FormEntry) GetInvestmentDiscretion() string`

GetInvestmentDiscretion returns the InvestmentDiscretion field if non-nil, zero value otherwise.

### GetInvestmentDiscretionOk

`func (o *ApiV1FormEntry) GetInvestmentDiscretionOk() (*string, bool)`

GetInvestmentDiscretionOk returns a tuple with the InvestmentDiscretion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentDiscretion

`func (o *ApiV1FormEntry) SetInvestmentDiscretion(v string)`

SetInvestmentDiscretion sets InvestmentDiscretion field to given value.

### HasInvestmentDiscretion

`func (o *ApiV1FormEntry) HasInvestmentDiscretion() bool`

HasInvestmentDiscretion returns a boolean if a field has been set.

### SetInvestmentDiscretionNil

`func (o *ApiV1FormEntry) SetInvestmentDiscretionNil(b bool)`

 SetInvestmentDiscretionNil sets the value for InvestmentDiscretion to be an explicit nil

### UnsetInvestmentDiscretion
`func (o *ApiV1FormEntry) UnsetInvestmentDiscretion()`

UnsetInvestmentDiscretion ensures that no value is present for InvestmentDiscretion, not even an explicit nil
### GetVotingAuthoritySole

`func (o *ApiV1FormEntry) GetVotingAuthoritySole() int64`

GetVotingAuthoritySole returns the VotingAuthoritySole field if non-nil, zero value otherwise.

### GetVotingAuthoritySoleOk

`func (o *ApiV1FormEntry) GetVotingAuthoritySoleOk() (*int64, bool)`

GetVotingAuthoritySoleOk returns a tuple with the VotingAuthoritySole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingAuthoritySole

`func (o *ApiV1FormEntry) SetVotingAuthoritySole(v int64)`

SetVotingAuthoritySole sets VotingAuthoritySole field to given value.

### HasVotingAuthoritySole

`func (o *ApiV1FormEntry) HasVotingAuthoritySole() bool`

HasVotingAuthoritySole returns a boolean if a field has been set.

### SetVotingAuthoritySoleNil

`func (o *ApiV1FormEntry) SetVotingAuthoritySoleNil(b bool)`

 SetVotingAuthoritySoleNil sets the value for VotingAuthoritySole to be an explicit nil

### UnsetVotingAuthoritySole
`func (o *ApiV1FormEntry) UnsetVotingAuthoritySole()`

UnsetVotingAuthoritySole ensures that no value is present for VotingAuthoritySole, not even an explicit nil
### GetVotingAuthorityShared

`func (o *ApiV1FormEntry) GetVotingAuthorityShared() int64`

GetVotingAuthorityShared returns the VotingAuthorityShared field if non-nil, zero value otherwise.

### GetVotingAuthoritySharedOk

`func (o *ApiV1FormEntry) GetVotingAuthoritySharedOk() (*int64, bool)`

GetVotingAuthoritySharedOk returns a tuple with the VotingAuthorityShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingAuthorityShared

`func (o *ApiV1FormEntry) SetVotingAuthorityShared(v int64)`

SetVotingAuthorityShared sets VotingAuthorityShared field to given value.

### HasVotingAuthorityShared

`func (o *ApiV1FormEntry) HasVotingAuthorityShared() bool`

HasVotingAuthorityShared returns a boolean if a field has been set.

### SetVotingAuthoritySharedNil

`func (o *ApiV1FormEntry) SetVotingAuthoritySharedNil(b bool)`

 SetVotingAuthoritySharedNil sets the value for VotingAuthorityShared to be an explicit nil

### UnsetVotingAuthorityShared
`func (o *ApiV1FormEntry) UnsetVotingAuthorityShared()`

UnsetVotingAuthorityShared ensures that no value is present for VotingAuthorityShared, not even an explicit nil
### GetVotingAuthorityNone

`func (o *ApiV1FormEntry) GetVotingAuthorityNone() int64`

GetVotingAuthorityNone returns the VotingAuthorityNone field if non-nil, zero value otherwise.

### GetVotingAuthorityNoneOk

`func (o *ApiV1FormEntry) GetVotingAuthorityNoneOk() (*int64, bool)`

GetVotingAuthorityNoneOk returns a tuple with the VotingAuthorityNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingAuthorityNone

`func (o *ApiV1FormEntry) SetVotingAuthorityNone(v int64)`

SetVotingAuthorityNone sets VotingAuthorityNone field to given value.

### HasVotingAuthorityNone

`func (o *ApiV1FormEntry) HasVotingAuthorityNone() bool`

HasVotingAuthorityNone returns a boolean if a field has been set.

### SetVotingAuthorityNoneNil

`func (o *ApiV1FormEntry) SetVotingAuthorityNoneNil(b bool)`

 SetVotingAuthorityNoneNil sets the value for VotingAuthorityNone to be an explicit nil

### UnsetVotingAuthorityNone
`func (o *ApiV1FormEntry) UnsetVotingAuthorityNone()`

UnsetVotingAuthorityNone ensures that no value is present for VotingAuthorityNone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


