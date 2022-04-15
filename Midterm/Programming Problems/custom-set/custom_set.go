package stringset

import (
	"reflect"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set struct {
	strings map[string]bool
}

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	var set Set
	res := make(map[string]bool)
	if len(l) != 0 {
		for _, s := range l {
			res[s] = true
		}
	}
	set.strings = res
	return set
}

func (s Set) String() string {
	if len(s.strings) == 0 {
		return "{}"
	}

	resultString := "{\""
	resultString += strings.Join(grabAllKeys(s), "\", \"")
	resultString += "\"}"

	return resultString
}

func grabAllKeys(s Set) []string {
	res := []string{}
	for k, _ := range s.strings {
		res = append(res, k)
	}
	return res
}

func (s Set) IsEmpty() bool {
	return len(s.strings) == 0
}

func (s Set) Has(elem string) bool {
	result := false

	if _, ok := s.strings[elem]; ok {
		result = true
	}
	return result
}

func (s Set) Add(elem string) {
	if _, ok := s.strings[elem]; !ok {
		s.strings[elem] = true
	}
}

func Subset(s1, s2 Set) bool {

	for k := range s1.strings {
		if !s2.Has(k) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1.strings {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func Intersection(s1, s2 Set) Set {

	resultSet := Set{}
	resultSet.strings = make(map[string]bool)

	for i := range s1.strings {
		for j := range s2.strings {
			if i == j {
				resultSet.strings[i] = true
			}
		}
	}
	return resultSet
}

func Difference(s1, s2 Set) Set {

	result := Set{}
	result.strings = make(map[string]bool)

	for k := range s1.strings {
		if !s2.Has(k) {
			result.strings[k] = true
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	res := Set{}
	res.strings = make(map[string]bool)
	res.strings = s1.strings

	for k, v := range s2.strings {
		if _, ok := res.strings[k]; !ok {
			res.strings[k] = v
		}
	}

	return res
}
