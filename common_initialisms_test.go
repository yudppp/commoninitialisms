package commoninitialisms_test

import (
	"testing"

	"github.com/yudppp/commoninitialisms"
)

func TestGetCommonInitialisms(t *testing.T) {
	commonInitialisms, err := commoninitialisms.GetCommonInitialisms()
	if err != nil {
		t.Errorf("GetCommonInitialisms should be not error, %v", err)
	}
	if _, ok := commonInitialisms["ID"]; !ok {
		t.Error("GetCommonInitialisms should contain ID")
	}
}

func TestMustGetCommonInitialisms(t *testing.T) {
	commonInitialisms := commoninitialisms.Must(commoninitialisms.GetCommonInitialisms())
	if _, ok := commonInitialisms["ID"]; !ok {
		t.Error("GetCommonInitialisms should contain ID")
	}
}
