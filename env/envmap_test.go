package env

import (
	"bytes"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/telnet2/pm/cmd"
)

var _ = Describe("envmap_test", func() {
	It("should load system envs", func() {
		val := "A\nB\nC"
		os.Setenv("NEWLINE_VALUE", "A\nB\nC")
		envMap := make(EnvMap)
		envMap.LoadSystem()
		Expect(envMap).NotTo(BeEmpty())
		Expect(envMap["NEWLINE_VALUE"]).To(Equal(val))
	})

	It("should generate env []string from EnvMap", func() {
		envMap := make(EnvMap)
		envMap.LoadSystem()
		envMap["HELLO"] = "hello"
		envMap["hello"] = "world"
		Expect(envMap.AsString()).To(ContainElements([]string{"HELLO=hello", "hello=world"}))

		// check if env is correctly applied
		By("check PATH env", func() {
			exeCmd := cmd.NewCmdOptions(cmd.Options{
				Buffered: true,
			}, "show-env", "go run . -env PATH")
			exeCmd.Dir = "./testdata/showenv"
			exeCmd.Env = envMap.AsString()
			status := <-exeCmd.Start()
			Expect(status.Error).NotTo(HaveOccurred())
			Expect(status.Stdout[0]).NotTo(BeEmpty())
		})

		By("check HELLO env", func() {
			exeCmd := cmd.NewCmdOptions(cmd.Options{
				Buffered: true,
			}, "show-env", "go run . -env HELLO")
			exeCmd.Dir = "./testdata/showenv"
			exeCmd.Env = envMap.AsString()
			status := <-exeCmd.Start()
			Expect(status.Stdout).To(ConsistOf([]string{"hello"}))
		})

		By("check hello env", func() {
			exeCmd := cmd.NewCmdOptions(cmd.Options{
				Buffered: true,
			}, "show-env", "go run . -env hello")
			exeCmd.Dir = "./testdata/showenv"
			exeCmd.Env = envMap.AsString()
			status := <-exeCmd.Start()
			Expect(status.Stdout).To(ConsistOf([]string{"world"}))
		})

		By("shell expansion", func() {
			expand := envMap.NewExpander()
			// Replace with default value
			Expect(expand("Hello ${HELLO:-WELCOME}")).To(Equal("Hello hello"))
			Expect(expand("Hello ${WORLD:-WELCOME}")).To(Equal("Hello WELCOME"))
			Expect(expand("/$HELLO/world")).To(Equal("/hello/world"))
			Expect(expand("$HELLO$WORLD$GYPSY")).To(Equal("hello"))
			Expect(expand("${HELLO}$WORLD${GYPSY}")).To(Equal("hello"))
			Expect(expand("${HELLO}${WORLD:-}${GYPSY:- GYPSY}")).To(Equal("hello GYPSY"))
			// it doesn't support recursive expansion
			Expect(expand("${HELLO}${WORLD:-}${ENV:-${HELLO}}")).To(Equal("hello${HELLO}"))
		})
	})

	It("should run a command with a variable", func() {
		cmd := exec.Command("bash", "-c", "for x in {0..3}; do echo $x; done")
		out, err := cmd.CombinedOutput()
		Expect(err).NotTo(HaveOccurred())
		lines := bytes.Split(out, []byte("\n"))
		Expect(lines).To(HaveLen(5)) // with last empty line
	})

	It("should expand string fields", func() {
		envmap := EnvMap{"KEY": "VALUE", "HOME": "/homedir"}
		type TestSubtype struct {
			LogDir  string
			WorkDir string
		}

		type TestStruct struct {
			Command    string
			SubType    TestSubtype
			SubTypePtr *TestSubtype
			NilPtr     *TestSubtype
		}

		x := TestStruct{
			Command: "Key=${KEY}",
			SubType: TestSubtype{
				LogDir:  "$HOME",
				WorkDir: "$WORK",
			},
			SubTypePtr: &TestSubtype{
				LogDir:  "v=${HOME}=/log",
				WorkDir: "$HOME/work",
			},
		}
		envmap.Expand(&x)
		Expect(x.Command).To(ContainSubstring("=VALUE"))
		Expect(x.SubType.LogDir).To(ContainSubstring("/homedir"))
		Expect(x.SubType.WorkDir).To(Equal(""))
		Expect(x.SubTypePtr.LogDir).To(ContainSubstring("/homedir=/log"))
		// This doesn't work as it shoud do.
		Expect(x.SubTypePtr.WorkDir).To(ContainSubstring("/homedir/work"))
	})
})
