// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

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

func Test_MongodbDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabase, MongodbDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabase runs a test to see if a specific instance of MongodbDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabase(subject MongodbDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabase
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

// Generator of MongodbDatabase instances for property testing - lazily instantiated by MongodbDatabaseGenerator()
var mongodbDatabaseGenerator gopter.Gen

// MongodbDatabaseGenerator returns a generator of MongodbDatabase instances for property testing.
func MongodbDatabaseGenerator() gopter.Gen {
	if mongodbDatabaseGenerator != nil {
		return mongodbDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMongodbDatabase(generators)
	mongodbDatabaseGenerator = gen.Struct(reflect.TypeOf(MongodbDatabase{}), generators)

	return mongodbDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForMongodbDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccounts_MongodbDatabase_SpecGenerator()
	gens["Status"] = DatabaseAccounts_MongodbDatabase_STATUSGenerator()
}

func Test_DatabaseAccounts_MongodbDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec, DatabaseAccounts_MongodbDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_Spec(subject DatabaseAccounts_MongodbDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabase_Spec
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

// Generator of DatabaseAccounts_MongodbDatabase_Spec instances for property testing - lazily instantiated by
// DatabaseAccounts_MongodbDatabase_SpecGenerator()
var databaseAccounts_MongodbDatabase_SpecGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabase_SpecGenerator returns a generator of DatabaseAccounts_MongodbDatabase_Spec instances for property testing.
// We first initialize databaseAccounts_MongodbDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabase_SpecGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabase_SpecGenerator != nil {
		return databaseAccounts_MongodbDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec(generators)
	databaseAccounts_MongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec(generators)
	databaseAccounts_MongodbDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_Spec{}), generators)

	return databaseAccounts_MongodbDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_Spec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseResourceGenerator())
}

func Test_DatabaseAccounts_MongodbDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccounts_MongodbDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_STATUS, DatabaseAccounts_MongodbDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_STATUS runs a test to see if a specific instance of DatabaseAccounts_MongodbDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccounts_MongodbDatabase_STATUS(subject DatabaseAccounts_MongodbDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccounts_MongodbDatabase_STATUS
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

// Generator of DatabaseAccounts_MongodbDatabase_STATUS instances for property testing - lazily instantiated by
// DatabaseAccounts_MongodbDatabase_STATUSGenerator()
var databaseAccounts_MongodbDatabase_STATUSGenerator gopter.Gen

// DatabaseAccounts_MongodbDatabase_STATUSGenerator returns a generator of DatabaseAccounts_MongodbDatabase_STATUS instances for property testing.
// We first initialize databaseAccounts_MongodbDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccounts_MongodbDatabase_STATUSGenerator() gopter.Gen {
	if databaseAccounts_MongodbDatabase_STATUSGenerator != nil {
		return databaseAccounts_MongodbDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS(generators)
	databaseAccounts_MongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS(generators)
	databaseAccounts_MongodbDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(DatabaseAccounts_MongodbDatabase_STATUS{}), generators)

	return databaseAccounts_MongodbDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccounts_MongodbDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(MongoDBDatabaseGetProperties_Resource_STATUSGenerator())
}

func Test_CreateUpdateOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CreateUpdateOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCreateUpdateOptions, CreateUpdateOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCreateUpdateOptions runs a test to see if a specific instance of CreateUpdateOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForCreateUpdateOptions(subject CreateUpdateOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CreateUpdateOptions
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

// Generator of CreateUpdateOptions instances for property testing - lazily instantiated by
// CreateUpdateOptionsGenerator()
var createUpdateOptionsGenerator gopter.Gen

// CreateUpdateOptionsGenerator returns a generator of CreateUpdateOptions instances for property testing.
// We first initialize createUpdateOptionsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CreateUpdateOptionsGenerator() gopter.Gen {
	if createUpdateOptionsGenerator != nil {
		return createUpdateOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCreateUpdateOptions(generators)
	AddRelatedPropertyGeneratorsForCreateUpdateOptions(generators)
	createUpdateOptionsGenerator = gen.Struct(reflect.TypeOf(CreateUpdateOptions{}), generators)

	return createUpdateOptionsGenerator
}

// AddIndependentPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForCreateUpdateOptions is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCreateUpdateOptions(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsGenerator())
}

