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

// Domain struct for Domain
type Domain struct {
	Id *int32 `json:"id,omitempty"`
	Domain *string `json:"domain,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	RegisteredDate *string `json:"registered_date,omitempty"`
	Renew *bool `json:"renew,omitempty"`
	Registrant *string `json:"registrant,omitempty"`
	// Domain status
	Status *string `json:"status,omitempty"`
	Nameservers *[]string `json:"nameservers,omitempty"`
	Services *DomainServices `json:"services,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Domain) SetId(v int32) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *Domain) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *Domain) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *Domain) SetDomain(v string) {
	o.Domain = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Domain) GetExpiryDate() string {
	if o == nil || o.ExpiryDate == nil {
		var ret string
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetExpiryDateOk() (*string, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Domain) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given string and assigns it to the ExpiryDate field.
func (o *Domain) SetExpiryDate(v string) {
	o.ExpiryDate = &v
}

// GetRegisteredDate returns the RegisteredDate field value if set, zero value otherwise.
func (o *Domain) GetRegisteredDate() string {
	if o == nil || o.RegisteredDate == nil {
		var ret string
		return ret
	}
	return *o.RegisteredDate
}

// GetRegisteredDateOk returns a tuple with the RegisteredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRegisteredDateOk() (*string, bool) {
	if o == nil || o.RegisteredDate == nil {
		return nil, false
	}
	return o.RegisteredDate, true
}

// HasRegisteredDate returns a boolean if a field has been set.
func (o *Domain) HasRegisteredDate() bool {
	if o != nil && o.RegisteredDate != nil {
		return true
	}

	return false
}

// SetRegisteredDate gets a reference to the given string and assigns it to the RegisteredDate field.
func (o *Domain) SetRegisteredDate(v string) {
	o.RegisteredDate = &v
}

// GetRenew returns the Renew field value if set, zero value otherwise.
func (o *Domain) GetRenew() bool {
	if o == nil || o.Renew == nil {
		var ret bool
		return ret
	}
	return *o.Renew
}

// GetRenewOk returns a tuple with the Renew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRenewOk() (*bool, bool) {
	if o == nil || o.Renew == nil {
		return nil, false
	}
	return o.Renew, true
}

// HasRenew returns a boolean if a field has been set.
func (o *Domain) HasRenew() bool {
	if o != nil && o.Renew != nil {
		return true
	}

	return false
}

// SetRenew gets a reference to the given bool and assigns it to the Renew field.
func (o *Domain) SetRenew(v bool) {
	o.Renew = &v
}

// GetRegistrant returns the Registrant field value if set, zero value otherwise.
func (o *Domain) GetRegistrant() string {
	if o == nil || o.Registrant == nil {
		var ret string
		return ret
	}
	return *o.Registrant
}

// GetRegistrantOk returns a tuple with the Registrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRegistrantOk() (*string, bool) {
	if o == nil || o.Registrant == nil {
		return nil, false
	}
	return o.Registrant, true
}

// HasRegistrant returns a boolean if a field has been set.
func (o *Domain) HasRegistrant() bool {
	if o != nil && o.Registrant != nil {
		return true
	}

	return false
}

// SetRegistrant gets a reference to the given string and assigns it to the Registrant field.
func (o *Domain) SetRegistrant(v string) {
	o.Registrant = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Domain) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Domain) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Domain) SetStatus(v string) {
	o.Status = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *Domain) GetNameservers() []string {
	if o == nil || o.Nameservers == nil {
		var ret []string
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameserversOk() (*[]string, bool) {
	if o == nil || o.Nameservers == nil {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *Domain) HasNameservers() bool {
	if o != nil && o.Nameservers != nil {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *Domain) SetNameservers(v []string) {
	o.Nameservers = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Domain) GetServices() DomainServices {
	if o == nil || o.Services == nil {
		var ret DomainServices
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetServicesOk() (*DomainServices, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Domain) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given DomainServices and assigns it to the Services field.
func (o *Domain) SetServices(v DomainServices) {
	o.Services = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.ExpiryDate != nil {
		toSerialize["expiry_date"] = o.ExpiryDate
	}
	if o.RegisteredDate != nil {
		toSerialize["registered_date"] = o.RegisteredDate
	}
	if o.Renew != nil {
		toSerialize["renew"] = o.Renew
	}
	if o.Registrant != nil {
		toSerialize["registrant"] = o.Registrant
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Nameservers != nil {
		toSerialize["nameservers"] = o.Nameservers
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

