# \ForwardsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdForwardsGet**](ForwardsApi.md#DomainsDomainIdForwardsGet) | **Get** /domains/{domainId}/forwards/ | List forwards
[**DomainsDomainIdForwardsHostDelete**](ForwardsApi.md#DomainsDomainIdForwardsHostDelete) | **Delete** /domains/{domainId}/forwards/{host} | Delete forward by host
[**DomainsDomainIdForwardsHostGet**](ForwardsApi.md#DomainsDomainIdForwardsHostGet) | **Get** /domains/{domainId}/forwards/{host} | Find forward by host
[**DomainsDomainIdForwardsHostPut**](ForwardsApi.md#DomainsDomainIdForwardsHostPut) | **Put** /domains/{domainId}/forwards/{host} | Update forward by host
[**DomainsDomainIdForwardsPost**](ForwardsApi.md#DomainsDomainIdForwardsPost) | **Post** /domains/{domainId}/forwards/ | Add forward



## DomainsDomainIdForwardsGet

> []HTTPForward DomainsDomainIdForwardsGet(ctx, domainId).Execute()

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
    resp, r, err := api_client.ForwardsApi.DomainsDomainIdForwardsGet(context.Background(), domainId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DomainsDomainIdForwardsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainIdForwardsGet`: []HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.DomainsDomainIdForwardsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdForwardsGetRequest struct via the builder pattern


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


## DomainsDomainIdForwardsHostDelete

> DomainsDomainIdForwardsHostDelete(ctx, domainId, host).Execute()

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
    resp, r, err := api_client.ForwardsApi.DomainsDomainIdForwardsHostDelete(context.Background(), domainId, host).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DomainsDomainIdForwardsHostDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDomainsDomainIdForwardsHostDeleteRequest struct via the builder pattern


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


## DomainsDomainIdForwardsHostGet

> HTTPForward DomainsDomainIdForwardsHostGet(ctx, domainId, host).Execute()

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
    resp, r, err := api_client.ForwardsApi.DomainsDomainIdForwardsHostGet(context.Background(), domainId, host).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DomainsDomainIdForwardsHostGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainIdForwardsHostGet`: HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.DomainsDomainIdForwardsHostGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**host** | **string** | Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdForwardsHostGetRequest struct via the builder pattern


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


## DomainsDomainIdForwardsHostPut

> HTTPForward DomainsDomainIdForwardsHostPut(ctx, domainId, host).HTTPForward(hTTPForward).Execute()

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
    hTTPForward := *openapiclient.NewHTTPForward() // HTTPForward |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.DomainsDomainIdForwardsHostPut(context.Background(), domainId, host).HTTPForward(hTTPForward).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DomainsDomainIdForwardsHostPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainIdForwardsHostPut`: HTTPForward
    fmt.Fprintf(os.Stdout, "Response from `ForwardsApi.DomainsDomainIdForwardsHostPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**host** | **string** | Subdomain of the forward, &#x60;@&#x60; for the root domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdForwardsHostPutRequest struct via the builder pattern


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


## DomainsDomainIdForwardsPost

> DomainsDomainIdForwardsPost(ctx, domainId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ForwardsApi.DomainsDomainIdForwardsPost(context.Background(), domainId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `ForwardsApi.DomainsDomainIdForwardsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDomainsDomainIdForwardsPostRequest struct via the builder pattern


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

