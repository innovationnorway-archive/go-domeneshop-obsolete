/*
 * Domeneshop API Documentation
 *
 * # Overview  Domeneshop offers a simple, REST-based API, which currently supports the following features:  - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including:  - Web hosting administration - Email address and email user/account administration  # Testing period  The API service is in version 0, which means it is possible that the interface will change rapidly during the testing period. For that reason, **the documentation on these pages may sometimes be outdated.**  Additionally, we make no guarantees about the stability of the API service during this testing period, and therefore ask customers to be careful with using the service for business critical purposes.  # Authentication  The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website:  <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a>  # Libraries  Domeneshop maintains multiple API libraries to simplify using the API. Please note that these libraries have the same stability guarantees to the API while the API is in version 0.  The libraries may be found in our [Github repository](https://github.com/domeneshop/).  Domeneshop also maintains a plugin for [EFF's Certbot](https://certbot.eff.org/), which automates issuance and renewal of SSL-certificates on your own servers for domains that use Domeneshop's DNS service. This plugin is found in our Github repository [here](https://github.com/domeneshop/certbot-dns-domeneshop).  <SecurityDefinitions /> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package domeneshop

import (
	"encoding/json"
)

// CAA struct for CAA
type CAA struct {
	// ID of DNS record
	Id int32 `json:"id"`
	// The host/subdomain the DNS record applies to
	Host string `json:"host"`
	// TTL of DNS record in seconds. Must be a multiple of 60.
	Ttl *int32 `json:"ttl,omitempty"`
	Type string `json:"type"`
	// The value of the record.
	Data string `json:"data"`
	// CAA record flags.
	Flags string `json:"flags"`
	// CAA record tag.
	Tag string `json:"tag"`
}

// NewCAA instantiates a new CAA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCAA(id int32, host string, type_ string, data string, flags string, tag string, ) *CAA {
	this := CAA{}
	this.Id = id
	this.Host = host
	var ttl int32 = 3600
	this.Ttl = &ttl
	this.Type = type_
	this.Data = data
	this.Flags = flags
	this.Tag = tag
	return &this
}

// NewCAAWithDefaults instantiates a new CAA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCAAWithDefaults() *CAA {
	this := CAA{}
	var ttl int32 = 3600
	this.Ttl = &ttl
	return &this
}

// GetId returns the Id field value
func (o *CAA) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CAA) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CAA) SetId(v int32) {
	o.Id = v
}

// GetHost returns the Host field value
func (o *CAA) GetHost() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CAA) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CAA) SetHost(v string) {
	o.Host = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CAA) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CAA) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CAA) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *CAA) SetTtl(v int32) {
	o.Ttl = &v
}

// GetType returns the Type field value
func (o *CAA) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CAA) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CAA) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *CAA) GetData() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CAA) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CAA) SetData(v string) {
	o.Data = v
}

// GetFlags returns the Flags field value
func (o *CAA) GetFlags() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *CAA) GetFlagsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *CAA) SetFlags(v string) {
	o.Flags = v
}

// GetTag returns the Tag field value
func (o *CAA) GetTag() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CAA) GetTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *CAA) SetTag(v string) {
	o.Tag = v
}

func (o CAA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["flags"] = o.Flags
	}
	if true {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableCAA struct {
	value *CAA
	isSet bool
}

func (v NullableCAA) Get() *CAA {
	return v.value
}

func (v *NullableCAA) Set(val *CAA) {
	v.value = val
	v.isSet = true
}

func (v NullableCAA) IsSet() bool {
	return v.isSet
}

func (v *NullableCAA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCAA(val *CAA) *NullableCAA {
	return &NullableCAA{value: val, isSet: true}
}

func (v NullableCAA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCAA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


