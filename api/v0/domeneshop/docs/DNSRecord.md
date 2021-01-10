# DNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of DNS record | [readonly] 
**Host** | **string** | The host/subdomain the DNS record applies to | 
**Ttl** | Pointer to **int32** | TTL of DNS record in seconds. Must be a multiple of 60. | [optional] [default to 3600]
**Type** | **string** |  | 
**Data** | **string** | The value of the record. | 
**Priority** | **string** | SRV record priority, also known as preference. Lower values are usually preferred first | 
**Weight** | **string** | SRV record weight. Relevant if multiple records have same preference | 
**Port** | **string** | SRV record port. The port where the service is found. | 
**Flags** | **string** | CAA record flags. | 
**Tag** | **string** | DS record tag. | 
**Alg** | **string** | DS record algorithm. | 
**Digest** | **string** | DS record digest type. | 
**Usage** | **string** | TLSA record certificate usage. | 
**Selector** | **string** | TLSA record selector. | 
**Dtype** | **string** | TLSA record matching type. | 

## Methods

### NewDNSRecord

`func NewDNSRecord(id int32, host string, type_ string, data string, priority string, weight string, port string, flags string, tag string, alg string, digest string, usage string, selector string, dtype string, ) *DNSRecord`

NewDNSRecord instantiates a new DNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordWithDefaults

`func NewDNSRecordWithDefaults() *DNSRecord`

NewDNSRecordWithDefaults instantiates a new DNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DNSRecord) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DNSRecord) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DNSRecord) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *DNSRecord) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DNSRecord) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DNSRecord) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *DNSRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *DNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRecord) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *DNSRecord) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DNSRecord) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DNSRecord) SetData(v string)`

SetData sets Data field to given value.


### GetPriority

`func (o *DNSRecord) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecord) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecord) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetWeight

`func (o *DNSRecord) GetWeight() string`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecord) GetWeightOk() (*string, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecord) SetWeight(v string)`

SetWeight sets Weight field to given value.


### GetPort

`func (o *DNSRecord) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecord) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecord) SetPort(v string)`

SetPort sets Port field to given value.


### GetFlags

`func (o *DNSRecord) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *DNSRecord) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *DNSRecord) SetFlags(v string)`

SetFlags sets Flags field to given value.


### GetTag

`func (o *DNSRecord) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DNSRecord) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DNSRecord) SetTag(v string)`

SetTag sets Tag field to given value.


### GetAlg

`func (o *DNSRecord) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DNSRecord) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DNSRecord) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetDigest

`func (o *DNSRecord) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DNSRecord) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DNSRecord) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetUsage

`func (o *DNSRecord) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *DNSRecord) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *DNSRecord) SetUsage(v string)`

SetUsage sets Usage field to given value.


### GetSelector

`func (o *DNSRecord) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *DNSRecord) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *DNSRecord) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetDtype

`func (o *DNSRecord) GetDtype() string`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *DNSRecord) GetDtypeOk() (*string, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *DNSRecord) SetDtype(v string)`

SetDtype sets Dtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


