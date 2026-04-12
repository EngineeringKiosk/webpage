package cmd

import (
	"testing"
)

func TestParseHeadlinesFromHTML(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedHTML      string
		expectedHeadlines []HTMLHeadline
	}{
		{
			name:              "empty string",
			input:             "",
			expectedHTML:      "",
			expectedHeadlines: []HTMLHeadline{},
		},
		{
			name:              "no headlines",
			input:             "<p>Just a paragraph</p>",
			expectedHTML:      "<p>Just a paragraph</p>",
			expectedHeadlines: []HTMLHeadline{},
		},
		{
			name:         "simple h3 without span",
			input:        "<h3>Links</h3>",
			expectedHTML: `<h3 id="links">Links</h3>`,
			expectedHeadlines: []HTMLHeadline{
				{Slug: "links", Text: "Links"},
			},
		},
		{
			name:         "h3 with span wrapper",
			input:        "<h3><span>Sprungmarken</span></h3>",
			expectedHTML: `<h3 id="sprungmarken">Sprungmarken</h3>`,
			expectedHeadlines: []HTMLHeadline{
				{Slug: "sprungmarken", Text: "Sprungmarken"},
			},
		},
		{
			name:         "multiple headlines",
			input:        "<h3>Links</h3><p>Some content</p><h3>Hosts</h3>",
			expectedHTML: `<h3 id="links">Links</h3><p>Some content</p><h3 id="hosts">Hosts</h3>`,
			expectedHeadlines: []HTMLHeadline{
				{Slug: "links", Text: "Links"},
				{Slug: "hosts", Text: "Hosts"},
			},
		},
		{
			name:              "h3 containing HTML tags should be skipped",
			input:             "<h3>Text with <br> line break</h3>",
			expectedHTML:      "<h3>Text with <br> line break</h3>",
			expectedHeadlines: []HTMLHeadline{},
		},
		{
			name:         "headline with surrounding content preserved",
			input:        "<p>Before</p><h3>Title</h3><p>After</p>",
			expectedHTML: `<p>Before</p><h3 id="title">Title</h3><p>After</p>`,
			expectedHeadlines: []HTMLHeadline{
				{Slug: "title", Text: "Title"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHTML, gotHeadlines := parseHeadlinesFromHTML(tt.input)
			if gotHTML != tt.expectedHTML {
				t.Errorf("parseHeadlinesFromHTML(%q)\n  got HTML:  %q\n  want HTML: %q", tt.input, gotHTML, tt.expectedHTML)
			}
			if len(gotHeadlines) != len(tt.expectedHeadlines) {
				t.Fatalf("parseHeadlinesFromHTML(%q)\n  got %d headlines, want %d", tt.input, len(gotHeadlines), len(tt.expectedHeadlines))
			}
			for i, want := range tt.expectedHeadlines {
				got := gotHeadlines[i]
				if got.Slug != want.Slug || got.Text != want.Text {
					t.Errorf("headline[%d] = {Slug: %q, Text: %q}, want {Slug: %q, Text: %q}", i, got.Slug, got.Text, want.Slug, want.Text)
				}
			}
		})
	}
}

func TestGetChapterFromDescription(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []map[string]string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "no chapters",
			input:    "<p>Just a description without timestamps</p>",
			expected: nil,
		},
		{
			name:  "pattern 1 with span wrapper",
			input: "<p><span>(00:01:30) Introduction</span></p><p><span>(00:15:00) Main Topic</span></p>",
			expected: []map[string]string{
				{"start": "00:01:30", "title": "Introduction"},
				{"start": "00:15:00", "title": "Main Topic"},
			},
		},
		{
			name:  "pattern 2 without span",
			input: "<p>(00:00:00) Intro</p><p>(00:05:30) Topic</p>",
			expected: []map[string]string{
				{"start": "00:00:00", "title": "Intro"},
				{"start": "00:05:30", "title": "Topic"},
			},
		},
		{
			name:  "MM:SS format gets prepended with 00:",
			input: "<p>(05:30) Short Format</p>",
			expected: []map[string]string{
				{"start": "00:05:30", "title": "Short Format"},
			},
		},
		{
			name:  "HTML entities in title are unescaped",
			input: "<p><span>(00:10:00) Q&amp;A Session</span></p>",
			expected: []map[string]string{
				{"start": "00:10:00", "title": "Q&A Session"},
			},
		},
		{
			name:  "single chapter",
			input: "<p><span>(00:00:00) Only Chapter</span></p>",
			expected: []map[string]string{
				{"start": "00:00:00", "title": "Only Chapter"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getChapterFromDescription(tt.input)
			if len(got) != len(tt.expected) {
				t.Fatalf("getChapterFromDescription()\n  got %d chapters, want %d\n  got: %v", len(got), len(tt.expected), got)
			}
			for i, want := range tt.expected {
				if got[i]["start"] != want["start"] || got[i]["title"] != want["title"] {
					t.Errorf("chapter[%d] = %v, want %v", i, got[i], want)
				}
			}
		})
	}
}
