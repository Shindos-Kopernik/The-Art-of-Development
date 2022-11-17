package main

import (
	"os"
	"os/exec"
)

func main() {
	// ssh c := exec.Command("ssh", "-t", "kart", "htop")  И КАЖДУЮ КОМАНДУ ЧЕРЕЗ ЗАПЯТУЮ
	c := exec.Command("cat", "test3.txt")
	c.Stdout = os.Stdout
	c.Run()
}
