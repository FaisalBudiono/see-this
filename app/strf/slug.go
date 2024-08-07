package strf

import (
	"regexp"
)

var (
	regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")
	regexpMultipleDashes     = regexp.MustCompile("-+")
)

func Slug(val string) string {
	slug := regexpNonAuthorizedChars.ReplaceAllString(val, "-")

	return regexpMultipleDashes.ReplaceAllString(slug, "-")
}
