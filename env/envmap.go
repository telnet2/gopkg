package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	shellexpand "github.com/ganbarodigital/go_shellexpand"
	"github.com/jinzhu/copier"
	"github.com/joho/godotenv"
	"github.com/mitchellh/reflectwalk"
)

type EnvMap map[string]string

// AsString returns a list of KEY=VALUE strings.
func (em EnvMap) AsString() (ret []string) {
	for k, v := range em {
		ret = append(ret, fmt.Sprintf("%s=%s", k, v))
	}
	return
}

// LoadSystem loads os.Environ() into this map
func (em EnvMap) LoadSystem() {
	em.Load(os.Environ())
}

// Load loads envs from a string array
func (em EnvMap) Load(envs []string) {
	for _, env := range envs {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) == 2 {
			em[kv[0]] = kv[1]
		}
	}
}

// LoadFile load env files
func (em EnvMap) LoadFile(f ...string) {
	other, _ := godotenv.Read(f...)
	em.Merge(other)
}

// Merge merges the other env map into this env map.
func (em EnvMap) Merge(other EnvMap) {
	_ = copier.Copy(&em, other)
}

// Expand performs shell expansion using this EnvMap.
func (em EnvMap) NewExpander() func(string) string {
	callbacks := shellexpand.ExpansionCallbacks{
		LookupVar: func(s string) (string, bool) {
			v, ok := em[s]
			return v, ok
		},
	}
	return func(input string) string {
		out, _ := shellexpand.Expand(input, callbacks)
		return out
	}
}

// Expand recursively expands string files from a struct.
func (em EnvMap) Expand(t interface{}) {
	expander := varExpander{expander: em.NewExpander(), environ: em.AsString()}
	_ = expander.Expand(t)
}

// varExpander is a helper type to recursively expand the struct's string fields
type varExpander struct {
	environ  []string
	expander func(string) string
}

func (em *varExpander) Expand(t interface{}) error {
	return reflectwalk.Walk(t, em)
}

// Struct is required to implement StructWalker
func (em *varExpander) Struct(v reflect.Value) error {
	return nil
}

func (em *varExpander) StructField(f reflect.StructField, v reflect.Value) error {
	if f.Type.Kind() == reflect.String {
		if v.CanSet() {
			org := v.String()
			v.SetString(em.expander(org))
		}
		return nil
	}
	return nil
}

// Map is necessary to implement MapWalker
func (em *varExpander) Map(m reflect.Value) error {
	return nil
}

func (em *varExpander) MapElem(m, k, v reflect.Value) error {
	if v.Type().Kind() == reflect.String {
		org := v.String()
		rep := em.expander(org)
		m.SetMapIndex(k, reflect.ValueOf(rep))
	} else if v.Type().Kind() == reflect.Interface {
		vv := v.Elem()
		if vv.Type().Kind() == reflect.String {
			org := vv.String()
			rep := em.expander(org)
			m.SetMapIndex(k, reflect.ValueOf(rep))
		}
	}
	return nil
}

// ExpandEnvVars expands environment variables in a struct recusively.
// It expands the string fields in the struct and its sub-structs including a map with string values.
func ExpandEnvVars(s interface{}) {
	envs := make(EnvMap)
	envs.LoadSystem()
	envs.Expand(s)
}
