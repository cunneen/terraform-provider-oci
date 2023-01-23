// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_job", StackMonitoringDiscoveryJobDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_job_logs", StackMonitoringDiscoveryJobLogsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_discovery_jobs", StackMonitoringDiscoveryJobsDataSource())
	tfresource.RegisterDatasource("oci_stack_monitoring_monitored_resource", StackMonitoringMonitoredResourceDataSource())
}
