# ApiV1HoldersGet200Response


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cusip** | **str** | The CUSIP of the securities. | [optional] 
**period_of_report** | **date** | The period of report date. | [optional] 
**ciks** | **List[str]** | An array of CIKs. | [optional] 

## Example

```python
from forms13f.models.api_v1_holders_get200_response import ApiV1HoldersGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of ApiV1HoldersGet200Response from a JSON string
api_v1_holders_get200_response_instance = ApiV1HoldersGet200Response.from_json(json)
# print the JSON string representation of the object
print(ApiV1HoldersGet200Response.to_json())

# convert the object into a dict
api_v1_holders_get200_response_dict = api_v1_holders_get200_response_instance.to_dict()
# create an instance of ApiV1HoldersGet200Response from a dict
api_v1_holders_get200_response_from_dict = ApiV1HoldersGet200Response.from_dict(api_v1_holders_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


