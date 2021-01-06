# \InvoicesApi

All URIs are relative to *https://api.domeneshop.no/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /invoices/{invoiceId} | Find invoice by invoice number
[**GetInvoices**](InvoicesApi.md#GetInvoices) | **Get** /invoices | List invoices



## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Find invoice by invoice number

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
    invoiceId := int32(56) // int32 | An invoice number

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GetInvoice(context.Background(), invoiceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** | An invoice number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> []Invoice GetInvoices(ctx).Status(status).Execute()

List invoices



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
    status := "paid" // string | Only return invoices with this status (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InvoicesApi.GetInvoices(context.Background()).Status(status).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoices`: []Invoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Only return invoices with this status | 

### Return type

[**[]Invoice**](Invoice.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

