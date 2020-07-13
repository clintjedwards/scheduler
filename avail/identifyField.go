package avail

import "regexp"

// list of regexs that we use to match against a single cron expression
var (
	spanRegex     = regexp.MustCompile(`^[0-9]+-[0-9]+$`)
	wildcardRegex = regexp.MustCompile(`^\*$`)
	listRegex     = regexp.MustCompile(`,+`)
	valueRegex    = regexp.MustCompile(`^([0-9]+)$`)
)

// fieldType is an enum which we use to note the type of an expression
type fieldType string

const (
	span     fieldType = "span"
	wildcard           = "wildcard"
	list               = "list"
	value              = "value"
	unknown            = "unknown"
)

// fieldRegexToType stores the mapping between a field's regex representation
// and the concrete type it is. This is used to help is identify the field
// and run the appropriate parsing handler.
var fieldRegexToType = map[*regexp.Regexp]fieldType{
	spanRegex:     span,
	wildcardRegex: wildcard,
	listRegex:     list,
	valueRegex:    value,
}

func identifyFieldType(expression string) fieldType {
	for regex := range fieldRegexToType {
		isMatch := regex.MatchString(expression)
		if isMatch {
			return fieldRegexToType[regex]
		}
	}

	return unknown
}
