# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ExpiryDate** | Pointer to **string** |  | [optional] 
**RegisteredDate** | Pointer to **string** |  | [optional] 
**Renew** | Pointer to **bool** |  | [optional] 
**Registrant** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Domain status | [optional] 
**Nameservers** | Pointer to **[]string** |  | [optional] 
**Services** | Pointer to [**DomainServices**](DomainServices.md) |  | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Domain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Domain) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Domain) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Domain) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Domain) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetRegisteredDate

`func (o *Domain) GetRegisteredDate() string`

GetRegisteredDate returns the RegisteredDate field if non-nil, zero value otherwise.

### GetRegisteredDateOk

`func (o *Domain) GetRegisteredDateOk() (*string, bool)`

GetRegisteredDateOk returns a tuple with the RegisteredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDate

`func (o *Domain) SetRegisteredDate(v string)`

SetRegisteredDate sets RegisteredDate field to given value.

### HasRegisteredDate

`func (o *Domain) HasRegisteredDate() bool`

HasRegisteredDate returns a boolean if a field has been set.

### GetRenew

`func (o *Domain) GetRenew() bool`

GetRenew returns the Renew field if non-nil, zero value otherwise.

### GetRenewOk

`func (o *Domain) GetRenewOk() (*bool, bool)`

GetRenewOk returns a tuple with the Renew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenew

`func (o *Domain) SetRenew(v bool)`

SetRenew sets Renew field to given value.

### HasRenew

`func (o *Domain) HasRenew() bool`

HasRenew returns a boolean if a field has been set.

### GetRegistrant

`func (o *Domain) GetRegistrant() string`

GetRegistrant returns the Registrant field if non-nil, zero value otherwise.

### GetRegistrantOk

`func (o *Domain) GetRegistrantOk() (*string, bool)`

GetRegistrantOk returns a tuple with the Registrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrant

`func (o *Domain) SetRegistrant(v string)`

SetRegistrant sets Registrant field to given value.

### HasRegistrant

`func (o *Domain) HasRegistrant() bool`

HasRegistrant returns a boolean if a field has been set.

### GetStatus

`func (o *Domain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Domain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Domain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Domain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNameservers

`func (o *Domain) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Domain) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Domain) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Domain) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetServices

`func (o *Domain) GetServices() DomainServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Domain) GetServicesOk() (*DomainServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Domain) SetServices(v DomainServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *Domain) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


