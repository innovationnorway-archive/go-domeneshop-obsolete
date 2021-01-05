# HTTPForward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The subdomain this forward applies to, without the domain part.  For instance, &#x60;www&#x60; in the context of the &#x60;example.com&#x60; domain signifies a forward for &#x60;www.example.com&#x60;.  | [optional] 
**Frame** | Pointer to **bool** | Whether to enable frame forwarding using an iframe embed. **NOT** recommended for a variety of reasons. | [optional] 
**Url** | Pointer to **string** | The URL to forward to. Must include scheme, e.g. &#x60;https://&#x60; or &#x60;ftp://&#x60;. | [optional] 

## Methods

### NewHTTPForward

`func NewHTTPForward() *HTTPForward`

NewHTTPForward instantiates a new HTTPForward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPForwardWithDefaults

`func NewHTTPForwardWithDefaults() *HTTPForward`

NewHTTPForwardWithDefaults instantiates a new HTTPForward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HTTPForward) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HTTPForward) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HTTPForward) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HTTPForward) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetFrame

`func (o *HTTPForward) GetFrame() bool`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *HTTPForward) GetFrameOk() (*bool, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *HTTPForward) SetFrame(v bool)`

SetFrame sets Frame field to given value.

### HasFrame

`func (o *HTTPForward) HasFrame() bool`

HasFrame returns a boolean if a field has been set.

### GetUrl

`func (o *HTTPForward) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HTTPForward) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HTTPForward) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HTTPForward) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


