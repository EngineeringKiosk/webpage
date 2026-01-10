package cmd

import "testing"

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
