// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

import (
	"fmt"
	"github.com/Azure/azure-service-operator/hack/generated/apis/microsoft.managedidentity/v1alpha1api20181130storage"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/reflecthelpers"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups=microsoft.managedidentity.azure.com,resources=userassignedidentities,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.managedidentity.azure.com,resources={userassignedidentities/status,userassignedidentities/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2018-11-30/Microsoft.ManagedIdentity.json#/resourceDefinitions/userAssignedIdentities
type UserAssignedIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentities_Spec `json:"spec,omitempty"`
	Status            Identity_Status             `json:"status,omitempty"`
}

var _ conditions.Conditioner = &UserAssignedIdentity{}

// GetConditions returns the conditions of the resource
func (userAssignedIdentity *UserAssignedIdentity) GetConditions() conditions.Conditions {
	return userAssignedIdentity.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (userAssignedIdentity *UserAssignedIdentity) SetConditions(conditions conditions.Conditions) {
	userAssignedIdentity.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-managedidentity-azure-com-v1alpha1api20181130-userassignedidentity,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1alpha1api20181130,name=default.v1alpha1api20181130.userassignedidentities.microsoft.managedidentity.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &UserAssignedIdentity{}

// Default applies defaults to the UserAssignedIdentity resource
func (userAssignedIdentity *UserAssignedIdentity) Default() {
	userAssignedIdentity.defaultImpl()
	var temp interface{} = userAssignedIdentity
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (userAssignedIdentity *UserAssignedIdentity) defaultAzureName() {
	if userAssignedIdentity.Spec.AzureName == "" {
		userAssignedIdentity.Spec.AzureName = userAssignedIdentity.Name
	}
}

// defaultImpl applies the code generated defaults to the UserAssignedIdentity resource
func (userAssignedIdentity *UserAssignedIdentity) defaultImpl() {
	userAssignedIdentity.defaultAzureName()
}

var _ genruntime.KubernetesResource = &UserAssignedIdentity{}

// AzureName returns the Azure name of the resource
func (userAssignedIdentity *UserAssignedIdentity) AzureName() string {
	return userAssignedIdentity.Spec.AzureName
}

// GetSpec returns the specification of this resource
func (userAssignedIdentity *UserAssignedIdentity) GetSpec() genruntime.ConvertibleSpec {
	return &userAssignedIdentity.Spec
}

// GetStatus returns the status of this resource
func (userAssignedIdentity *UserAssignedIdentity) GetStatus() genruntime.ConvertibleStatus {
	return &userAssignedIdentity.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (userAssignedIdentity *UserAssignedIdentity) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (userAssignedIdentity *UserAssignedIdentity) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(userAssignedIdentity.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: userAssignedIdentity.Namespace, Name: userAssignedIdentity.Spec.Owner.Name}
}

// +kubebuilder:webhook:path=/validate-microsoft-managedidentity-azure-com-v1alpha1api20181130-userassignedidentity,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1alpha1api20181130,name=validate.v1alpha1api20181130.userassignedidentities.microsoft.managedidentity.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &UserAssignedIdentity{}

// ValidateCreate validates the creation of the resource
func (userAssignedIdentity *UserAssignedIdentity) ValidateCreate() error {
	validations := userAssignedIdentity.createValidations()
	var temp interface{} = userAssignedIdentity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (userAssignedIdentity *UserAssignedIdentity) ValidateDelete() error {
	validations := userAssignedIdentity.deleteValidations()
	var temp interface{} = userAssignedIdentity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (userAssignedIdentity *UserAssignedIdentity) ValidateUpdate(old runtime.Object) error {
	validations := userAssignedIdentity.updateValidations()
	var temp interface{} = userAssignedIdentity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (userAssignedIdentity *UserAssignedIdentity) createValidations() []func() error {
	return []func() error{userAssignedIdentity.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (userAssignedIdentity *UserAssignedIdentity) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (userAssignedIdentity *UserAssignedIdentity) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return userAssignedIdentity.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (userAssignedIdentity *UserAssignedIdentity) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&userAssignedIdentity.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromUserAssignedIdentity populates our UserAssignedIdentity from the provided source UserAssignedIdentity
func (userAssignedIdentity *UserAssignedIdentity) AssignPropertiesFromUserAssignedIdentity(source *v1alpha1api20181130storage.UserAssignedIdentity) error {

	// Spec
	var spec UserAssignedIdentities_Spec
	err := spec.AssignPropertiesFromUserAssignedIdentitiesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromUserAssignedIdentitiesSpec()")
	}
	userAssignedIdentity.Spec = spec

	// Status
	var status Identity_Status
	err = status.AssignPropertiesFromIdentityStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromIdentityStatus()")
	}
	userAssignedIdentity.Status = status

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentity populates the provided destination UserAssignedIdentity from our UserAssignedIdentity
func (userAssignedIdentity *UserAssignedIdentity) AssignPropertiesToUserAssignedIdentity(destination *v1alpha1api20181130storage.UserAssignedIdentity) error {

	// Spec
	var spec v1alpha1api20181130storage.UserAssignedIdentities_Spec
	err := userAssignedIdentity.Spec.AssignPropertiesToUserAssignedIdentitiesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToUserAssignedIdentitiesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20181130storage.Identity_Status
	err = userAssignedIdentity.Status.AssignPropertiesToIdentityStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToIdentityStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (userAssignedIdentity *UserAssignedIdentity) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: userAssignedIdentity.Spec.OriginalVersion(),
		Kind:    "UserAssignedIdentity",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2018-11-30/Microsoft.ManagedIdentity.json#/resourceDefinitions/userAssignedIdentities
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

//Generated from:
type Identity_Status struct {
	//ClientId: The id of the app associated with the identity. This is a random
	//generated UUID by MSI.
	ClientId *string `json:"clientId,omitempty"`

	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//PrincipalId: The id of the service principal object associated with the created
	//identity.
	PrincipalId *string `json:"principalId,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//TenantId: The id of the tenant which the identity belongs to.
	TenantId *string `json:"tenantId,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Identity_Status{}

// ConvertStatusFrom populates our Identity_Status from the provided source
func (identityStatus *Identity_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20181130storage.Identity_Status)
	if ok {
		// Populate our instance from source
		return identityStatus.AssignPropertiesFromIdentityStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20181130storage.Identity_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = identityStatus.AssignPropertiesFromIdentityStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Identity_Status
func (identityStatus *Identity_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20181130storage.Identity_Status)
	if ok {
		// Populate destination from our instance
		return identityStatus.AssignPropertiesToIdentityStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20181130storage.Identity_Status{}
	err := identityStatus.AssignPropertiesToIdentityStatus(dst)
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

var _ genruntime.FromARMConverter = &Identity_Status{}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (identityStatus *Identity_Status) CreateEmptyARMValue() interface{} {
	return Identity_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (identityStatus *Identity_Status) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(Identity_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Identity_StatusARM, got %T", armInput)
	}

	// Set property ‘ClientId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ClientId != nil {
			clientId := *typedInput.Properties.ClientId
			identityStatus.ClientId = &clientId
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		identityStatus.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		identityStatus.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		identityStatus.Name = &name
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			identityStatus.PrincipalId = &principalId
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		identityStatus.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			identityStatus.Tags[key] = value
		}
	}

	// Set property ‘TenantId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			identityStatus.TenantId = &tenantId
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		identityStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromIdentityStatus populates our Identity_Status from the provided source Identity_Status
func (identityStatus *Identity_Status) AssignPropertiesFromIdentityStatus(source *v1alpha1api20181130storage.Identity_Status) error {

	// ClientId
	if source.ClientId != nil {
		clientId := *source.ClientId
		identityStatus.ClientId = &clientId
	} else {
		identityStatus.ClientId = nil
	}

	// Conditions
	conditionList := make([]conditions.Condition, len(source.Conditions))
	for conditionIndex, conditionItem := range source.Conditions {
		// Shadow the loop variable to avoid aliasing
		conditionItem := conditionItem
		conditionList[conditionIndex] = conditionItem.Copy()
	}
	identityStatus.Conditions = conditionList

	// Id
	if source.Id != nil {
		id := *source.Id
		identityStatus.Id = &id
	} else {
		identityStatus.Id = nil
	}

	// Location
	if source.Location != nil {
		location := *source.Location
		identityStatus.Location = &location
	} else {
		identityStatus.Location = nil
	}

	// Name
	if source.Name != nil {
		name := *source.Name
		identityStatus.Name = &name
	} else {
		identityStatus.Name = nil
	}

	// PrincipalId
	if source.PrincipalId != nil {
		principalId := *source.PrincipalId
		identityStatus.PrincipalId = &principalId
	} else {
		identityStatus.PrincipalId = nil
	}

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range source.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	identityStatus.Tags = tagMap

	// TenantId
	if source.TenantId != nil {
		tenantId := *source.TenantId
		identityStatus.TenantId = &tenantId
	} else {
		identityStatus.TenantId = nil
	}

	// Type
	if source.Type != nil {
		typeVar := *source.Type
		identityStatus.Type = &typeVar
	} else {
		identityStatus.Type = nil
	}

	// No error
	return nil
}

// AssignPropertiesToIdentityStatus populates the provided destination Identity_Status from our Identity_Status
func (identityStatus *Identity_Status) AssignPropertiesToIdentityStatus(destination *v1alpha1api20181130storage.Identity_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ClientId
	if identityStatus.ClientId != nil {
		clientId := *identityStatus.ClientId
		destination.ClientId = &clientId
	} else {
		destination.ClientId = nil
	}

	// Conditions
	conditionList := make([]conditions.Condition, len(identityStatus.Conditions))
	for conditionIndex, conditionItem := range identityStatus.Conditions {
		// Shadow the loop variable to avoid aliasing
		conditionItem := conditionItem
		conditionList[conditionIndex] = conditionItem.Copy()
	}
	destination.Conditions = conditionList

	// Id
	if identityStatus.Id != nil {
		id := *identityStatus.Id
		destination.Id = &id
	} else {
		destination.Id = nil
	}

	// Location
	if identityStatus.Location != nil {
		location := *identityStatus.Location
		destination.Location = &location
	} else {
		destination.Location = nil
	}

	// Name
	if identityStatus.Name != nil {
		name := *identityStatus.Name
		destination.Name = &name
	} else {
		destination.Name = nil
	}

	// PrincipalId
	if identityStatus.PrincipalId != nil {
		principalId := *identityStatus.PrincipalId
		destination.PrincipalId = &principalId
	} else {
		destination.PrincipalId = nil
	}

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range identityStatus.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	destination.Tags = tagMap

	// TenantId
	if identityStatus.TenantId != nil {
		tenantId := *identityStatus.TenantId
		destination.TenantId = &tenantId
	} else {
		destination.TenantId = nil
	}

	// Type
	if identityStatus.Type != nil {
		typeVar := *identityStatus.Type
		destination.Type = &typeVar
	} else {
		destination.Type = nil
	}

	// Update the property bag
	destination.PropertyBag = propertyBag

	// No error
	return nil
}

type UserAssignedIdentities_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: The Azure region where the identity lives.
	Location string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.resources.azure.com" json:"owner" kind:"ResourceGroup"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &UserAssignedIdentities_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) ConvertToARM(name string, resolvedReferences genruntime.ResolvedReferences) (interface{}, error) {
	if userAssignedIdentitiesSpec == nil {
		return nil, nil
	}
	var result UserAssignedIdentities_SpecARM

	// Set property ‘APIVersion’:
	result.APIVersion = UserAssignedIdentitiesSpecAPIVersion20181130

	// Set property ‘Location’:
	result.Location = userAssignedIdentitiesSpec.Location

	// Set property ‘Name’:
	result.Name = name

	// Set property ‘Tags’:
	if userAssignedIdentitiesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range userAssignedIdentitiesSpec.Tags {
			result.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	result.Type = UserAssignedIdentitiesSpecTypeMicrosoftManagedIdentityUserAssignedIdentities
	return result, nil
}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) CreateEmptyARMValue() interface{} {
	return UserAssignedIdentities_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentities_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentities_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	userAssignedIdentitiesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	userAssignedIdentitiesSpec.Location = typedInput.Location

	// Set property ‘Owner’:
	userAssignedIdentitiesSpec.Owner = owner

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		userAssignedIdentitiesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			userAssignedIdentitiesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentities_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentities_Spec from the provided source
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20181130storage.UserAssignedIdentities_Spec)
	if ok {
		// Populate our instance from source
		return userAssignedIdentitiesSpec.AssignPropertiesFromUserAssignedIdentitiesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20181130storage.UserAssignedIdentities_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = userAssignedIdentitiesSpec.AssignPropertiesFromUserAssignedIdentitiesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentities_Spec
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20181130storage.UserAssignedIdentities_Spec)
	if ok {
		// Populate destination from our instance
		return userAssignedIdentitiesSpec.AssignPropertiesToUserAssignedIdentitiesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20181130storage.UserAssignedIdentities_Spec{}
	err := userAssignedIdentitiesSpec.AssignPropertiesToUserAssignedIdentitiesSpec(dst)
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

// AssignPropertiesFromUserAssignedIdentitiesSpec populates our UserAssignedIdentities_Spec from the provided source UserAssignedIdentities_Spec
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) AssignPropertiesFromUserAssignedIdentitiesSpec(source *v1alpha1api20181130storage.UserAssignedIdentities_Spec) error {

	// AzureName
	userAssignedIdentitiesSpec.AzureName = source.AzureName

	// Location
	if source.Location != nil {
		userAssignedIdentitiesSpec.Location = *source.Location
	} else {
		userAssignedIdentitiesSpec.Location = ""
	}

	// Owner
	userAssignedIdentitiesSpec.Owner = source.Owner.Copy()

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range source.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	userAssignedIdentitiesSpec.Tags = tagMap

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentitiesSpec populates the provided destination UserAssignedIdentities_Spec from our UserAssignedIdentities_Spec
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) AssignPropertiesToUserAssignedIdentitiesSpec(destination *v1alpha1api20181130storage.UserAssignedIdentities_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = userAssignedIdentitiesSpec.AzureName

	// Location
	location := userAssignedIdentitiesSpec.Location
	destination.Location = &location

	// OriginalVersion
	destination.OriginalVersion = userAssignedIdentitiesSpec.OriginalVersion()

	// Owner
	destination.Owner = userAssignedIdentitiesSpec.Owner.Copy()

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range userAssignedIdentitiesSpec.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	destination.Tags = tagMap

	// Update the property bag
	destination.PropertyBag = propertyBag

	// No error
	return nil
}

func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (userAssignedIdentitiesSpec *UserAssignedIdentities_Spec) SetAzureName(azureName string) {
	userAssignedIdentitiesSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
