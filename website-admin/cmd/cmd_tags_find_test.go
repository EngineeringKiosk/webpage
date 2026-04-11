package cmd

import (
	"testing"
)

func TestConvertToStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "[]any with strings",
			input:    []any{"Go", "Rust", "Python"},
			expected: []string{"Go", "Rust", "Python"},
		},
		{
			name:     "[]any with mixed types filters non-strings",
			input:    []any{"Go", 42, "Rust", true, "Python"},
			expected: []string{"Go", "Rust", "Python"},
		},
		{
			name:     "[]any empty slice",
			input:    []any{},
			expected: []string{},
		},
		{
			name:     "[]string passthrough",
			input:    []string{"one", "two"},
			expected: []string{"one", "two"},
		},
		{
			name:     "unexpected type returns nil",
			input:    "just a string",
			expected: nil,
		},
		{
			name:     "int type returns nil",
			input:    42,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToStringSlice(tt.input)
			if tt.expected == nil {
				if got != nil {
					t.Errorf("convertToStringSlice(%v) = %v, want nil", tt.input, got)
				}
				return
			}
			if len(got) != len(tt.expected) {
				t.Fatalf("convertToStringSlice(%v) returned %d items, want %d", tt.input, len(got), len(tt.expected))
			}
			for i, want := range tt.expected {
				if got[i] != want {
					t.Errorf("convertToStringSlice(%v)[%d] = %q, want %q", tt.input, i, got[i], want)
				}
			}
		})
	}
}

func TestGetAllTagsWithoutDescription(t *testing.T) {
	tests := []struct {
		name            string
		tagDescriptions map[string]*TagDescription
		tags            map[string]int
		expected        map[string]bool
	}{
		{
			name:            "all tags have complete descriptions",
			tagDescriptions: map[string]*TagDescription{"Go": {ShortDesc: "short", LongDesc: "long"}},
			tags:            map[string]int{"Go": 5},
			expected:        map[string]bool{},
		},
		{
			name:            "tag missing from description file",
			tagDescriptions: map[string]*TagDescription{},
			tags:            map[string]int{"Rust": 3},
			expected:        map[string]bool{"Rust": true},
		},
		{
			name:            "tag with empty ShortDesc",
			tagDescriptions: map[string]*TagDescription{"Go": {ShortDesc: "", LongDesc: "long"}},
			tags:            map[string]int{"Go": 5},
			expected:        map[string]bool{"Go": true},
		},
		{
			name:            "tag with empty LongDesc",
			tagDescriptions: map[string]*TagDescription{"Go": {ShortDesc: "short", LongDesc: ""}},
			tags:            map[string]int{"Go": 5},
			expected:        map[string]bool{"Go": true},
		},
		{
			name:            "both empty inputs",
			tagDescriptions: map[string]*TagDescription{},
			tags:            map[string]int{},
			expected:        map[string]bool{},
		},
		{
			name: "mixed: some described, some missing, some incomplete",
			tagDescriptions: map[string]*TagDescription{
				"Go":     {ShortDesc: "short", LongDesc: "long"},
				"Rust":   {ShortDesc: "", LongDesc: "long"},
				"Python": {ShortDesc: "short", LongDesc: ""},
			},
			tags:     map[string]int{"Go": 5, "Rust": 3, "Python": 2, "Java": 1},
			expected: map[string]bool{"Rust": true, "Python": true, "Java": true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getAllTagsWithoutDescription(tt.tagDescriptions, tt.tags)
			if len(got) != len(tt.expected) {
				t.Fatalf("getAllTagsWithoutDescription() returned %d items, want %d\n  got: %v\n  want: %v", len(got), len(tt.expected), got, tt.expected)
			}
			for tag := range tt.expected {
				if !got[tag] {
					t.Errorf("expected tag %q to be in result, but it was not", tag)
				}
			}
		})
	}
}
