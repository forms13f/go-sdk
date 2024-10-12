# ApiV1Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URL of the form. | [optional] 
**AccessionNumber** | Pointer to **string** | The accession number of the form. | [optional] 
**SubmissionType** | Pointer to **string** | The submission type of the form. | [optional] 
**PublicDocumentCount** | Pointer to **NullableInt32** | The public document count. | [optional] 
**PeriodOfReport** | Pointer to **string** | The period of the report. | [optional] 
**FiledAsOfDate** | Pointer to **string** | The filed as of date. | [optional] 
**DateAsOfChange** | Pointer to **string** | The date as of change. | [optional] 
**EffectivenessDate** | Pointer to **string** | The effectiveness date. | [optional] 
**Cik** | Pointer to **string** | The Central Index Key (CIK). | [optional] 
**CompanyName** | Pointer to **string** | The company name. | [optional] 
**IrsNumber** | Pointer to **string** | The IRS number. | [optional] 
**StateOfIncorporation** | Pointer to **string** | The state of incorporation. | [optional] 
**FiscalYearEnd** | Pointer to **string** | The fiscal year end. | [optional] 
**FormType** | Pointer to **string** | The form type. | [optional] 
**SecAct** | Pointer to **string** | The SEC act. | [optional] 
**SecFileNumber** | Pointer to **string** | The SEC file number. | [optional] 
**FilmNumber** | Pointer to **string** | The film number. | [optional] 
**BusinessAddress** | Pointer to **NullableString** | The business address. | [optional] 
**BusinessPhone** | Pointer to **NullableString** | The business phone. | [optional] 
**TableValueTotal** | Pointer to **NullableInt64** | The total value of the table. | [optional] 
**TableEntryTotal** | Pointer to **NullableInt64** | The total entry count of the table. | [optional] 
**IsAmendment** | Pointer to **bool** | Indicates if the form is an amendment. | [optional] 
**AmendmentType** | Pointer to **string** | The type of amendment. | [optional] 
**ConfDeniedExpired** | Pointer to **bool** | Indicates if the confidentiality request was denied or expired. | [optional] 
**ConfDateDeniedExpired** | Pointer to **string** | The date when the confidentiality request was denied or expired. | [optional] 
**AmendmentDateReported** | Pointer to **string** | The date when the amendment was reported. | [optional] 

## Methods

### NewApiV1Form

`func NewApiV1Form() *ApiV1Form`

NewApiV1Form instantiates a new ApiV1Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1FormWithDefaults

`func NewApiV1FormWithDefaults() *ApiV1Form`

NewApiV1FormWithDefaults instantiates a new ApiV1Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ApiV1Form) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiV1Form) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiV1Form) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiV1Form) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAccessionNumber

`func (o *ApiV1Form) GetAccessionNumber() string`

GetAccessionNumber returns the AccessionNumber field if non-nil, zero value otherwise.

### GetAccessionNumberOk

`func (o *ApiV1Form) GetAccessionNumberOk() (*string, bool)`

GetAccessionNumberOk returns a tuple with the AccessionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionNumber

`func (o *ApiV1Form) SetAccessionNumber(v string)`

SetAccessionNumber sets AccessionNumber field to given value.

### HasAccessionNumber

`func (o *ApiV1Form) HasAccessionNumber() bool`

HasAccessionNumber returns a boolean if a field has been set.

### GetSubmissionType

