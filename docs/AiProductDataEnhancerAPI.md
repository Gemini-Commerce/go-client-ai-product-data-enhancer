# GeminiCommerce\AiProductDataEnhancer\AiProductDataEnhancerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiProductDataEnhancerFillProductData**](AiProductDataEnhancerAPI.md#AiProductDataEnhancerFillProductData) | **Post** /aiproductdataenhancer.AiProductDataEnhancer/FillProductData | 
[**AiProductDataEnhancerFillProductDataCheck**](AiProductDataEnhancerAPI.md#AiProductDataEnhancerFillProductDataCheck) | **Post** /aiproductdataenhancer.AiProductDataEnhancer/FillProductDataCheck | 
[**AiProductDataEnhancerTranslateData**](AiProductDataEnhancerAPI.md#AiProductDataEnhancerTranslateData) | **Post** /aiproductdataenhancer.AiProductDataEnhancer/TranslateData | 



## AiProductDataEnhancerFillProductData

> AiproductdataenhancerFillProductDataResponse AiProductDataEnhancerFillProductData(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-ai-product-data-enhancer"
)

func main() {
	body := *openapiclient.NewAiproductdataenhancerFillProductDataRequest() // AiproductdataenhancerFillProductDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProductDataEnhancerFillProductData`: AiproductdataenhancerFillProductDataResponse
	fmt.Fprintf(os.Stdout, "Response from `AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProductDataEnhancerFillProductDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AiproductdataenhancerFillProductDataRequest**](AiproductdataenhancerFillProductDataRequest.md) |  | 

### Return type

[**AiproductdataenhancerFillProductDataResponse**](AiproductdataenhancerFillProductDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProductDataEnhancerFillProductDataCheck

> AiproductdataenhancerFillProductDataCheckResponse AiProductDataEnhancerFillProductDataCheck(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-ai-product-data-enhancer"
)

func main() {
	body := *openapiclient.NewAiproductdataenhancerFillProductDataCheckRequest() // AiproductdataenhancerFillProductDataCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductDataCheck(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductDataCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProductDataEnhancerFillProductDataCheck`: AiproductdataenhancerFillProductDataCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `AiProductDataEnhancerAPI.AiProductDataEnhancerFillProductDataCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProductDataEnhancerFillProductDataCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AiproductdataenhancerFillProductDataCheckRequest**](AiproductdataenhancerFillProductDataCheckRequest.md) |  | 

### Return type

[**AiproductdataenhancerFillProductDataCheckResponse**](AiproductdataenhancerFillProductDataCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiProductDataEnhancerTranslateData

> AiproductdataenhancerTranslateDataResponse AiProductDataEnhancerTranslateData(ctx).Body(body).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-ai-product-data-enhancer"
)

func main() {
	body := *openapiclient.NewAiproductdataenhancerTranslateDataRequest() // AiproductdataenhancerTranslateDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AiProductDataEnhancerAPI.AiProductDataEnhancerTranslateData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AiProductDataEnhancerAPI.AiProductDataEnhancerTranslateData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiProductDataEnhancerTranslateData`: AiproductdataenhancerTranslateDataResponse
	fmt.Fprintf(os.Stdout, "Response from `AiProductDataEnhancerAPI.AiProductDataEnhancerTranslateData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiProductDataEnhancerTranslateDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AiproductdataenhancerTranslateDataRequest**](AiproductdataenhancerTranslateDataRequest.md) |  | 

### Return type

[**AiproductdataenhancerTranslateDataResponse**](AiproductdataenhancerTranslateDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

