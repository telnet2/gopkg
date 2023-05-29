package ginkgo

import (
	"encoding/json"
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/onsi/gomega/types"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

var _ types.GomegaMatcher = (*equalJsonMatcher)(nil)

func EqualJson(expected interface{}) types.GomegaMatcher {

	matcher := &equalJsonMatcher{
		expected: expected.([]byte),
	}
	return matcher
}

type equalJsonMatcher struct {
	expected, actual []byte
	diffRes          diff.Diff
}

func (matcher *equalJsonMatcher) Match(actual interface{}) (success bool, err error) {
	matcher.actual = actual.([]byte)
	differ := diff.New()
	res, err := differ.Compare(matcher.actual, matcher.expected)
	if err != nil {
		return false, err
	}
	if !res.Modified() {
		return true, nil
	}
	matcher.diffRes = res
	return false, nil
}

func (matcher *equalJsonMatcher) FailureMessage(actual interface{}) (message string) {
	msg := formatDiff(matcher.actual, matcher.diffRes)
	return fmt.Sprintf("JSON Diff expect to actual\r\n%s", msg)
}

func (matcher *equalJsonMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Two JSON are equivalent and no diff found"
}

func formatDiff(actual []byte, res diff.Diff) string {
	if res != nil && res.Modified() {
		var actualJson map[string]interface{}
		json.Unmarshal(actual, &actualJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: false,
			Coloring:       true,
		}

		formatter := formatter.NewAsciiFormatter(actualJson, config)
		diffString, _ := formatter.Format(res)
		return color.White + diffString
	}
	return ""
}
