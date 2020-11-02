/*
 * AppGate SDP Controller REST API
 *
 * # About   This specification documents the REST API calls for the AppGate SDP Controller.    Please refer to the Integration chapter in the manual or contact AppGate support with any questions about   this functionality. # Getting Started   Requirements for API scripting:   - Access to the peer interface (default port 444) of a Controller appliance.     (https://sdphelp.appgate.com/adminguide/appliances-configure.html?anchor=peer)   - An API user with relevant permissions.     (https://sdphelp.appgate.com/adminguide/administrative-roles-configure.html)   - In order to use the simple login API, Admin MFA must be disabled or the API user must be excluded.     (https://sdphelp.appgate.com/adminguide/mfa-for-admins.html) # Base path   HTTPS requests must be sent to the Peer Interface hostname and port, with **_/admin** path.    For example: **https://appgate.company.com:444/admin**    All requests must have the **Accept** header as:    **application/vnd.appgate.peer-v13+json** # API Conventions   API conventions are  important to understand and follow strictly.    - While updating objects (via PUT), entire object must be sent with all fields.     - For example, if in order to add a remedy method to the condition below:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": []       }       ```     - send the entire object with updated and non-updated fields:       ```       {         \"id\": \"12699e27-b584-464a-81ee-5b4784b6d425\",         \"name\": \"Test\",         \"notes\": \"Making a point\",         \"tags\": [\"test\", \"tag\"],         \"expression\": \"return true;\",         \"remedyMethods\": [{\"type\": \"DisplayMessage\", \"message\": \"test message\"}]       }       ```    - In case Controller returns an error (non-2xx HTTP status code), response body is JSON.     The \"message\" field contains information about the error.     HTTP 422 \"Unprocessable Entity\" has extra `errors` field to list all the issues with specific fields.    - Empty string (\"\") is considered a different value than \"null\" or field being omitted from JSON.     Omitting the field is recommend if no value is intended.     Empty string (\"\") will be almost always rejected as invalid value.    - There are common pattern between many objects:     - **Configuration Objects**: There are many objects with common fields, namely \"id\", \"name\", \"notes\", \"created\"       and \"updated\". These entities are listed, queried, created, updated and deleted in a similar fashion.     - **Distinguished Name**: Users and Devices are identified with what is called Distinguished Names, as used in        LDAP. The distinguished format that identifies a device and a user combination is        \"CN=\\<Device ID\\>,CN=\\<username\\>,OU=\\<Identity Provider Name\\>\". Some objects have the        \"userDistinguishedName\" field, which does not include the CN for Device ID.        This identifies a user on every device.
 *
 * API version: API version 13
 * Contact: appgatesdp.support@appgate.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
)

// IpPool struct for IpPool
type IpPool struct {
	BaseEntity
	// Whether the IP pool is for v4 or v6.
	IpVersion6 *bool `json:"ipVersion6,omitempty"`
	// List of (non-conflicting) IP address ranges to allocate IPs in order.
	Ranges *[]IpPoolAllOfRanges `json:"ranges,omitempty"`
	// Number of days Allocated IPs will be reserved for device&users before they are reclaimable by others.
	LeaseTimeDays *int32 `json:"leaseTimeDays,omitempty"`
	// The total size of the IP Pool.
	Total *int32 `json:"total,omitempty"`
	// Number of IPs in the pool are currently in use by device&users.
	CurrentlyUsed *int32 `json:"currentlyUsed,omitempty"`
	// Number of IPs in the pool are not currently in use but reserved for device&users according to the \"leaseTimeDays\" setting.
	Reserved *int32 `json:"reserved,omitempty"`
}

// NewIpPool instantiates a new IpPool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpPool() *IpPool {
	this := IpPool{}
	var ipVersion6 bool = false
	this.IpVersion6 = &ipVersion6
	var leaseTimeDays int32 = 30
	this.LeaseTimeDays = &leaseTimeDays
	return &this
}

// NewIpPoolWithDefaults instantiates a new IpPool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpPoolWithDefaults() *IpPool {
	this := IpPool{}
	var ipVersion6 bool = false
	this.IpVersion6 = &ipVersion6
	var leaseTimeDays int32 = 30
	this.LeaseTimeDays = &leaseTimeDays
	return &this
}

// GetIpVersion6 returns the IpVersion6 field value if set, zero value otherwise.
func (o *IpPool) GetIpVersion6() bool {
	if o == nil || o.IpVersion6 == nil {
		var ret bool
		return ret
	}
	return *o.IpVersion6
}

// GetIpVersion6Ok returns a tuple with the IpVersion6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetIpVersion6Ok() (*bool, bool) {
	if o == nil || o.IpVersion6 == nil {
		return nil, false
	}
	return o.IpVersion6, true
}

// HasIpVersion6 returns a boolean if a field has been set.
func (o *IpPool) HasIpVersion6() bool {
	if o != nil && o.IpVersion6 != nil {
		return true
	}

	return false
}

// SetIpVersion6 gets a reference to the given bool and assigns it to the IpVersion6 field.
func (o *IpPool) SetIpVersion6(v bool) {
	o.IpVersion6 = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *IpPool) GetRanges() []IpPoolAllOfRanges {
	if o == nil || o.Ranges == nil {
		var ret []IpPoolAllOfRanges
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetRangesOk() (*[]IpPoolAllOfRanges, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *IpPool) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []IpPoolAllOfRanges and assigns it to the Ranges field.
func (o *IpPool) SetRanges(v []IpPoolAllOfRanges) {
	o.Ranges = &v
}

// GetLeaseTimeDays returns the LeaseTimeDays field value if set, zero value otherwise.
func (o *IpPool) GetLeaseTimeDays() int32 {
	if o == nil || o.LeaseTimeDays == nil {
		var ret int32
		return ret
	}
	return *o.LeaseTimeDays
}

// GetLeaseTimeDaysOk returns a tuple with the LeaseTimeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetLeaseTimeDaysOk() (*int32, bool) {
	if o == nil || o.LeaseTimeDays == nil {
		return nil, false
	}
	return o.LeaseTimeDays, true
}

// HasLeaseTimeDays returns a boolean if a field has been set.
func (o *IpPool) HasLeaseTimeDays() bool {
	if o != nil && o.LeaseTimeDays != nil {
		return true
	}

	return false
}

// SetLeaseTimeDays gets a reference to the given int32 and assigns it to the LeaseTimeDays field.
func (o *IpPool) SetLeaseTimeDays(v int32) {
	o.LeaseTimeDays = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *IpPool) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *IpPool) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *IpPool) SetTotal(v int32) {
	o.Total = &v
}

// GetCurrentlyUsed returns the CurrentlyUsed field value if set, zero value otherwise.
func (o *IpPool) GetCurrentlyUsed() int32 {
	if o == nil || o.CurrentlyUsed == nil {
		var ret int32
		return ret
	}
	return *o.CurrentlyUsed
}

// GetCurrentlyUsedOk returns a tuple with the CurrentlyUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetCurrentlyUsedOk() (*int32, bool) {
	if o == nil || o.CurrentlyUsed == nil {
		return nil, false
	}
	return o.CurrentlyUsed, true
}

// HasCurrentlyUsed returns a boolean if a field has been set.
func (o *IpPool) HasCurrentlyUsed() bool {
	if o != nil && o.CurrentlyUsed != nil {
		return true
	}

	return false
}

// SetCurrentlyUsed gets a reference to the given int32 and assigns it to the CurrentlyUsed field.
func (o *IpPool) SetCurrentlyUsed(v int32) {
	o.CurrentlyUsed = &v
}

// GetReserved returns the Reserved field value if set, zero value otherwise.
func (o *IpPool) GetReserved() int32 {
	if o == nil || o.Reserved == nil {
		var ret int32
		return ret
	}
	return *o.Reserved
}

// GetReservedOk returns a tuple with the Reserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpPool) GetReservedOk() (*int32, bool) {
	if o == nil || o.Reserved == nil {
		return nil, false
	}
	return o.Reserved, true
}

// HasReserved returns a boolean if a field has been set.
func (o *IpPool) HasReserved() bool {
	if o != nil && o.Reserved != nil {
		return true
	}

	return false
}

// SetReserved gets a reference to the given int32 and assigns it to the Reserved field.
func (o *IpPool) SetReserved(v int32) {
	o.Reserved = &v
}

func (o IpPool) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseEntity, errBaseEntity := json.Marshal(o.BaseEntity)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	errBaseEntity = json.Unmarshal([]byte(serializedBaseEntity), &toSerialize)
	if errBaseEntity != nil {
		return []byte{}, errBaseEntity
	}
	if o.IpVersion6 != nil {
		toSerialize["ipVersion6"] = o.IpVersion6
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	if o.LeaseTimeDays != nil {
		toSerialize["leaseTimeDays"] = o.LeaseTimeDays
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.CurrentlyUsed != nil {
		toSerialize["currentlyUsed"] = o.CurrentlyUsed
	}
	if o.Reserved != nil {
		toSerialize["reserved"] = o.Reserved
	}
	return json.Marshal(toSerialize)
}

type NullableIpPool struct {
	value *IpPool
	isSet bool
}

func (v NullableIpPool) Get() *IpPool {
	return v.value
}

func (v *NullableIpPool) Set(val *IpPool) {
	v.value = val
	v.isSet = true
}

func (v NullableIpPool) IsSet() bool {
	return v.isSet
}

func (v *NullableIpPool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpPool(val *IpPool) *NullableIpPool {
	return &NullableIpPool{value: val, isSet: true}
}

func (v NullableIpPool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpPool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
