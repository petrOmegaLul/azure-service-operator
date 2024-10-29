// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_CosmosDbSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CosmosDbSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCosmosDbSettings, CosmosDbSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCosmosDbSettings runs a test to see if a specific instance of CosmosDbSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForCosmosDbSettings(subject CosmosDbSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CosmosDbSettings
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CosmosDbSettings instances for property testing - lazily instantiated by CosmosDbSettingsGenerator()
var cosmosDbSettingsGenerator gopter.Gen

// CosmosDbSettingsGenerator returns a generator of CosmosDbSettings instances for property testing.
func CosmosDbSettingsGenerator() gopter.Gen {
	if cosmosDbSettingsGenerator != nil {
		return cosmosDbSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCosmosDbSettings(generators)
	cosmosDbSettingsGenerator = gen.Struct(reflect.TypeOf(CosmosDbSettings{}), generators)

	return cosmosDbSettingsGenerator
}

// AddIndependentPropertyGeneratorsForCosmosDbSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCosmosDbSettings(gens map[string]gopter.Gen) {
	gens["CollectionsThroughput"] = gen.PtrOf(gen.Int())
}

func Test_EncryptionProperty_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionProperty via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionProperty, EncryptionPropertyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionProperty runs a test to see if a specific instance of EncryptionProperty round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionProperty(subject EncryptionProperty) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionProperty
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of EncryptionProperty instances for property testing - lazily instantiated by EncryptionPropertyGenerator()
var encryptionPropertyGenerator gopter.Gen

// EncryptionPropertyGenerator returns a generator of EncryptionProperty instances for property testing.
func EncryptionPropertyGenerator() gopter.Gen {
	if encryptionPropertyGenerator != nil {
		return encryptionPropertyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForEncryptionProperty(generators)
	encryptionPropertyGenerator = gen.Struct(reflect.TypeOf(EncryptionProperty{}), generators)

	return encryptionPropertyGenerator
}

// AddRelatedPropertyGeneratorsForEncryptionProperty is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionProperty(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityForCmkGenerator())
	gens["KeyVaultProperties"] = gen.PtrOf(KeyVaultPropertiesGenerator())
}

func Test_Identity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentity, IdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentity runs a test to see if a specific instance of Identity round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentity(subject Identity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Identity instances for property testing - lazily instantiated by IdentityGenerator()
var identityGenerator gopter.Gen

// IdentityGenerator returns a generator of Identity instances for property testing.
// We first initialize identityGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func IdentityGenerator() gopter.Gen {
	if identityGenerator != nil {
		return identityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity(generators)
	identityGenerator = gen.Struct(reflect.TypeOf(Identity{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentity(generators)
	AddRelatedPropertyGeneratorsForIdentity(generators)
	identityGenerator = gen.Struct(reflect.TypeOf(Identity{}), generators)

	return identityGenerator
}

// AddIndependentPropertyGeneratorsForIdentity is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentity(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		Identity_Type_None,
		Identity_Type_SystemAssigned,
		Identity_Type_SystemAssignedUserAssigned,
		Identity_Type_UserAssigned))
}

// AddRelatedPropertyGeneratorsForIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForIdentity(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetailsGenerator())
}

func Test_IdentityForCmk_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of IdentityForCmk via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityForCmk, IdentityForCmkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityForCmk runs a test to see if a specific instance of IdentityForCmk round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityForCmk(subject IdentityForCmk) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual IdentityForCmk
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of IdentityForCmk instances for property testing - lazily instantiated by IdentityForCmkGenerator()
var identityForCmkGenerator gopter.Gen

