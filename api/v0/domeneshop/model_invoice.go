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

// Invoice struct for Invoice
type Invoice struct {
	// Invoice ID/number
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Amount *int32 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
	// The invoice due date. Only available for type `invoice`.
	DueDate *string `json:"due_date,omitempty"`
	// The date when the invoice was issued.
	IssuedDate *string `json:"issued_date,omitempty"`
	// The payment date. Only available if the invoice has status `paid`.
	PaidDate *string `json:"paid_date,omitempty"`
	// `settled` is only applicable to credit notes. These are usually created if  domains have been 
	Status *string `json:"status,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice() *Invoice {
	this := Invoice{}
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invoice) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invoice) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Invoice) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Invoice) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Invoice) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Invoice) SetType(v string) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Invoice) GetAmount() int32 {
	if o == nil || o.Amount == nil {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAmountOk() (*int32, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Invoice) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *Invoice) SetAmount(v int32) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Invoice) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Invoice) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Invoice) SetCurrency(v string) {
	o.Currency = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Invoice) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Invoice) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Invoice) SetDueDate(v string) {
	o.DueDate = &v
}

// GetIssuedDate returns the IssuedDate field value if set, zero value otherwise.
func (o *Invoice) GetIssuedDate() string {
	if o == nil || o.IssuedDate == nil {
		var ret string
		return ret
	}
	return *o.IssuedDate
}

// GetIssuedDateOk returns a tuple with the IssuedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIssuedDateOk() (*string, bool) {
	if o == nil || o.IssuedDate == nil {
		return nil, false
	}
	return o.IssuedDate, true
}

// HasIssuedDate returns a boolean if a field has been set.
func (o *Invoice) HasIssuedDate() bool {
	if o != nil && o.IssuedDate != nil {
		return true
	}

	return false
}

// SetIssuedDate gets a reference to the given string and assigns it to the IssuedDate field.
func (o *Invoice) SetIssuedDate(v string) {
	o.IssuedDate = &v
}

// GetPaidDate returns the PaidDate field value if set, zero value otherwise.
func (o *Invoice) GetPaidDate() string {
	if o == nil || o.PaidDate == nil {
		var ret string
		return ret
	}
	return *o.PaidDate
}

// GetPaidDateOk returns a tuple with the PaidDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPaidDateOk() (*string, bool) {
	if o == nil || o.PaidDate == nil {
		return nil, false
	}
	return o.PaidDate, true
}

// HasPaidDate returns a boolean if a field has been set.
func (o *Invoice) HasPaidDate() bool {
	if o != nil && o.PaidDate != nil {
		return true
	}

	return false
}

// SetPaidDate gets a reference to the given string and assigns it to the PaidDate field.
func (o *Invoice) SetPaidDate(v string) {
	o.PaidDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Invoice) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Invoice) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Invoice) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Invoice) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Invoice) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Invoice) SetUrl(v string) {
	o.Url = &v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if o.IssuedDate != nil {
		toSerialize["issued_date"] = o.IssuedDate
	}
	if o.PaidDate != nil {
		toSerialize["paid_date"] = o.PaidDate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


