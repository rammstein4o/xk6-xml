package xml

import (
	"strings"

	"github.com/sbabiv/xml2map"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/xml", new(XML))
}

// XML is the k6 xml extension.
type XML struct{}

// Parse parses xml
func (*XML) Parse(body string) (map[string]interface{}, error) {
	decoder := xml2map.NewDecoder(strings.NewReader(body))
	return decoder.Decode()
}
