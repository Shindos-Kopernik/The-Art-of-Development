package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) { // TODO для всего пакета main
	fmt.Println("setup")
	res := m.Run()
	fmt.Println("tear-down")
	os.Exit(res) //Это код ответа тестов
}

// TODO запуск групп тестов
func TestMultiple(t *testing.T) {
	t.Run("groupA", func(t *testing.T) {
		t.Log("A")
		t.Run("simple", func(t *testing.T) {
			t.Parallel() // TODO для паралельного запуска тестов(если тесты не завсисят друг от друга)
			t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			t.Parallel() // TODO для паралельного запуска тестов
			t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			t.Parallel() // TODO для паралельного запуска тестов
			t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
		})
	})
	t.Run("groupB", func(t *testing.T) {
		t.Log("B")
		t.Run("simple", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов(если тесты не завсисят друг от друга)
			//t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
				// tearDown
				// delete test data within db
			}
		})
	})
	t.Run("groupC", func(t *testing.T) {
		t.Log("C")
		t.Run("simple", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов(если тесты не завсисят друг от друга)
			//t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
				// tearDown
				// delete test data within db
			}
		})
	})
	t.Run("groupD", func(t *testing.T) {
		t.Log("D")
		t.Run("simple", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов(если тесты не завсисят друг от друга)
			//t.Log("simple")
			var x, y, result = 2, 2, 4
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
			t.Run("1", func(t *testing.T) {
				r := Multiple(1, 1)
				if r != 1 {
					t.Errorf("failed")
				}
			})
		})

		t.Run("medium", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("medium")
			var x, y, result = 222, 222, 49284
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
			}
		})
		t.Run("negative", func(t *testing.T) {
			//t.Parallel()   TODO для паралельного запуска тестов
			//t.Log("negative")
			var x, y, result = -2, 4, -8
			realResult := Multiple(x, y)
			if realResult != result {
				t.Errorf("%d =! %d", result, realResult)
				// tearDown
				// delete test data within db
			}
		})
	})
}
