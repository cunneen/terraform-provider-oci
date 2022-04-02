// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// GoldenGateConnection Represents the metadata of a GoldenGate Connection.
type GoldenGateConnection struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the connection being
	// referenced.
	Id *string `mandatory:"true" json:"id"`

	// An object's Display Name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the resource was created. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle
	// Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more
	// information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Describes the object's current state in detail. For example, it can be used to provide
	// actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer vault being
	// referenced.
	// If provided, this will reference a vault which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to manage secrets contained
	// within this vault.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer "Master" key being
	// referenced.
	// If provided, this will reference a key which the customer will be required to ensure
	// the policies are established to permit the GoldenGate Service to utilize this key to
	// manage secrets.
	KeyId *string `mandatory:"false" json:"keyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet being referenced.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// A Private Endpoint IP Address created in the customer's subnet.  A customer
	// database can expect network traffic initiated by GGS from this IP address and send network traffic
	// to this IP address, typically in response to requests from GGS (OGG).  The customer may utilize
	// this IP address in Security Lists or Network Security Groups (NSG) as needed.
	RcePrivateIp *string `mandatory:"false" json:"rcePrivateIp"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"false" json:"deploymentId"`

	// The name or address of a host.
	Host *string `mandatory:"false" json:"host"`

	// The port of an endpoint usually specified for a connection.
	Port *int `mandatory:"false" json:"port"`

	// The private IP address of the connection's endpoint in the customer's VCN, typically a
	// database endpoint or a big data endpoint (e.g. Kafka bootstrap server).
	// In case the privateIp is provided the subnetId must also be provided.
	// In case the privateIp (and the subnetId) is not provided it is assumed the datasource is publicly accessible.
	// In case the connection is accessible only privately, the lack of privateIp will result in not being able to access the connection.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The GoldenGate technology type.
	TechnologyType GoldenGateConnectionTechnologyTypeEnum `mandatory:"true" json:"technologyType"`

	// Possible lifecycle states for connection.
	LifecycleState ConnectionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m GoldenGateConnection) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m GoldenGateConnection) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m GoldenGateConnection) GetDescription() *string {
	return m.Description
}

//GetCompartmentId returns CompartmentId
func (m GoldenGateConnection) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m GoldenGateConnection) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GoldenGateConnection) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m GoldenGateConnection) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetLifecycleState returns LifecycleState
func (m GoldenGateConnection) GetLifecycleState() ConnectionLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m GoldenGateConnection) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetTimeCreated returns TimeCreated
func (m GoldenGateConnection) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m GoldenGateConnection) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetVaultId returns VaultId
func (m GoldenGateConnection) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m GoldenGateConnection) GetKeyId() *string {
	return m.KeyId
}

//GetSubnetId returns SubnetId
func (m GoldenGateConnection) GetSubnetId() *string {
	return m.SubnetId
}

//GetRcePrivateIp returns RcePrivateIp
func (m GoldenGateConnection) GetRcePrivateIp() *string {
	return m.RcePrivateIp
}

func (m GoldenGateConnection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GoldenGateConnection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGoldenGateConnectionTechnologyTypeEnum(string(m.TechnologyType)); !ok && m.TechnologyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TechnologyType: %s. Supported values are: %s.", m.TechnologyType, strings.Join(GetGoldenGateConnectionTechnologyTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingConnectionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetConnectionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GoldenGateConnection) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGoldenGateConnection GoldenGateConnection
	s := struct {
		DiscriminatorParam string `json:"connectionType"`
		MarshalTypeGoldenGateConnection
	}{
		"GOLDENGATE",
		(MarshalTypeGoldenGateConnection)(m),
	}

	return json.Marshal(&s)
}

// GoldenGateConnectionTechnologyTypeEnum Enum with underlying type: string
type GoldenGateConnectionTechnologyTypeEnum string

// Set of constants representing the allowable values for GoldenGateConnectionTechnologyTypeEnum
const (
	GoldenGateConnectionTechnologyTypeGoldengate GoldenGateConnectionTechnologyTypeEnum = "GOLDENGATE"
)

var mappingGoldenGateConnectionTechnologyTypeEnum = map[string]GoldenGateConnectionTechnologyTypeEnum{
	"GOLDENGATE": GoldenGateConnectionTechnologyTypeGoldengate,
}

var mappingGoldenGateConnectionTechnologyTypeEnumLowerCase = map[string]GoldenGateConnectionTechnologyTypeEnum{
	"goldengate": GoldenGateConnectionTechnologyTypeGoldengate,
}

// GetGoldenGateConnectionTechnologyTypeEnumValues Enumerates the set of values for GoldenGateConnectionTechnologyTypeEnum
func GetGoldenGateConnectionTechnologyTypeEnumValues() []GoldenGateConnectionTechnologyTypeEnum {
	values := make([]GoldenGateConnectionTechnologyTypeEnum, 0)
	for _, v := range mappingGoldenGateConnectionTechnologyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGoldenGateConnectionTechnologyTypeEnumStringValues Enumerates the set of values in String for GoldenGateConnectionTechnologyTypeEnum
func GetGoldenGateConnectionTechnologyTypeEnumStringValues() []string {
	return []string{
		"GOLDENGATE",
	}
}

// GetMappingGoldenGateConnectionTechnologyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGoldenGateConnectionTechnologyTypeEnum(val string) (GoldenGateConnectionTechnologyTypeEnum, bool) {
	enum, ok := mappingGoldenGateConnectionTechnologyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
