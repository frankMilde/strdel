// strdel provides routines to delete words or regexp from strings
package strdel

import (
	"fmt"
	"regexp"
	"strings"
)

func Word(s string, wordToDelete string) string {
	regWordToDelete := regexp.MustCompile(`\b` + wordToDelete + `\b`)
	s = regWordToDelete.ReplaceAllString(s, "")
	return s
}

func RegExp(s string, regExp string) string {
	regWordToDelete := regexp.MustCompile(regExp)
	s = regWordToDelete.ReplaceAllString(s, "")
	return s
}

// Duplicates deletes duplicate from a string slice.
func Duplicates(strings []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]struct{}{}
	result := []string{}

	for _, s := range strings {
		if _, ok := encountered[s]; !ok {
			encountered[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}

// Numbering removes leading enumerations at the begin of a line from string
// s. Example: 3. Heading --> Heading
func Numbering(s string) string {
	digitsAndDotAtStartOfLine := `(?m)^[ \t\r\f]*\d+\.[ \t\r\f]*`
	return RegExp(s, digitsAndDotAtStartOfLine)
}

// TrailingSpaces removes trailing non-line breaking white spaces from
// string s
func TrailingSpaces(s string) string {
	// convert unicode char \u00A0 = &nbsp = 'non-breaking space' to space
	//s = strings.Replace(s, " \n", "\n", -1)

	regSpace := regexp.MustCompile(`[ \t\r\f]+\n`)
	s = regSpace.ReplaceAllString(s, "\n")

	return s
}

// LeadingSpaces removes leading non-line breaking white spaces from string
// s
func LeadingSpaces(s string) string {
	regSpace := regexp.MustCompile(`^[ \t\r\f]+`)
	s = regSpace.ReplaceAllString(s, "")
	regSpace = regexp.MustCompile(`\n[ \t\r\f]+`)
	s = regSpace.ReplaceAllString(s, "\n")
	return s
}

// EmptyBrackets changes multiline empty `{\n\n}` into `{}`
func EmptyBrackets(s string) string {
	reg := regexp.MustCompile(`\{(\s+)\}`)
	s = reg.ReplaceAllString(s, "{}")

	reg = regexp.MustCompile(`\{\\\\\}`)
	//fmt.Println(reg.FindAllString(s, -1))
	s = reg.ReplaceAllString(s, "{}")
	return s
}

func EmptyLinesInMacros(s string) string {

	pattern := `(\\\b[a-z]+\b\{(?s).*?)[\r\n]((?m)^\s*$[\r\n]*)[\r\n]\s*(.*?\})`

	replace := `${1} 
	${3}`

	emptyLineInMacro := regexp.MustCompile(pattern)

	matches := emptyLineInMacro.FindAllStringSubmatch(s, -1)
	fmt.Println(matches)
	fmt.Println("\n\n\n\nNew Match:\n\n")
	return emptyLineInMacro.ReplaceAllString(s, replace)
}

func EmptyMacros(s string, nestingDepth int) string {

	simpleMacro := regexp.MustCompile(`\\\b([a-z]+)\b\{\}(\s*\\\\)?`)
	for i := 0; i < nestingDepth; i++ {
		s = simpleMacro.ReplaceAllString(s, "")
	}
	return s
}

// SpaceBeforeClosingBrackets deletes linebreaks and spaces before closing
// brackets "}".
func SpaceBeforeClosingBrackets(s string) string {

	linebreaks := regexp.MustCompile(`(\\\\)+\}`)
	s = linebreaks.ReplaceAllString(s, `}\\`)

	spaces := regexp.MustCompile(`\s+\}`)
	s = spaces.ReplaceAllString(s, "} ")

	spacesAndLineBreaks := regexp.MustCompile(`(\s|\\\\)+\}`)
	s = spacesAndLineBreaks.ReplaceAllString(s, `} $1`)
	return s
}

func EmptyLine(s string) string {
	re := regexp.MustCompile("(?m)^\\s*$[\r\n]*")

	return strings.Trim(re.ReplaceAllString(s, ""), "\r\n")
}

func SpaceAfterOpeningBrackets(s string) string {

	reg := regexp.MustCompile(`(\S)\{\s+`)
	s = reg.ReplaceAllString(s, "$1{")

	return s
}