// IdentityForCmkGenerator returns a generator of IdentityForCmk instances for property testing.
func IdentityForCmkGenerator() gopter.Gen {
	if identityForCmkGenerator != nil {
		return identityForCmkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityForCmk(generators)
	identityForCmkGenerator = gen.Struct(reflect.TypeOf(IdentityForCmk{}), generators)

	return identityForCmkGenerator
}

// AddIndependentPropertyGeneratorsForIdentityForCmk is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityForCmk(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyVaultProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyVaultProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyVaultProperties, KeyVaultPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyVaultProperties runs a test to see if a specific instance of KeyVaultProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyVaultProperties(subject KeyVaultProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyVaultProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of KeyVaultProperties instances for property testing - lazily instantiated by KeyVaultPropertiesGenerator()
var keyVaultPropertiesGenerator gopter.Gen

// KeyVaultPropertiesGenerator returns a generator of KeyVaultProperties instances for property testing.
func KeyVaultPropertiesGenerator() gopter.Gen {
	if keyVaultPropertiesGenerator != nil {
		return keyVaultPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyVaultProperties(generators)
	keyVaultPropertiesGenerator = gen.Struct(reflect.TypeOf(KeyVaultProperties{}), generators)

	return keyVaultPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForKeyVaultProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyVaultProperties(gens map[string]gopter.Gen) {
	gens["IdentityClientId"] = gen.PtrOf(gen.AlphaString())
	gens["KeyIdentifier"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVaultArmId"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServiceManagedResourcesSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServiceManagedResourcesSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServiceManagedResourcesSettings, ServiceManagedResourcesSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServiceManagedResourcesSettings runs a test to see if a specific instance of ServiceManagedResourcesSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForServiceManagedResourcesSettings(subject ServiceManagedResourcesSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServiceManagedResourcesSettings
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ServiceManagedResourcesSettings instances for property testing - lazily instantiated by
// ServiceManagedResourcesSettingsGenerator()
var serviceManagedResourcesSettingsGenerator gopter.Gen

// ServiceManagedResourcesSettingsGenerator returns a generator of ServiceManagedResourcesSettings instances for property testing.
func ServiceManagedResourcesSettingsGenerator() gopter.Gen {
	if serviceManagedResourcesSettingsGenerator != nil {
		return serviceManagedResourcesSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings(generators)
	serviceManagedResourcesSettingsGenerator = gen.Struct(reflect.TypeOf(ServiceManagedResourcesSettings{}), generators)

	return serviceManagedResourcesSettingsGenerator
}

// AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForServiceManagedResourcesSettings(gens map[string]gopter.Gen) {
	gens["CosmosDb"] = gen.PtrOf(CosmosDbSettingsGenerator())
}

func Test_SharedPrivateLinkResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResource, SharedPrivateLinkResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResource runs a test to see if a specific instance of SharedPrivateLinkResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResource(subject SharedPrivateLinkResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResource
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SharedPrivateLinkResource instances for property testing - lazily instantiated by
// SharedPrivateLinkResourceGenerator()
var sharedPrivateLinkResourceGenerator gopter.Gen

// SharedPrivateLinkResourceGenerator returns a generator of SharedPrivateLinkResource instances for property testing.
// We first initialize sharedPrivateLinkResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SharedPrivateLinkResourceGenerator() gopter.Gen {
	if sharedPrivateLinkResourceGenerator != nil {
		return sharedPrivateLinkResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource(generators)
	sharedPrivateLinkResourceGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResource(generators)
	AddRelatedPropertyGeneratorsForSharedPrivateLinkResource(generators)
	sharedPrivateLinkResourceGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResource{}), generators)

	return sharedPrivateLinkResourceGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResource(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSharedPrivateLinkResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSharedPrivateLinkResource(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SharedPrivateLinkResourcePropertyGenerator())
}

func Test_SharedPrivateLinkResourceProperty_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SharedPrivateLinkResourceProperty via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSharedPrivateLinkResourceProperty, SharedPrivateLinkResourcePropertyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSharedPrivateLinkResourceProperty runs a test to see if a specific instance of SharedPrivateLinkResourceProperty round trips to JSON and back losslessly
func RunJSONSerializationTestForSharedPrivateLinkResourceProperty(subject SharedPrivateLinkResourceProperty) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SharedPrivateLinkResourceProperty
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SharedPrivateLinkResourceProperty instances for property testing - lazily instantiated by
// SharedPrivateLinkResourcePropertyGenerator()
var sharedPrivateLinkResourcePropertyGenerator gopter.Gen

// SharedPrivateLinkResourcePropertyGenerator returns a generator of SharedPrivateLinkResourceProperty instances for property testing.
func SharedPrivateLinkResourcePropertyGenerator() gopter.Gen {
	if sharedPrivateLinkResourcePropertyGenerator != nil {
		return sharedPrivateLinkResourcePropertyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty(generators)
	sharedPrivateLinkResourcePropertyGenerator = gen.Struct(reflect.TypeOf(SharedPrivateLinkResourceProperty{}), generators)

	return sharedPrivateLinkResourcePropertyGenerator
}

// AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSharedPrivateLinkResourceProperty(gens map[string]gopter.Gen) {
	gens["GroupId"] = gen.PtrOf(gen.AlphaString())
	gens["PrivateLinkResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["RequestMessage"] = gen.PtrOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		PrivateEndpointServiceConnectionStatus_Approved,
		PrivateEndpointServiceConnectionStatus_Disconnected,
		PrivateEndpointServiceConnectionStatus_Pending,
		PrivateEndpointServiceConnectionStatus_Rejected,
		PrivateEndpointServiceConnectionStatus_Timeout))
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_SystemData_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData, SystemDataGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData runs a test to see if a specific instance of SystemData round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData(subject SystemData) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData instances for property testing - lazily instantiated by SystemDataGenerator()
var systemDataGenerator gopter.Gen

// SystemDataGenerator returns a generator of SystemData instances for property testing.
func SystemDataGenerator() gopter.Gen {
	if systemDataGenerator != nil {
		return systemDataGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData(generators)
	systemDataGenerator = gen.Struct(reflect.TypeOf(SystemData{}), generators)

	return systemDataGenerator
}

// AddIndependentPropertyGeneratorsForSystemData is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_CreatedByType_Application,
		SystemData_CreatedByType_Key,
		SystemData_CreatedByType_ManagedIdentity,
		SystemData_CreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemData_LastModifiedByType_Application,
		SystemData_LastModifiedByType_Key,
		SystemData_LastModifiedByType_ManagedIdentity,
		SystemData_LastModifiedByType_User))
}

func Test_UserAssignedIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails, UserAssignedIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails runs a test to see if a specific instance of UserAssignedIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails(subject UserAssignedIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of UserAssignedIdentityDetails instances for property testing - lazily instantiated by
// UserAssignedIdentityDetailsGenerator()
var userAssignedIdentityDetailsGenerator gopter.Gen

// UserAssignedIdentityDetailsGenerator returns a generator of UserAssignedIdentityDetails instances for property testing.
func UserAssignedIdentityDetailsGenerator() gopter.Gen {
	if userAssignedIdentityDetailsGenerator != nil {
		return userAssignedIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails{}), generators)

	return userAssignedIdentityDetailsGenerator
}

func Test_WorkspaceProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceProperties, WorkspacePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceProperties runs a test to see if a specific instance of WorkspaceProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceProperties(subject WorkspaceProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of WorkspaceProperties instances for property testing - lazily instantiated by
// WorkspacePropertiesGenerator()
var workspacePropertiesGenerator gopter.Gen

// WorkspacePropertiesGenerator returns a generator of WorkspaceProperties instances for property testing.
// We first initialize workspacePropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacePropertiesGenerator() gopter.Gen {
	if workspacePropertiesGenerator != nil {
		return workspacePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties(generators)
	workspacePropertiesGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceProperties(generators)
	AddRelatedPropertyGeneratorsForWorkspaceProperties(generators)
	workspacePropertiesGenerator = gen.Struct(reflect.TypeOf(WorkspaceProperties{}), generators)

	return workspacePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceProperties(gens map[string]gopter.Gen) {
	gens["AllowPublicAccessWhenBehindVnet"] = gen.PtrOf(gen.Bool())
	gens["ApplicationInsights"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerRegistry"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DiscoveryUrl"] = gen.PtrOf(gen.AlphaString())
	gens["FriendlyName"] = gen.PtrOf(gen.AlphaString())
	gens["HbiWorkspace"] = gen.PtrOf(gen.Bool())
	gens["ImageBuildCompute"] = gen.PtrOf(gen.AlphaString())
	gens["KeyVault"] = gen.PtrOf(gen.AlphaString())
	gens["PrimaryUserAssignedIdentity"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(WorkspaceProperties_PublicNetworkAccess_Disabled, WorkspaceProperties_PublicNetworkAccess_Enabled))
	gens["StorageAccount"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspaceProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspaceProperties(gens map[string]gopter.Gen) {
	gens["Encryption"] = gen.PtrOf(EncryptionPropertyGenerator())
	gens["ServiceManagedResourcesSettings"] = gen.PtrOf(ServiceManagedResourcesSettingsGenerator())
	gens["SharedPrivateLinkResources"] = gen.SliceOf(SharedPrivateLinkResourceGenerator())
}

func Test_Workspace_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Workspace_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspace_Spec, Workspace_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspace_Spec runs a test to see if a specific instance of Workspace_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspace_Spec(subject Workspace_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Workspace_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Workspace_Spec instances for property testing - lazily instantiated by Workspace_SpecGenerator()
var workspace_SpecGenerator gopter.Gen

// Workspace_SpecGenerator returns a generator of Workspace_Spec instances for property testing.
// We first initialize workspace_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Workspace_SpecGenerator() gopter.Gen {
	if workspace_SpecGenerator != nil {
		return workspace_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec(generators)
	workspace_SpecGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspace_Spec(generators)
	AddRelatedPropertyGeneratorsForWorkspace_Spec(generators)
	workspace_SpecGenerator = gen.Struct(reflect.TypeOf(Workspace_Spec{}), generators)

	return workspace_SpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspace_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspace_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForWorkspace_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspace_Spec(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(IdentityGenerator())
	gens["Properties"] = gen.PtrOf(WorkspacePropertiesGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataGenerator())
}
