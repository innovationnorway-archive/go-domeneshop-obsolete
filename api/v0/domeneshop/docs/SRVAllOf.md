# SRVAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Data** | **string** | The target hostname | 
**Priority** | **int32** | SRV record priority, also known as preference. Lower values are usually preferred first | 
**Weight** | **int32** | SRV record weight. Relevant if multiple records have same preference | 
**Port** | **int32** | SRV record port. The port where the service is found. | 

## Methods

### NewSRVAllOf

`func NewSRVAllOf(type_ string, data string, priority int32, weight int32, port int32, ) *SRVAllOf`

NewSRVAllOf instantiates a new SRVAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVAllOfWithDefaults

`func NewSRVAllOfWithDefaults() *SRVAllOf`

NewSRVAllOfWithDefaults instantiates a new SRVAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SRVAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SRVAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SRVAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *SRVAllOf) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SRVAllOf) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SRVAllOf) SetData(v string)`

SetData sets Data field to given value.


### GetPriority

`func (o *SRVAllOf) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SRVAllOf) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SRVAllOf) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetWeight

`func (o *SRVAllOf) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SRVAllOf) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SRVAllOf) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetPort

`func (o *SRVAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SRVAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SRVAllOf) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


