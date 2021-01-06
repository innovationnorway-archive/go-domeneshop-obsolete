# \DnsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRecord**](DnsApi.md#CreateRecord) | **Post** /domains/{domainId}/dns | Add DNS record
[**DeleteRecord**](DnsApi.md#DeleteRecord) | **Delete** /domains/{domainId}/dns/{recordId} | Delete DNS record by ID
[**GetRecord**](DnsApi.md#GetRecord) | **Get** /domains/{domainId}/dns/{recordId} | Find DNS record by ID
[**GetRecords**](DnsApi.md#GetRecords) | **Get** /domains/{domainId}/dns | List DNS records
[**ModifyRecord**](DnsApi.md#ModifyRecord) | **Put** /domains/{domainId}/dns/{recordId} | Update DNS record by ID



## CreateRecord

> InlineResponse201 CreateRecord(ctx, domainId).DNSRecord(dNSRecord).Execute()

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
    resp, r, err := api_client.DnsApi.CreateRecord(context.Background(), domainId).DNSRecord(dNSRecord).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.CreateRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRecord`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.CreateRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


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


## DeleteRecord

> DeleteRecord(ctx, domainId, recordId).Execute()

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
    resp, r, err := api_client.DnsApi.DeleteRecord(context.Background(), domainId, recordId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DeleteRecord``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


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


## GetRecord

> DNSRecord GetRecord(ctx, domainId, recordId).Execute()

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
    resp, r, err := api_client.DnsApi.GetRecord(context.Background(), domainId, recordId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.GetRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecord`: DNSRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.GetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 
**recordId** | **int32** | ID of DNS record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordRequest struct via the builder pattern


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


## GetRecords

> []DNSRecord GetRecords(ctx, domainId).Host(host).Type_(type_).Execute()

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
    resp, r, err := api_client.DnsApi.GetRecords(context.Background(), domainId).Host(host).Type_(type_).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.GetRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecords`: []DNSRecord
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.GetRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainId** | **int32** | ID of the domain | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordsRequest struct via the builder pattern


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


## ModifyRecord

> ModifyRecord(ctx, domainId, recordId).DNSRecord(dNSRecord).Execute()

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
    resp, r, err := api_client.DnsApi.ModifyRecord(context.Background(), domainId, recordId).DNSRecord(dNSRecord).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.ModifyRecord``: %v\n", err)
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

Other parameters are passed through a pointer to a apiModifyRecordRequest struct via the builder pattern


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

