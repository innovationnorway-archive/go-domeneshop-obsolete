# TLSAAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** | **string** | The value of the record. | 
**Usage** | **int32** | TLSA record certificate usage. | 
**Selector** | **int32** | TLSA record selector. | 
**Dtype** | **int32** | TLSA record matching type. | 

## Methods

### NewTLSAAllOf

`func NewTLSAAllOf(type_ string, data string, usage int32, selector int32, dtype int32, ) *TLSAAllOf`

NewTLSAAllOf instantiates a new TLSAAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTLSAAllOfWithDefaults

`func NewTLSAAllOfWithDefaults() *TLSAAllOf`

NewTLSAAllOfWithDefaults instantiates a new TLSAAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TLSAAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TLSAAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TLSAAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *TLSAAllOf) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TLSAAllOf) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TLSAAllOf) SetData(v string)`

SetData sets Data field to given value.


### GetUsage

`func (o *TLSAAllOf) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TLSAAllOf) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TLSAAllOf) SetUsage(v int32)`

SetUsage sets Usage field to given value.


### GetSelector

`func (o *TLSAAllOf) GetSelector() int32`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *TLSAAllOf) GetSelectorOk() (*int32, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *TLSAAllOf) SetSelector(v int32)`

SetSelector sets Selector field to given value.


### GetDtype

`func (o *TLSAAllOf) GetDtype() int32`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *TLSAAllOf) GetDtypeOk() (*int32, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *TLSAAllOf) SetDtype(v int32)`

SetDtype sets Dtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