`func (o *ApiV1Form) GetSubmissionType() string`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *ApiV1Form) GetSubmissionTypeOk() (*string, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *ApiV1Form) SetSubmissionType(v string)`

SetSubmissionType sets SubmissionType field to given value.

### HasSubmissionType

`func (o *ApiV1Form) HasSubmissionType() bool`

HasSubmissionType returns a boolean if a field has been set.

### GetPublicDocumentCount

`func (o *ApiV1Form) GetPublicDocumentCount() int32`

GetPublicDocumentCount returns the PublicDocumentCount field if non-nil, zero value otherwise.

### GetPublicDocumentCountOk

`func (o *ApiV1Form) GetPublicDocumentCountOk() (*int32, bool)`

GetPublicDocumentCountOk returns a tuple with the PublicDocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDocumentCount

`func (o *ApiV1Form) SetPublicDocumentCount(v int32)`

SetPublicDocumentCount sets PublicDocumentCount field to given value.

### HasPublicDocumentCount

`func (o *ApiV1Form) HasPublicDocumentCount() bool`

HasPublicDocumentCount returns a boolean if a field has been set.

### SetPublicDocumentCountNil

`func (o *ApiV1Form) SetPublicDocumentCountNil(b bool)`

 SetPublicDocumentCountNil sets the value for PublicDocumentCount to be an explicit nil

### UnsetPublicDocumentCount
`func (o *ApiV1Form) UnsetPublicDocumentCount()`

UnsetPublicDocumentCount ensures that no value is present for PublicDocumentCount, not even an explicit nil
### GetPeriodOfReport

`func (o *ApiV1Form) GetPeriodOfReport() string`

GetPeriodOfReport returns the PeriodOfReport field if non-nil, zero value otherwise.

### GetPeriodOfReportOk

`func (o *ApiV1Form) GetPeriodOfReportOk() (*string, bool)`

GetPeriodOfReportOk returns a tuple with the PeriodOfReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodOfReport

`func (o *ApiV1Form) SetPeriodOfReport(v string)`

SetPeriodOfReport sets PeriodOfReport field to given value.

### HasPeriodOfReport

`func (o *ApiV1Form) HasPeriodOfReport() bool`

HasPeriodOfReport returns a boolean if a field has been set.

### GetFiledAsOfDate

`func (o *ApiV1Form) GetFiledAsOfDate() string`

GetFiledAsOfDate returns the FiledAsOfDate field if non-nil, zero value otherwise.

### GetFiledAsOfDateOk

`func (o *ApiV1Form) GetFiledAsOfDateOk() (*string, bool)`

GetFiledAsOfDateOk returns a tuple with the FiledAsOfDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiledAsOfDate

`func (o *ApiV1Form) SetFiledAsOfDate(v string)`

SetFiledAsOfDate sets FiledAsOfDate field to given value.

### HasFiledAsOfDate

`func (o *ApiV1Form) HasFiledAsOfDate() bool`

HasFiledAsOfDate returns a boolean if a field has been set.

### GetDateAsOfChange

`func (o *ApiV1Form) GetDateAsOfChange() string`

GetDateAsOfChange returns the DateAsOfChange field if non-nil, zero value otherwise.

### GetDateAsOfChangeOk

`func (o *ApiV1Form) GetDateAsOfChangeOk() (*string, bool)`

GetDateAsOfChangeOk returns a tuple with the DateAsOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAsOfChange

`func (o *ApiV1Form) SetDateAsOfChange(v string)`

SetDateAsOfChange sets DateAsOfChange field to given value.

### HasDateAsOfChange

`func (o *ApiV1Form) HasDateAsOfChange() bool`

HasDateAsOfChange returns a boolean if a field has been set.

### GetEffectivenessDate

`func (o *ApiV1Form) GetEffectivenessDate() string`

GetEffectivenessDate returns the EffectivenessDate field if non-nil, zero value otherwise.

### GetEffectivenessDateOk

`func (o *ApiV1Form) GetEffectivenessDateOk() (*string, bool)`

GetEffectivenessDateOk returns a tuple with the EffectivenessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivenessDate

`func (o *ApiV1Form) SetEffectivenessDate(v string)`

SetEffectivenessDate sets EffectivenessDate field to given value.

### HasEffectivenessDate

`func (o *ApiV1Form) HasEffectivenessDate() bool`

HasEffectivenessDate returns a boolean if a field has been set.

### GetCik

`func (o *ApiV1Form) GetCik() string`

GetCik returns the Cik field if non-nil, zero value otherwise.

### GetCikOk

`func (o *ApiV1Form) GetCikOk() (*string, bool)`

GetCikOk returns a tuple with the Cik field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCik

`func (o *ApiV1Form) SetCik(v string)`

SetCik sets Cik field to given value.

### HasCik

`func (o *ApiV1Form) HasCik() bool`

HasCik returns a boolean if a field has been set.

### GetCompanyName

`func (o *ApiV1Form) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ApiV1Form) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ApiV1Form) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ApiV1Form) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetIrsNumber

