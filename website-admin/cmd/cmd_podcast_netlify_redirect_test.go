package cmd

import "testing"

func TestRedirectExistence_BothExist(t *testing.T) {
	tests := []struct {
		name     string
		input    RedirectExistence
		expected bool
	}{
		{
			name:     "both redirects exist",
			input:    RedirectExistence{EpisodesRedirect: true, EpRedirect: true},
			expected: true,
		},
		{
			name:     "only episodes redirect exists",
			input:    RedirectExistence{EpisodesRedirect: true, EpRedirect: false},
			expected: false,
		},
		{
			name:     "only ep redirect exists",
			input:    RedirectExistence{EpisodesRedirect: false, EpRedirect: true},
			expected: false,
		},
		{
			name:     "neither redirect exists",
			input:    RedirectExistence{EpisodesRedirect: false, EpRedirect: false},
			expected: false,
		},
		{
			name:     "zero value struct",
			input:    RedirectExistence{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.BothExist()
			if got != tt.expected {
				t.Errorf("RedirectExistence%+v.BothExist() = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestCheckBothRedirectsExist(t *testing.T) {
	tests := []struct {
		name              string
		redirectExistence map[string]RedirectExistence
		episodeNumber     string
		expected          bool
	}{
		{
			name: "episode has both redirects",
			redirectExistence: map[string]RedirectExistence{
				"42": {EpisodesRedirect: true, EpRedirect: true},
			},
			episodeNumber: "42",
			expected:      true,
		},
		{
			name: "episode has only episodes redirect",
			redirectExistence: map[string]RedirectExistence{
				"42": {EpisodesRedirect: true, EpRedirect: false},
			},
			episodeNumber: "42",
			expected:      false,
		},
		{
			name: "episode has only ep redirect",
			redirectExistence: map[string]RedirectExistence{
				"42": {EpisodesRedirect: false, EpRedirect: true},
			},
			episodeNumber: "42",
			expected:      false,
		},
		{
			name: "episode has neither redirect",
			redirectExistence: map[string]RedirectExistence{
				"42": {EpisodesRedirect: false, EpRedirect: false},
			},
			episodeNumber: "42",
			expected:      false,
		},
		{
			name: "episode not in map",
			redirectExistence: map[string]RedirectExistence{
				"1": {EpisodesRedirect: true, EpRedirect: true},
			},
			episodeNumber: "42",
			expected:      false,
		},
		{
			name:              "empty map",
			redirectExistence: map[string]RedirectExistence{},
			episodeNumber:     "42",
			expected:          false,
		},
		{
			name:              "nil map",
			redirectExistence: nil,
			episodeNumber:     "42",
			expected:          false,
		},
		{
			name: "multiple episodes with mixed states",
			redirectExistence: map[string]RedirectExistence{
				"1":   {EpisodesRedirect: true, EpRedirect: true},
				"42":  {EpisodesRedirect: true, EpRedirect: false},
				"100": {EpisodesRedirect: false, EpRedirect: true},
				"200": {EpisodesRedirect: true, EpRedirect: true},
			},
			episodeNumber: "42",
			expected:      false,
		},
		{
			name: "episode 0",
			redirectExistence: map[string]RedirectExistence{
				"0": {EpisodesRedirect: true, EpRedirect: true},
			},
			episodeNumber: "0",
			expected:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckBothRedirectsExist(tt.redirectExistence, tt.episodeNumber)
			if got != tt.expected {
				t.Errorf("CheckBothRedirectsExist() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestExtractEpisodeNumber(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantNumber string
		wantOk     bool
	}{
		{
			name:       "episode 170 with 404 in title (bug fix case)",
			filename:   "170-404-not-found.md",
			wantNumber: "170",
			wantOk:     true,
		},
		{
			name:       "episode with leading zero",
			filename:   "01-side-projects-fluch-oder-segen-für-die-karriere.md",
			wantNumber: "1",
			wantOk:     true,
		},
		{
			name:       "episode zero with double zero prefix",
			filename:   "00-developer-fangen-bei-0-an-zu-zählen.md",
			wantNumber: "0",
			wantOk:     true,
		},
		{
			name:       "episode 99 two digit number",
			filename:   "99-modernes-sql-ist-mehr-als-select-from-mit-markus-winand.md",
			wantNumber: "99",
			wantOk:     true,
		},
		{
			name:       "episode 100 three digit number",
			filename:   "100-episoden-ein-tech-rückblick-auf-202223-predictions-2024-und-viel-tech-trivia.md",
			wantNumber: "100",
			wantOk:     true,
		},
		{
			name:       "episode 131 with special characters in title",
			filename:   "131-equity-in-tech-startups-mehr-als-nur-gehalt-mit-philipp-pip-klöckner.md",
			wantNumber: "131",
			wantOk:     true,
		},
		{
			name:       "episode 209 high episode number",
			filename:   "209-in-der-besenkammer-jenseits-der-cloud-mittelstands-it-mit-patrick-terlisten.md",
			wantNumber: "209",
			wantOk:     true,
		},
		{
			name:       "negative episode number should not match",
			filename:   "-1-wrap-up-2022-und-1-geburtstag-learnings-statistiken-und-was-2023-geplant-ist.md",
			wantNumber: "",
			wantOk:     false,
		},
		{
			name:       "filename without episode number pattern",
			filename:   "invalid.md",
			wantNumber: "",
			wantOk:     false,
		},
		{
			name:       "filename with only text and dash",
			filename:   "wrap-up-special.md",
			wantNumber: "",
			wantOk:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumber, gotOk := ExtractEpisodeNumber(tt.filename)
			if gotNumber != tt.wantNumber {
				t.Errorf("ExtractEpisodeNumber(%q) number = %q, want %q", tt.filename, gotNumber, tt.wantNumber)
			}
			if gotOk != tt.wantOk {
				t.Errorf("ExtractEpisodeNumber(%q) ok = %v, want %v", tt.filename, gotOk, tt.wantOk)
			}
		})
	}
}
