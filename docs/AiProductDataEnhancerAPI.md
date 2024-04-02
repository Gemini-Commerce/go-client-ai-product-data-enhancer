# GeminiCommerce\AiProductDataEnhancer\AiProductDataEnhancerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiProductDataEnhancerFillProductData**](AiProductDataEnhancerAPI.md#AiProductDataEnhancerFillProductData) | **Post** /aiproductdataenhancer.AiProductDataEnhancer/FillProductData | 



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

