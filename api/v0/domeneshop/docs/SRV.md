# SRV

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of DNS record | [readonly] 
**Host** | **string** | The host/subdomain the DNS record applies to | 
**Ttl** | Pointer to **int32** | TTL of DNS record in seconds. Must be a multiple of 60. | [optional] [default to 3600]
**Type** | **string** |  | 
**Data** | **string** | The target hostname | 
**Priority** | **string** | SRV record priority, also known as preference. Lower values are usually preferred first | 
**Weight** | **string** | SRV record weight. Relevant if multiple records have same preference | 
**Port** | **string** | SRV record port. The port where the service is found. | 

## Methods

### NewSRV

`func NewSRV(id int32, host string, type_ string, data string, priority string, weight string, port string, ) *SRV`

NewSRV instantiates a new SRV object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVWithDefaults

`func NewSRVWithDefaults() *SRV`

NewSRVWithDefaults instantiates a new SRV object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SRV) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SRV) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SRV) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *SRV) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SRV) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SRV) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *SRV) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *SRV) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *SRV) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *SRV) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *SRV) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SRV) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SRV) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *SRV) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SRV) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SRV) SetData(v string)`

SetData sets Data field to given value.


### GetPriority

`func (o *SRV) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SRV) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SRV) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetWeight

`func (o *SRV) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SRV) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SRV) SetWeight(v string)`

SetWeight sets Weight field to given value.


### GetPort

`func (o *SRV) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SRV) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SRV) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


