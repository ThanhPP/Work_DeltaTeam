package rebrandly

import (
	"testing"
)

func TestCreateShortLinkRebrandly(t *testing.T) {
	inputFwdSlice := []string{"1704fm004pxrac.deltapodvn.com"}
	_, _, errCount, _ := CreateShortLinkRebrandly("4031:4031", inputFwdSlice)
	if errCount != 1 {
		t.Error(errCount)
	}
}
