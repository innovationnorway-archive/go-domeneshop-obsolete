# \DdnsApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DyndnsUpdateGet**](DdnsApi.md#DyndnsUpdateGet) | **Get** /dyndns/update | Update



## DyndnsUpdateGet

> DyndnsUpdateGet(ctx).Hostname(hostname).Myip(myip).Execute()

Update



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
    hostname := "www.example.com" // string | The fully qualified domain (FQDN) to be updated, without trailing dot.
    myip := "myip_example" // string | The new IPv4 or IPv6 address to set. If not provided, the IP of the client making the API request will be used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DdnsApi.DyndnsUpdateGet(context.Background()).Hostname(hostname).Myip(myip).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DdnsApi.DyndnsUpdateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDyndnsUpdateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostname** | **string** | The fully qualified domain (FQDN) to be updated, without trailing dot. | 
 **myip** | **string** | The new IPv4 or IPv6 address to set. If not provided, the IP of the client making the API request will be used. | 

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

