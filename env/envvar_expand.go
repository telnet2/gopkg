package env

import (
	"fmt"
	"regexp"
)

var (
	RegexEnvVar = regexp.MustCompile(`\$([a-zA-Z_][a-zA-Z0-9_]*)|\$\{([a-zA-Z_][a-zA-Z0-9_]*)(:-([^}]*))?\}`)
)

func init() {
	// must not use the longest match
	// RegexEnvVar.Longest()
}

type EnvVarRef struct {
	Name         string
	DefaultValue string
	Index        [2]int
}

func ExpandEnvVar(s string) (string, error) {
	values := RegexEnvVar.FindAllStringSubmatch(s, -1)
	fmt.Println(values)
	return "", nil
}

func CaptureEnvNameAndDefaultValue(s string) (refs []EnvVarRef, ok bool) {
	values := RegexEnvVar.FindAllStringSubmatchIndex(s, -1)
	if len(values) == 0 {
		return nil, false
	}
	for _, capture := range values {
		ref := EnvVarRef{}
		ref.Index = [2]int{capture[0], capture[1]}
		if capture[2] > 0 {
			ref.Name = s[capture[2]:capture[3]]
		}
		if capture[4] > 0 {
			ref.Name = s[capture[4]:capture[5]]
		}
		if capture[8] > 0 {
			ref.DefaultValue = s[capture[8]:capture[9]]
		}
		if ref.Name != "" {
			refs = append(refs, ref)
		}
	}
	return refs, true
}

func ReplaceEnvVars(s string, refs []EnvVarRef, em EnvMap) string {
	if em == nil {
		em = make(EnvMap)
	}
	i := 0
	var replaced string
	for j := 0; j < len(refs) && i < len(s); j++ {
		// if the index is invalid, return the original string
		if refs[j].Index[0] < 0 || refs[j].Index[0] >= refs[j].Index[1] || refs[j].Index[1] > len(s) {
			return s
		}
		if refs[j].Index[0] > i {
			replaced += s[i:refs[j].Index[0]]
		}
		// capture := s[refs[j].Index[0]:refs[j].Index[1]]
		value := em[refs[j].Name]
		if value == "" {
			value = refs[j].DefaultValue
			// we cannot do recursive replacement.
			// refs, ok := CaptureEnvNameAndDefaultValue(value)
			// if ok {
			// 	value = ReplaceEnvVars(value, refs, em)
			// }
		}
		replaced += value
		i = refs[j].Index[1]
	}
	if i < len(s) {
		replaced += s[i:]
	}
	return replaced
}
