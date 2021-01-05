# \DomainsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdGet**](DomainsApi.md#DomainsDomainIdGet) | **Get** /domains/{domainId} | Find domain by ID
[**GetDomains**](DomainsApi.md#GetDomains) | **Get** /domains | List domains



## DomainsDomainIdGet

> Domain DomainsDomainIdGet(ctx, domainId).Execute()

Find domain by ID

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
    resp, r, err := api_client.DomainsApi.DomainsDomainIdGet(context.Background(), domainId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.DomainsDomainIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainIdGet`: Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.DomainsDomainIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Domain**](Domain.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> []Domain GetDomains(ctx).Domain(domain).Execute()

List domains

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
    domain := ".no" // string | Only return domains whose `domain` field includes this string (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DomainsApi.GetDomains(context.Background()).Domain(domain).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsApi.GetDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDomains`: []Domain
    fmt.Fprintf(os.Stdout, "Response from `DomainsApi.GetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Only return domains whose &#x60;domain&#x60; field includes this string | 

### Return type

[**[]Domain**](Domain.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

