package xml

import (
	"github.com/clbanning/mxj"
	"go.k6.io/k6/js/modules"
)

// XML is the k6 xml extension.
type XML struct {
	vu modules.VU
}

// Parse parses xml
func (*XML) Parse(body string) (map[string]interface{}, error) {
	return mxj.NewMapXml([]byte(body))
}

func newXml(vu modules.VU) *XML {
	return &XML{vu: vu}
}
