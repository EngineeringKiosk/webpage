package utils

import "testing"

func TestRemoveHTMLTags(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "plain text without tags",
			input:    "Hello World",
			expected: "Hello World",
		},
		{
			name:     "simple paragraph",
			input:    "<p>Hello World</p>",
			expected: "Hello World ",
		},
		{
			name:     "closing p tag adds space",
			input:    "<p>First</p><p>Second</p>",
			expected: "First Second ",
		},
		{
			name:     "nested tags",
			input:    "<div><p><strong>Bold</strong> text</p></div>",
			expected: "Bold text ",
		},
		{
			name:     "self-closing tags",
			input:    "Line one<br/>Line two",
			expected: "Line oneLine two",
		},
		{
			name:     "tags with attributes",
			input:    `<a href="https://example.com" class="link">Click here</a>`,
			expected: "Click here",
		},
		{
			name:     "multiple p tags for readability",
			input:    "<p>Paragraph one.</p><p>Paragraph two.</p><p>Paragraph three.</p>",
			expected: "Paragraph one. Paragraph two. Paragraph three. ",
		},
		{
			name:     "img tag removed entirely",
			input:    `<p>Text <img src="photo.jpg" alt="photo"> more text</p>`,
			expected: "Text  more text ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveHTMLTags(tt.input)
			if got != tt.expected {
				t.Errorf("RemoveHTMLTags(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestRemoveRelNofollowFromInternalLinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "no links",
			input:    "Just plain text",
			expected: "Just plain text",
		},
		{
			name:     "engineeringkiosk.dev link with nofollow removed",
			input:    `<a href="https://engineeringkiosk.dev/podcast/" rel="nofollow">Podcast</a>`,
			expected: `<a href="https://engineeringkiosk.dev/podcast/">Podcast</a>`,
		},
		{
			name:     "jump.engineeringkiosk.dev link with nofollow removed",
			input:    `<a href="https://jump.engineeringkiosk.dev/some-link" rel="nofollow">Link</a>`,
			expected: `<a href="https://jump.engineeringkiosk.dev/some-link">Link</a>`,
		},
		{
			name:     "external link nofollow preserved",
			input:    `<a href="https://example.com/page" rel="nofollow">External</a>`,
			expected: `<a href="https://example.com/page" rel="nofollow">External</a>`,
		},
		{
			name:     "internal link without nofollow unchanged",
			input:    `<a href="https://engineeringkiosk.dev/blog/">Blog</a>`,
			expected: `<a href="https://engineeringkiosk.dev/blog/">Blog</a>`,
		},
		{
			name:     "mixed internal and external links",
			input:    `<a href="https://engineeringkiosk.dev/ep1" rel="nofollow">Ep1</a> and <a href="https://github.com/test" rel="nofollow">GH</a>`,
			expected: `<a href="https://engineeringkiosk.dev/ep1">Ep1</a> and <a href="https://github.com/test" rel="nofollow">GH</a>`,
		},
		{
			name:     "multiple internal links",
			input:    `<a href="https://engineeringkiosk.dev/a" rel="nofollow">A</a> <a href="https://jump.engineeringkiosk.dev/b" rel="nofollow">B</a>`,
			expected: `<a href="https://engineeringkiosk.dev/a">A</a> <a href="https://jump.engineeringkiosk.dev/b">B</a>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveRelNofollowFromInternalLinks(tt.input)
			if got != tt.expected {
				t.Errorf("RemoveRelNofollowFromInternalLinks(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
