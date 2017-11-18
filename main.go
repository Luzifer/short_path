package main

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	cur, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	// Shorten homedir
	cur = strings.Replace(cur, home, "~", 1)

	out := strings.Split(cur, "/")
	for i := 0; i < len(out)-2; i++ {
		if len(out[i]) > 0 {
			out[i] = string(out[i][0])
		}
	}

	fmt.Print(strings.Join(out, "/"))
}
