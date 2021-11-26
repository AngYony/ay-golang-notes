package prose

import "strings"

func JoinWithCommas(phrasse []string) string {
	result := strings.Join(phrasse[:len(phrasse)-1], ", ")
	result += " and"
	result += phrasse[len(phrasse)-1]
	return result
}
