// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrPlanGroup Details of a plan group in a DR Plan.
type DrPlanGroup struct {

	// The unique id of this group. Must not be modified by user.
	// Example: `sgid1.group..&lt;unique_id&gt;`
	Id *string `mandatory:"true" json:"id"`

	// The plan group type.
	Type DrPlanGroupTypeEnum `mandatory:"true" json:"type"`

	// The display name of this DR Plan Group.
	// Example: `DATABASE_SWITCHOVER`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The list of steps in this plan group.
	Steps []DrPlanStep `mandatory:"true" json:"steps"`
}

func (m DrPlanGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrPlanGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDrPlanGroupTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDrPlanGroupTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
