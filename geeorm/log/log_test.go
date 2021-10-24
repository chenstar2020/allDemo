package log

import "testing"

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	DEBUG("test debug")
	DEBUGF("test debug:%s", "abc")
	INFOF("test info")
}
