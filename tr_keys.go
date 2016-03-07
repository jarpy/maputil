// The maputil packages contains general-purpose functions for manipulating
// maps.
package maputil

import "strings"

// TrKeys transliterates the key names of string-keyed maps.
// It replaces all occurances of 'old' rune with 'new' rune.
func TrKeys(m map[string]interface{}, old rune, new rune) map[string]interface{} {
	for key, value := range m {
		// Recurse into sub-trees.
		switch value := value.(type) {
		case map[string]interface{}:
			m[key] = TrKeys(value, old, new)
		}

		if strings.ContainsRune(key, old) {
			delete(m, key)
			key = strings.Replace(key, string(old), string(new), -1)
			m[key] = value
		}
	}
	return m
}
