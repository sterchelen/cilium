// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/sni_cluster/v3/sni_cluster.proto

package sni_clusterv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SniCluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SniCluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SniCluster with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SniClusterMultiError, or
// nil if none found.
func (m *SniCluster) ValidateAll() error {
	return m.validate(true)
}

func (m *SniCluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SniClusterMultiError(errors)
	}
	return nil
}

// SniClusterMultiError is an error wrapping multiple validation errors
// returned by SniCluster.ValidateAll() if the designated constraints aren't met.
type SniClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SniClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SniClusterMultiError) AllErrors() []error { return m }

// SniClusterValidationError is the validation error returned by
// SniCluster.Validate if the designated constraints aren't met.
type SniClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SniClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SniClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SniClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SniClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SniClusterValidationError) ErrorName() string { return "SniClusterValidationError" }

// Error satisfies the builtin error interface
func (e SniClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSniCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SniClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SniClusterValidationError{}