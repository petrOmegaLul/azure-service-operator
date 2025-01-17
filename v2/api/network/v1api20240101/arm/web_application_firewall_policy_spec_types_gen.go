// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type WebApplicationFirewallPolicy_Spec struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the web application firewall policy.
	Properties *WebApplicationFirewallPolicyPropertiesFormat `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &WebApplicationFirewallPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (policy WebApplicationFirewallPolicy_Spec) GetAPIVersion() string {
	return "2024-01-01"
}

// GetName returns the Name of the resource
func (policy *WebApplicationFirewallPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
func (policy *WebApplicationFirewallPolicy_Spec) GetType() string {
	return "Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
}

// Defines web application firewall policy properties.
type WebApplicationFirewallPolicyPropertiesFormat struct {
	// CustomRules: The custom rules inside the policy.
	CustomRules []WebApplicationFirewallCustomRule `json:"customRules,omitempty"`

	// ManagedRules: Describes the managedRules structure.
	ManagedRules *ManagedRulesDefinition `json:"managedRules,omitempty"`

	// PolicySettings: The PolicySettings for policy.
	PolicySettings *PolicySettings `json:"policySettings,omitempty"`
}

// Allow to exclude some variable satisfy the condition for the WAF check.
type ManagedRulesDefinition struct {
	// Exclusions: The Exclusions that are applied on the policy.
	Exclusions []OwaspCrsExclusionEntry `json:"exclusions,omitempty"`

	// ManagedRuleSets: The managed rule sets that are associated with the policy.
	ManagedRuleSets []ManagedRuleSet `json:"managedRuleSets,omitempty"`
}

// Defines contents of a web application firewall global configuration.
type PolicySettings struct {
	// CustomBlockResponseBody: If the action type is block, customer can override the response body. The body must be
	// specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty"`

	// CustomBlockResponseStatusCode: If the action type is block, customer can override the response status code.
	CustomBlockResponseStatusCode *int `json:"customBlockResponseStatusCode,omitempty"`

	// FileUploadEnforcement: Whether allow WAF to enforce file upload limits.
	FileUploadEnforcement *bool `json:"fileUploadEnforcement,omitempty"`

	// FileUploadLimitInMb: Maximum file upload size in Mb for WAF.
	FileUploadLimitInMb *int `json:"fileUploadLimitInMb,omitempty"`

	// JsChallengeCookieExpirationInMins: Web Application Firewall JavaScript Challenge Cookie Expiration time in minutes.
	JsChallengeCookieExpirationInMins *int `json:"jsChallengeCookieExpirationInMins,omitempty"`

	// LogScrubbing: To scrub sensitive log fields
	LogScrubbing *PolicySettings_LogScrubbing `json:"logScrubbing,omitempty"`

	// MaxRequestBodySizeInKb: Maximum request body size in Kb for WAF.
	MaxRequestBodySizeInKb *int `json:"maxRequestBodySizeInKb,omitempty"`

	// Mode: The mode of the policy.
	Mode *PolicySettings_Mode `json:"mode,omitempty"`

	// RequestBodyCheck: Whether to allow WAF to check request Body.
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty"`

	// RequestBodyEnforcement: Whether allow WAF to enforce request body limits.
	RequestBodyEnforcement *bool `json:"requestBodyEnforcement,omitempty"`

	// RequestBodyInspectLimitInKB: Max inspection limit in KB for request body inspection for WAF.
	RequestBodyInspectLimitInKB *int `json:"requestBodyInspectLimitInKB,omitempty"`

	// State: The state of the policy.
	State *PolicySettings_State `json:"state,omitempty"`
}

// Defines contents of a web application rule.
type WebApplicationFirewallCustomRule struct {
	// Action: Type of Actions.
	Action *WebApplicationFirewallCustomRule_Action `json:"action,omitempty"`

	// GroupByUserSession: List of user session identifier group by clauses.
	GroupByUserSession []GroupByUserSession `json:"groupByUserSession,omitempty"`

	// MatchConditions: List of match conditions.
	MatchConditions []MatchCondition `json:"matchConditions,omitempty"`

	// Name: The name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Priority: Priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *int `json:"priority,omitempty"`

	// RateLimitDuration: Duration over which Rate Limit policy will be applied. Applies only when ruleType is RateLimitRule.
	RateLimitDuration *WebApplicationFirewallCustomRule_RateLimitDuration `json:"rateLimitDuration,omitempty"`

	// RateLimitThreshold: Rate Limit threshold to apply in case ruleType is RateLimitRule. Must be greater than or equal to 1
	RateLimitThreshold *int `json:"rateLimitThreshold,omitempty"`

	// RuleType: The rule type.
	RuleType *WebApplicationFirewallCustomRule_RuleType `json:"ruleType,omitempty"`

	// State: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
	State *WebApplicationFirewallCustomRule_State `json:"state,omitempty"`
}