`func (o *ApiV1Form) GetIrsNumber() string`

GetIrsNumber returns the IrsNumber field if non-nil, zero value otherwise.

### GetIrsNumberOk

`func (o *ApiV1Form) GetIrsNumberOk() (*string, bool)`

GetIrsNumberOk returns a tuple with the IrsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrsNumber

`func (o *ApiV1Form) SetIrsNumber(v string)`

SetIrsNumber sets IrsNumber field to given value.

### HasIrsNumber

`func (o *ApiV1Form) HasIrsNumber() bool`

HasIrsNumber returns a boolean if a field has been set.

### GetStateOfIncorporation

`func (o *ApiV1Form) GetStateOfIncorporation() string`

GetStateOfIncorporation returns the StateOfIncorporation field if non-nil, zero value otherwise.

### GetStateOfIncorporationOk

`func (o *ApiV1Form) GetStateOfIncorporationOk() (*string, bool)`

GetStateOfIncorporationOk returns a tuple with the StateOfIncorporation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOfIncorporation

`func (o *ApiV1Form) SetStateOfIncorporation(v string)`

SetStateOfIncorporation sets StateOfIncorporation field to given value.

### HasStateOfIncorporation

`func (o *ApiV1Form) HasStateOfIncorporation() bool`

HasStateOfIncorporation returns a boolean if a field has been set.

### GetFiscalYearEnd

`func (o *ApiV1Form) GetFiscalYearEnd() string`

GetFiscalYearEnd returns the FiscalYearEnd field if non-nil, zero value otherwise.

### GetFiscalYearEndOk

`func (o *ApiV1Form) GetFiscalYearEndOk() (*string, bool)`

GetFiscalYearEndOk returns a tuple with the FiscalYearEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearEnd

`func (o *ApiV1Form) SetFiscalYearEnd(v string)`

SetFiscalYearEnd sets FiscalYearEnd field to given value.

### HasFiscalYearEnd

`func (o *ApiV1Form) HasFiscalYearEnd() bool`

HasFiscalYearEnd returns a boolean if a field has been set.

### GetFormType

