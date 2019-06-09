package astgen

import "strings"

type Importer func(name string) bool

func ImportAll(name string) bool {
	return true
}

func ImportNone(name string) bool {
	return false
}

func ImportFrom(names []string) Importer {
	return func(name string) bool {
		for _, n := range names {
			if strings.EqualFold(name, n) {
				return true
			}
		}
		return false
	}
}
