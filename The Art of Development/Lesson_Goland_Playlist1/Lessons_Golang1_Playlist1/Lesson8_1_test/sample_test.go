package main

// TODO
//Запуск в консоли go test
// go test -v -json
// go test -v -run /sample - какой тест мы хотим запустит, если тестов несколько
// go test -v -run "Название любого не существующего теста"команда для проверки компиляции тестов
import "testing"

func TestMultiple(t *testing.T) {
	var x, y, result = 2, 2, 4

	realResult := Multiple(x, y)
	if realResult != result {
		t.Errorf("%d != %d", result, realResult)
	}
}
