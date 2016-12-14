// Package strutil implements extra string methods that are popular
// here at SkilStak in courses and projects.

package strutil

import (
	"bytes"
	"strings"
	u "unicode"
)

// ToSlug produces a traditional SLUG string suitable for use in URLs
// as is done, for example, by GitHub when creating header anchors from
// GHM documents.
func ToSlug(s string) string {
	var buffer bytes.Buffer
	lastWasSpace := false
	s = strings.TrimSpace(s)
	for _, r := range s {
		if u.IsSpace(r) {
			if !lastWasSpace {
				buffer.WriteRune('-')
			}
			lastWasSpace = true
			continue
		} else {
			lastWasSpace = false
		}
		if u.IsNumber(r) || u.IsLetter(r) {
			buffer.WriteRune(u.ToLower(r))
		}
	}
	return buffer.String()
}
