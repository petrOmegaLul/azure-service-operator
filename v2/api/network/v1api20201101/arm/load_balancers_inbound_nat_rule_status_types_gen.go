// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type LoadBalancersInboundNatRule_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name
	// can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of load balancer inbound nat rule.
	Properties *InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded `json:"properties,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

// Properties of the inbound NAT rule.
type InboundNatRulePropertiesFormat_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded struct {
	// BackendIPConfiguration: A reference to a private IP address defined on a network interface of a VM. Traffic sent to the
	// frontend port of each of the frontend IP configurations is forwarded to the backend IP.
	BackendIPConfiguration *NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded `json:"backendIPConfiguration,omitempty"`

	// BackendPort: The port used for the internal endpoint. Acceptable values range from 1 to 65535.
	BackendPort *int `json:"backendPort,omitempty"`

	// EnableFloatingIP: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL
	// AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
	// This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty"`

	// EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	// element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	// FrontendIPConfiguration: A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResource_STATUS `json:"frontendIPConfiguration,omitempty"`

	// FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
	// Acceptable values range from 1 to 65534.
	FrontendPort *int `json:"frontendPort,omitempty"`

	// IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
	// default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// Protocol: The reference to the transport protocol used by the load balancing rule.
	Protocol *TransportProtocol_STATUS `json:"protocol,omitempty"`

	// ProvisioningState: The provisioning state of the inbound NAT rule resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// IPConfiguration in a network interface.
type NetworkInterfaceIPConfiguration_STATUS_LoadBalancers_InboundNatRule_SubResourceEmbedded struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Reference to another subresource.
type SubResource_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// The transport protocol for the endpoint.
type TransportProtocol_STATUS string

const (
	TransportProtocol_STATUS_All = TransportProtocol_STATUS("All")
	TransportProtocol_STATUS_Tcp = TransportProtocol_STATUS("Tcp")
	TransportProtocol_STATUS_Udp = TransportProtocol_STATUS("Udp")
)

// Mapping from string to TransportProtocol_STATUS
var transportProtocol_STATUS_Values = map[string]TransportProtocol_STATUS{
	"all": TransportProtocol_STATUS_All,
	"tcp": TransportProtocol_STATUS_Tcp,
	"udp": TransportProtocol_STATUS_Udp,
}
