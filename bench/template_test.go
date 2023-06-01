package bench

import (
	"fmt"
	"strings"
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/require"
)

func BenchmarkTemplate(b *testing.B) {
	// Create a new template
	tmpl := template.Must(template.New("").Parse(`{{.YYYY}}{{.MM}}{{.DD}}/{{.hh}}{{.mm}}/filename_{{.uuid}}`))

	formatter := map[string]string{
		"YYYY": "2006",
		"MM":   "01",
		"DD":   "02",
		"hh":   "15",
		"mm":   "04",
		"uuid": "a-b-c-d",
	}

	sb := strings.Builder{}
	sb.Grow(256)
	for i := 0; i < b.N; i++ {
		sb.Reset()
		// Execute the template with the data
		_ = tmpl.Execute(&sb, formatter)
		require.NotEmpty(b, sb.String())
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		name := fmt.Sprintf(time.Now().Format("20060102/1504/%s"), "a-b-c-d")
		require.NotEmpty(b, name)
	}
}
