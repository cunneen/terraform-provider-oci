// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Anomaly Detection API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UnivariateSignalMetricDetails This contains signal metrics data.
type UnivariateSignalMetricDetails struct {

	// The minimum value for signal.
	Min *float64 `mandatory:"false" json:"min"`

	// The maximum value for signal.
	Max *float64 `mandatory:"false" json:"max"`

	// The standard deviation of the signal value
	Std *float64 `mandatory:"false" json:"std"`

	// The total number of rows for the signal.
	RowCount *int `mandatory:"false" json:"rowCount"`

	// The total number of null values for the signal.
	NullCount *int `mandatory:"false" json:"nullCount"`
}

func (m UnivariateSignalMetricDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnivariateSignalMetricDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
