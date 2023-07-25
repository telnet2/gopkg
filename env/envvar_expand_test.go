package env

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("envvar_expand_test", func() {
	var (
		refs []EnvVarRef
		ok   bool
	)

	It("should not allow a brace within the default value", func() {
		// This cannot be done in golang's regexp package.
		matches := RegexEnvVar.FindAllString(`${WORK:-${HOME:-/home}/work}`, -1)
		Expect(matches).To(HaveLen(1))
		Expect(matches[0]).To(Equal(`${WORK:-${HOME:-/home}`))
	})

	It("should capture the envvar name and value", func() {
		refs, ok = CaptureEnvNameAndDefaultValue(`$HOME`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(1))
		Expect(refs[0]).To(BeEquivalentTo(EnvVarRef{Name: "HOME", Index: [2]int{0, 5}}))

		refs, ok = CaptureEnvNameAndDefaultValue(`${HOME}`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(1))
		Expect(refs[0]).To(BeEquivalentTo(EnvVarRef{Name: "HOME", Index: [2]int{0, 7}}))

		refs, ok = CaptureEnvNameAndDefaultValue(`${HOME:-/home}`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(1))
		Expect(refs[0]).To(BeEquivalentTo(EnvVarRef{Name: "HOME", DefaultValue: "/home", Index: [2]int{0, 14}}))

		// ENV_WITH_SPACE is ignored
		refs, ok = CaptureEnvNameAndDefaultValue(`${HOME:-/home} ${PWD} $HOME ${drs_url:-http://www.google.com}` +
			"\r\n" + `${SECOND:-true} ${ENV_WITH_SPACE :-false}`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(5))
		Expect(refs[0]).To(BeEquivalentTo(EnvVarRef{Name: "HOME", DefaultValue: "/home", Index: [2]int{0, 14}}))
		Expect(refs[1]).To(BeEquivalentTo(EnvVarRef{Name: "PWD", DefaultValue: "", Index: [2]int{15, 21}}))
		Expect(refs[2]).To(BeEquivalentTo(EnvVarRef{Name: "HOME", DefaultValue: "", Index: [2]int{22, 27}}))
		Expect(refs[3]).To(BeEquivalentTo(EnvVarRef{Name: "drs_url", DefaultValue: "http://www.google.com", Index: [2]int{28, 61}}))
		Expect(refs[4]).To(BeEquivalentTo(EnvVarRef{Name: "SECOND", DefaultValue: "true", Index: [2]int{63, 78}}))
	})

	It("should capture recursive expression", func() {
		refs, ok := CaptureEnvNameAndDefaultValue(`${WORK_DIR:-${HOME:-/home}/work}`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(1))
		Expect(refs[0].DefaultValue).To(Equal("${HOME:-/home"))
	})

	It("should return an error if the env var is invalid", func() {
		refs, ok = CaptureEnvNameAndDefaultValue(`${HOME:-/home`)
		Expect(ok).To(BeFalse())

		refs, ok = CaptureEnvNameAndDefaultValue(`$ HOME`)
		Expect(ok).To(BeFalse())

		refs, ok = CaptureEnvNameAndDefaultValue(`${ENV_WITH_SPACE }`)
		Expect(ok).To(BeFalse())

		refs, ok = CaptureEnvNameAndDefaultValue(`${ENV_NOT_CLOSED`)
		Expect(ok).To(BeFalse())

		refs, ok = CaptureEnvNameAndDefaultValue(`${ENV_NOT_CLOSED $HOME`)
		Expect(ok).To(BeTrue())
		Expect(refs).To(HaveLen(1))
	})

	It("should replace the environment variable references to values", func() {
		str := `${GREET:-hello!} My home directory was "${HOME:-/home}", but now "/home"`
		refs, ok = CaptureEnvNameAndDefaultValue(str)
		replaced := ReplaceEnvVars(str, refs, EnvMap{"HOME": "/world"})
		Expect(replaced).To(Equal(`hello! My home directory was "/world", but now "/home"`))
	})

	XIt("should RECURSIVELY replace the environment variable references to values", func() {
		str := `${WORK_DIR:-${HOME:-/home}/work}`
		refs, ok = CaptureEnvNameAndDefaultValue(str)
		replaced := ReplaceEnvVars(str, refs, nil)
		Expect(replaced).To(Equal(`/home/work`))
	})

	It("should not panic", func() {
		x := "012345"
		Expect(x[0:6]).To(Equal("012345"))
	})
})
