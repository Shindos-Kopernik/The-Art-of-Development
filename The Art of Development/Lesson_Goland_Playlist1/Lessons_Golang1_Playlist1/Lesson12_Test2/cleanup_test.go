package main

import (
	"fmt"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	fmt.Println("START")
	t.Cleanup(func() { // TODO правильный способ подчистки за собой хвостов использование CleanUp func!
		fmt.Println("TEARDOWN ON CLEANUP")
	})
	t.Run("FIRST", func(t *testing.T) {
		fmt.Println("ok")
	})
	t.Run("SECOND", func(t *testing.T) {
		fmt.Println("ok")
	})

	t.Run("THIRD", func(t *testing.T) {
		t.Fatal("fatal test")
	})

	fmt.Println("TEARDOWN AT END")
}
