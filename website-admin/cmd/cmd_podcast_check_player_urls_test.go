package cmd

import (
	"sort"
	"testing"

	"github.com/EngineeringKiosk/website/website-admin/episode"
)

func TestCheckMissingPlayerURLs(t *testing.T) {
	tests := []struct {
		name     string
		fm       episode.EpisodeFrontmatter
		expected []string
	}{
		{
			name: "all URLs present",
			fm: episode.EpisodeFrontmatter{
				Spotify:      "https://spotify.com/ep1",
				ApplePodcasts: "https://apple.com/ep1",
				AmazonMusic:  "https://amazon.com/ep1",
				Deezer:       "https://deezer.com/ep1",
				YouTube:      "https://youtube.com/ep1",
			},
			expected: nil,
		},
		{
			name: "all URLs missing",
			fm:   episode.EpisodeFrontmatter{},
			expected: []string{
				"amazon_music",
				"apple_podcasts",
				"deezer",
				"spotify",
				"youtube",
			},
		},
		{
			name: "only spotify missing",
			fm: episode.EpisodeFrontmatter{
				ApplePodcasts: "https://apple.com/ep1",
				AmazonMusic:  "https://amazon.com/ep1",
				Deezer:       "https://deezer.com/ep1",
				YouTube:      "https://youtube.com/ep1",
			},
			expected: []string{"spotify"},
		},
		{
			name: "only deezer missing",
			fm: episode.EpisodeFrontmatter{
				Spotify:      "https://spotify.com/ep1",
				ApplePodcasts: "https://apple.com/ep1",
				AmazonMusic:  "https://amazon.com/ep1",
				YouTube:      "https://youtube.com/ep1",
			},
			expected: []string{"deezer"},
		},
		{
			name: "multiple missing",
			fm: episode.EpisodeFrontmatter{
				Spotify: "https://spotify.com/ep1",
				YouTube: "https://youtube.com/ep1",
			},
			expected: []string{
				"amazon_music",
				"apple_podcasts",
				"deezer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkMissingPlayerURLs(tt.fm)

			// Sort both for stable comparison since map iteration is non-deterministic
			sort.Strings(got)
			sort.Strings(tt.expected)

			if len(got) != len(tt.expected) {
				t.Fatalf("checkMissingPlayerURLs() returned %d items %v, want %d items %v", len(got), got, len(tt.expected), tt.expected)
			}
			for i, want := range tt.expected {
				if got[i] != want {
					t.Errorf("checkMissingPlayerURLs()[%d] = %q, want %q", i, got[i], want)
				}
			}
		})
	}
}
