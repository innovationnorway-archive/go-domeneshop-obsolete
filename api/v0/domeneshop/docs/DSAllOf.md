# DSAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** | **string** | The value of the record. | 
**Tag** | **int32** | DS record tag. | 
**Alg** | **int32** | DS record algorithm. | 
**Digest** | **int32** | DS record digest type. | 

## Methods

### NewDSAllOf

`func NewDSAllOf(type_ string, data string, tag int32, alg int32, digest int32, ) *DSAllOf`

NewDSAllOf instantiates a new DSAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSAllOfWithDefaults

`func NewDSAllOfWithDefaults() *DSAllOf`

NewDSAllOfWithDefaults instantiates a new DSAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DSAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DSAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DSAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *DSAllOf) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DSAllOf) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DSAllOf) SetData(v string)`

SetData sets Data field to given value.


### GetTag

`func (o *DSAllOf) GetTag() int32`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DSAllOf) GetTagOk() (*int32, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DSAllOf) SetTag(v int32)`

SetTag sets Tag field to given value.


### GetAlg

`func (o *DSAllOf) GetAlg() int32`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DSAllOf) GetAlgOk() (*int32, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DSAllOf) SetAlg(v int32)`

SetAlg sets Alg field to given value.


### GetDigest

`func (o *DSAllOf) GetDigest() int32`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DSAllOf) GetDigestOk() (*int32, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DSAllOf) SetDigest(v int32)`

SetDigest sets Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


