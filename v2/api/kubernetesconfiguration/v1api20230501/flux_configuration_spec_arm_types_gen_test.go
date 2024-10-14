// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_AzureBlobDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureBlobDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureBlobDefinition_ARM, AzureBlobDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureBlobDefinition_ARM runs a test to see if a specific instance of AzureBlobDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureBlobDefinition_ARM(subject AzureBlobDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureBlobDefinition_ARM
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

// Generator of AzureBlobDefinition_ARM instances for property testing - lazily instantiated by
// AzureBlobDefinition_ARMGenerator()
var azureBlobDefinition_ARMGenerator gopter.Gen

// AzureBlobDefinition_ARMGenerator returns a generator of AzureBlobDefinition_ARM instances for property testing.
// We first initialize azureBlobDefinition_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AzureBlobDefinition_ARMGenerator() gopter.Gen {
	if azureBlobDefinition_ARMGenerator != nil {
		return azureBlobDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureBlobDefinition_ARM(generators)
	azureBlobDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(AzureBlobDefinition_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureBlobDefinition_ARM(generators)
	AddRelatedPropertyGeneratorsForAzureBlobDefinition_ARM(generators)
	azureBlobDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(AzureBlobDefinition_ARM{}), generators)

	return azureBlobDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAzureBlobDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureBlobDefinition_ARM(gens map[string]gopter.Gen) {
	gens["AccountKey"] = gen.PtrOf(gen.AlphaString())
	gens["ContainerName"] = gen.PtrOf(gen.AlphaString())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SasToken"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAzureBlobDefinition_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAzureBlobDefinition_ARM(gens map[string]gopter.Gen) {
	gens["ManagedIdentity"] = gen.PtrOf(ManagedIdentityDefinition_ARMGenerator())
	gens["ServicePrincipal"] = gen.PtrOf(ServicePrincipalDefinition_ARMGenerator())
}

func Test_BucketDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BucketDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBucketDefinition_ARM, BucketDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBucketDefinition_ARM runs a test to see if a specific instance of BucketDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForBucketDefinition_ARM(subject BucketDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BucketDefinition_ARM
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

// Generator of BucketDefinition_ARM instances for property testing - lazily instantiated by
// BucketDefinition_ARMGenerator()
var bucketDefinition_ARMGenerator gopter.Gen

// BucketDefinition_ARMGenerator returns a generator of BucketDefinition_ARM instances for property testing.
func BucketDefinition_ARMGenerator() gopter.Gen {
	if bucketDefinition_ARMGenerator != nil {
		return bucketDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBucketDefinition_ARM(generators)
	bucketDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(BucketDefinition_ARM{}), generators)

	return bucketDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForBucketDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBucketDefinition_ARM(gens map[string]gopter.Gen) {
	gens["AccessKey"] = gen.PtrOf(gen.AlphaString())
	gens["BucketName"] = gen.PtrOf(gen.AlphaString())
	gens["Insecure"] = gen.PtrOf(gen.Bool())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

func Test_FluxConfiguration_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FluxConfiguration_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFluxConfiguration_Properties_Spec_ARM, FluxConfiguration_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFluxConfiguration_Properties_Spec_ARM runs a test to see if a specific instance of FluxConfiguration_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFluxConfiguration_Properties_Spec_ARM(subject FluxConfiguration_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FluxConfiguration_Properties_Spec_ARM
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

// Generator of FluxConfiguration_Properties_Spec_ARM instances for property testing - lazily instantiated by
// FluxConfiguration_Properties_Spec_ARMGenerator()
var fluxConfiguration_Properties_Spec_ARMGenerator gopter.Gen

// FluxConfiguration_Properties_Spec_ARMGenerator returns a generator of FluxConfiguration_Properties_Spec_ARM instances for property testing.
// We first initialize fluxConfiguration_Properties_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FluxConfiguration_Properties_Spec_ARMGenerator() gopter.Gen {
	if fluxConfiguration_Properties_Spec_ARMGenerator != nil {
		return fluxConfiguration_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM(generators)
	fluxConfiguration_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Properties_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM(generators)
	fluxConfiguration_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Properties_Spec_ARM{}), generators)

	return fluxConfiguration_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["ConfigurationProtectedSettings"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Namespace"] = gen.PtrOf(gen.AlphaString())
	gens["ReconciliationWaitDuration"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.OneConstOf(ScopeDefinition_ARM_Cluster, ScopeDefinition_ARM_Namespace))
	gens["SourceKind"] = gen.PtrOf(gen.OneConstOf(SourceKindDefinition_ARM_AzureBlob, SourceKindDefinition_ARM_Bucket, SourceKindDefinition_ARM_GitRepository))
	gens["Suspend"] = gen.PtrOf(gen.Bool())
	gens["WaitForReconciliation"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFluxConfiguration_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["AzureBlob"] = gen.PtrOf(AzureBlobDefinition_ARMGenerator())
	gens["Bucket"] = gen.PtrOf(BucketDefinition_ARMGenerator())
	gens["GitRepository"] = gen.PtrOf(GitRepositoryDefinition_ARMGenerator())
	gens["Kustomizations"] = gen.MapOf(
		gen.AlphaString(),
		KustomizationDefinition_ARMGenerator())
}

func Test_FluxConfiguration_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FluxConfiguration_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFluxConfiguration_Spec_ARM, FluxConfiguration_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFluxConfiguration_Spec_ARM runs a test to see if a specific instance of FluxConfiguration_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFluxConfiguration_Spec_ARM(subject FluxConfiguration_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FluxConfiguration_Spec_ARM
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

// Generator of FluxConfiguration_Spec_ARM instances for property testing - lazily instantiated by
// FluxConfiguration_Spec_ARMGenerator()
var fluxConfiguration_Spec_ARMGenerator gopter.Gen

// FluxConfiguration_Spec_ARMGenerator returns a generator of FluxConfiguration_Spec_ARM instances for property testing.
// We first initialize fluxConfiguration_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FluxConfiguration_Spec_ARMGenerator() gopter.Gen {
	if fluxConfiguration_Spec_ARMGenerator != nil {
		return fluxConfiguration_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Spec_ARM(generators)
	fluxConfiguration_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFluxConfiguration_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForFluxConfiguration_Spec_ARM(generators)
	fluxConfiguration_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(FluxConfiguration_Spec_ARM{}), generators)

	return fluxConfiguration_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForFluxConfiguration_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFluxConfiguration_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForFluxConfiguration_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFluxConfiguration_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FluxConfiguration_Properties_Spec_ARMGenerator())
}

func Test_GitRepositoryDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of GitRepositoryDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForGitRepositoryDefinition_ARM, GitRepositoryDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForGitRepositoryDefinition_ARM runs a test to see if a specific instance of GitRepositoryDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForGitRepositoryDefinition_ARM(subject GitRepositoryDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual GitRepositoryDefinition_ARM
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

// Generator of GitRepositoryDefinition_ARM instances for property testing - lazily instantiated by
// GitRepositoryDefinition_ARMGenerator()
var gitRepositoryDefinition_ARMGenerator gopter.Gen

// GitRepositoryDefinition_ARMGenerator returns a generator of GitRepositoryDefinition_ARM instances for property testing.
// We first initialize gitRepositoryDefinition_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func GitRepositoryDefinition_ARMGenerator() gopter.Gen {
	if gitRepositoryDefinition_ARMGenerator != nil {
		return gitRepositoryDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitRepositoryDefinition_ARM(generators)
	gitRepositoryDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(GitRepositoryDefinition_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForGitRepositoryDefinition_ARM(generators)
	AddRelatedPropertyGeneratorsForGitRepositoryDefinition_ARM(generators)
	gitRepositoryDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(GitRepositoryDefinition_ARM{}), generators)

	return gitRepositoryDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForGitRepositoryDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForGitRepositoryDefinition_ARM(gens map[string]gopter.Gen) {
	gens["HttpsCACert"] = gen.PtrOf(gen.AlphaString())
	gens["HttpsUser"] = gen.PtrOf(gen.AlphaString())
	gens["LocalAuthRef"] = gen.PtrOf(gen.AlphaString())
	gens["SshKnownHosts"] = gen.PtrOf(gen.AlphaString())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForGitRepositoryDefinition_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForGitRepositoryDefinition_ARM(gens map[string]gopter.Gen) {
	gens["RepositoryRef"] = gen.PtrOf(RepositoryRefDefinition_ARMGenerator())
}

func Test_KustomizationDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KustomizationDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKustomizationDefinition_ARM, KustomizationDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKustomizationDefinition_ARM runs a test to see if a specific instance of KustomizationDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKustomizationDefinition_ARM(subject KustomizationDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KustomizationDefinition_ARM
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

// Generator of KustomizationDefinition_ARM instances for property testing - lazily instantiated by
// KustomizationDefinition_ARMGenerator()
var kustomizationDefinition_ARMGenerator gopter.Gen

// KustomizationDefinition_ARMGenerator returns a generator of KustomizationDefinition_ARM instances for property testing.
// We first initialize kustomizationDefinition_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KustomizationDefinition_ARMGenerator() gopter.Gen {
	if kustomizationDefinition_ARMGenerator != nil {
		return kustomizationDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKustomizationDefinition_ARM(generators)
	kustomizationDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(KustomizationDefinition_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKustomizationDefinition_ARM(generators)
	AddRelatedPropertyGeneratorsForKustomizationDefinition_ARM(generators)
	kustomizationDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(KustomizationDefinition_ARM{}), generators)

	return kustomizationDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKustomizationDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKustomizationDefinition_ARM(gens map[string]gopter.Gen) {
	gens["DependsOn"] = gen.SliceOf(gen.AlphaString())
	gens["Force"] = gen.PtrOf(gen.Bool())
	gens["Path"] = gen.PtrOf(gen.AlphaString())
	gens["Prune"] = gen.PtrOf(gen.Bool())
	gens["RetryIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SyncIntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["TimeoutInSeconds"] = gen.PtrOf(gen.Int())
	gens["Wait"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForKustomizationDefinition_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKustomizationDefinition_ARM(gens map[string]gopter.Gen) {
	gens["PostBuild"] = gen.PtrOf(PostBuildDefinition_ARMGenerator())
}

func Test_ManagedIdentityDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedIdentityDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedIdentityDefinition_ARM, ManagedIdentityDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedIdentityDefinition_ARM runs a test to see if a specific instance of ManagedIdentityDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedIdentityDefinition_ARM(subject ManagedIdentityDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedIdentityDefinition_ARM
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

// Generator of ManagedIdentityDefinition_ARM instances for property testing - lazily instantiated by
// ManagedIdentityDefinition_ARMGenerator()
var managedIdentityDefinition_ARMGenerator gopter.Gen

// ManagedIdentityDefinition_ARMGenerator returns a generator of ManagedIdentityDefinition_ARM instances for property testing.
func ManagedIdentityDefinition_ARMGenerator() gopter.Gen {
	if managedIdentityDefinition_ARMGenerator != nil {
		return managedIdentityDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedIdentityDefinition_ARM(generators)
	managedIdentityDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(ManagedIdentityDefinition_ARM{}), generators)

	return managedIdentityDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedIdentityDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedIdentityDefinition_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
}

func Test_PostBuildDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PostBuildDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPostBuildDefinition_ARM, PostBuildDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPostBuildDefinition_ARM runs a test to see if a specific instance of PostBuildDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPostBuildDefinition_ARM(subject PostBuildDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PostBuildDefinition_ARM
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

// Generator of PostBuildDefinition_ARM instances for property testing - lazily instantiated by
// PostBuildDefinition_ARMGenerator()
var postBuildDefinition_ARMGenerator gopter.Gen

// PostBuildDefinition_ARMGenerator returns a generator of PostBuildDefinition_ARM instances for property testing.
// We first initialize postBuildDefinition_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PostBuildDefinition_ARMGenerator() gopter.Gen {
	if postBuildDefinition_ARMGenerator != nil {
		return postBuildDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPostBuildDefinition_ARM(generators)
	postBuildDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(PostBuildDefinition_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPostBuildDefinition_ARM(generators)
	AddRelatedPropertyGeneratorsForPostBuildDefinition_ARM(generators)
	postBuildDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(PostBuildDefinition_ARM{}), generators)

	return postBuildDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPostBuildDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPostBuildDefinition_ARM(gens map[string]gopter.Gen) {
	gens["Substitute"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPostBuildDefinition_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPostBuildDefinition_ARM(gens map[string]gopter.Gen) {
	gens["SubstituteFrom"] = gen.SliceOf(SubstituteFromDefinition_ARMGenerator())
}

func Test_RepositoryRefDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RepositoryRefDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRepositoryRefDefinition_ARM, RepositoryRefDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRepositoryRefDefinition_ARM runs a test to see if a specific instance of RepositoryRefDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRepositoryRefDefinition_ARM(subject RepositoryRefDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RepositoryRefDefinition_ARM
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

// Generator of RepositoryRefDefinition_ARM instances for property testing - lazily instantiated by
// RepositoryRefDefinition_ARMGenerator()
var repositoryRefDefinition_ARMGenerator gopter.Gen

// RepositoryRefDefinition_ARMGenerator returns a generator of RepositoryRefDefinition_ARM instances for property testing.
func RepositoryRefDefinition_ARMGenerator() gopter.Gen {
	if repositoryRefDefinition_ARMGenerator != nil {
		return repositoryRefDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRepositoryRefDefinition_ARM(generators)
	repositoryRefDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(RepositoryRefDefinition_ARM{}), generators)

	return repositoryRefDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRepositoryRefDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRepositoryRefDefinition_ARM(gens map[string]gopter.Gen) {
	gens["Branch"] = gen.PtrOf(gen.AlphaString())
	gens["Commit"] = gen.PtrOf(gen.AlphaString())
	gens["Semver"] = gen.PtrOf(gen.AlphaString())
	gens["Tag"] = gen.PtrOf(gen.AlphaString())
}

func Test_ServicePrincipalDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServicePrincipalDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServicePrincipalDefinition_ARM, ServicePrincipalDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServicePrincipalDefinition_ARM runs a test to see if a specific instance of ServicePrincipalDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForServicePrincipalDefinition_ARM(subject ServicePrincipalDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServicePrincipalDefinition_ARM
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

// Generator of ServicePrincipalDefinition_ARM instances for property testing - lazily instantiated by
// ServicePrincipalDefinition_ARMGenerator()
var servicePrincipalDefinition_ARMGenerator gopter.Gen

// ServicePrincipalDefinition_ARMGenerator returns a generator of ServicePrincipalDefinition_ARM instances for property testing.
func ServicePrincipalDefinition_ARMGenerator() gopter.Gen {
	if servicePrincipalDefinition_ARMGenerator != nil {
		return servicePrincipalDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServicePrincipalDefinition_ARM(generators)
	servicePrincipalDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(ServicePrincipalDefinition_ARM{}), generators)

	return servicePrincipalDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForServicePrincipalDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServicePrincipalDefinition_ARM(gens map[string]gopter.Gen) {
	gens["ClientCertificate"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificatePassword"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificateSendChain"] = gen.PtrOf(gen.Bool())
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["ClientSecret"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
}

func Test_SubstituteFromDefinition_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubstituteFromDefinition_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubstituteFromDefinition_ARM, SubstituteFromDefinition_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubstituteFromDefinition_ARM runs a test to see if a specific instance of SubstituteFromDefinition_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubstituteFromDefinition_ARM(subject SubstituteFromDefinition_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubstituteFromDefinition_ARM
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

// Generator of SubstituteFromDefinition_ARM instances for property testing - lazily instantiated by
// SubstituteFromDefinition_ARMGenerator()
var substituteFromDefinition_ARMGenerator gopter.Gen

// SubstituteFromDefinition_ARMGenerator returns a generator of SubstituteFromDefinition_ARM instances for property testing.
func SubstituteFromDefinition_ARMGenerator() gopter.Gen {
	if substituteFromDefinition_ARMGenerator != nil {
		return substituteFromDefinition_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubstituteFromDefinition_ARM(generators)
	substituteFromDefinition_ARMGenerator = gen.Struct(reflect.TypeOf(SubstituteFromDefinition_ARM{}), generators)

	return substituteFromDefinition_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubstituteFromDefinition_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubstituteFromDefinition_ARM(gens map[string]gopter.Gen) {
	gens["Kind"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Optional"] = gen.PtrOf(gen.Bool())
}
