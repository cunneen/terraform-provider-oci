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
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// DatabaseUpgradeHistoryEntry The Database service supports the upgrade history of databases.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DatabaseUpgradeHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database upgrade history.
	Id *string `mandatory:"true" json:"id"`

	// The database upgrade action.
	Action DatabaseUpgradeHistoryEntryActionEnum `mandatory:"true" json:"action"`

	// Status of database upgrade history SUCCEEDED|IN_PROGRESS|FAILED.
	LifecycleState DatabaseUpgradeHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the database upgrade started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The source of the Oracle Database software to be used for the upgrade.
	//  - Use `DB_HOME` to specify an existing Database Home to upgrade the database. The database is moved to the target Database Home and makes use of the Oracle Database software version of the target Database Home.
	//  - Use `DB_VERSION` to specify a generally-available Oracle Database software version to upgrade the database.
	//  - Use `DB_SOFTWARE_IMAGE` to specify a database software image (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databasesoftwareimage.htm) to upgrade the database.
	Source DatabaseUpgradeHistoryEntrySourceEnum `mandatory:"false" json:"source,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	TargetDBVersion *string `mandatory:"false" json:"targetDBVersion"`

	// the database software image used for upgrading database.
	TargetDatabaseSoftwareImageId *string `mandatory:"false" json:"targetDatabaseSoftwareImageId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	TargetDbHomeId *string `mandatory:"false" json:"targetDbHomeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	SourceDbHomeId *string `mandatory:"false" json:"sourceDbHomeId"`

	// The date and time when the database upgrade ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional upgrade options supported by DBUA(Database Upgrade Assistant).
	// Example: "-upgradeTimezone false -keepEvents"
	Options *string `mandatory:"false" json:"options"`
}

func (m DatabaseUpgradeHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseUpgradeHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingDatabaseUpgradeHistoryEntryActionEnum[string(m.Action)]; !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDatabaseUpgradeHistoryEntryActionEnumStringValues(), ",")))
	}
	if _, ok := mappingDatabaseUpgradeHistoryEntryLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabaseUpgradeHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := mappingDatabaseUpgradeHistoryEntrySourceEnum[string(m.Source)]; !ok && m.Source != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Source: %s. Supported values are: %s.", m.Source, strings.Join(GetDatabaseUpgradeHistoryEntrySourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseUpgradeHistoryEntryActionEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntryActionEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntryActionEnum
const (
	DatabaseUpgradeHistoryEntryActionPrecheck DatabaseUpgradeHistoryEntryActionEnum = "PRECHECK"
	DatabaseUpgradeHistoryEntryActionUpgrade  DatabaseUpgradeHistoryEntryActionEnum = "UPGRADE"
	DatabaseUpgradeHistoryEntryActionRollback DatabaseUpgradeHistoryEntryActionEnum = "ROLLBACK"
)

var mappingDatabaseUpgradeHistoryEntryActionEnum = map[string]DatabaseUpgradeHistoryEntryActionEnum{
	"PRECHECK": DatabaseUpgradeHistoryEntryActionPrecheck,
	"UPGRADE":  DatabaseUpgradeHistoryEntryActionUpgrade,
	"ROLLBACK": DatabaseUpgradeHistoryEntryActionRollback,
}

// GetDatabaseUpgradeHistoryEntryActionEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntryActionEnum
func GetDatabaseUpgradeHistoryEntryActionEnumValues() []DatabaseUpgradeHistoryEntryActionEnum {
	values := make([]DatabaseUpgradeHistoryEntryActionEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseUpgradeHistoryEntryActionEnumStringValues Enumerates the set of values in String for DatabaseUpgradeHistoryEntryActionEnum
func GetDatabaseUpgradeHistoryEntryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"UPGRADE",
		"ROLLBACK",
	}
}

// DatabaseUpgradeHistoryEntrySourceEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntrySourceEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntrySourceEnum
const (
	DatabaseUpgradeHistoryEntrySourceHome          DatabaseUpgradeHistoryEntrySourceEnum = "DB_HOME"
	DatabaseUpgradeHistoryEntrySourceVersion       DatabaseUpgradeHistoryEntrySourceEnum = "DB_VERSION"
	DatabaseUpgradeHistoryEntrySourceSoftwareImage DatabaseUpgradeHistoryEntrySourceEnum = "DB_SOFTWARE_IMAGE"
)

var mappingDatabaseUpgradeHistoryEntrySourceEnum = map[string]DatabaseUpgradeHistoryEntrySourceEnum{
	"DB_HOME":           DatabaseUpgradeHistoryEntrySourceHome,
	"DB_VERSION":        DatabaseUpgradeHistoryEntrySourceVersion,
	"DB_SOFTWARE_IMAGE": DatabaseUpgradeHistoryEntrySourceSoftwareImage,
}

// GetDatabaseUpgradeHistoryEntrySourceEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntrySourceEnum
func GetDatabaseUpgradeHistoryEntrySourceEnumValues() []DatabaseUpgradeHistoryEntrySourceEnum {
	values := make([]DatabaseUpgradeHistoryEntrySourceEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntrySourceEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseUpgradeHistoryEntrySourceEnumStringValues Enumerates the set of values in String for DatabaseUpgradeHistoryEntrySourceEnum
func GetDatabaseUpgradeHistoryEntrySourceEnumStringValues() []string {
	return []string{
		"DB_HOME",
		"DB_VERSION",
		"DB_SOFTWARE_IMAGE",
	}
}

// DatabaseUpgradeHistoryEntryLifecycleStateEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntryLifecycleStateEnum
const (
	DatabaseUpgradeHistoryEntryLifecycleStateSucceeded  DatabaseUpgradeHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	DatabaseUpgradeHistoryEntryLifecycleStateFailed     DatabaseUpgradeHistoryEntryLifecycleStateEnum = "FAILED"
	DatabaseUpgradeHistoryEntryLifecycleStateInProgress DatabaseUpgradeHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
)

var mappingDatabaseUpgradeHistoryEntryLifecycleStateEnum = map[string]DatabaseUpgradeHistoryEntryLifecycleStateEnum{
	"SUCCEEDED":   DatabaseUpgradeHistoryEntryLifecycleStateSucceeded,
	"FAILED":      DatabaseUpgradeHistoryEntryLifecycleStateFailed,
	"IN_PROGRESS": DatabaseUpgradeHistoryEntryLifecycleStateInProgress,
}

// GetDatabaseUpgradeHistoryEntryLifecycleStateEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntryLifecycleStateEnum
func GetDatabaseUpgradeHistoryEntryLifecycleStateEnumValues() []DatabaseUpgradeHistoryEntryLifecycleStateEnum {
	values := make([]DatabaseUpgradeHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseUpgradeHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for DatabaseUpgradeHistoryEntryLifecycleStateEnum
func GetDatabaseUpgradeHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}