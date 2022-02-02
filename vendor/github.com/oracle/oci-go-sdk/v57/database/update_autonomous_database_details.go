// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v57/common"
	"strings"
)

// UpdateAutonomousDatabaseDetails Details to update an Oracle Autonomous Database.
// **Notes**
// - To specify OCPU core count, you must use either `ocpuCount` or `cpuCoreCount`. You cannot use both parameters at the same time.
// - To specify a storage allocation, you must use  either `dataStorageSizeInGBs` or `dataStorageSizeInTBs`.
// - See the individual parameter discriptions for more information on the OCPU and storage value parameters.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateAutonomousDatabaseDetails struct {

	// The number of OCPU cores to be made available to the Autonomous Database.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	CpuCoreCount *int `mandatory:"false" json:"cpuCoreCount"`

	// The number of OCPU cores to be made available to the Autonomous Database. To provision less than 1 core, enter a fractional value in an increment of 0.1. To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available to the infrastructure shape. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. Likewise, you can provision 2 cores or 3 cores, but not 2.5 cores. The maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// The size, in terabytes, of the data volume that will be created and attached to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	DataStorageSizeInTBs *int `mandatory:"false" json:"dataStorageSizeInTBs"`

	// Applies to dedicated Exadata infrastructure only.
	// The size, in gigabytes, of the data volume that will be created and attached to the database. The maximum storage value depends on the system shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	DataStorageSizeInGBs *int `mandatory:"false" json:"dataStorageSizeInGBs"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique. The display name can only be updated for Autonomous Databases
	// using dedicated Exadata infrastructure.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing. It must be different from the last four passwords and it must not be a password used within the last 24 hours.
	AdminPassword *string `mandatory:"false" json:"adminPassword"`

	// New name for this Autonomous Database.
	// For databases using dedicated Exadata infrastructure, the name must begin with an alphabetic character, and can contain a maximum of eight alphanumeric characters. Special characters are not permitted.
	// For databases using shared Exadata infrastructure, the name must begin with an alphabetic character, and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	DbName *string `mandatory:"false" json:"dbName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	DbWorkload UpdateAutonomousDatabaseDetailsDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	LicenseModel UpdateAutonomousDatabaseDetailsLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	IsAccessControlEnabled *bool `mandatory:"false" json:"isAccessControlEnabled"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// `TRUE` if the Autonomous Database has Data Guard and Access Control enabled, and the Autonomous Database uses the primary's IP access control list (ACL) for standby.
	// `FALSE` if the Autonomous Database has Data Guard and Access Control enabled, and the Autonomous Database uses a different IP access control list (ACL) for standby compared to primary.
	ArePrimaryWhitelistedIpsUsed *bool `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	StandbyWhitelistedIps []string `mandatory:"false" json:"standbyWhitelistedIps"`

	// Indicates whether auto scaling is enabled for the Autonomous Database OCPU core count. Setting to `TRUE` enables auto scaling. Setting to `FALSE` disables auto scaling. The default value is true. Auto scaling is available for databases on shared Exadata infrastructure (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Indicates whether the Autonomous Database is a refreshable clone.
	IsRefreshableClone *bool `mandatory:"false" json:"isRefreshableClone"`

	// The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.
	RefreshableMode UpdateAutonomousDatabaseDetailsRefreshableModeEnum `mandatory:"false" json:"refreshableMode,omitempty"`

	// Indicates whether the Autonomous Database has a local (in-region) standby database. Not applicable when creating a cross-region Autonomous Data Guard associations, or to
	// Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	// To create a local standby, set to `TRUE`. To delete a local standby, set to `FALSE`. For more information on using Autonomous Data Guard on shared Exadata infrastructure (local and cross-region) , see About Standby Databases (https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-about.html#GUID-045AD017-8120-4BDC-AF58-7430FFE28D2B)
	// To enable cross-region Autonomous Data Guard on shared Exadata infrastructure, see CreateCrossRegionAutonomousDatabaseDataGuardDetails.
	IsLocalDataGuardEnabled *bool `mandatory:"false" json:"isLocalDataGuardEnabled"`

	// ** Deprecated. ** Indicates whether the Autonomous Database has a local (in-region) standby database. Not applicable when creating a cross-region Autonomous Data Guard associations, or to
	// Autonomous Databases using dedicated Exadata infrastructure or Exadata Cloud@Customer infrastructure.
	// To create a local standby, set to `TRUE`. To delete a local standby, set to `FALSE`. For more information on using Autonomous Data Guard on shared Exadata infrastructure (local and cross-region) , see About Standby Databases (https://docs.oracle.com/en/cloud/paas/autonomous-database/adbsa/autonomous-data-guard-about.html#GUID-045AD017-8120-4BDC-AF58-7430FFE28D2B)
	// To enable cross-region Autonomous Data Guard on shared Exadata infrastructure, see CreateCrossRegionAutonomousDatabaseDataGuardDetails.
	// To delete a cross-region standby database, provide the `peerDbId` for the standby database in a remote region, and set `isDataGuardEnabled` to `FALSE`.
	IsDataGuardEnabled *bool `mandatory:"false" json:"isDataGuardEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Data Guard standby database located in a different (remote) region from the source primary Autonomous Database.
	// To create or delete a local (in-region) standby, see the `isDataGuardEnabled` parameter.
	PeerDbId *string `mandatory:"false" json:"peerDbId"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.
	OpenMode UpdateAutonomousDatabaseDetailsOpenModeEnum `mandatory:"false" json:"openMode,omitempty"`

	// The Autonomous Database permission level. Restricted mode allows access only to admin users.
	PermissionLevel UpdateAutonomousDatabaseDetailsPermissionLevelEnum `mandatory:"false" json:"permissionLevel,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	PrivateEndpointLabel *string `mandatory:"false" json:"privateEndpointLabel"`

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The auto-refresh policy for the Autonomous Database refreshable clone. You can specify continuous refreshing or a custom refresh schedule.
	AutoRefreshPolicy UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum `mandatory:"false" json:"autoRefreshPolicy,omitempty"`

	// The frequency at which the data is refreshed for a refreshable clone after auto-refresh is enabled. The minimum is 1 minute. The maximum is 1 day. The date and time that auto-refresh is enabled is controlled by the `timeOfAutoRefreshStart` parameter.
	AutoRefreshFrequencyInSeconds *int `mandatory:"false" json:"autoRefreshFrequencyInSeconds"`

	// The amount of time, in seconds, that the data of the refreshable clone lags the data of the primary database at the point of refresh. The minimum is 1 minute. The maximum is 7 days. The lag time increases after refreshing until the next data refresh happens.
	AutoRefreshPointInSeconds *int `mandatory:"false" json:"autoRefreshPointInSeconds"`

	// The the date and time that auto-refreshing will begin for an Autonomous Database refreshable clone. This value controls only the start time for the first refresh operation. Subsequent (ongoing) refresh operations have start times controlled by the value of the `autoRefreshFrequencyInSeconds` parameter.
	TimeOfAutoRefreshStart *common.SDKTime `mandatory:"false" json:"timeOfAutoRefreshStart"`

	// Customer Contacts. Setting this to an empty list removes all customer contacts of an Oracle Autonomous Database.
	CustomerContacts []CustomerContact `mandatory:"false" json:"customerContacts"`

	// Indicates whether the Autonomous Database requires mTLS connections.
	IsMtlsConnectionRequired *bool `mandatory:"false" json:"isMtlsConnectionRequired"`

	// True if allow Oracle services to use the Service Gateway to connect to the Autonomous Database.
	IsOracleServiceGatewayAllowed *bool `mandatory:"false" json:"isOracleServiceGatewayAllowed"`

	// list of scheduled operations
	ScheduledOperations []ScheduledOperationDetails `mandatory:"false" json:"scheduledOperations"`

	// Indicates if auto scaling is enabled for the Autonomous Database storage. The default value is `FALSE`.
	IsAutoScalingForStorageEnabled *bool `mandatory:"false" json:"isAutoScalingForStorageEnabled"`

	// The number of Max OCPU cores to be made available to the autonomous database with auto scaling of cpu enabled.
	MaxCpuCoreCount *int `mandatory:"false" json:"maxCpuCoreCount"`

	// The Oracle Database Edition that applies to the Autonomous databases.
	DatabaseEdition AutonomousDatabaseSummaryDatabaseEditionEnum `mandatory:"false" json:"databaseEdition,omitempty"`
}

