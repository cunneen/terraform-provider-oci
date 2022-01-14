// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v55/common"
	"strings"
)

// CalculateAuditVolumeAvailableDetails The details for calculating audit data volume on target.
type CalculateAuditVolumeAvailableDetails struct {

	// The trail locations for which the audit data volume has to be calculated.
	TrailLocations []string `mandatory:"true" json:"trailLocations"`

	// The date from which the audit trail must start collecting data in UTC, in the format defined by RFC3339. If not specified, this will default to the date based on the retention period.
	AuditCollectionStartTime *common.SDKTime `mandatory:"false" json:"auditCollectionStartTime"`
}

func (m CalculateAuditVolumeAvailableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CalculateAuditVolumeAvailableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}