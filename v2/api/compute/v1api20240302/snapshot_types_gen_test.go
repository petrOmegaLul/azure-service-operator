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

func Test_CopyCompletionError_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CopyCompletionError to CopyCompletionError via AssignProperties_To_CopyCompletionError & AssignProperties_From_CopyCompletionError returns original",
		prop.ForAll(RunPropertyAssignmentTestForCopyCompletionError, CopyCompletionErrorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCopyCompletionError tests if a specific instance of CopyCompletionError can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForCopyCompletionError(subject CopyCompletionError) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.CopyCompletionError
	err := copied.AssignProperties_To_CopyCompletionError(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CopyCompletionError
	err = actual.AssignProperties_From_CopyCompletionError(&other)
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

func Test_CopyCompletionError_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CopyCompletionError via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCopyCompletionError, CopyCompletionErrorGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCopyCompletionError runs a test to see if a specific instance of CopyCompletionError round trips to JSON and back losslessly
func RunJSONSerializationTestForCopyCompletionError(subject CopyCompletionError) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CopyCompletionError
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

// Generator of CopyCompletionError instances for property testing - lazily instantiated by
// CopyCompletionErrorGenerator()
var copyCompletionErrorGenerator gopter.Gen

// CopyCompletionErrorGenerator returns a generator of CopyCompletionError instances for property testing.
func CopyCompletionErrorGenerator() gopter.Gen {
	if copyCompletionErrorGenerator != nil {
		return copyCompletionErrorGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCopyCompletionError(generators)
	copyCompletionErrorGenerator = gen.Struct(reflect.TypeOf(CopyCompletionError{}), generators)

	return copyCompletionErrorGenerator
}

// AddIndependentPropertyGeneratorsForCopyCompletionError is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCopyCompletionError(gens map[string]gopter.Gen) {
	gens["ErrorCode"] = gen.PtrOf(gen.OneConstOf(CopyCompletionError_ErrorCode_CopySourceNotFound))
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
}

func Test_CopyCompletionError_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from CopyCompletionError_STATUS to CopyCompletionError_STATUS via AssignProperties_To_CopyCompletionError_STATUS & AssignProperties_From_CopyCompletionError_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForCopyCompletionError_STATUS, CopyCompletionError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForCopyCompletionError_STATUS tests if a specific instance of CopyCompletionError_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForCopyCompletionError_STATUS(subject CopyCompletionError_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.CopyCompletionError_STATUS
	err := copied.AssignProperties_To_CopyCompletionError_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual CopyCompletionError_STATUS
	err = actual.AssignProperties_From_CopyCompletionError_STATUS(&other)
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

func Test_CopyCompletionError_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CopyCompletionError_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCopyCompletionError_STATUS, CopyCompletionError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCopyCompletionError_STATUS runs a test to see if a specific instance of CopyCompletionError_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForCopyCompletionError_STATUS(subject CopyCompletionError_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CopyCompletionError_STATUS
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

// Generator of CopyCompletionError_STATUS instances for property testing - lazily instantiated by
// CopyCompletionError_STATUSGenerator()
var copyCompletionError_STATUSGenerator gopter.Gen

// CopyCompletionError_STATUSGenerator returns a generator of CopyCompletionError_STATUS instances for property testing.
func CopyCompletionError_STATUSGenerator() gopter.Gen {
	if copyCompletionError_STATUSGenerator != nil {
		return copyCompletionError_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS(generators)
	copyCompletionError_STATUSGenerator = gen.Struct(reflect.TypeOf(CopyCompletionError_STATUS{}), generators)

	return copyCompletionError_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCopyCompletionError_STATUS(gens map[string]gopter.Gen) {
	gens["ErrorCode"] = gen.PtrOf(gen.OneConstOf(CopyCompletionError_ErrorCode_STATUS_CopySourceNotFound))
	gens["ErrorMessage"] = gen.PtrOf(gen.AlphaString())
}

func Test_Snapshot_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot to hub returns original",
		prop.ForAll(RunResourceConversionTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSnapshot tests if a specific instance of Snapshot round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSnapshot(subject Snapshot) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.Snapshot
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Snapshot
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

func Test_Snapshot_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot to Snapshot via AssignProperties_To_Snapshot & AssignProperties_From_Snapshot returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshot tests if a specific instance of Snapshot can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSnapshot(subject Snapshot) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Snapshot
	err := copied.AssignProperties_To_Snapshot(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshot
	err = actual.AssignProperties_From_Snapshot(&other)
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

func Test_Snapshot_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot, SnapshotGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot runs a test to see if a specific instance of Snapshot round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot(subject Snapshot) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot
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

// Generator of Snapshot instances for property testing - lazily instantiated by SnapshotGenerator()
var snapshotGenerator gopter.Gen

// SnapshotGenerator returns a generator of Snapshot instances for property testing.
func SnapshotGenerator() gopter.Gen {
	if snapshotGenerator != nil {
		return snapshotGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSnapshot(generators)
	snapshotGenerator = gen.Struct(reflect.TypeOf(Snapshot{}), generators)

	return snapshotGenerator
}

// AddRelatedPropertyGeneratorsForSnapshot is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot(gens map[string]gopter.Gen) {
	gens["Spec"] = Snapshot_SpecGenerator()
	gens["Status"] = Snapshot_STATUSGenerator()
}

func Test_SnapshotSku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SnapshotSku to SnapshotSku via AssignProperties_To_SnapshotSku & AssignProperties_From_SnapshotSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotSku tests if a specific instance of SnapshotSku can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSnapshotSku(subject SnapshotSku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SnapshotSku
	err := copied.AssignProperties_To_SnapshotSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SnapshotSku
	err = actual.AssignProperties_From_SnapshotSku(&other)
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

func Test_SnapshotSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku, SnapshotSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku runs a test to see if a specific instance of SnapshotSku round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku(subject SnapshotSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku
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

// Generator of SnapshotSku instances for property testing - lazily instantiated by SnapshotSkuGenerator()
var snapshotSkuGenerator gopter.Gen

// SnapshotSkuGenerator returns a generator of SnapshotSku instances for property testing.
func SnapshotSkuGenerator() gopter.Gen {
	if snapshotSkuGenerator != nil {
		return snapshotSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku(generators)
	snapshotSkuGenerator = gen.Struct(reflect.TypeOf(SnapshotSku{}), generators)

	return snapshotSkuGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSku_Name_Premium_LRS, SnapshotSku_Name_Standard_LRS, SnapshotSku_Name_Standard_ZRS))
}

func Test_SnapshotSku_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SnapshotSku_STATUS to SnapshotSku_STATUS via AssignProperties_To_SnapshotSku_STATUS & AssignProperties_From_SnapshotSku_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshotSku_STATUS, SnapshotSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshotSku_STATUS tests if a specific instance of SnapshotSku_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSnapshotSku_STATUS(subject SnapshotSku_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SnapshotSku_STATUS
	err := copied.AssignProperties_To_SnapshotSku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SnapshotSku_STATUS
	err = actual.AssignProperties_From_SnapshotSku_STATUS(&other)
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

func Test_SnapshotSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SnapshotSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshotSku_STATUS, SnapshotSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshotSku_STATUS runs a test to see if a specific instance of SnapshotSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshotSku_STATUS(subject SnapshotSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SnapshotSku_STATUS
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

// Generator of SnapshotSku_STATUS instances for property testing - lazily instantiated by SnapshotSku_STATUSGenerator()
var snapshotSku_STATUSGenerator gopter.Gen

// SnapshotSku_STATUSGenerator returns a generator of SnapshotSku_STATUS instances for property testing.
func SnapshotSku_STATUSGenerator() gopter.Gen {
	if snapshotSku_STATUSGenerator != nil {
		return snapshotSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshotSku_STATUS(generators)
	snapshotSku_STATUSGenerator = gen.Struct(reflect.TypeOf(SnapshotSku_STATUS{}), generators)

	return snapshotSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshotSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshotSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(SnapshotSku_Name_STATUS_Premium_LRS, SnapshotSku_Name_STATUS_Standard_LRS, SnapshotSku_Name_STATUS_Standard_ZRS))
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_Snapshot_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot_STATUS to Snapshot_STATUS via AssignProperties_To_Snapshot_STATUS & AssignProperties_From_Snapshot_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshot_STATUS, Snapshot_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshot_STATUS tests if a specific instance of Snapshot_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSnapshot_STATUS(subject Snapshot_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Snapshot_STATUS
	err := copied.AssignProperties_To_Snapshot_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshot_STATUS
	err = actual.AssignProperties_From_Snapshot_STATUS(&other)
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

func Test_Snapshot_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_STATUS, Snapshot_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_STATUS runs a test to see if a specific instance of Snapshot_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_STATUS(subject Snapshot_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_STATUS
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

// Generator of Snapshot_STATUS instances for property testing - lazily instantiated by Snapshot_STATUSGenerator()
var snapshot_STATUSGenerator gopter.Gen

// Snapshot_STATUSGenerator returns a generator of Snapshot_STATUS instances for property testing.
// We first initialize snapshot_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_STATUSGenerator() gopter.Gen {
	if snapshot_STATUSGenerator != nil {
		return snapshot_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUS(generators)
	snapshot_STATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_STATUS(generators)
	AddRelatedPropertyGeneratorsForSnapshot_STATUS(generators)
	snapshot_STATUSGenerator = gen.Struct(reflect.TypeOf(Snapshot_STATUS{}), generators)

	return snapshot_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_STATUS(gens map[string]gopter.Gen) {
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.OneConstOf(DataAccessAuthMode_STATUS_AzureActiveDirectory, DataAccessAuthMode_STATUS_None))
	gens["DiskAccessId"] = gen.PtrOf(gen.AlphaString())
	gens["DiskSizeBytes"] = gen.PtrOf(gen.Int())
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_STATUS_ActiveSAS,
		DiskState_STATUS_ActiveSASFrozen,
		DiskState_STATUS_ActiveUpload,
		DiskState_STATUS_Attached,
		DiskState_STATUS_Frozen,
		DiskState_STATUS_ReadyToUpload,
		DiskState_STATUS_Reserved,
		DiskState_STATUS_Unattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_HyperVGeneration_STATUS_V1, SnapshotProperties_HyperVGeneration_STATUS_V2))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["IncrementalSnapshotFamilyId"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["ManagedBy"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_STATUS_AllowAll, NetworkAccessPolicy_STATUS_AllowPrivate, NetworkAccessPolicy_STATUS_DenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_OsType_STATUS_Linux, SnapshotProperties_OsType_STATUS_Windows))
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccess_STATUS_Disabled, PublicNetworkAccess_STATUS_Enabled))
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["TimeCreated"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UniqueId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_STATUS(gens map[string]gopter.Gen) {
	gens["CopyCompletionError"] = gen.PtrOf(CopyCompletionError_STATUSGenerator())
	gens["CreationData"] = gen.PtrOf(CreationData_STATUSGenerator())
	gens["Encryption"] = gen.PtrOf(Encryption_STATUSGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollection_STATUSGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlan_STATUSGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfile_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSku_STATUSGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilities_STATUSGenerator())
}

func Test_Snapshot_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Snapshot_Spec to Snapshot_Spec via AssignProperties_To_Snapshot_Spec & AssignProperties_From_Snapshot_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSnapshot_Spec, Snapshot_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSnapshot_Spec tests if a specific instance of Snapshot_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSnapshot_Spec(subject Snapshot_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.Snapshot_Spec
	err := copied.AssignProperties_To_Snapshot_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Snapshot_Spec
	err = actual.AssignProperties_From_Snapshot_Spec(&other)
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

func Test_Snapshot_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Snapshot_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSnapshot_Spec, Snapshot_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSnapshot_Spec runs a test to see if a specific instance of Snapshot_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSnapshot_Spec(subject Snapshot_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Snapshot_Spec
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

// Generator of Snapshot_Spec instances for property testing - lazily instantiated by Snapshot_SpecGenerator()
var snapshot_SpecGenerator gopter.Gen

// Snapshot_SpecGenerator returns a generator of Snapshot_Spec instances for property testing.
// We first initialize snapshot_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Snapshot_SpecGenerator() gopter.Gen {
	if snapshot_SpecGenerator != nil {
		return snapshot_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec(generators)
	snapshot_SpecGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSnapshot_Spec(generators)
	AddRelatedPropertyGeneratorsForSnapshot_Spec(generators)
	snapshot_SpecGenerator = gen.Struct(reflect.TypeOf(Snapshot_Spec{}), generators)

	return snapshot_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSnapshot_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSnapshot_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["CompletionPercent"] = gen.PtrOf(gen.Float64())
	gens["DataAccessAuthMode"] = gen.PtrOf(gen.OneConstOf(DataAccessAuthMode_AzureActiveDirectory, DataAccessAuthMode_None))
	gens["DiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["DiskState"] = gen.PtrOf(gen.OneConstOf(
		DiskState_ActiveSAS,
		DiskState_ActiveSASFrozen,
		DiskState_ActiveUpload,
		DiskState_Attached,
		DiskState_Frozen,
		DiskState_ReadyToUpload,
		DiskState_Reserved,
		DiskState_Unattached))
	gens["HyperVGeneration"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_HyperVGeneration_V1, SnapshotProperties_HyperVGeneration_V2))
	gens["Incremental"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["NetworkAccessPolicy"] = gen.PtrOf(gen.OneConstOf(NetworkAccessPolicy_AllowAll, NetworkAccessPolicy_AllowPrivate, NetworkAccessPolicy_DenyAll))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(SnapshotProperties_OsType_Linux, SnapshotProperties_OsType_Windows))
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.OneConstOf(PublicNetworkAccess_Disabled, PublicNetworkAccess_Enabled))
	gens["SupportsHibernation"] = gen.PtrOf(gen.Bool())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSnapshot_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSnapshot_Spec(gens map[string]gopter.Gen) {
	gens["CopyCompletionError"] = gen.PtrOf(CopyCompletionErrorGenerator())
	gens["CreationData"] = gen.PtrOf(CreationDataGenerator())
	gens["Encryption"] = gen.PtrOf(EncryptionGenerator())
	gens["EncryptionSettingsCollection"] = gen.PtrOf(EncryptionSettingsCollectionGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["PurchasePlan"] = gen.PtrOf(PurchasePlanGenerator())
	gens["SecurityProfile"] = gen.PtrOf(DiskSecurityProfileGenerator())
	gens["Sku"] = gen.PtrOf(SnapshotSkuGenerator())
	gens["SupportedCapabilities"] = gen.PtrOf(SupportedCapabilitiesGenerator())
}
