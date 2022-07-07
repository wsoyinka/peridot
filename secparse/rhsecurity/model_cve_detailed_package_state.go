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

// CVEDetailedPackageState struct for CVEDetailedPackageState
type CVEDetailedPackageState struct {
	ProductName string `json:"product_name"`
	FixState string `json:"fix_state"`
	PackageName string `json:"package_name"`
	Cpe string `json:"cpe"`
}

// NewCVEDetailedPackageState instantiates a new CVEDetailedPackageState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCVEDetailedPackageState(productName string, fixState string, packageName string, cpe string) *CVEDetailedPackageState {
	this := CVEDetailedPackageState{}
	this.ProductName = productName
	this.FixState = fixState
	this.PackageName = packageName
	this.Cpe = cpe
	return &this
}

// NewCVEDetailedPackageStateWithDefaults instantiates a new CVEDetailedPackageState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCVEDetailedPackageStateWithDefaults() *CVEDetailedPackageState {
	this := CVEDetailedPackageState{}
	return &this
}

// GetProductName returns the ProductName field value
func (o *CVEDetailedPackageState) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedPackageState) GetProductNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *CVEDetailedPackageState) SetProductName(v string) {
	o.ProductName = v
}

// GetFixState returns the FixState field value
func (o *CVEDetailedPackageState) GetFixState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FixState
}

// GetFixStateOk returns a tuple with the FixState field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedPackageState) GetFixStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FixState, true
}

// SetFixState sets field value
func (o *CVEDetailedPackageState) SetFixState(v string) {
	o.FixState = v
}

// GetPackageName returns the PackageName field value
func (o *CVEDetailedPackageState) GetPackageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedPackageState) GetPackageNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PackageName, true
}

// SetPackageName sets field value
func (o *CVEDetailedPackageState) SetPackageName(v string) {
	o.PackageName = v
}

// GetCpe returns the Cpe field value
func (o *CVEDetailedPackageState) GetCpe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpe
}

// GetCpeOk returns a tuple with the Cpe field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedPackageState) GetCpeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cpe, true
}

// SetCpe sets field value
func (o *CVEDetailedPackageState) SetCpe(v string) {
	o.Cpe = v
}

func (o CVEDetailedPackageState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product_name"] = o.ProductName
	}
	if true {
		toSerialize["fix_state"] = o.FixState
	}
	if true {
		toSerialize["package_name"] = o.PackageName
	}
	if true {
		toSerialize["cpe"] = o.Cpe
	}
	return json.Marshal(toSerialize)
}

type NullableCVEDetailedPackageState struct {
	value *CVEDetailedPackageState
	isSet bool
}

func (v NullableCVEDetailedPackageState) Get() *CVEDetailedPackageState {
	return v.value
}

func (v *NullableCVEDetailedPackageState) Set(val *CVEDetailedPackageState) {
	v.value = val
	v.isSet = true
}

func (v NullableCVEDetailedPackageState) IsSet() bool {
	return v.isSet
}

func (v *NullableCVEDetailedPackageState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCVEDetailedPackageState(val *CVEDetailedPackageState) *NullableCVEDetailedPackageState {
	return &NullableCVEDetailedPackageState{value: val, isSet: true}
}

func (v NullableCVEDetailedPackageState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCVEDetailedPackageState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


