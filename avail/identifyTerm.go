package avail

// identifyTerm is used to create a concrete type for a single cron term.
// A cron term is a single field in a complete cron expression.
//
// Ex. in the expression: "0 15 10 * * *", "15" would be a term of type "value".

import "regexp"

// List of regexs that we use to match against a single cron term for identification.
var (
	spanRegex     = regexp.MustCompile(`^[0-9]+-[0-9]+$`)
	wildcardRegex = regexp.MustCompile(`^\*$`)
	listRegex     = regexp.MustCompile(`,+`)
	valueRegex    = regexp.MustCompile(`^([0-9]+)$`)
)

// termType is an enum which represents different term kinds
type termType string

const (
	span     termType = "span"
	wildcard          = "wildcard"
	list              = "list"
	value             = "value"
	unknown           = "unknown"
)

// termRegexToType stores the mapping between a term's regex representation
// and the concrete type it is. This is used to help is identify the term type
// so that we can run the correct parser later.
var termRegexToType = map[*regexp.Regexp]termType{
	spanRegex:     span,
	wildcardRegex: wildcard,
	listRegex:     list,
	valueRegex:    value,
}

func identifyTermType(term string) termType {
	for regex := range termRegexToType {
		isMatch := regex.MatchString(term)
		if isMatch {
			return termRegexToType[regex]
		}
	}

	return unknown
}
