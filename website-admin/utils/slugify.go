package utils

import (
	"fmt"
	"strings"

	"github.com/gosimple/slug"
)

// Slugify converts a string into a URL-friendly slug.
// This is the exact same logic as used in our website (originally from Django, https://docs.djangoproject.com/en/2.1/_modules/django/utils/text/#slugify).
// The file names in https://github.com/EngineeringKiosk/webpage/tree/main/src/content/podcast
// are generated using this slug version.
// This is the Go version of https://github.com/EngineeringKiosk/webpage/blob/706159224361c91ac718095b3098dd5d49bcabc7/scripts/podcast_feed_to_content.py#L53-L67
func Slugify(value string) string {
	// Some episodes start with a negative number.
	// The slug library strips the negative sign (-), as this is the separator character.
	// We need to preserve it, so we can add it back in later.
	// Example: #-1 Wrap Up 2022 und 1. Geburtstag: Learnings, Statistiken und was 2023 geplant ist
	preserveMinusChar := false
	if strings.HasPrefix(value, "#-") {
		preserveMinusChar = true
	}

	// A few characters are replaced differently in the two slug libraries.
	// We
	slug.CustomSub = map[string]string{
		// Characters we just want to skip
		// Example: #177 Stream Processing & Kafka: Die Basis moderner Datenpipelines mit Stefan Sprenger
		"&": "",
		// Example: #10 Das Karriere Booster Meeting 1:1s
		":": "",
		// Example: #100 Episoden: ein Tech Rückblick auf 2022/23, Predictions 2024 und viel Tech Trivia
		"/": "",
		// Example: #76 Mit Open Source 100.000$ verdienen, Sponsorware und Plattform-Risiken bei GitHub mit Martin Donath
		".": "",
		// Example: #28 O(1), O(log n), O(n^2) - Ist die Komplexität von Algorithmen im Entwickler-Alltag relevant?
		"(": "",
		")": "",
		"^": "",
		// Example: #181 Von Code zu Value: Wie Entwickler·innen Business-Mehrwert schaffen
		"·": "",
		// Example: #196 Star Wars auf GitHub: 4,5 Mio. Fake-Sterne entdeckt
		",": "",

		// Characters we want to preserve by changing it to something unique
		// and then changing it back after slugification
		// Example: #63 Spaß mit Zahlen: Under- und Overflows, Rückwärtslaufende Zeit, Negative Modulos und Währungsbeträge
		"ß": "0ss0",
		// Example: #115 Die Shift Left Philosophie: Mehr Verantwortung für Devs
		"ü": "0ue0",
		// Example: #131 Equity in Tech-Startups: Mehr als nur Gehalt mit Philipp "Pip" Klöckner
		"ö": "0oe0",
		// Example: #104 Präsentieren mit Wirkung: Public Speaking und Storytelling mit Anna Momber
		"ä": "0ae0",
		// Example: #19 Datenbank-Deepdive (oder das Ende einer Ära): von Redis bis ClickHouse
		"Ä": "0aae0",

		// No case yet, but just in case we need it
		"Ü": "0uue0",
		"Ö": "0ooe0",
	}
	text := slug.Make(value)

	// Substituation map to convert the unique characters back to their original form
	substitueMap := map[string]string{
		"0ue0":  "ü",
		"0oe0":  "ö",
		"0ae0":  "ä",
		"0uue0": "ü",
		"0ooe0": "ö",
		"0aae0": "ä",
		"0ss0":  "ß",
	}
	text = slug.Substitute(text, substitueMap)

	// Re-add the minus sign if we preserved it
	if preserveMinusChar {
		text = fmt.Sprintf("-%s", text)
	}
	return text
}