func (m UpdateAutonomousDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutonomousDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum[string(m.DbWorkload)]; !ok && m.DbWorkload != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbWorkload: %s. Supported values are: %s.", m.DbWorkload, strings.Join(GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues(), ",")))
	}
	if _, ok := mappingUpdateAutonomousDatabaseDetailsLicenseModelEnum[string(m.LicenseModel)]; !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum[string(m.RefreshableMode)]; !ok && m.RefreshableMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefreshableMode: %s. Supported values are: %s.", m.RefreshableMode, strings.Join(GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues(), ",")))
	}
	if _, ok := mappingUpdateAutonomousDatabaseDetailsOpenModeEnum[string(m.OpenMode)]; !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues(), ",")))
	}
	if _, ok := mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum[string(m.PermissionLevel)]; !ok && m.PermissionLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PermissionLevel: %s. Supported values are: %s.", m.PermissionLevel, strings.Join(GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues(), ",")))
	}
	if _, ok := mappingUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum[string(m.AutoRefreshPolicy)]; !ok && m.AutoRefreshPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoRefreshPolicy: %s. Supported values are: %s.", m.AutoRefreshPolicy, strings.Join(GetUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnumStringValues(), ",")))
	}
	if _, ok := mappingAutonomousDatabaseSummaryDatabaseEditionEnum[string(m.DatabaseEdition)]; !ok && m.DatabaseEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseEdition: %s. Supported values are: %s.", m.DatabaseEdition, strings.Join(GetAutonomousDatabaseSummaryDatabaseEditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAutonomousDatabaseDetailsDbWorkloadEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsDbWorkloadEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
