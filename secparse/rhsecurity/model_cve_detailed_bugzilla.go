/*
 * Red Hat Security Data API
 *
 * Unofficial OpenAPI definitions for Red Hat Security Data API
 *
 * API version: 1.0
 * Contact: mustafa@ctrliq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rhsecurity

import (
	"encoding/json"
)

// CVEDetailedBugzilla struct for CVEDetailedBugzilla
type CVEDetailedBugzilla struct {
	Description string `json:"description"`
	Id string `json:"id"`
	Url string `json:"url"`
}

// NewCVEDetailedBugzilla instantiates a new CVEDetailedBugzilla object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCVEDetailedBugzilla(description string, id string, url string) *CVEDetailedBugzilla {
	this := CVEDetailedBugzilla{}
	this.Description = description
	this.Id = id
	this.Url = url
	return &this
}

// NewCVEDetailedBugzillaWithDefaults instantiates a new CVEDetailedBugzilla object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCVEDetailedBugzillaWithDefaults() *CVEDetailedBugzilla {
	this := CVEDetailedBugzilla{}
	return &this
}

// GetDescription returns the Description field value
func (o *CVEDetailedBugzilla) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedBugzilla) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CVEDetailedBugzilla) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *CVEDetailedBugzilla) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedBugzilla) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CVEDetailedBugzilla) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CVEDetailedBugzilla) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedBugzilla) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CVEDetailedBugzilla) SetUrl(v string) {
	o.Url = v
}

func (o CVEDetailedBugzilla) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableCVEDetailedBugzilla struct {
	value *CVEDetailedBugzilla
	isSet bool
}

func (v NullableCVEDetailedBugzilla) Get() *CVEDetailedBugzilla {
	return v.value
}

func (v *NullableCVEDetailedBugzilla) Set(val *CVEDetailedBugzilla) {
	v.value = val
	v.isSet = true
}

func (v NullableCVEDetailedBugzilla) IsSet() bool {
	return v.isSet
}

func (v *NullableCVEDetailedBugzilla) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCVEDetailedBugzilla(val *CVEDetailedBugzilla) *NullableCVEDetailedBugzilla {
	return &NullableCVEDetailedBugzilla{value: val, isSet: true}
}

func (v NullableCVEDetailedBugzilla) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCVEDetailedBugzilla) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


