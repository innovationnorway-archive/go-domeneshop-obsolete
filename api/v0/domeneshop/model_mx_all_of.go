/*
 * Domeneshop API Documentation
 *
 * # Overview  Domeneshop offers a simple, REST-based API, which currently supports the following features:  - List domains - List invoices - Create, read, update and delete DNS records for domains - Create, read, update and delete HTTP forwards (\"WWW forwarding\") for domains - Dynamic DNS (DDNS) update endpoints for use in consumer routers  More features are planned, including:  - Web hosting administration - Email address and email user/account administration  # Testing period  The API service is in version 0, which means it is possible that the interface will change rapidly during the testing period. For that reason, **the documentation on these pages may sometimes be outdated.**  Additionally, we make no guarantees about the stability of the API service during this testing period, and therefore ask customers to be careful with using the service for business critical purposes.  # Authentication  The Domeneshop API currently supports only one method of authentication, **HTTP Basic Auth**. More authentication methods may be added in the future.  To generate credentials, visit <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">this page</a> after logging in to the control panel on our website:  <a href=\"https://www.domeneshop.no/admin?view=api\" target=\"_blank\">https://www.domeneshop.no/admin?view=api</a>  # Libraries  Domeneshop maintains multiple API libraries to simplify using the API. Please note that these libraries have the same stability guarantees to the API while the API is in version 0.  The libraries may be found in our [Github repository](https://github.com/domeneshop/).  Domeneshop also maintains a plugin for [EFF's Certbot](https://certbot.eff.org/), which automates issuance and renewal of SSL-certificates on your own servers for domains that use Domeneshop's DNS service. This plugin is found in our Github repository [here](https://github.com/domeneshop/certbot-dns-domeneshop).  <SecurityDefinitions /> 
 *
 * API version: v0
 * Contact: kundeservice@domeneshop.no
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MXAllOf struct for MXAllOf
type MXAllOf struct {
	Type string `json:"type"`
	// The target MX host.
	Data string `json:"data"`
	// MX record priority, also known as preference. Lower values are usually preferred first, but this is not guaranteed
	Priority int32 `json:"priority"`
}

// NewMXAllOf instantiates a new MXAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMXAllOf(type_ string, data string, priority int32, ) *MXAllOf {
	this := MXAllOf{}
	this.Type = type_
	this.Data = data
	this.Priority = priority
	return &this
}

// NewMXAllOfWithDefaults instantiates a new MXAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMXAllOfWithDefaults() *MXAllOf {
	this := MXAllOf{}
	return &this
}

// GetType returns the Type field value
func (o *MXAllOf) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MXAllOf) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MXAllOf) SetType(v string) {
	o.Type = v
}

// GetData returns the Data field value
func (o *MXAllOf) GetData() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MXAllOf) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MXAllOf) SetData(v string) {
	o.Data = v
}

// GetPriority returns the Priority field value
func (o *MXAllOf) GetPriority() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *MXAllOf) GetPriorityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *MXAllOf) SetPriority(v int32) {
	o.Priority = v
}

func (o MXAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableMXAllOf struct {
	value *MXAllOf
	isSet bool
}

func (v NullableMXAllOf) Get() *MXAllOf {
	return v.value
}

func (v *NullableMXAllOf) Set(val *MXAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMXAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMXAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMXAllOf(val *MXAllOf) *NullableMXAllOf {
	return &NullableMXAllOf{value: val, isSet: true}
}

func (v NullableMXAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMXAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


