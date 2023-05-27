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
			// This works differently from bash
			Expect(expand("/$HELLO/world")).To(Equal("/$HELLO/world"))
		})
	})

	It("should run a command with a variable", func() {
		cmd := exec.Command("bash", "-c", "for x in {0..3}; do echo $x; done")
		out, err := cmd.CombinedOutput()
		Expect(err).NotTo(HaveOccurred())
		lines := bytes.Split(out, []byte("\n"))
		Expect(lines).To(HaveLen(5)) // with last empty line
	})

	It("should expand string fields with `envexp` tag", func() {
		envmap := EnvMap{"KEY": "VALUE", "HOME": "/homedir"}
		type TestSubtype struct {
			LogDir  string `envexp:""`
			WorkDir string `envexp:""`
		}

		type TestStruct struct {
			Command    string `envexp:""`
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
		Expect(x.SubTypePtr.WorkDir).To(ContainSubstring("$HOME/work"))
	})
})