// Define user session identifier group by clauses.
type GroupByUserSession struct {
	// GroupByVariables: List of group by clause variables.
	GroupByVariables []GroupByVariable `json:"groupByVariables,omitempty"`
}

// Defines a managed rule set.
type ManagedRuleSet struct {
	// RuleGroupOverrides: Defines the rule group overrides to apply to the rule set.
	RuleGroupOverrides []ManagedRuleGroupOverride `json:"ruleGroupOverrides,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

// Define match conditions.
type MatchCondition struct {
	// MatchValues: Match value.
	MatchValues []string `json:"matchValues,omitempty"`

	// MatchVariables: List of match variables.
	MatchVariables []MatchVariable `json:"matchVariables,omitempty"`

	// NegationConditon: Whether this is negate condition or not.
	NegationConditon *bool `json:"negationConditon,omitempty"`

	// Operator: The operator to be matched.
	Operator *MatchCondition_Operator `json:"operator,omitempty"`

	// Transforms: List of transforms.
	Transforms []Transform `json:"transforms,omitempty"`
}

// Allow to exclude some variable satisfy the condition for the WAF check.
type OwaspCrsExclusionEntry struct {
	// ExclusionManagedRuleSets: The managed rule sets that are associated with the exclusion.
	ExclusionManagedRuleSets []ExclusionManagedRuleSet `json:"exclusionManagedRuleSets,omitempty"`

	// MatchVariable: The variable to be excluded.
	MatchVariable *OwaspCrsExclusionEntry_MatchVariable `json:"matchVariable,omitempty"`

	// Selector: When matchVariable is a collection, operator used to specify which elements in the collection this exclusion
	// applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: When matchVariable is a collection, operate on the selector to specify which elements in the
	// collection this exclusion applies to.
	SelectorMatchOperator *OwaspCrsExclusionEntry_SelectorMatchOperator `json:"selectorMatchOperator,omitempty"`
}

type PolicySettings_LogScrubbing struct {
	// ScrubbingRules: The rules that are applied to the logs for scrubbing.
	ScrubbingRules []WebApplicationFirewallScrubbingRules `json:"scrubbingRules,omitempty"`

	// State: State of the log scrubbing config. Default value is Enabled.
	State *PolicySettings_LogScrubbing_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"Detection","Prevention"}
type PolicySettings_Mode string

const (
	PolicySettings_Mode_Detection  = PolicySettings_Mode("Detection")
	PolicySettings_Mode_Prevention = PolicySettings_Mode("Prevention")
)

// Mapping from string to PolicySettings_Mode
var policySettings_Mode_Values = map[string]PolicySettings_Mode{
	"detection":  PolicySettings_Mode_Detection,
	"prevention": PolicySettings_Mode_Prevention,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PolicySettings_State string

const (
	PolicySettings_State_Disabled = PolicySettings_State("Disabled")
	PolicySettings_State_Enabled  = PolicySettings_State("Enabled")
)

// Mapping from string to PolicySettings_State
var policySettings_State_Values = map[string]PolicySettings_State{
	"disabled": PolicySettings_State_Disabled,
	"enabled":  PolicySettings_State_Enabled,
}

// +kubebuilder:validation:Enum={"Allow","Block","JSChallenge","Log"}
type WebApplicationFirewallCustomRule_Action string

const (
	WebApplicationFirewallCustomRule_Action_Allow       = WebApplicationFirewallCustomRule_Action("Allow")
	WebApplicationFirewallCustomRule_Action_Block       = WebApplicationFirewallCustomRule_Action("Block")
	WebApplicationFirewallCustomRule_Action_JSChallenge = WebApplicationFirewallCustomRule_Action("JSChallenge")
	WebApplicationFirewallCustomRule_Action_Log         = WebApplicationFirewallCustomRule_Action("Log")
)

// Mapping from string to WebApplicationFirewallCustomRule_Action
var webApplicationFirewallCustomRule_Action_Values = map[string]WebApplicationFirewallCustomRule_Action{
	"allow":       WebApplicationFirewallCustomRule_Action_Allow,
	"block":       WebApplicationFirewallCustomRule_Action_Block,
	"jschallenge": WebApplicationFirewallCustomRule_Action_JSChallenge,
	"log":         WebApplicationFirewallCustomRule_Action_Log,
}

// +kubebuilder:validation:Enum={"FiveMins","OneMin"}
type WebApplicationFirewallCustomRule_RateLimitDuration string

const (
	WebApplicationFirewallCustomRule_RateLimitDuration_FiveMins = WebApplicationFirewallCustomRule_RateLimitDuration("FiveMins")
	WebApplicationFirewallCustomRule_RateLimitDuration_OneMin   = WebApplicationFirewallCustomRule_RateLimitDuration("OneMin")
)

// Mapping from string to WebApplicationFirewallCustomRule_RateLimitDuration
var webApplicationFirewallCustomRule_RateLimitDuration_Values = map[string]WebApplicationFirewallCustomRule_RateLimitDuration{
	"fivemins": WebApplicationFirewallCustomRule_RateLimitDuration_FiveMins,
	"onemin":   WebApplicationFirewallCustomRule_RateLimitDuration_OneMin,
}

// +kubebuilder:validation:Enum={"Invalid","MatchRule","RateLimitRule"}
type WebApplicationFirewallCustomRule_RuleType string

const (
	WebApplicationFirewallCustomRule_RuleType_Invalid       = WebApplicationFirewallCustomRule_RuleType("Invalid")
	WebApplicationFirewallCustomRule_RuleType_MatchRule     = WebApplicationFirewallCustomRule_RuleType("MatchRule")
	WebApplicationFirewallCustomRule_RuleType_RateLimitRule = WebApplicationFirewallCustomRule_RuleType("RateLimitRule")
)

// Mapping from string to WebApplicationFirewallCustomRule_RuleType
var webApplicationFirewallCustomRule_RuleType_Values = map[string]WebApplicationFirewallCustomRule_RuleType{
	"invalid":       WebApplicationFirewallCustomRule_RuleType_Invalid,
	"matchrule":     WebApplicationFirewallCustomRule_RuleType_MatchRule,
	"ratelimitrule": WebApplicationFirewallCustomRule_RuleType_RateLimitRule,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type WebApplicationFirewallCustomRule_State string

const (
	WebApplicationFirewallCustomRule_State_Disabled = WebApplicationFirewallCustomRule_State("Disabled")
	WebApplicationFirewallCustomRule_State_Enabled  = WebApplicationFirewallCustomRule_State("Enabled")
)

// Mapping from string to WebApplicationFirewallCustomRule_State
var webApplicationFirewallCustomRule_State_Values = map[string]WebApplicationFirewallCustomRule_State{
	"disabled": WebApplicationFirewallCustomRule_State_Disabled,
	"enabled":  WebApplicationFirewallCustomRule_State_Enabled,
}

// Defines a managed rule set for Exclusions.
type ExclusionManagedRuleSet struct {
	// RuleGroups: Defines the rule groups to apply to the rule set.
	RuleGroups []ExclusionManagedRuleGroup `json:"ruleGroups,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

// Define user session group by clause variables.
type GroupByVariable struct {
	// VariableName: User Session clause variable.
	VariableName *GroupByVariable_VariableName `json:"variableName,omitempty"`
}

// Defines a managed rule group override setting.
type ManagedRuleGroupOverride struct {
	// RuleGroupName: The managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
	Rules []ManagedRuleOverride `json:"rules,omitempty"`
}

// +kubebuilder:validation:Enum={"Any","BeginsWith","Contains","EndsWith","Equal","GeoMatch","GreaterThan","GreaterThanOrEqual","IPMatch","LessThan","LessThanOrEqual","Regex"}
type MatchCondition_Operator string

const (
	MatchCondition_Operator_Any                = MatchCondition_Operator("Any")
	MatchCondition_Operator_BeginsWith         = MatchCondition_Operator("BeginsWith")
	MatchCondition_Operator_Contains           = MatchCondition_Operator("Contains")
	MatchCondition_Operator_EndsWith           = MatchCondition_Operator("EndsWith")
	MatchCondition_Operator_Equal              = MatchCondition_Operator("Equal")
	MatchCondition_Operator_GeoMatch           = MatchCondition_Operator("GeoMatch")
	MatchCondition_Operator_GreaterThan        = MatchCondition_Operator("GreaterThan")
	MatchCondition_Operator_GreaterThanOrEqual = MatchCondition_Operator("GreaterThanOrEqual")
	MatchCondition_Operator_IPMatch            = MatchCondition_Operator("IPMatch")
	MatchCondition_Operator_LessThan           = MatchCondition_Operator("LessThan")
	MatchCondition_Operator_LessThanOrEqual    = MatchCondition_Operator("LessThanOrEqual")
	MatchCondition_Operator_Regex              = MatchCondition_Operator("Regex")
)

// Mapping from string to MatchCondition_Operator
var matchCondition_Operator_Values = map[string]MatchCondition_Operator{
	"any":                MatchCondition_Operator_Any,
	"beginswith":         MatchCondition_Operator_BeginsWith,
	"contains":           MatchCondition_Operator_Contains,
	"endswith":           MatchCondition_Operator_EndsWith,
	"equal":              MatchCondition_Operator_Equal,
	"geomatch":           MatchCondition_Operator_GeoMatch,
	"greaterthan":        MatchCondition_Operator_GreaterThan,
	"greaterthanorequal": MatchCondition_Operator_GreaterThanOrEqual,
	"ipmatch":            MatchCondition_Operator_IPMatch,
	"lessthan":           MatchCondition_Operator_LessThan,
	"lessthanorequal":    MatchCondition_Operator_LessThanOrEqual,
	"regex":              MatchCondition_Operator_Regex,
}

// Define match variables.
type MatchVariable struct {
	// Selector: The selector of match variable.
	Selector *string `json:"selector,omitempty"`

	// VariableName: Match Variable.
	VariableName *MatchVariable_VariableName `json:"variableName,omitempty"`
}

// +kubebuilder:validation:Enum={"RequestArgKeys","RequestArgNames","RequestArgValues","RequestCookieKeys","RequestCookieNames","RequestCookieValues","RequestHeaderKeys","RequestHeaderNames","RequestHeaderValues"}
type OwaspCrsExclusionEntry_MatchVariable string

const (
	OwaspCrsExclusionEntry_MatchVariable_RequestArgKeys      = OwaspCrsExclusionEntry_MatchVariable("RequestArgKeys")
	OwaspCrsExclusionEntry_MatchVariable_RequestArgNames     = OwaspCrsExclusionEntry_MatchVariable("RequestArgNames")
	OwaspCrsExclusionEntry_MatchVariable_RequestArgValues    = OwaspCrsExclusionEntry_MatchVariable("RequestArgValues")
	OwaspCrsExclusionEntry_MatchVariable_RequestCookieKeys   = OwaspCrsExclusionEntry_MatchVariable("RequestCookieKeys")
	OwaspCrsExclusionEntry_MatchVariable_RequestCookieNames  = OwaspCrsExclusionEntry_MatchVariable("RequestCookieNames")
	OwaspCrsExclusionEntry_MatchVariable_RequestCookieValues = OwaspCrsExclusionEntry_MatchVariable("RequestCookieValues")
	OwaspCrsExclusionEntry_MatchVariable_RequestHeaderKeys   = OwaspCrsExclusionEntry_MatchVariable("RequestHeaderKeys")
	OwaspCrsExclusionEntry_MatchVariable_RequestHeaderNames  = OwaspCrsExclusionEntry_MatchVariable("RequestHeaderNames")
	OwaspCrsExclusionEntry_MatchVariable_RequestHeaderValues = OwaspCrsExclusionEntry_MatchVariable("RequestHeaderValues")
)

// Mapping from string to OwaspCrsExclusionEntry_MatchVariable
var owaspCrsExclusionEntry_MatchVariable_Values = map[string]OwaspCrsExclusionEntry_MatchVariable{
	"requestargkeys":      OwaspCrsExclusionEntry_MatchVariable_RequestArgKeys,
	"requestargnames":     OwaspCrsExclusionEntry_MatchVariable_RequestArgNames,
	"requestargvalues":    OwaspCrsExclusionEntry_MatchVariable_RequestArgValues,
	"requestcookiekeys":   OwaspCrsExclusionEntry_MatchVariable_RequestCookieKeys,
	"requestcookienames":  OwaspCrsExclusionEntry_MatchVariable_RequestCookieNames,
	"requestcookievalues": OwaspCrsExclusionEntry_MatchVariable_RequestCookieValues,
	"requestheaderkeys":   OwaspCrsExclusionEntry_MatchVariable_RequestHeaderKeys,
	"requestheadernames":  OwaspCrsExclusionEntry_MatchVariable_RequestHeaderNames,
	"requestheadervalues": OwaspCrsExclusionEntry_MatchVariable_RequestHeaderValues,
}

// +kubebuilder:validation:Enum={"Contains","EndsWith","Equals","EqualsAny","StartsWith"}
type OwaspCrsExclusionEntry_SelectorMatchOperator string

const (
	OwaspCrsExclusionEntry_SelectorMatchOperator_Contains   = OwaspCrsExclusionEntry_SelectorMatchOperator("Contains")
	OwaspCrsExclusionEntry_SelectorMatchOperator_EndsWith   = OwaspCrsExclusionEntry_SelectorMatchOperator("EndsWith")
	OwaspCrsExclusionEntry_SelectorMatchOperator_Equals     = OwaspCrsExclusionEntry_SelectorMatchOperator("Equals")
	OwaspCrsExclusionEntry_SelectorMatchOperator_EqualsAny  = OwaspCrsExclusionEntry_SelectorMatchOperator("EqualsAny")
	OwaspCrsExclusionEntry_SelectorMatchOperator_StartsWith = OwaspCrsExclusionEntry_SelectorMatchOperator("StartsWith")
)

// Mapping from string to OwaspCrsExclusionEntry_SelectorMatchOperator
var owaspCrsExclusionEntry_SelectorMatchOperator_Values = map[string]OwaspCrsExclusionEntry_SelectorMatchOperator{
	"contains":   OwaspCrsExclusionEntry_SelectorMatchOperator_Contains,
	"endswith":   OwaspCrsExclusionEntry_SelectorMatchOperator_EndsWith,
	"equals":     OwaspCrsExclusionEntry_SelectorMatchOperator_Equals,
	"equalsany":  OwaspCrsExclusionEntry_SelectorMatchOperator_EqualsAny,
	"startswith": OwaspCrsExclusionEntry_SelectorMatchOperator_StartsWith,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PolicySettings_LogScrubbing_State string

const (
	PolicySettings_LogScrubbing_State_Disabled = PolicySettings_LogScrubbing_State("Disabled")
	PolicySettings_LogScrubbing_State_Enabled  = PolicySettings_LogScrubbing_State("Enabled")
)

// Mapping from string to PolicySettings_LogScrubbing_State
var policySettings_LogScrubbing_State_Values = map[string]PolicySettings_LogScrubbing_State{
	"disabled": PolicySettings_LogScrubbing_State_Disabled,
	"enabled":  PolicySettings_LogScrubbing_State_Enabled,
}

// Transforms applied before matching.
// +kubebuilder:validation:Enum={"HtmlEntityDecode","Lowercase","RemoveNulls","Trim","Uppercase","UrlDecode","UrlEncode"}
type Transform string

const (
	Transform_HtmlEntityDecode = Transform("HtmlEntityDecode")
	Transform_Lowercase        = Transform("Lowercase")
	Transform_RemoveNulls      = Transform("RemoveNulls")
	Transform_Trim             = Transform("Trim")
	Transform_Uppercase        = Transform("Uppercase")
	Transform_UrlDecode        = Transform("UrlDecode")
	Transform_UrlEncode        = Transform("UrlEncode")
)

// Mapping from string to Transform
var transform_Values = map[string]Transform{
	"htmlentitydecode": Transform_HtmlEntityDecode,
	"lowercase":        Transform_Lowercase,
	"removenulls":      Transform_RemoveNulls,
	"trim":             Transform_Trim,
	"uppercase":        Transform_Uppercase,
	"urldecode":        Transform_UrlDecode,
	"urlencode":        Transform_UrlEncode,
}

// Allow certain variables to be scrubbed on WAF logs
type WebApplicationFirewallScrubbingRules struct {
	// MatchVariable: The variable to be scrubbed from the logs.
	MatchVariable *WebApplicationFirewallScrubbingRules_MatchVariable `json:"matchVariable,omitempty"`

	// Selector: When matchVariable is a collection, operator used to specify which elements in the collection this rule
	// applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: When matchVariable is a collection, operate on the selector to specify which elements in the
	// collection this rule applies to.
	SelectorMatchOperator *WebApplicationFirewallScrubbingRules_SelectorMatchOperator `json:"selectorMatchOperator,omitempty"`

	// State: Defines the state of log scrubbing rule. Default value is Enabled.
	State *WebApplicationFirewallScrubbingRules_State `json:"state,omitempty"`
}

// Defines a managed rule group to use for exclusion.
type ExclusionManagedRuleGroup struct {
	// RuleGroupName: The managed rule group for exclusion.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be excluded. If none specified, all rules in the group will be excluded.
	Rules []ExclusionManagedRule `json:"rules,omitempty"`
}

// +kubebuilder:validation:Enum={"ClientAddr","GeoLocation","None"}
type GroupByVariable_VariableName string

const (
	GroupByVariable_VariableName_ClientAddr  = GroupByVariable_VariableName("ClientAddr")
	GroupByVariable_VariableName_GeoLocation = GroupByVariable_VariableName("GeoLocation")
	GroupByVariable_VariableName_None        = GroupByVariable_VariableName("None")
)

// Mapping from string to GroupByVariable_VariableName
var groupByVariable_VariableName_Values = map[string]GroupByVariable_VariableName{
	"clientaddr":  GroupByVariable_VariableName_ClientAddr,
	"geolocation": GroupByVariable_VariableName_GeoLocation,
	"none":        GroupByVariable_VariableName_None,
}

// Defines a managed rule group override setting.
type ManagedRuleOverride struct {
	// Action: Describes the override action to be applied when rule matches.
	Action *ActionType `json:"action,omitempty"`

	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`

	// State: The state of the managed rule. Defaults to Disabled if not specified.
	State *ManagedRuleOverride_State `json:"state,omitempty"`
}

// +kubebuilder:validation:Enum={"PostArgs","QueryString","RemoteAddr","RequestBody","RequestCookies","RequestHeaders","RequestMethod","RequestUri"}
type MatchVariable_VariableName string

const (
	MatchVariable_VariableName_PostArgs       = MatchVariable_VariableName("PostArgs")
	MatchVariable_VariableName_QueryString    = MatchVariable_VariableName("QueryString")
	MatchVariable_VariableName_RemoteAddr     = MatchVariable_VariableName("RemoteAddr")
	MatchVariable_VariableName_RequestBody    = MatchVariable_VariableName("RequestBody")
	MatchVariable_VariableName_RequestCookies = MatchVariable_VariableName("RequestCookies")
	MatchVariable_VariableName_RequestHeaders = MatchVariable_VariableName("RequestHeaders")
	MatchVariable_VariableName_RequestMethod  = MatchVariable_VariableName("RequestMethod")
	MatchVariable_VariableName_RequestUri     = MatchVariable_VariableName("RequestUri")
)

// Mapping from string to MatchVariable_VariableName
var matchVariable_VariableName_Values = map[string]MatchVariable_VariableName{
	"postargs":       MatchVariable_VariableName_PostArgs,
	"querystring":    MatchVariable_VariableName_QueryString,
	"remoteaddr":     MatchVariable_VariableName_RemoteAddr,
	"requestbody":    MatchVariable_VariableName_RequestBody,
	"requestcookies": MatchVariable_VariableName_RequestCookies,
	"requestheaders": MatchVariable_VariableName_RequestHeaders,
	"requestmethod":  MatchVariable_VariableName_RequestMethod,
	"requesturi":     MatchVariable_VariableName_RequestUri,
}

// +kubebuilder:validation:Enum={"RequestArgNames","RequestCookieNames","RequestHeaderNames","RequestIPAddress","RequestJSONArgNames","RequestPostArgNames"}
type WebApplicationFirewallScrubbingRules_MatchVariable string

const (
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestArgNames     = WebApplicationFirewallScrubbingRules_MatchVariable("RequestArgNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestCookieNames  = WebApplicationFirewallScrubbingRules_MatchVariable("RequestCookieNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestHeaderNames  = WebApplicationFirewallScrubbingRules_MatchVariable("RequestHeaderNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestIPAddress    = WebApplicationFirewallScrubbingRules_MatchVariable("RequestIPAddress")
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestJSONArgNames = WebApplicationFirewallScrubbingRules_MatchVariable("RequestJSONArgNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_RequestPostArgNames = WebApplicationFirewallScrubbingRules_MatchVariable("RequestPostArgNames")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_MatchVariable
var webApplicationFirewallScrubbingRules_MatchVariable_Values = map[string]WebApplicationFirewallScrubbingRules_MatchVariable{
	"requestargnames":     WebApplicationFirewallScrubbingRules_MatchVariable_RequestArgNames,
	"requestcookienames":  WebApplicationFirewallScrubbingRules_MatchVariable_RequestCookieNames,
	"requestheadernames":  WebApplicationFirewallScrubbingRules_MatchVariable_RequestHeaderNames,
	"requestipaddress":    WebApplicationFirewallScrubbingRules_MatchVariable_RequestIPAddress,
	"requestjsonargnames": WebApplicationFirewallScrubbingRules_MatchVariable_RequestJSONArgNames,
	"requestpostargnames": WebApplicationFirewallScrubbingRules_MatchVariable_RequestPostArgNames,
}

// +kubebuilder:validation:Enum={"Equals","EqualsAny"}
type WebApplicationFirewallScrubbingRules_SelectorMatchOperator string

const (
	WebApplicationFirewallScrubbingRules_SelectorMatchOperator_Equals    = WebApplicationFirewallScrubbingRules_SelectorMatchOperator("Equals")
	WebApplicationFirewallScrubbingRules_SelectorMatchOperator_EqualsAny = WebApplicationFirewallScrubbingRules_SelectorMatchOperator("EqualsAny")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_SelectorMatchOperator
var webApplicationFirewallScrubbingRules_SelectorMatchOperator_Values = map[string]WebApplicationFirewallScrubbingRules_SelectorMatchOperator{
	"equals":    WebApplicationFirewallScrubbingRules_SelectorMatchOperator_Equals,
	"equalsany": WebApplicationFirewallScrubbingRules_SelectorMatchOperator_EqualsAny,
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type WebApplicationFirewallScrubbingRules_State string

const (
	WebApplicationFirewallScrubbingRules_State_Disabled = WebApplicationFirewallScrubbingRules_State("Disabled")
	WebApplicationFirewallScrubbingRules_State_Enabled  = WebApplicationFirewallScrubbingRules_State("Enabled")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_State
var webApplicationFirewallScrubbingRules_State_Values = map[string]WebApplicationFirewallScrubbingRules_State{
	"disabled": WebApplicationFirewallScrubbingRules_State_Disabled,
	"enabled":  WebApplicationFirewallScrubbingRules_State_Enabled,
}

// Defines the action to take on rule match.
// +kubebuilder:validation:Enum={"Allow","AnomalyScoring","Block","JSChallenge","Log"}
type ActionType string

const (
	ActionType_Allow          = ActionType("Allow")
	ActionType_AnomalyScoring = ActionType("AnomalyScoring")
	ActionType_Block          = ActionType("Block")
	ActionType_JSChallenge    = ActionType("JSChallenge")
	ActionType_Log            = ActionType("Log")
)

// Mapping from string to ActionType
var actionType_Values = map[string]ActionType{
	"allow":          ActionType_Allow,
	"anomalyscoring": ActionType_AnomalyScoring,
	"block":          ActionType_Block,
	"jschallenge":    ActionType_JSChallenge,
	"log":            ActionType_Log,
}

// Defines a managed rule to use for exclusion.
type ExclusionManagedRule struct {
	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ManagedRuleOverride_State string

const (
	ManagedRuleOverride_State_Disabled = ManagedRuleOverride_State("Disabled")
	ManagedRuleOverride_State_Enabled  = ManagedRuleOverride_State("Enabled")
)

// Mapping from string to ManagedRuleOverride_State
var managedRuleOverride_State_Values = map[string]ManagedRuleOverride_State{
	"disabled": ManagedRuleOverride_State_Disabled,
	"enabled":  ManagedRuleOverride_State_Enabled,
}
