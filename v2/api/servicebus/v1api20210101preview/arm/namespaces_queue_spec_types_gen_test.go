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

func Test_NamespacesQueue_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesQueue_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesQueue_Spec, NamespacesQueue_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesQueue_Spec runs a test to see if a specific instance of NamespacesQueue_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesQueue_Spec(subject NamespacesQueue_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesQueue_Spec
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

// Generator of NamespacesQueue_Spec instances for property testing - lazily instantiated by
// NamespacesQueue_SpecGenerator()
var namespacesQueue_SpecGenerator gopter.Gen

// NamespacesQueue_SpecGenerator returns a generator of NamespacesQueue_Spec instances for property testing.
// We first initialize namespacesQueue_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesQueue_SpecGenerator() gopter.Gen {
	if namespacesQueue_SpecGenerator != nil {
		return namespacesQueue_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_Spec(generators)
	namespacesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesQueue_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesQueue_Spec(generators)
	namespacesQueue_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesQueue_Spec{}), generators)

	return namespacesQueue_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesQueue_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesQueue_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesQueue_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesQueue_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SBQueuePropertiesGenerator())
}

func Test_SBQueueProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SBQueueProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSBQueueProperties, SBQueuePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSBQueueProperties runs a test to see if a specific instance of SBQueueProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForSBQueueProperties(subject SBQueueProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SBQueueProperties
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

// Generator of SBQueueProperties instances for property testing - lazily instantiated by SBQueuePropertiesGenerator()
var sbQueuePropertiesGenerator gopter.Gen

// SBQueuePropertiesGenerator returns a generator of SBQueueProperties instances for property testing.
func SBQueuePropertiesGenerator() gopter.Gen {
	if sbQueuePropertiesGenerator != nil {
		return sbQueuePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSBQueueProperties(generators)
	sbQueuePropertiesGenerator = gen.Struct(reflect.TypeOf(SBQueueProperties{}), generators)

	return sbQueuePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForSBQueueProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSBQueueProperties(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["DeadLetteringOnMessageExpiration"] = gen.PtrOf(gen.Bool())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["ForwardDeadLetteredMessagesTo"] = gen.PtrOf(gen.AlphaString())
	gens["ForwardTo"] = gen.PtrOf(gen.AlphaString())
	gens["LockDuration"] = gen.PtrOf(gen.AlphaString())
	gens["MaxDeliveryCount"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["RequiresSession"] = gen.PtrOf(gen.Bool())
}
