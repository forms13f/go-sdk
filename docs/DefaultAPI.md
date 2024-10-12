# \DefaultAPI

All URIs are relative to *https://forms13f.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1FilerGet**](DefaultAPI.md#ApiV1FilerGet) | **Get** /api/v1/filer | Retrieve a filer by CIK.
[**ApiV1FilersGet**](DefaultAPI.md#ApiV1FilersGet) | **Get** /api/v1/filers | Retrieve unique filers.
[**ApiV1FilingsGet**](DefaultAPI.md#ApiV1FilingsGet) | **Get** /api/v1/filings | Retrieve 13F filings for all filers in the time range.
[**ApiV1FormGet**](DefaultAPI.md#ApiV1FormGet) | **Get** /api/v1/form | Get SEC Form 13F.
[**ApiV1FormsGet**](DefaultAPI.md#ApiV1FormsGet) | **Get** /api/v1/forms | Retrieve SEC forms 13F for a filer.
[**ApiV1FundsGet**](DefaultAPI.md#ApiV1FundsGet) | **Get** /api/v1/funds | Retrieve unique companies.
[**ApiV1TickersGet**](DefaultAPI.md#ApiV1TickersGet) | **Get** /api/v1/tickers | Returns cusip, ticker, and company name for provided matching cusips or tickers parameters



## ApiV1FilerGet

> ApiV1Filer ApiV1FilerGet(ctx).Cik(cik).Execute()

Retrieve a filer by CIK.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	cik := "1067983" // string | The Central Index Key (CIK) of the filer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FilerGet(context.Background()).Cik(cik).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FilerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FilerGet`: ApiV1Filer
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FilerGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FilerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cik** | **string** | The Central Index Key (CIK) of the filer. | 

### Return type

[**ApiV1Filer**](ApiV1Filer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FilersGet

> []ApiV1Filer ApiV1FilersGet(ctx).Offset(offset).Limit(limit).Execute()

Retrieve unique filers.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	offset := int32(0) // int32 | Skip previous offset CIKs (optional) (default to 0)
	limit := int32(100) // int32 | Return max limit CIKs (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FilersGet(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FilersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FilersGet`: []ApiV1Filer
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FilersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FilersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Skip previous offset CIKs | [default to 0]
 **limit** | **int32** | Return max limit CIKs | [default to 100]

### Return type

[**[]ApiV1Filer**](ApiV1Filer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FilingsGet

> []ApiV1Form ApiV1FilingsGet(ctx).From(from).To(to).Offset(offset).Limit(limit).Execute()

Retrieve 13F filings for all filers in the time range.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	from := time.Now() // string | All filings returned will be on or after this filing date.
	to := time.Now() // string | All filings returned will be on or before this filing date.
	offset := int32(0) // int32 | Skip the first offset filings. (optional) (default to 0)
	limit := int32(100) // int32 | Return at most limit filings. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FilingsGet(context.Background()).From(from).To(to).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FilingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FilingsGet`: []ApiV1Form
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FilingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FilingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | All filings returned will be on or after this filing date. | 
 **to** | **string** | All filings returned will be on or before this filing date. | 
 **offset** | **int32** | Skip the first offset filings. | [default to 0]
 **limit** | **int32** | Return at most limit filings. | [default to 100]

### Return type

[**[]ApiV1Form**](ApiV1Form.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FormGet

> []ApiV1FormEntry ApiV1FormGet(ctx).AccessionNumber(accessionNumber).Cik(cik).Offset(offset).Limit(limit).Execute()

Get SEC Form 13F.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	accessionNumber := "0000950123-24-008740" // string | The accession number of the form entry.
	cik := "0001067983" // string | The Central Index Key (CIK) of the form entry.
	offset := int32(0) // int32 | The offset for pagination. (optional) (default to 0)
	limit := int32(100) // int32 | The limit for pagination. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FormGet(context.Background()).AccessionNumber(accessionNumber).Cik(cik).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FormGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FormGet`: []ApiV1FormEntry
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FormGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FormGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessionNumber** | **string** | The accession number of the form entry. | 
 **cik** | **string** | The Central Index Key (CIK) of the form entry. | 
 **offset** | **int32** | The offset for pagination. | [default to 0]
 **limit** | **int32** | The limit for pagination. | [default to 100]

### Return type

[**[]ApiV1FormEntry**](ApiV1FormEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FormsGet

> []ApiV1Form ApiV1FormsGet(ctx).Cik(cik).From(from).To(to).Offset(offset).Limit(limit).Execute()

Retrieve SEC forms 13F for a filer.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	cik := "0001067983" // string | The Central Index Key (CIK) of the filer.
	from := "2023-12-31" // string | All forms returned will be on or after this period of report date. (optional) (default to "2010-01-01")
	to := "2024-12-31" // string | All forms returned will be on or before this period of report date. (optional) (default to "2030-01-01")
	offset := int32(0) // int32 | Skip the first offset forms. (optional) (default to 0)
	limit := int32(100) // int32 | Return at most limit forms. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FormsGet(context.Background()).Cik(cik).From(from).To(to).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FormsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FormsGet`: []ApiV1Form
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FormsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FormsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cik** | **string** | The Central Index Key (CIK) of the filer. | 
 **from** | **string** | All forms returned will be on or after this period of report date. | [default to &quot;2010-01-01&quot;]
 **to** | **string** | All forms returned will be on or before this period of report date. | [default to &quot;2030-01-01&quot;]
 **offset** | **int32** | Skip the first offset forms. | [default to 0]
 **limit** | **int32** | Return at most limit forms. | [default to 100]

### Return type

[**[]ApiV1Form**](ApiV1Form.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1FundsGet

> []ApiV1Fund ApiV1FundsGet(ctx).Name(name).Offset(offset).Limit(limit).Execute()

Retrieve unique companies.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	name := "berkshire" // string | Filter companies by by a substring of their name (optional)
	offset := int32(0) // int32 | Skip previous offset companies (optional) (default to 0)
	limit := int32(100) // int32 | Return max limit companies (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1FundsGet(context.Background()).Name(name).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1FundsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1FundsGet`: []ApiV1Fund
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1FundsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1FundsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter companies by by a substring of their name | 
 **offset** | **int32** | Skip previous offset companies | [default to 0]
 **limit** | **int32** | Return max limit companies | [default to 100]

### Return type

[**[]ApiV1Fund**](ApiV1Fund.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1TickersGet

> []TickerInfo ApiV1TickersGet(ctx).Cusips(cusips).Tickers(tickers).Execute()

Returns cusip, ticker, and company name for provided matching cusips or tickers parameters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/forms13f/go-sdk"
)

func main() {
	cusips := []string{"Inner_example"} // []string | Array of CUSIPs to match (optional)
	tickers := []string{"Inner_example"} // []string | Array of tickers to match (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1TickersGet(context.Background()).Cusips(cusips).Tickers(tickers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1TickersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1TickersGet`: []TickerInfo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1TickersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1TickersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cusips** | **[]string** | Array of CUSIPs to match | 
 **tickers** | **[]string** | Array of tickers to match | 

### Return type

[**[]TickerInfo**](TickerInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

