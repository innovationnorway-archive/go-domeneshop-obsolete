# TLSA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of DNS record | [readonly] 
**Host** | **string** | The host/subdomain the DNS record applies to | 
**Ttl** | Pointer to **int32** | TTL of DNS record in seconds. Must be a multiple of 60. | [optional] [default to 3600]
**Type** | **string** |  | 
**Data** | **string** | The value of the record. | 
**Usage** | **int32** | TLSA record certificate usage. | 
**Selector** | **int32** | TLSA record selector. | 
**Dtype** | **int32** | TLSA record matching type. | 

## Methods

### NewTLSA

`func NewTLSA(id int32, host string, type_ string, data string, usage int32, selector int32, dtype int32, ) *TLSA`

NewTLSA instantiates a new TLSA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSAWithDefaults

`func NewTLSAWithDefaults() *TLSA`

NewTLSAWithDefaults instantiates a new TLSA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TLSA) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TLSA) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TLSA) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *TLSA) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *TLSA) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *TLSA) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *TLSA) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *TLSA) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *TLSA) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *TLSA) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *TLSA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSA) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *TLSA) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSA) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSA) SetData(v string)`

SetData sets Data field to given value.


### GetUsage

`func (o *TLSA) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TLSA) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TLSA) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetSelector

`func (o *TLSA) GetSelector() int32`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TLSA) GetSelectorOk() (*int32, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TLSA) SetSelector(v int32)`

SetSelector sets Selector field to given value.


### GetDtype

`func (o *TLSA) GetDtype() int32`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *TLSA) GetDtypeOk() (*int32, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *TLSA) SetDtype(v int32)`

SetDtype sets Dtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


