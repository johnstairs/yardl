// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package dsl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeParsing_Valid(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "Foo", expected: `{"name":"Foo"}`},
		{input: "Foo ", expected: `{"name":"Foo"}`},
		{input: " Foo ", expected: `{"name":"Foo"}`},
		{input: "Foo?", expected: `{"name":"Foo","optional":true}`},
		{input: "Foo ? ", expected: `{"name":"Foo","optional":true}`},
		{input: "Foo<int>", expected: `{"name":"Foo","args":[{"name":"int"}]}`},
		{input: "Foo< int >", expected: `{"name":"Foo","args":[{"name":"int"}]}`},
		{input: "Foo<int?>", expected: `{"name":"Foo","args":[{"name":"int","optional":true}]}`},
		{input: "Foo<int>?", expected: `{"name":"Foo","args":[{"name":"int"}],"optional":true}`},
		{input: "Foo<int,float>", expected: `{"name":"Foo","args":[{"name":"int"},{"name":"float"}]}`},
		{input: " Foo < int , float > ", expected: `{"name":"Foo","args":[{"name":"int"},{"name":"float"}]}`},
		{input: "Foo<Bar<int>>", expected: `{"name":"Foo","args":[{"name":"Bar","args":[{"name":"int"}]}]}`},
		{input: "Foo<Bar<int>,Baz<long>>", expected: `{"name":"Foo","args":[{"name":"Bar","args":[{"name":"int"}]},{"name":"Baz","args":[{"name":"long"}]}]}`},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			ts, err := parseSimpleTypeString(tc.input)
			assert.Nil(t, err)
			bytes, err := json.Marshal(ts)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected, string(bytes))
		})
	}
}

func TestTypeParsing_Invalid(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{input: "Foo<", expected: `missing '>'`},
		{input: "Foo?<int>", expected: `'?' at position 4 must appear after generic type arguments`},
		{input: "Foo??", expected: `unexpected trailing '?' in type string`},
		{input: "Foo<int>>", expected: `unexpected trailing '>' in type string`},
		{input: "Foo<int>,bar>", expected: `unexpected trailing ',bar>' in type string`},
		{input: "<int>", expected: `the type name cannot be empty`},
		{input: "Foo<>", expected: `the type parameter name cannot be empty at position 5`},
		{input: "Foo<int,>", expected: `the type parameter name cannot be empty at position 9`},
	}
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			_, err := parseSimpleTypeString(tc.input)
			assert.ErrorContains(t, err, tc.expected)
		})
	}
}