func Test_MongoDBDatabaseGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS, MongoDBDatabaseGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS runs a test to see if a specific instance of MongoDBDatabaseGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseGetProperties_Resource_STATUS(subject MongoDBDatabaseGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseGetProperties_Resource_STATUS
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

// Generator of MongoDBDatabaseGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// MongoDBDatabaseGetProperties_Resource_STATUSGenerator()
var mongoDBDatabaseGetProperties_Resource_STATUSGenerator gopter.Gen

// MongoDBDatabaseGetProperties_Resource_STATUSGenerator returns a generator of MongoDBDatabaseGetProperties_Resource_STATUS instances for property testing.
// We first initialize mongoDBDatabaseGetProperties_Resource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBDatabaseGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if mongoDBDatabaseGetProperties_Resource_STATUSGenerator != nil {
		return mongoDBDatabaseGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	mongoDBDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(generators)
	mongoDBDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseGetProperties_Resource_STATUS{}), generators)

	return mongoDBDatabaseGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["RestoreParameters"] = gen.PtrOf(RestoreParametersBase_STATUSGenerator())
}

func Test_MongoDBDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongoDBDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongoDBDatabaseResource, MongoDBDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongoDBDatabaseResource runs a test to see if a specific instance of MongoDBDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForMongoDBDatabaseResource(subject MongoDBDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongoDBDatabaseResource
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

// Generator of MongoDBDatabaseResource instances for property testing - lazily instantiated by
// MongoDBDatabaseResourceGenerator()
var mongoDBDatabaseResourceGenerator gopter.Gen

// MongoDBDatabaseResourceGenerator returns a generator of MongoDBDatabaseResource instances for property testing.
// We first initialize mongoDBDatabaseResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongoDBDatabaseResourceGenerator() gopter.Gen {
	if mongoDBDatabaseResourceGenerator != nil {
		return mongoDBDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(generators)
	mongoDBDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(generators)
	AddRelatedPropertyGeneratorsForMongoDBDatabaseResource(generators)
	mongoDBDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(MongoDBDatabaseResource{}), generators)

	return mongoDBDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForMongoDBDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongoDBDatabaseResource(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongoDBDatabaseResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongoDBDatabaseResource(gens map[string]gopter.Gen) {
	gens["RestoreParameters"] = gen.PtrOf(RestoreParametersBaseGenerator())
}

func Test_OptionsResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OptionsResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOptionsResource_STATUS, OptionsResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOptionsResource_STATUS runs a test to see if a specific instance of OptionsResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForOptionsResource_STATUS(subject OptionsResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OptionsResource_STATUS
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

// Generator of OptionsResource_STATUS instances for property testing - lazily instantiated by
// OptionsResource_STATUSGenerator()
var optionsResource_STATUSGenerator gopter.Gen

// OptionsResource_STATUSGenerator returns a generator of OptionsResource_STATUS instances for property testing.
// We first initialize optionsResource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OptionsResource_STATUSGenerator() gopter.Gen {
	if optionsResource_STATUSGenerator != nil {
		return optionsResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUS(generators)
	optionsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOptionsResource_STATUS(generators)
	AddRelatedPropertyGeneratorsForOptionsResource_STATUS(generators)
	optionsResource_STATUSGenerator = gen.Struct(reflect.TypeOf(OptionsResource_STATUS{}), generators)

	return optionsResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForOptionsResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOptionsResource_STATUS(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForOptionsResource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOptionsResource_STATUS(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettings_STATUSGenerator())
}

func Test_AutoscaleSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings, AutoscaleSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings runs a test to see if a specific instance of AutoscaleSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings(subject AutoscaleSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings
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

// Generator of AutoscaleSettings instances for property testing - lazily instantiated by AutoscaleSettingsGenerator()
var autoscaleSettingsGenerator gopter.Gen

// AutoscaleSettingsGenerator returns a generator of AutoscaleSettings instances for property testing.
func AutoscaleSettingsGenerator() gopter.Gen {
	if autoscaleSettingsGenerator != nil {
		return autoscaleSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings(generators)
	autoscaleSettingsGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings{}), generators)

	return autoscaleSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_AutoscaleSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettings_STATUS, AutoscaleSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettings_STATUS runs a test to see if a specific instance of AutoscaleSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettings_STATUS(subject AutoscaleSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettings_STATUS
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

// Generator of AutoscaleSettings_STATUS instances for property testing - lazily instantiated by
// AutoscaleSettings_STATUSGenerator()
var autoscaleSettings_STATUSGenerator gopter.Gen

// AutoscaleSettings_STATUSGenerator returns a generator of AutoscaleSettings_STATUS instances for property testing.
func AutoscaleSettings_STATUSGenerator() gopter.Gen {
	if autoscaleSettings_STATUSGenerator != nil {
		return autoscaleSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(generators)
	autoscaleSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettings_STATUS{}), generators)

	return autoscaleSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettings_STATUS(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

func Test_RestoreParametersBase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestoreParametersBase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestoreParametersBase, RestoreParametersBaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestoreParametersBase runs a test to see if a specific instance of RestoreParametersBase round trips to JSON and back losslessly
func RunJSONSerializationTestForRestoreParametersBase(subject RestoreParametersBase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestoreParametersBase
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

// Generator of RestoreParametersBase instances for property testing - lazily instantiated by
// RestoreParametersBaseGenerator()
var restoreParametersBaseGenerator gopter.Gen

// RestoreParametersBaseGenerator returns a generator of RestoreParametersBase instances for property testing.
func RestoreParametersBaseGenerator() gopter.Gen {
	if restoreParametersBaseGenerator != nil {
		return restoreParametersBaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestoreParametersBase(generators)
	restoreParametersBaseGenerator = gen.Struct(reflect.TypeOf(RestoreParametersBase{}), generators)

	return restoreParametersBaseGenerator
}

// AddIndependentPropertyGeneratorsForRestoreParametersBase is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestoreParametersBase(gens map[string]gopter.Gen) {
	gens["RestoreSource"] = gen.PtrOf(gen.AlphaString())
	gens["RestoreTimestampInUtc"] = gen.PtrOf(gen.AlphaString())
}

func Test_RestoreParametersBase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RestoreParametersBase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRestoreParametersBase_STATUS, RestoreParametersBase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRestoreParametersBase_STATUS runs a test to see if a specific instance of RestoreParametersBase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRestoreParametersBase_STATUS(subject RestoreParametersBase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RestoreParametersBase_STATUS
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

// Generator of RestoreParametersBase_STATUS instances for property testing - lazily instantiated by
// RestoreParametersBase_STATUSGenerator()
var restoreParametersBase_STATUSGenerator gopter.Gen

// RestoreParametersBase_STATUSGenerator returns a generator of RestoreParametersBase_STATUS instances for property testing.
func RestoreParametersBase_STATUSGenerator() gopter.Gen {
	if restoreParametersBase_STATUSGenerator != nil {
		return restoreParametersBase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRestoreParametersBase_STATUS(generators)
	restoreParametersBase_STATUSGenerator = gen.Struct(reflect.TypeOf(RestoreParametersBase_STATUS{}), generators)

	return restoreParametersBase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRestoreParametersBase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRestoreParametersBase_STATUS(gens map[string]gopter.Gen) {
	gens["RestoreSource"] = gen.PtrOf(gen.AlphaString())
	gens["RestoreTimestampInUtc"] = gen.PtrOf(gen.AlphaString())
}
