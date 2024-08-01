// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302/storage"
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

func Test_DiskAccess_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DiskAccess to hub returns original",
		prop.ForAll(RunResourceConversionTestForDiskAccess, DiskAccessGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForDiskAccess tests if a specific instance of DiskAccess round trips to the hub storage version and back losslessly
func RunResourceConversionTestForDiskAccess(subject DiskAccess) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.DiskAccess
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual DiskAccess
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DiskAccess_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DiskAccess to DiskAccess via AssignProperties_To_DiskAccess & AssignProperties_From_DiskAccess returns original",
		prop.ForAll(RunPropertyAssignmentTestForDiskAccess, DiskAccessGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDiskAccess tests if a specific instance of DiskAccess can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDiskAccess(subject DiskAccess) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DiskAccess
	err := copied.AssignProperties_To_DiskAccess(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DiskAccess
	err = actual.AssignProperties_From_DiskAccess(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DiskAccess_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess, DiskAccessGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess runs a test to see if a specific instance of DiskAccess round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess(subject DiskAccess) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess
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

// Generator of DiskAccess instances for property testing - lazily instantiated by DiskAccessGenerator()
var diskAccessGenerator gopter.Gen

// DiskAccessGenerator returns a generator of DiskAccess instances for property testing.
func DiskAccessGenerator() gopter.Gen {
	if diskAccessGenerator != nil {
		return diskAccessGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDiskAccess(generators)
	diskAccessGenerator = gen.Struct(reflect.TypeOf(DiskAccess{}), generators)

	return diskAccessGenerator
}

// AddRelatedPropertyGeneratorsForDiskAccess is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess(gens map[string]gopter.Gen) {
	gens["Spec"] = DiskAccess_SpecGenerator()
	gens["Status"] = DiskAccess_STATUSGenerator()
}

func Test_DiskAccess_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DiskAccess_STATUS to DiskAccess_STATUS via AssignProperties_To_DiskAccess_STATUS & AssignProperties_From_DiskAccess_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForDiskAccess_STATUS, DiskAccess_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDiskAccess_STATUS tests if a specific instance of DiskAccess_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDiskAccess_STATUS(subject DiskAccess_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DiskAccess_STATUS
	err := copied.AssignProperties_To_DiskAccess_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DiskAccess_STATUS
	err = actual.AssignProperties_From_DiskAccess_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DiskAccess_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess_STATUS, DiskAccess_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess_STATUS runs a test to see if a specific instance of DiskAccess_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess_STATUS(subject DiskAccess_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess_STATUS
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

// Generator of DiskAccess_STATUS instances for property testing - lazily instantiated by DiskAccess_STATUSGenerator()
var diskAccess_STATUSGenerator gopter.Gen

// DiskAccess_STATUSGenerator returns a generator of DiskAccess_STATUS instances for property testing.
// We first initialize diskAccess_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskAccess_STATUSGenerator() gopter.Gen {
	if diskAccess_STATUSGenerator != nil {
		return diskAccess_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_STATUS(generators)
	diskAccess_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskAccess_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiskAccess_STATUS(generators)
	diskAccess_STATUSGenerator = gen.Struct(reflect.TypeOf(DiskAccess_STATUS{}), generators)

	return diskAccess_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiskAccess_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskAccess_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskAccess_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess_STATUS(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnection_STATUSGenerator())
}

func Test_DiskAccess_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DiskAccess_Spec to DiskAccess_Spec via AssignProperties_To_DiskAccess_Spec & AssignProperties_From_DiskAccess_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDiskAccess_Spec, DiskAccess_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDiskAccess_Spec tests if a specific instance of DiskAccess_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForDiskAccess_Spec(subject DiskAccess_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.DiskAccess_Spec
	err := copied.AssignProperties_To_DiskAccess_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DiskAccess_Spec
	err = actual.AssignProperties_From_DiskAccess_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DiskAccess_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskAccess_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskAccess_Spec, DiskAccess_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskAccess_Spec runs a test to see if a specific instance of DiskAccess_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskAccess_Spec(subject DiskAccess_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskAccess_Spec
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

// Generator of DiskAccess_Spec instances for property testing - lazily instantiated by DiskAccess_SpecGenerator()
var diskAccess_SpecGenerator gopter.Gen

// DiskAccess_SpecGenerator returns a generator of DiskAccess_Spec instances for property testing.
// We first initialize diskAccess_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskAccess_SpecGenerator() gopter.Gen {
	if diskAccess_SpecGenerator != nil {
		return diskAccess_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec(generators)
	diskAccess_SpecGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskAccess_Spec(generators)
	AddRelatedPropertyGeneratorsForDiskAccess_Spec(generators)
	diskAccess_SpecGenerator = gen.Struct(reflect.TypeOf(DiskAccess_Spec{}), generators)

	return diskAccess_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDiskAccess_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskAccess_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskAccess_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskAccess_Spec(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
}

func Test_PrivateEndpointConnection_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointConnection_STATUS to PrivateEndpointConnection_STATUS via AssignProperties_To_PrivateEndpointConnection_STATUS & AssignProperties_From_PrivateEndpointConnection_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS tests if a specific instance of PrivateEndpointConnection_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.PrivateEndpointConnection_STATUS
	err := copied.AssignProperties_To_PrivateEndpointConnection_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointConnection_STATUS
	err = actual.AssignProperties_From_PrivateEndpointConnection_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PrivateEndpointConnection_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnection_STATUS, PrivateEndpointConnection_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnection_STATUS runs a test to see if a specific instance of PrivateEndpointConnection_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnection_STATUS(subject PrivateEndpointConnection_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS
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

// Generator of PrivateEndpointConnection_STATUS instances for property testing - lazily instantiated by
// PrivateEndpointConnection_STATUSGenerator()
var privateEndpointConnection_STATUSGenerator gopter.Gen

// PrivateEndpointConnection_STATUSGenerator returns a generator of PrivateEndpointConnection_STATUS instances for property testing.
func PrivateEndpointConnection_STATUSGenerator() gopter.Gen {
	if privateEndpointConnection_STATUSGenerator != nil {
		return privateEndpointConnection_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(generators)
	privateEndpointConnection_STATUSGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS{}), generators)

	return privateEndpointConnection_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnection_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
