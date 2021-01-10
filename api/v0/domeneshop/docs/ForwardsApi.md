# \ForwardsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForward**](ForwardsApi.md#CreateForward) | **Post** /domains/{domainId}/forwards | Add forward
[**DeleteForward**](ForwardsApi.md#DeleteForward) | **Delete** /domains/{domainId}/forwards/{host} | Delete forward by host
[**GetForward**](ForwardsApi.md#GetForward) | **Get** /domains/{domainId}/forwards/{host} | Find forward by host
[**GetForwards**](ForwardsApi.md#GetForwards) | **Get** /domains/{domainId}/forwards | List forwards
[**ModifyForward**](ForwardsApi.md#ModifyForward) | **Put** /domains/{domainId}/forwards/{host} | Update forward by host



## CreateForward

> CreateForward(ctx, domainId).HTTPForward(hTTPForward).Execute()

Add forward



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | ID of the domain
    hTTPForward := *openapiclient.NewHTTPForward("@", "https://www.example.com") // HTTPForward |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.CreateForward(context.Background(), domainId).HTTPForward(hTTPForward).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.CreateForward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateForwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hTTPForward** | [**HTTPForward**](HTTPForward.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForward

> DeleteForward(ctx, domainId, host).Execute()

Delete forward by host

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | ID of the domain
    host := "www" // string | Subdomain of the forward, `@` for the root domain

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.DeleteForward(context.Background(), domainId, host).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DeleteForward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**host** | **string** | Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForward

> HTTPForward GetForward(ctx, domainId, host).Execute()

Find forward by host

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | ID of the domain
    host := "www" // string | Subdomain of the forward, `@` for the root domain

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.GetForward(context.Background(), domainId, host).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.GetForward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetForward`: HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.GetForward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**host** | **string** | Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HTTPForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForwards

> []HTTPForward GetForwards(ctx, domainId).Execute()

List forwards



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | ID of the domain

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.GetForwards(context.Background(), domainId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.GetForwards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetForwards`: []HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.GetForwards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForwardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]HTTPForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyForward

> HTTPForward ModifyForward(ctx, domainId, host).HTTPForward(hTTPForward).Execute()

Update forward by host



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | ID of the domain
    host := "www" // string | Subdomain of the forward, `@` for the root domain
    hTTPForward := *openapiclient.NewHTTPForward("@", "https://www.example.com") // HTTPForward |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.ModifyForward(context.Background(), domainId, host).HTTPForward(hTTPForward).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.ModifyForward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyForward`: HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.ModifyForward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**host** | **string** | Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyForwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hTTPForward** | [**HTTPForward**](HTTPForward.md) |  | 

### Return type

[**HTTPForward**](HTTPForward.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

