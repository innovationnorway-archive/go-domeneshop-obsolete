# DS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of DNS record | [readonly] 
**Host** | **string** | The host/subdomain the DNS record applies to | 
**Ttl** | Pointer to **int32** | TTL of DNS record in seconds. Must be a multiple of 60. | [optional] [default to 3600]
**Type** | **string** |  | 
**Data** | **string** | The value of the record. | 
**Tag** | **string** | DS record tag. | 
**Alg** | **string** | DS record algorithm. | 
**Digest** | **string** | DS record digest type. | 

## Methods

### NewDS

`func NewDS(id int32, host string, type_ string, data string, tag string, alg string, digest string, ) *DS`

NewDS instantiates a new DS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSWithDefaults

`func NewDSWithDefaults() *DS`

NewDSWithDefaults instantiates a new DS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DS) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DS) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DS) SetId(v int32)`

SetId sets Id field to given value.


### GetHost

`func (o *DS) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DS) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DS) SetHost(v string)`

SetHost sets Host field to given value.


### GetTtl

`func (o *DS) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DS) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DS) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DS) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *DS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DS) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *DS) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DS) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DS) SetData(v string)`

SetData sets Data field to given value.


### GetTag

`func (o *DS) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DS) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DS) SetTag(v string)`

SetTag sets Tag field to given value.


### GetAlg

`func (o *DS) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DS) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DS) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetDigest

`func (o *DS) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DS) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DS) SetDigest(v string)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


