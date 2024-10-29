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

func Test_WorkspaceConnectionProps_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspaceConnectionProps via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspaceConnectionProps, WorkspaceConnectionPropsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspaceConnectionProps runs a test to see if a specific instance of WorkspaceConnectionProps round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspaceConnectionProps(subject WorkspaceConnectionProps) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspaceConnectionProps
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

// Generator of WorkspaceConnectionProps instances for property testing - lazily instantiated by
// WorkspaceConnectionPropsGenerator()
var workspaceConnectionPropsGenerator gopter.Gen

// WorkspaceConnectionPropsGenerator returns a generator of WorkspaceConnectionProps instances for property testing.
func WorkspaceConnectionPropsGenerator() gopter.Gen {
	if workspaceConnectionPropsGenerator != nil {
		return workspaceConnectionPropsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspaceConnectionProps(generators)
	workspaceConnectionPropsGenerator = gen.Struct(reflect.TypeOf(WorkspaceConnectionProps{}), generators)

	return workspaceConnectionPropsGenerator
}

// AddIndependentPropertyGeneratorsForWorkspaceConnectionProps is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspaceConnectionProps(gens map[string]gopter.Gen) {
	gens["AuthType"] = gen.PtrOf(gen.AlphaString())
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
	gens["ValueFormat"] = gen.PtrOf(gen.OneConstOf(WorkspaceConnectionProps_ValueFormat_JSON))
}

func Test_WorkspacesConnection_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of WorkspacesConnection_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForWorkspacesConnection_Spec, WorkspacesConnection_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForWorkspacesConnection_Spec runs a test to see if a specific instance of WorkspacesConnection_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForWorkspacesConnection_Spec(subject WorkspacesConnection_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual WorkspacesConnection_Spec
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

// Generator of WorkspacesConnection_Spec instances for property testing - lazily instantiated by
// WorkspacesConnection_SpecGenerator()
var workspacesConnection_SpecGenerator gopter.Gen

// WorkspacesConnection_SpecGenerator returns a generator of WorkspacesConnection_Spec instances for property testing.
// We first initialize workspacesConnection_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func WorkspacesConnection_SpecGenerator() gopter.Gen {
	if workspacesConnection_SpecGenerator != nil {
		return workspacesConnection_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnection_Spec(generators)
	workspacesConnection_SpecGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnection_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForWorkspacesConnection_Spec(generators)
	AddRelatedPropertyGeneratorsForWorkspacesConnection_Spec(generators)
	workspacesConnection_SpecGenerator = gen.Struct(reflect.TypeOf(WorkspacesConnection_Spec{}), generators)

	return workspacesConnection_SpecGenerator
}

// AddIndependentPropertyGeneratorsForWorkspacesConnection_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForWorkspacesConnection_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForWorkspacesConnection_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForWorkspacesConnection_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(WorkspaceConnectionPropsGenerator())
}
