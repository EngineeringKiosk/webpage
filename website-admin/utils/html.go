package utils

import (
	"regexp"
	"strings"
)

// RemoveHTMLTags removes all HTML tags from a string
func RemoveHTMLTags(rawHTML string) string {
	// Replace </p> with </p> followed by a space
	cleanText := strings.ReplaceAll(rawHTML, "</p>", "</p> ")

	// Remove all HTML tags
	htmlTagsRegex := regexp.MustCompile(`<.*?>`)
	cleanText = htmlTagsRegex.ReplaceAllString(cleanText, "")

	return cleanText
}

// RemoveRelNofollowFromInternalLinks removes rel="nofollow" from internal links
func RemoveRelNofollowFromInternalLinks(htmlContent string) string {
	// Remove nofollow from engineeringkiosk.dev links
	pattern1 := regexp.MustCompile(`<a href="(https://engineeringkiosk\.dev/[^\s]*)" rel="nofollow">`)
	htmlContent = pattern1.ReplaceAllString(htmlContent, `<a href="$1">`)

	// Remove nofollow from jump.engineeringkiosk.dev links
	pattern2 := regexp.MustCompile(`<a href="(https://jump\.engineeringkiosk\.dev/[^\s]*)" rel="nofollow">`)
	htmlContent = pattern2.ReplaceAllString(htmlContent, `<a href="$1">`)

	return htmlContent
}
