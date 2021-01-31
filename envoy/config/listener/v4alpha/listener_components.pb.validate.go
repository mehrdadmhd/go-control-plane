// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v4alpha/listener_components.proto

package envoy_config_listener_v4alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _listener_components_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Filter) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return FilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	switch m.ConfigType.(type) {

	case *Filter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on FilterChainMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *FilterChainMatch) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetDestinationPort(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 65535 {
			return FilterChainMatchValidationError{
				field:  "DestinationPort",
				reason: "value must be inside range [1, 65535]",
			}
		}

	}

	for idx, item := range m.GetPrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					field:  fmt.Sprintf("PrefixRanges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AddressSuffix

	if v, ok := interface{}(m.GetSuffixLen()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainMatchValidationError{
				field:  "SuffixLen",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := FilterChainMatch_ConnectionSourceType_name[int32(m.GetSourceType())]; !ok {
		return FilterChainMatchValidationError{
			field:  "SourceType",
			reason: "value must be one of the defined enum values",
		}
	}

	for idx, item := range m.GetSourcePrefixRanges() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainMatchValidationError{
					field:  fmt.Sprintf("SourcePrefixRanges[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSourcePorts() {
		_, _ = idx, item

		if val := item; val < 1 || val > 65535 {
			return FilterChainMatchValidationError{
				field:  fmt.Sprintf("SourcePorts[%v]", idx),
				reason: "value must be inside range [1, 65535]",
			}
		}

	}

	// no validation rules for TransportProtocol

	return nil
}

// FilterChainMatchValidationError is the validation error returned by
// FilterChainMatch.Validate if the designated constraints aren't met.
type FilterChainMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterChainMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterChainMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterChainMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterChainMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterChainMatchValidationError) ErrorName() string { return "FilterChainMatchValidationError" }

// Error satisfies the builtin error interface
func (e FilterChainMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChainMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterChainMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterChainMatchValidationError{}

// Validate checks the field values on FilterChain with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilterChain) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFilterChainMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "FilterChainMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FilterChainValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedUseProxyProto()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "HiddenEnvoyDeprecatedUseProxyProto",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTransportSocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "TransportSocket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTransportSocketConnectTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "TransportSocketConnectTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetOnDemandConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChainValidationError{
				field:  "OnDemandConfiguration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FilterChainValidationError is the validation error returned by
// FilterChain.Validate if the designated constraints aren't met.
type FilterChainValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterChainValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterChainValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterChainValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterChainValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterChainValidationError) ErrorName() string { return "FilterChainValidationError" }

// Error satisfies the builtin error interface
func (e FilterChainValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChain.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterChainValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterChainValidationError{}

// Validate checks the field values on ListenerFilterChainMatchPredicate with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListenerFilterChainMatchPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *ListenerFilterChainMatchPredicate_OrMatch:

		if v, ok := interface{}(m.GetOrMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterChainMatchPredicateValidationError{
					field:  "OrMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListenerFilterChainMatchPredicate_AndMatch:

		if v, ok := interface{}(m.GetAndMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterChainMatchPredicateValidationError{
					field:  "AndMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListenerFilterChainMatchPredicate_NotMatch:

		if v, ok := interface{}(m.GetNotMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterChainMatchPredicateValidationError{
					field:  "NotMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ListenerFilterChainMatchPredicate_AnyMatch:

		if m.GetAnyMatch() != true {
			return ListenerFilterChainMatchPredicateValidationError{
				field:  "AnyMatch",
				reason: "value must equal true",
			}
		}

	case *ListenerFilterChainMatchPredicate_DestinationPortRange:

		if v, ok := interface{}(m.GetDestinationPortRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterChainMatchPredicateValidationError{
					field:  "DestinationPortRange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ListenerFilterChainMatchPredicateValidationError{
			field:  "Rule",
			reason: "value is required",
		}

	}

	return nil
}

// ListenerFilterChainMatchPredicateValidationError is the validation error
// returned by ListenerFilterChainMatchPredicate.Validate if the designated
// constraints aren't met.
type ListenerFilterChainMatchPredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerFilterChainMatchPredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerFilterChainMatchPredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerFilterChainMatchPredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerFilterChainMatchPredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerFilterChainMatchPredicateValidationError) ErrorName() string {
	return "ListenerFilterChainMatchPredicateValidationError"
}

// Error satisfies the builtin error interface
func (e ListenerFilterChainMatchPredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerFilterChainMatchPredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerFilterChainMatchPredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerFilterChainMatchPredicateValidationError{}

// Validate checks the field values on ListenerFilter with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListenerFilter) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return ListenerFilterValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetFilterDisabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerFilterValidationError{
				field:  "FilterDisabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.ConfigType.(type) {

	case *ListenerFilter_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListenerFilterValidationError is the validation error returned by
// ListenerFilter.Validate if the designated constraints aren't met.
type ListenerFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerFilterValidationError) ErrorName() string { return "ListenerFilterValidationError" }

// Error satisfies the builtin error interface
func (e ListenerFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerFilterValidationError{}

// Validate checks the field values on FilterChain_OnDemandConfiguration with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *FilterChain_OnDemandConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRebuildTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FilterChain_OnDemandConfigurationValidationError{
				field:  "RebuildTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FilterChain_OnDemandConfigurationValidationError is the validation error
// returned by FilterChain_OnDemandConfiguration.Validate if the designated
// constraints aren't met.
type FilterChain_OnDemandConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterChain_OnDemandConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterChain_OnDemandConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterChain_OnDemandConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterChain_OnDemandConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterChain_OnDemandConfigurationValidationError) ErrorName() string {
	return "FilterChain_OnDemandConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e FilterChain_OnDemandConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterChain_OnDemandConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterChain_OnDemandConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterChain_OnDemandConfigurationValidationError{}

// Validate checks the field values on
// ListenerFilterChainMatchPredicate_MatchSet with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ListenerFilterChainMatchPredicate_MatchSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 2 {
		return ListenerFilterChainMatchPredicate_MatchSetValidationError{
			field:  "Rules",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerFilterChainMatchPredicate_MatchSetValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListenerFilterChainMatchPredicate_MatchSetValidationError is the validation
// error returned by ListenerFilterChainMatchPredicate_MatchSet.Validate if
// the designated constraints aren't met.
type ListenerFilterChainMatchPredicate_MatchSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) ErrorName() string {
	return "ListenerFilterChainMatchPredicate_MatchSetValidationError"
}

// Error satisfies the builtin error interface
func (e ListenerFilterChainMatchPredicate_MatchSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerFilterChainMatchPredicate_MatchSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerFilterChainMatchPredicate_MatchSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerFilterChainMatchPredicate_MatchSetValidationError{}
