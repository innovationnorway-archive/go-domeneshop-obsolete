# MXAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** | **string** | The target MX host. | 
**Priority** | **string** | MX record priority, also known as preference. Lower values are usually preferred first, but this is not guaranteed | 

## Methods

### NewMXAllOf

`func NewMXAllOf(type_ string, data string, priority string, ) *MXAllOf`

NewMXAllOf instantiates a new MXAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMXAllOfWithDefaults

`func NewMXAllOfWithDefaults() *MXAllOf`

NewMXAllOfWithDefaults instantiates a new MXAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MXAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MXAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MXAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *MXAllOf) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MXAllOf) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MXAllOf) SetData(v string)`

SetData sets Data field to given value.


### GetPriority

`func (o *MXAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MXAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MXAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


