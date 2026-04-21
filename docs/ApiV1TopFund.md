# ApiV1TopFund


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cik** | **str** | The Central Index Key (CIK) of the fund. | [optional] 
**name** | **str** | The name of the fund. | [optional] 
**period_of_report** | **date** | The period of report date (end of quarter). | [optional] 
**pnl** | **float** | The portfolio PnL percentage for the given quarter. | [optional] 

## Example

```python
from forms13f.models.api_v1_top_fund import ApiV1TopFund

# TODO update the JSON string below
json = "{}"
# create an instance of ApiV1TopFund from a JSON string
api_v1_top_fund_instance = ApiV1TopFund.from_json(json)
# print the JSON string representation of the object
print(ApiV1TopFund.to_json())

# convert the object into a dict
api_v1_top_fund_dict = api_v1_top_fund_instance.to_dict()
# create an instance of ApiV1TopFund from a dict
api_v1_top_fund_from_dict = ApiV1TopFund.from_dict(api_v1_top_fund_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


