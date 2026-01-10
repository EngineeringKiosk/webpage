package episode

import (
	"testing"
)

func TestGetEpisodeNumberFromFilename(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		leadingZero bool
		want        string
		wantErr     bool
	}{
		// Real episode filenames - single digit with leading zero
		{
			name:        "episode 00 with leading zero preserved",
			filename:    "00-developer-fangen-bei-0-an-zu-zählen.md",
			leadingZero: true,
			want:        "00",
			wantErr:     false,
		},
		{
			name:        "episode 00 without leading zero",
			filename:    "00-developer-fangen-bei-0-an-zu-zählen.md",
			leadingZero: false,
			want:        "",
			wantErr:     false,
		},
		{
			name:        "episode 01 with leading zero preserved",
			filename:    "01-side-projects-fluch-oder-segen-für-die-karriere.md",
			leadingZero: true,
			want:        "01",
			wantErr:     false,
		},
		{
			name:        "episode 01 without leading zero",
			filename:    "01-side-projects-fluch-oder-segen-für-die-karriere.md",
			leadingZero: false,
			want:        "1",
			wantErr:     false,
		},
		{
			name:        "episode 09 with leading zero preserved",
			filename:    "09-ukraine.md",
			leadingZero: true,
			want:        "09",
			wantErr:     false,
		},
		{
			name:        "episode 09 without leading zero",
			filename:    "09-ukraine.md",
			leadingZero: false,
			want:        "9",
			wantErr:     false,
		},

		// Real episode filenames - double digit
		{
			name:        "episode 10 with leading zero flag",
			filename:    "10-das-karriere-booster-meeting-11s.md",
			leadingZero: true,
			want:        "10",
			wantErr:     false,
		},
		{
			name:        "episode 10 without leading zero flag",
			filename:    "10-das-karriere-booster-meeting-11s.md",
			leadingZero: false,
			want:        "10",
			wantErr:     false,
		},
		{
			name:        "episode 99 with leading zero flag",
			filename:    "99-modernes-sql-ist-mehr-als-select-from-mit-markus-winand.md",
			leadingZero: true,
			want:        "99",
			wantErr:     false,
		},
		{
			name:        "episode 99 without leading zero flag",
			filename:    "99-modernes-sql-ist-mehr-als-select-from-mit-markus-winand.md",
			leadingZero: false,
			want:        "99",
			wantErr:     false,
		},

		// Real episode filenames - triple digit
		{
			name:        "episode 100 with leading zero flag",
			filename:    "100-episoden-ein-tech-rückblick-auf-202223-predictions-2024-und-viel-tech-trivia.md",
			leadingZero: true,
			want:        "100",
			wantErr:     false,
		},
		{
			name:        "episode 100 without leading zero flag",
			filename:    "100-episoden-ein-tech-rückblick-auf-202223-predictions-2024-und-viel-tech-trivia.md",
			leadingZero: false,
			want:        "100",
			wantErr:     false,
		},
		{
			name:        "episode 101 with leading zero flag",
			filename:    "101-observability-und-opentelemetry-mit-severin-neumann.md",
			leadingZero: true,
			want:        "101",
			wantErr:     false,
		},
		{
			name:        "episode 131 with special characters in title",
			filename:    "131-equity-in-tech-startups-mehr-als-nur-gehalt-mit-philipp-pip-klöckner.md",
			leadingZero: false,
			want:        "131",
			wantErr:     false,
		},
		{
			name:        "episode 170 with 404 in title (bug fix case)",
			filename:    "170-404-not-found.md",
			leadingZero: false,
			want:        "170",
			wantErr:     false,
		},
		{
			name:        "episode 209 high episode number",
			filename:    "209-in-der-besenkammer-jenseits-der-cloud-mittelstands-it-mit-patrick-terlisten.md",
			leadingZero: false,
			want:        "209",
			wantErr:     false,
		},

		// Real episode filenames - negative episode number
		{
			name:        "negative episode -1 with leading zero flag",
			filename:    "-1-wrap-up-2022-und-1-geburtstag-learnings-statistiken-und-was-2023-geplant-ist.md",
			leadingZero: true,
			want:        "-1",
			wantErr:     false,
		},
		{
			name:        "negative episode -1 without leading zero flag",
			filename:    "-1-wrap-up-2022-und-1-geburtstag-learnings-statistiken-und-was-2023-geplant-ist.md",
			leadingZero: false,
			want:        "-1",
			wantErr:     false,
		},

		// Real episode filenames - image files
		{
			name:        "image file episode 05",
			filename:    "05-team-lead-der-einzige-ausweg.jpg",
			leadingZero: true,
			want:        "05",
			wantErr:     false,
		},
		{
			name:        "image file episode 100",
			filename:    "100-episoden-ein-tech-rückblick-auf-202223-predictions-2024-und-viel-tech-trivia.jpg",
			leadingZero: false,
			want:        "100",
			wantErr:     false,
		},

		// Full path handling
		{
			name:        "full path with episode file",
			filename:    "/Users/test/src/content/podcast/42-some-episode-title.md",
			leadingZero: true,
			want:        "42",
			wantErr:     false,
		},
		{
			name:        "relative path with episode file",
			filename:    "src/content/podcast/07-die-freelance-freiheit.md",
			leadingZero: false,
			want:        "7",
			wantErr:     false,
		},

		// Error cases
		{
			name:        "filename without dash",
			filename:    "episode42.md",
			leadingZero: true,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "empty filename",
			filename:    "",
			leadingZero: false,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "only extension",
			filename:    ".md",
			leadingZero: true,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "negative number without second dash",
			filename:    "-42.md",
			leadingZero: true,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "non-numeric prefix with dash",
			filename:    "abc-episode-title.md",
			leadingZero: true,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "filename with only text and dash",
			filename:    "wrap-up-special.md",
			leadingZero: true,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "filename without episode number pattern",
			filename:    "invalid.md",
			leadingZero: false,
			want:        "",
			wantErr:     true,
		},
		{
			name:        "dash only filename",
			filename:    "-",
			leadingZero: false,
			want:        "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEpisodeNumberFromFilename(tt.filename, tt.leadingZero)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpisodeNumberFromFilename() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpisodeNumberFromFilename() = %q, want %q", got, tt.want)
			}
		})
	}
}