const (
	UpdateAutonomousDatabaseDetailsDbWorkloadOltp UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "OLTP"
	UpdateAutonomousDatabaseDetailsDbWorkloadDw   UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "DW"
	UpdateAutonomousDatabaseDetailsDbWorkloadAjd  UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "AJD"
	UpdateAutonomousDatabaseDetailsDbWorkloadApex UpdateAutonomousDatabaseDetailsDbWorkloadEnum = "APEX"
)

var mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum = map[string]UpdateAutonomousDatabaseDetailsDbWorkloadEnum{
	"OLTP": UpdateAutonomousDatabaseDetailsDbWorkloadOltp,
	"DW":   UpdateAutonomousDatabaseDetailsDbWorkloadDw,
	"AJD":  UpdateAutonomousDatabaseDetailsDbWorkloadAjd,
	"APEX": UpdateAutonomousDatabaseDetailsDbWorkloadApex,
}

// GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
func GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumValues() []UpdateAutonomousDatabaseDetailsDbWorkloadEnum {
	values := make([]UpdateAutonomousDatabaseDetailsDbWorkloadEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsDbWorkloadEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsDbWorkloadEnum
func GetUpdateAutonomousDatabaseDetailsDbWorkloadEnumStringValues() []string {
	return []string{
		"OLTP",
		"DW",
		"AJD",
		"APEX",
	}
}

// UpdateAutonomousDatabaseDetailsLicenseModelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsLicenseModelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
const (
	UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded     UpdateAutonomousDatabaseDetailsLicenseModelEnum = "LICENSE_INCLUDED"
	UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense UpdateAutonomousDatabaseDetailsLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingUpdateAutonomousDatabaseDetailsLicenseModelEnum = map[string]UpdateAutonomousDatabaseDetailsLicenseModelEnum{
	"LICENSE_INCLUDED":       UpdateAutonomousDatabaseDetailsLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": UpdateAutonomousDatabaseDetailsLicenseModelBringYourOwnLicense,
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumValues() []UpdateAutonomousDatabaseDetailsLicenseModelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsLicenseModelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsLicenseModelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsLicenseModelEnum
func GetUpdateAutonomousDatabaseDetailsLicenseModelEnumStringValues() []string {
	return []string{
		"LICENSE_INCLUDED",
		"BRING_YOUR_OWN_LICENSE",
	}
}

// UpdateAutonomousDatabaseDetailsRefreshableModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsRefreshableModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
const (
	UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "AUTOMATIC"
	UpdateAutonomousDatabaseDetailsRefreshableModeManual    UpdateAutonomousDatabaseDetailsRefreshableModeEnum = "MANUAL"
)

var mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum = map[string]UpdateAutonomousDatabaseDetailsRefreshableModeEnum{
	"AUTOMATIC": UpdateAutonomousDatabaseDetailsRefreshableModeAutomatic,
	"MANUAL":    UpdateAutonomousDatabaseDetailsRefreshableModeManual,
}

// GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
func GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumValues() []UpdateAutonomousDatabaseDetailsRefreshableModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsRefreshableModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsRefreshableModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsRefreshableModeEnum
func GetUpdateAutonomousDatabaseDetailsRefreshableModeEnumStringValues() []string {
	return []string{
		"AUTOMATIC",
		"MANUAL",
	}
}

// UpdateAutonomousDatabaseDetailsOpenModeEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsOpenModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsOpenModeEnum
const (
	UpdateAutonomousDatabaseDetailsOpenModeOnly  UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_ONLY"
	UpdateAutonomousDatabaseDetailsOpenModeWrite UpdateAutonomousDatabaseDetailsOpenModeEnum = "READ_WRITE"
)

var mappingUpdateAutonomousDatabaseDetailsOpenModeEnum = map[string]UpdateAutonomousDatabaseDetailsOpenModeEnum{
	"READ_ONLY":  UpdateAutonomousDatabaseDetailsOpenModeOnly,
	"READ_WRITE": UpdateAutonomousDatabaseDetailsOpenModeWrite,
}

// GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsOpenModeEnum
func GetUpdateAutonomousDatabaseDetailsOpenModeEnumValues() []UpdateAutonomousDatabaseDetailsOpenModeEnum {
	values := make([]UpdateAutonomousDatabaseDetailsOpenModeEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsOpenModeEnum
func GetUpdateAutonomousDatabaseDetailsOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
	}
}

// UpdateAutonomousDatabaseDetailsPermissionLevelEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsPermissionLevelEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
const (
	UpdateAutonomousDatabaseDetailsPermissionLevelRestricted   UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "RESTRICTED"
	UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted UpdateAutonomousDatabaseDetailsPermissionLevelEnum = "UNRESTRICTED"
)

var mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum = map[string]UpdateAutonomousDatabaseDetailsPermissionLevelEnum{
	"RESTRICTED":   UpdateAutonomousDatabaseDetailsPermissionLevelRestricted,
	"UNRESTRICTED": UpdateAutonomousDatabaseDetailsPermissionLevelUnrestricted,
}

// GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
func GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumValues() []UpdateAutonomousDatabaseDetailsPermissionLevelEnum {
	values := make([]UpdateAutonomousDatabaseDetailsPermissionLevelEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsPermissionLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsPermissionLevelEnum
func GetUpdateAutonomousDatabaseDetailsPermissionLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"UNRESTRICTED",
	}
}

// UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum Enum with underlying type: string
type UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum string

// Set of constants representing the allowable values for UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum
const (
	UpdateAutonomousDatabaseDetailsAutoRefreshPolicyContinuous UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum = "CONTINUOUS"
	UpdateAutonomousDatabaseDetailsAutoRefreshPolicyCustom     UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum = "CUSTOM"
)

var mappingUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum = map[string]UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum{
	"CONTINUOUS": UpdateAutonomousDatabaseDetailsAutoRefreshPolicyContinuous,
	"CUSTOM":     UpdateAutonomousDatabaseDetailsAutoRefreshPolicyCustom,
}

// GetUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnumValues Enumerates the set of values for UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum
func GetUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnumValues() []UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum {
	values := make([]UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum, 0)
	for _, v := range mappingUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnumStringValues Enumerates the set of values in String for UpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnum
func GetUpdateAutonomousDatabaseDetailsAutoRefreshPolicyEnumStringValues() []string {
	return []string{
		"CONTINUOUS",
		"CUSTOM",
	}
}