`func (o *ApiV1Form) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *ApiV1Form) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *ApiV1Form) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *ApiV1Form) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetSecAct

`func (o *ApiV1Form) GetSecAct() string`

GetSecAct returns the SecAct field if non-nil, zero value otherwise.

### GetSecActOk

`func (o *ApiV1Form) GetSecActOk() (*string, bool)`

GetSecActOk returns a tuple with the SecAct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAct

`func (o *ApiV1Form) SetSecAct(v string)`

SetSecAct sets SecAct field to given value.

### HasSecAct

`func (o *ApiV1Form) HasSecAct() bool`

HasSecAct returns a boolean if a field has been set.

### GetSecFileNumber

`func (o *ApiV1Form) GetSecFileNumber() string`

GetSecFileNumber returns the SecFileNumber field if non-nil, zero value otherwise.

### GetSecFileNumberOk

`func (o *ApiV1Form) GetSecFileNumberOk() (*string, bool)`

GetSecFileNumberOk returns a tuple with the SecFileNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecFileNumber

`func (o *ApiV1Form) SetSecFileNumber(v string)`

SetSecFileNumber sets SecFileNumber field to given value.

### HasSecFileNumber

`func (o *ApiV1Form) HasSecFileNumber() bool`

HasSecFileNumber returns a boolean if a field has been set.

### GetFilmNumber

`func (o *ApiV1Form) GetFilmNumber() string`

GetFilmNumber returns the FilmNumber field if non-nil, zero value otherwise.

### GetFilmNumberOk

`func (o *ApiV1Form) GetFilmNumberOk() (*string, bool)`

GetFilmNumberOk returns a tuple with the FilmNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilmNumber

`func (o *ApiV1Form) SetFilmNumber(v string)`

SetFilmNumber sets FilmNumber field to given value.

### HasFilmNumber

`func (o *ApiV1Form) HasFilmNumber() bool`

HasFilmNumber returns a boolean if a field has been set.

### GetBusinessAddress

`func (o *ApiV1Form) GetBusinessAddress() string`

GetBusinessAddress returns the BusinessAddress field if non-nil, zero value otherwise.

### GetBusinessAddressOk

`func (o *ApiV1Form) GetBusinessAddressOk() (*string, bool)`

GetBusinessAddressOk returns a tuple with the BusinessAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddress

`func (o *ApiV1Form) SetBusinessAddress(v string)`

SetBusinessAddress sets BusinessAddress field to given value.

### HasBusinessAddress

`func (o *ApiV1Form) HasBusinessAddress() bool`

HasBusinessAddress returns a boolean if a field has been set.

### SetBusinessAddressNil

`func (o *ApiV1Form) SetBusinessAddressNil(b bool)`

 SetBusinessAddressNil sets the value for BusinessAddress to be an explicit nil

### UnsetBusinessAddress
`func (o *ApiV1Form) UnsetBusinessAddress()`

UnsetBusinessAddress ensures that no value is present for BusinessAddress, not even an explicit nil
### GetBusinessPhone

`func (o *ApiV1Form) GetBusinessPhone() string`

GetBusinessPhone returns the BusinessPhone field if non-nil, zero value otherwise.

### GetBusinessPhoneOk

`func (o *ApiV1Form) GetBusinessPhoneOk() (*string, bool)`

GetBusinessPhoneOk returns a tuple with the BusinessPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessPhone

`func (o *ApiV1Form) SetBusinessPhone(v string)`

SetBusinessPhone sets BusinessPhone field to given value.

### HasBusinessPhone

`func (o *ApiV1Form) HasBusinessPhone() bool`

HasBusinessPhone returns a boolean if a field has been set.

### SetBusinessPhoneNil

`func (o *ApiV1Form) SetBusinessPhoneNil(b bool)`

 SetBusinessPhoneNil sets the value for BusinessPhone to be an explicit nil

### UnsetBusinessPhone
`func (o *ApiV1Form) UnsetBusinessPhone()`

UnsetBusinessPhone ensures that no value is present for BusinessPhone, not even an explicit nil
### GetTableValueTotal

`func (o *ApiV1Form) GetTableValueTotal() int64`

GetTableValueTotal returns the TableValueTotal field if non-nil, zero value otherwise.

### GetTableValueTotalOk

`func (o *ApiV1Form) GetTableValueTotalOk() (*int64, bool)`

GetTableValueTotalOk returns a tuple with the TableValueTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableValueTotal

`func (o *ApiV1Form) SetTableValueTotal(v int64)`

SetTableValueTotal sets TableValueTotal field to given value.

### HasTableValueTotal

`func (o *ApiV1Form) HasTableValueTotal() bool`

HasTableValueTotal returns a boolean if a field has been set.

### SetTableValueTotalNil

`func (o *ApiV1Form) SetTableValueTotalNil(b bool)`

 SetTableValueTotalNil sets the value for TableValueTotal to be an explicit nil

### UnsetTableValueTotal
`func (o *ApiV1Form) UnsetTableValueTotal()`

UnsetTableValueTotal ensures that no value is present for TableValueTotal, not even an explicit nil
### GetTableEntryTotal

`func (o *ApiV1Form) GetTableEntryTotal() int64`

GetTableEntryTotal returns the TableEntryTotal field if non-nil, zero value otherwise.

### GetTableEntryTotalOk

`func (o *ApiV1Form) GetTableEntryTotalOk() (*int64, bool)`

GetTableEntryTotalOk returns a tuple with the TableEntryTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableEntryTotal

`func (o *ApiV1Form) SetTableEntryTotal(v int64)`

SetTableEntryTotal sets TableEntryTotal field to given value.

### HasTableEntryTotal

`func (o *ApiV1Form) HasTableEntryTotal() bool`

HasTableEntryTotal returns a boolean if a field has been set.

### SetTableEntryTotalNil

`func (o *ApiV1Form) SetTableEntryTotalNil(b bool)`

 SetTableEntryTotalNil sets the value for TableEntryTotal to be an explicit nil

### UnsetTableEntryTotal
`func (o *ApiV1Form) UnsetTableEntryTotal()`

UnsetTableEntryTotal ensures that no value is present for TableEntryTotal, not even an explicit nil
### GetIsAmendment

`func (o *ApiV1Form) GetIsAmendment() bool`

GetIsAmendment returns the IsAmendment field if non-nil, zero value otherwise.

### GetIsAmendmentOk

`func (o *ApiV1Form) GetIsAmendmentOk() (*bool, bool)`

GetIsAmendmentOk returns a tuple with the IsAmendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmendment

`func (o *ApiV1Form) SetIsAmendment(v bool)`

SetIsAmendment sets IsAmendment field to given value.

### HasIsAmendment

`func (o *ApiV1Form) HasIsAmendment() bool`

HasIsAmendment returns a boolean if a field has been set.

### GetAmendmentType

`func (o *ApiV1Form) GetAmendmentType() string`

GetAmendmentType returns the AmendmentType field if non-nil, zero value otherwise.

### GetAmendmentTypeOk

`func (o *ApiV1Form) GetAmendmentTypeOk() (*string, bool)`

GetAmendmentTypeOk returns a tuple with the AmendmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentType

`func (o *ApiV1Form) SetAmendmentType(v string)`

SetAmendmentType sets AmendmentType field to given value.

### HasAmendmentType

`func (o *ApiV1Form) HasAmendmentType() bool`

HasAmendmentType returns a boolean if a field has been set.

### GetConfDeniedExpired

`func (o *ApiV1Form) GetConfDeniedExpired() bool`

GetConfDeniedExpired returns the ConfDeniedExpired field if non-nil, zero value otherwise.

### GetConfDeniedExpiredOk

`func (o *ApiV1Form) GetConfDeniedExpiredOk() (*bool, bool)`

GetConfDeniedExpiredOk returns a tuple with the ConfDeniedExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfDeniedExpired

`func (o *ApiV1Form) SetConfDeniedExpired(v bool)`

SetConfDeniedExpired sets ConfDeniedExpired field to given value.

### HasConfDeniedExpired

`func (o *ApiV1Form) HasConfDeniedExpired() bool`

HasConfDeniedExpired returns a boolean if a field has been set.

### GetConfDateDeniedExpired

`func (o *ApiV1Form) GetConfDateDeniedExpired() string`

GetConfDateDeniedExpired returns the ConfDateDeniedExpired field if non-nil, zero value otherwise.

### GetConfDateDeniedExpiredOk

`func (o *ApiV1Form) GetConfDateDeniedExpiredOk() (*string, bool)`

GetConfDateDeniedExpiredOk returns a tuple with the ConfDateDeniedExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfDateDeniedExpired

`func (o *ApiV1Form) SetConfDateDeniedExpired(v string)`

SetConfDateDeniedExpired sets ConfDateDeniedExpired field to given value.

### HasConfDateDeniedExpired

`func (o *ApiV1Form) HasConfDateDeniedExpired() bool`

HasConfDateDeniedExpired returns a boolean if a field has been set.

### GetAmendmentDateReported

`func (o *ApiV1Form) GetAmendmentDateReported() string`

GetAmendmentDateReported returns the AmendmentDateReported field if non-nil, zero value otherwise.

### GetAmendmentDateReportedOk

`func (o *ApiV1Form) GetAmendmentDateReportedOk() (*string, bool)`

GetAmendmentDateReportedOk returns a tuple with the AmendmentDateReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentDateReported

`func (o *ApiV1Form) SetAmendmentDateReported(v string)`

SetAmendmentDateReported sets AmendmentDateReported field to given value.

### HasAmendmentDateReported

`func (o *ApiV1Form) HasAmendmentDateReported() bool`

HasAmendmentDateReported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


