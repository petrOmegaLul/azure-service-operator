// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import (
	"fmt"
	v20231115s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabaseThroughputSetting{}

// ConvertFrom populates our MongodbDatabaseThroughputSetting from the provided hub MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20231115s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_MongodbDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20231115s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_MongodbDatabaseThroughputSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20231115-mongodbdatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1api20231115,name=default.v1api20231115.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &MongodbDatabaseThroughputSetting{}

// Default applies defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp any = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) defaultImpl() {}

var _ genruntime.ImportableResource = &MongodbDatabaseThroughputSetting{}

// InitializeSpec initializes the spec for this resource from the given status
func (setting *MongodbDatabaseThroughputSetting) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS); ok {
		return setting.Spec.Initialize_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(s)
	}

	return fmt.Errorf("expected Status of type DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (setting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *MongodbDatabaseThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20231115-mongodbdatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1api20231115,name=validate.v1api20231115.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &MongodbDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateCreate() (admission.Warnings, error) {
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateDelete() (admission.Warnings, error) {
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *MongodbDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (setting *MongodbDatabaseThroughputSetting) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(setting)
}

// validateResourceReferences validates all resource references
func (setting *MongodbDatabaseThroughputSetting) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *MongodbDatabaseThroughputSetting) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*MongodbDatabaseThroughputSetting)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignProperties_From_MongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_From_MongodbDatabaseThroughputSetting(source *v20231115s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
	err = status.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_To_MongodbDatabaseThroughputSetting(destination *v20231115s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-11-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

type DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`

	// +kubebuilder:validation:Required
	// Resource: The standard JSON format of a resource throughput
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_ARM{}

	// Set property "Location":
	if setting.Location != nil {
		location := *setting.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if setting.Resource != nil {
		result.Properties = &ThroughputSettingsUpdateProperties_ARM{}
	}
	if setting.Resource != nil {
		resource_ARM, err := (*setting.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resource_ARM.(*ThroughputSettingsResource_ARM)
		result.Properties.Resource = &resource
	}

	// Set property "Tags":
	if setting.Tags != nil {
		result.Tags = make(map[string]string, len(setting.Tags))
		for key, value := range setting.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec_ARM, got %T", armInput)
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property "Owner":
	setting.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec from the provided source
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec{}
	err := setting.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec populates our DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec from the provided source DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(source *v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) error {

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec populates the provided destination DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec from our DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec(destination *v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource v20231115s.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS populates our DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec from the provided source DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) Initialize_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(source *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) error {

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.Initialize_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling Initialize_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name     *string                                          `json:"name,omitempty"`
	Resource *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags     map[string]string                                `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS from the provided source
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsGetProperties_Resource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS populates our DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS from the provided source DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) AssignProperties_From_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(source *v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS populates the provided destination DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS from our DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) AssignProperties_To_DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS(destination *v20231115s.DatabaseAccounts_MongodbDatabases_ThroughputSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource v20231115s.ThroughputSettingsGetProperties_Resource_STATUS
		err := setting.Resource.AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
