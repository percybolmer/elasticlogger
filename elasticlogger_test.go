package elasticlogger

import (
	"io"
	"testing"
)

// TestWriter tests that elasticlogger successfully implements the Io.Writer interface, this is done by checking a type assertion
func TestWriter(t *testing.T) {
	var _ io.Writer = &ElasticLog{}
}
