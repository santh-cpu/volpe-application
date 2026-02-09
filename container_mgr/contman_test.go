// Package container_mgr_test tests container_mgr.
package container_mgr_test

import "testing"

import (
	cman "volpe-framework/container_mgr"
)

// TestContman runs basic problem test
func TestContman(t *testing.T) {
	cm := cman.NewContainerManager(false)
	err := cm.AddProblem("p1", "img.tar", 1)
	if err != nil {
		t.Error(err)
	}
}
