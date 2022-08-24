// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of PublicIPAddresses_Spec. Use v1beta20201101.PublicIPAddresses_Spec instead
type PublicIPAddresses_SpecARM struct {
	ExtendedLocation *ExtendedLocationARM                `json:"extendedLocation,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             string                              `json:"name,omitempty"`
	Properties       *PublicIPAddressPropertiesFormatARM `json:"properties,omitempty"`
	Sku              *PublicIPAddressSkuARM              `json:"sku,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Zones            []string                            `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PublicIPAddresses_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (addresses PublicIPAddresses_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (addresses *PublicIPAddresses_SpecARM) GetName() string {
	return addresses.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/publicIPAddresses"
func (addresses *PublicIPAddresses_SpecARM) GetType() string {
	return "Microsoft.Network/publicIPAddresses"
}

// Deprecated version of PublicIPAddressPropertiesFormat. Use v1beta20201101.PublicIPAddressPropertiesFormat instead
type PublicIPAddressPropertiesFormatARM struct {
	DdosSettings             *DdosSettingsARM                                          `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettingsARM                            `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int                                                      `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                                                   `json:"ipAddress,omitempty"`
	IpTags                   []IpTagARM                                                `json:"ipTags,omitempty"`
	PublicIPAddressVersion   *PublicIPAddressPropertiesFormat_PublicIPAddressVersion   `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *PublicIPAddressPropertiesFormat_PublicIPAllocationMethod `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResourceARM                                           `json:"publicIPPrefix,omitempty"`
}

// Deprecated version of PublicIPAddressSku. Use v1beta20201101.PublicIPAddressSku instead
type PublicIPAddressSkuARM struct {
	Name *PublicIPAddressSku_Name `json:"name,omitempty"`
	Tier *PublicIPAddressSku_Tier `json:"tier,omitempty"`
}

// Deprecated version of DdosSettings. Use v1beta20201101.DdosSettings instead
type DdosSettingsARM struct {
	DdosCustomPolicy   *SubResourceARM                  `json:"ddosCustomPolicy,omitempty"`
	ProtectedIP        *bool                            `json:"protectedIP,omitempty"`
	ProtectionCoverage *DdosSettings_ProtectionCoverage `json:"protectionCoverage,omitempty"`
}

// Deprecated version of IpTag. Use v1beta20201101.IpTag instead
type IpTagARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of PublicIPAddressDnsSettings. Use v1beta20201101.PublicIPAddressDnsSettings instead
type PublicIPAddressDnsSettingsARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	Fqdn            *string `json:"fqdn,omitempty"`
	ReverseFqdn     *string `json:"reverseFqdn,omitempty"`
}

// Deprecated version of PublicIPAddressSku_Name. Use v1beta20201101.PublicIPAddressSku_Name instead
// +kubebuilder:validation:Enum={"Basic","Standard"}
type PublicIPAddressSku_Name string

const (
	PublicIPAddressSku_Name_Basic    = PublicIPAddressSku_Name("Basic")
	PublicIPAddressSku_Name_Standard = PublicIPAddressSku_Name("Standard")
)

// Deprecated version of PublicIPAddressSku_Tier. Use v1beta20201101.PublicIPAddressSku_Tier instead
// +kubebuilder:validation:Enum={"Global","Regional"}
type PublicIPAddressSku_Tier string

const (
	PublicIPAddressSku_Tier_Global   = PublicIPAddressSku_Tier("Global")
	PublicIPAddressSku_Tier_Regional = PublicIPAddressSku_Tier("Regional")
)