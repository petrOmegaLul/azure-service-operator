// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=web.azure.com,resources=serverfarms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=web.azure.com,resources={serverfarms/status,serverfarms/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220301.ServerFarm
// Generator information:
// - Generated from: /web/resource-manager/Microsoft.Web/stable/2022-03-01/AppServicePlans.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}
type ServerFarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerFarm_Spec   `json:"spec,omitempty"`
	Status            ServerFarm_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServerFarm{}

// GetConditions returns the conditions of the resource
func (farm *ServerFarm) GetConditions() conditions.Conditions {
	return farm.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (farm *ServerFarm) SetConditions(conditions conditions.Conditions) {
	farm.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ServerFarm{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (farm *ServerFarm) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if farm.Spec.OperatorSpec == nil {
		return nil
	}
	return farm.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServerFarm{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (farm *ServerFarm) SecretDestinationExpressions() []*core.DestinationExpression {
	if farm.Spec.OperatorSpec == nil {
		return nil
	}
	return farm.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ServerFarm{}

// AzureName returns the Azure name of the resource
func (farm *ServerFarm) AzureName() string {
	return farm.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (farm ServerFarm) GetAPIVersion() string {
	return "2022-03-01"
}

// GetResourceScope returns the scope of the resource
func (farm *ServerFarm) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (farm *ServerFarm) GetSpec() genruntime.ConvertibleSpec {
	return &farm.Spec
}

// GetStatus returns the status of this resource
func (farm *ServerFarm) GetStatus() genruntime.ConvertibleStatus {
	return &farm.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (farm *ServerFarm) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Web/serverfarms"
func (farm *ServerFarm) GetType() string {
	return "Microsoft.Web/serverfarms"
}

// NewEmptyStatus returns a new empty (blank) status
func (farm *ServerFarm) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ServerFarm_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (farm *ServerFarm) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(farm.Spec)
	return farm.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (farm *ServerFarm) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServerFarm_STATUS); ok {
		farm.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServerFarm_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	farm.Status = st
	return nil
}

// Hub marks that this ServerFarm is the hub type for conversion
func (farm *ServerFarm) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (farm *ServerFarm) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: farm.Spec.OriginalVersion,
		Kind:    "ServerFarm",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220301.ServerFarm
// Generator information:
// - Generated from: /web/resource-manager/Microsoft.Web/stable/2022-03-01/AppServicePlans.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/serverfarms/{name}
type ServerFarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerFarm `json:"items"`
}

// Storage version of v1api20220301.APIVersion
// +kubebuilder:validation:Enum={"2022-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-03-01")

// Storage version of v1api20220301.ServerFarm_Spec
type ServerFarm_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                 string                     `json:"azureName,omitempty"`
	ElasticScaleEnabled       *bool                      `json:"elasticScaleEnabled,omitempty"`
	ExtendedLocation          *ExtendedLocation          `json:"extendedLocation,omitempty"`
	FreeOfferExpirationTime   *string                    `json:"freeOfferExpirationTime,omitempty"`
	HostingEnvironmentProfile *HostingEnvironmentProfile `json:"hostingEnvironmentProfile,omitempty"`
	HyperV                    *bool                      `json:"hyperV,omitempty"`
	IsSpot                    *bool                      `json:"isSpot,omitempty"`
	IsXenon                   *bool                      `json:"isXenon,omitempty"`
	Kind                      *string                    `json:"kind,omitempty"`
	KubeEnvironmentProfile    *KubeEnvironmentProfile    `json:"kubeEnvironmentProfile,omitempty"`
	Location                  *string                    `json:"location,omitempty"`
	MaximumElasticWorkerCount *int                       `json:"maximumElasticWorkerCount,omitempty"`
	OperatorSpec              *ServerFarmOperatorSpec    `json:"operatorSpec,omitempty"`
	OriginalVersion           string                     `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner              *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PerSiteScaling     *bool                              `json:"perSiteScaling,omitempty"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Reserved           *bool                              `json:"reserved,omitempty"`
	Sku                *SkuDescription                    `json:"sku,omitempty"`
	SpotExpirationTime *string                            `json:"spotExpirationTime,omitempty"`
	Tags               map[string]string                  `json:"tags,omitempty"`
	TargetWorkerCount  *int                               `json:"targetWorkerCount,omitempty"`
	TargetWorkerSizeId *int                               `json:"targetWorkerSizeId,omitempty"`
	WorkerTierName     *string                            `json:"workerTierName,omitempty"`
	ZoneRedundant      *bool                              `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ServerFarm_Spec{}

// ConvertSpecFrom populates our ServerFarm_Spec from the provided source
func (farm *ServerFarm_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == farm {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(farm)
}

// ConvertSpecTo populates the provided destination from our ServerFarm_Spec
func (farm *ServerFarm_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == farm {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(farm)
}

// Storage version of v1api20220301.ServerFarm_STATUS
type ServerFarm_STATUS struct {
	Conditions                []conditions.Condition            `json:"conditions,omitempty"`
	ElasticScaleEnabled       *bool                             `json:"elasticScaleEnabled,omitempty"`
	ExtendedLocation          *ExtendedLocation_STATUS          `json:"extendedLocation,omitempty"`
	FreeOfferExpirationTime   *string                           `json:"freeOfferExpirationTime,omitempty"`
	GeoRegion                 *string                           `json:"geoRegion,omitempty"`
	HostingEnvironmentProfile *HostingEnvironmentProfile_STATUS `json:"hostingEnvironmentProfile,omitempty"`
	HyperV                    *bool                             `json:"hyperV,omitempty"`
	Id                        *string                           `json:"id,omitempty"`
	IsSpot                    *bool                             `json:"isSpot,omitempty"`
	IsXenon                   *bool                             `json:"isXenon,omitempty"`
	Kind                      *string                           `json:"kind,omitempty"`
	KubeEnvironmentProfile    *KubeEnvironmentProfile_STATUS    `json:"kubeEnvironmentProfile,omitempty"`
	Location                  *string                           `json:"location,omitempty"`
	MaximumElasticWorkerCount *int                              `json:"maximumElasticWorkerCount,omitempty"`
	MaximumNumberOfWorkers    *int                              `json:"maximumNumberOfWorkers,omitempty"`
	Name                      *string                           `json:"name,omitempty"`
	NumberOfSites             *int                              `json:"numberOfSites,omitempty"`
	NumberOfWorkers           *int                              `json:"numberOfWorkers,omitempty"`
	PerSiteScaling            *bool                             `json:"perSiteScaling,omitempty"`
	PropertyBag               genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                           `json:"provisioningState,omitempty"`
	Reserved                  *bool                             `json:"reserved,omitempty"`
	ResourceGroup             *string                           `json:"resourceGroup,omitempty"`
	Sku                       *SkuDescription_STATUS            `json:"sku,omitempty"`
	SpotExpirationTime        *string                           `json:"spotExpirationTime,omitempty"`
	Status                    *string                           `json:"status,omitempty"`
	Subscription              *string                           `json:"subscription,omitempty"`
	Tags                      map[string]string                 `json:"tags,omitempty"`
	TargetWorkerCount         *int                              `json:"targetWorkerCount,omitempty"`
	TargetWorkerSizeId        *int                              `json:"targetWorkerSizeId,omitempty"`
	Type                      *string                           `json:"type,omitempty"`
	WorkerTierName            *string                           `json:"workerTierName,omitempty"`
	ZoneRedundant             *bool                             `json:"zoneRedundant,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ServerFarm_STATUS{}

// ConvertStatusFrom populates our ServerFarm_STATUS from the provided source
func (farm *ServerFarm_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == farm {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(farm)
}

// ConvertStatusTo populates the provided destination from our ServerFarm_STATUS
func (farm *ServerFarm_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == farm {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(farm)
}

// Storage version of v1api20220301.ExtendedLocation
// Extended Location.
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220301.ExtendedLocation_STATUS
// Extended Location.
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220301.HostingEnvironmentProfile
// Specification for an App Service Environment to use for this resource.
type HostingEnvironmentProfile struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID of the App Service Environment.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220301.HostingEnvironmentProfile_STATUS
// Specification for an App Service Environment to use for this resource.
type HostingEnvironmentProfile_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220301.KubeEnvironmentProfile
// Specification for a Kubernetes Environment to use for this resource.
type KubeEnvironmentProfile struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID of the Kubernetes Environment.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220301.KubeEnvironmentProfile_STATUS
// Specification for a Kubernetes Environment to use for this resource.
type KubeEnvironmentProfile_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// Storage version of v1api20220301.ServerFarmOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServerFarmOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20220301.SkuDescription
// Description of a SKU for a scalable resource.
type SkuDescription struct {
	Capabilities []Capability           `json:"capabilities,omitempty"`
	Capacity     *int                   `json:"capacity,omitempty"`
	Family       *string                `json:"family,omitempty"`
	Locations    []string               `json:"locations,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size         *string                `json:"size,omitempty"`
	SkuCapacity  *SkuCapacity           `json:"skuCapacity,omitempty"`
	Tier         *string                `json:"tier,omitempty"`
}

// Storage version of v1api20220301.SkuDescription_STATUS
// Description of a SKU for a scalable resource.
type SkuDescription_STATUS struct {
	Capabilities []Capability_STATUS    `json:"capabilities,omitempty"`
	Capacity     *int                   `json:"capacity,omitempty"`
	Family       *string                `json:"family,omitempty"`
	Locations    []string               `json:"locations,omitempty"`
	Name         *string                `json:"name,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size         *string                `json:"size,omitempty"`
	SkuCapacity  *SkuCapacity_STATUS    `json:"skuCapacity,omitempty"`
	Tier         *string                `json:"tier,omitempty"`
}

// Storage version of v1api20220301.Capability
// Describes the capabilities/features allowed for a specific SKU.
type Capability struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Reason      *string                `json:"reason,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20220301.Capability_STATUS
// Describes the capabilities/features allowed for a specific SKU.
type Capability_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Reason      *string                `json:"reason,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1api20220301.SkuCapacity
// Description of the App Service plan scale options.
type SkuCapacity struct {
	Default        *int                   `json:"default,omitempty"`
	ElasticMaximum *int                   `json:"elasticMaximum,omitempty"`
	Maximum        *int                   `json:"maximum,omitempty"`
	Minimum        *int                   `json:"minimum,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScaleType      *string                `json:"scaleType,omitempty"`
}

// Storage version of v1api20220301.SkuCapacity_STATUS
// Description of the App Service plan scale options.
type SkuCapacity_STATUS struct {
	Default        *int                   `json:"default,omitempty"`
	ElasticMaximum *int                   `json:"elasticMaximum,omitempty"`
	Maximum        *int                   `json:"maximum,omitempty"`
	Minimum        *int                   `json:"minimum,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScaleType      *string                `json:"scaleType,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ServerFarm{}, &ServerFarmList{})
}
