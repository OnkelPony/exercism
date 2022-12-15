package ledger

import (
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sort"
	"strings"
	"time"
)

// Defined constants for repeatedly used values and time manipulating.
const (
	changeWidth = 13
	datumWidth  = 11
	descWidth   = 25
	dutchLocale = "nl-NL"
	engLocale   = "en-US"
	layoutISO   = "2006-01-02"
	layoutUS    = "01/02/2006"
	layoutNL    = "02-01-2006"
)

var locales = map[string]Locale{
	dutchLocale: {
		date:        "Datum",
		description: "Omschrijving",
		change:      "Verandering",
		layout:      layoutNL,
		tag:         language.Dutch,
		separator:   " ",
		procNeg: func(s *string) {
			*s += "-"
		},
	},
	engLocale: {
		date:        "Date",
		description: "Description",
		change:      "Change",
		layout:      layoutUS,
		tag:         language.English,
		separator:   "",
		procNeg: func(s *string) {
			*s = "(" + *s + ")"
		},
	},
}

// Entry represents account record containing date, description and amount in cents.
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// Locale is for processing strings in various locales.
type Locale struct {
	date        string
	description string
	change      string
	layout      string
	tag         language.Tag
	separator   string
	procNeg     func(*string)
}

// FormatLedger returns nice sorted ledger based on locale, date and entries slice.
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// Less complicated error checking using map lookup.
	curSymbol := map[string]rune{
		"EUR": '€',
		"USD": '$',
	}
	if _, ok := curSymbol[currency]; !ok {
		return "", errors.New("wrong currency")
	}

	// Shorter and more efficient entries copy.
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	// More apparent result name and Sprintf to format it using map of locales.
	var result string
	l, ok := locales[locale]
	if ok {
		result = fmt.Sprintf(
			"%-*v| %-*v | %v\n", datumWidth, l.date, descWidth, l.description, l.change,
		)
	} else {
		return result, errors.New("wrong locale")
	}

	// Sort entries by Date, Description and Change ascending order.
	sort.SliceStable(entriesCopy, func(i, j int) bool {
		ei, ej := entriesCopy[i], entriesCopy[j]
		switch {
		case ei.Date != ej.Date:
			return ei.Date < ej.Date
		case ei.Description != ej.Description:
			return ei.Description < ej.Description
		default:
			return ei.Change < ej.Change
		}
	})

	// Parallelism, not always a great idea ;-)
	// Removed go anon function and channel.
	for _, entry := range entriesCopy {
		de := entry.Description
		if len(de) > descWidth {
			de = de[:descWidth-3] + "..."
		} else {
			de += strings.Repeat(" ", descWidth-len(de))
		}

		// Parse date using time package and map of locales.
		var datum string
		t, err := time.Parse(layoutISO, entry.Date)
		if err != nil {
			return "", errors.New("wrong date format")
		}
		datum = t.Format(locales[locale].layout)

		negative := false
		cents := entry.Change
		if cents < 0 {
			cents = -cents
			negative = true
		}

		// Formatting amount using golang.org/x/text/language and golang.org/x/text/message and map of locales.
		var amount string
		p := message.NewPrinter(l.tag)
		amount = p.Sprintf("%c%s%.2f", curSymbol[currency], l.separator, float64(cents)/100)
		if negative {
			l.procNeg(&amount)
		} else {
			amount += " "
		}

		// Format using Sprintf..
		result += fmt.Sprintf("%s | %s | %*v\n", datum, de, changeWidth, amount)
	}
	return result, nil
}
