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

func Test_ProductPolicy_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProductPolicy_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProductPolicy_Spec, ProductPolicy_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProductPolicy_Spec runs a test to see if a specific instance of ProductPolicy_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForProductPolicy_Spec(subject ProductPolicy_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProductPolicy_Spec
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

// Generator of ProductPolicy_Spec instances for property testing - lazily instantiated by ProductPolicy_SpecGenerator()
var productPolicy_SpecGenerator gopter.Gen

// ProductPolicy_SpecGenerator returns a generator of ProductPolicy_Spec instances for property testing.
// We first initialize productPolicy_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ProductPolicy_SpecGenerator() gopter.Gen {
	if productPolicy_SpecGenerator != nil {
		return productPolicy_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec(generators)
	productPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProductPolicy_Spec(generators)
	AddRelatedPropertyGeneratorsForProductPolicy_Spec(generators)
	productPolicy_SpecGenerator = gen.Struct(reflect.TypeOf(ProductPolicy_Spec{}), generators)

	return productPolicy_SpecGenerator
}

// AddIndependentPropertyGeneratorsForProductPolicy_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProductPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForProductPolicy_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProductPolicy_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PolicyContractPropertiesGenerator())
}
