// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package ptrace

import "go.opentelemetry.io/collector/pdata/internal"

// ResourceSpansSlice logically represents a slice of ResourceSpans.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewResourceSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpansSlice = internal.ResourceSpansSlice

// NewResourceSpansSlice creates a ResourceSpansSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewResourceSpansSlice = internal.NewResourceSpansSlice

// ResourceSpans is a collection of spans from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceSpans = internal.ResourceSpans 

// NewResourceSpans is an alias for a function to create a new empty ResourceSpans.
var NewResourceSpans = internal.NewResourceSpans

// ScopeSpansSlice logically represents a slice of ScopeSpans.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeSpansSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeSpansSlice = internal.ScopeSpansSlice

// NewScopeSpansSlice creates a ScopeSpansSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewScopeSpansSlice = internal.NewScopeSpansSlice

// ScopeSpans is a collection of spans from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeSpans function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeSpans = internal.ScopeSpans 

// NewScopeSpans is an alias for a function to create a new empty ScopeSpans.
var NewScopeSpans = internal.NewScopeSpans

// SpanSlice logically represents a slice of Span.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewSpanSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanSlice = internal.SpanSlice

// NewSpanSlice creates a SpanSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewSpanSlice = internal.NewSpanSlice

// Span represents a single operation within a trace.
// See Span definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpan function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Span = internal.Span 

// NewSpan is an alias for a function to create a new empty Span.
var NewSpan = internal.NewSpan

// SpanEventSlice logically represents a slice of SpanEvent.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewSpanEventSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEventSlice = internal.SpanEventSlice

// NewSpanEventSlice creates a SpanEventSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewSpanEventSlice = internal.NewSpanEventSlice

// SpanEvent is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs. See OTLP for event definition.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanEvent function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanEvent = internal.SpanEvent 

// NewSpanEvent is an alias for a function to create a new empty SpanEvent.
var NewSpanEvent = internal.NewSpanEvent

// SpanLinkSlice logically represents a slice of SpanLink.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewSpanLinkSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLinkSlice = internal.SpanLinkSlice

// NewSpanLinkSlice creates a SpanLinkSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
var NewSpanLinkSlice = internal.NewSpanLinkSlice

// SpanLink is a pointer from the current span to another span in the same trace or in a
// different trace.
// See Link definition in OTLP: https://github.com/open-telemetry/opentelemetry-proto/blob/main/opentelemetry/proto/trace/v1/trace.proto
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanLink function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanLink = internal.SpanLink 

// NewSpanLink is an alias for a function to create a new empty SpanLink.
var NewSpanLink = internal.NewSpanLink

// SpanStatus is an optional final status for this span. Semantically, when Status was not
// set, that means the span ended without errors and to assume Status.Ok (code = 0).
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewSpanStatus function to create new instances.
// Important: zero-initialized instance is not valid for use.
type SpanStatus = internal.SpanStatus 

// NewSpanStatus is an alias for a function to create a new empty SpanStatus.
var NewSpanStatus = internal.NewSpanStatus


