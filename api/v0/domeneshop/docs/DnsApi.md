# \DnsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsDomainIdDnsPost**](DnsApi.md#DomainsDomainIdDnsPost) | **Post** /domains/{domainId}/dns | Add DNS record
[**DomainsDomainIdDnsRecordIdDelete**](DnsApi.md#DomainsDomainIdDnsRecordIdDelete) | **Delete** /domains/{domainId}/dns/{recordId} | Delete DNS record by ID
[**DomainsDomainIdDnsRecordIdPut**](DnsApi.md#DomainsDomainIdDnsRecordIdPut) | **Put** /domains/{domainId}/dns/{recordId} | Update DNS record by ID
[**GetDnsRecords**](DnsApi.md#GetDnsRecords) | **Get** /domains/{domainId}/dns | List DNS records
[**GetRecordById**](DnsApi.md#GetRecordById) | **Get** /domains/{domainId}/dns/{recordId} | Find DNS record by ID



## DomainsDomainIdDnsPost

> InlineResponse201 DomainsDomainIdDnsPost(ctx, domainId).DNSRecord(dNSRecord).Execute()

Add DNS record

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
    dNSRecord := openapiclient.DNSRecord{A: openapiclient.NewA(int32(1), "@", "Type_example", "Data_example")} // DNSRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DomainsDomainIdDnsPost(context.Background(), domainId).DNSRecord(dNSRecord).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DomainsDomainIdDnsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsDomainIdDnsPost`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DomainsDomainIdDnsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdDnsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dNSRecord** | [**DNSRecord**](DNSRecord.md) |  | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsDomainIdDnsRecordIdDelete

> DomainsDomainIdDnsRecordIdDelete(ctx, domainId, recordId).Execute()

Delete DNS record by ID

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
    recordId := int32(56) // int32 | ID of DNS record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DomainsDomainIdDnsRecordIdDelete(context.Background(), domainId, recordId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DomainsDomainIdDnsRecordIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**recordId** | **int32** | ID of DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdDnsRecordIdDeleteRequest struct via the builder pattern


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


## DomainsDomainIdDnsRecordIdPut

> DomainsDomainIdDnsRecordIdPut(ctx, domainId, recordId).DNSRecord(dNSRecord).Execute()

Update DNS record by ID

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
    recordId := int32(56) // int32 | ID of DNS record
    dNSRecord := openapiclient.DNSRecord{A: openapiclient.NewA(int32(1), "@", "Type_example", "Data_example")} // DNSRecord |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DomainsDomainIdDnsRecordIdPut(context.Background(), domainId, recordId).DNSRecord(dNSRecord).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DomainsDomainIdDnsRecordIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**recordId** | **int32** | ID of DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsDomainIdDnsRecordIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dNSRecord** | [**DNSRecord**](DNSRecord.md) |  | 

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


## GetDnsRecords

> []DNSRecord GetDnsRecords(ctx, domainId).Host(host).Type_(type_).Execute()

List DNS records

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
    host := "www" // string | Only return records whose `host` field matches this string (optional)
    type_ := "A" // string | Only return records whose `type` field matches this string (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.GetDnsRecords(context.Background(), domainId).Host(host).Type_(type_).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.GetDnsRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDnsRecords`: []DNSRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.GetDnsRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **host** | **string** | Only return records whose &#x60;host&#x60; field matches this string | 
 **type_** | **string** | Only return records whose &#x60;type&#x60; field matches this string | 

### Return type

[**[]DNSRecord**](DNSRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecordById

> DNSRecord GetRecordById(ctx, domainId, recordId).Execute()

Find DNS record by ID

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
    recordId := int32(56) // int32 | ID of DNS record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.GetRecordById(context.Background(), domainId, recordId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.GetRecordById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecordById`: DNSRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.GetRecordById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**recordId** | **int32** | ID of DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DNSRecord**](DNSRecord.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

