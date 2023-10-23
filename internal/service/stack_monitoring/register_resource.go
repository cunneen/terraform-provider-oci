// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_stack_monitoring_config", StackMonitoringConfigResource())
	tfresource.RegisterResource("oci_stack_monitoring_discovery_job", StackMonitoringDiscoveryJobResource())
	tfresource.RegisterResource("oci_stack_monitoring_metric_extension", StackMonitoringMetricExtensionResource())
	tfresource.RegisterResource("oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management", StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource())
	tfresource.RegisterResource("oci_stack_monitoring_metric_extensions_test_management", StackMonitoringMetricExtensionsTestManagementResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resource", StackMonitoringMonitoredResourceResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resource_task", StackMonitoringMonitoredResourceTaskResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resource_type", StackMonitoringMonitoredResourceTypeResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_associate_monitored_resource", StackMonitoringMonitoredResourcesAssociateMonitoredResourceResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_list_member", StackMonitoringMonitoredResourcesListMemberResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_search", StackMonitoringMonitoredResourcesSearchResource())
	tfresource.RegisterResource("oci_stack_monitoring_monitored_resources_search_association", StackMonitoringMonitoredResourcesSearchAssociationResource())
}
