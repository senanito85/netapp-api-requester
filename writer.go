package main

import (
	"fmt"
	"strings"
)

func write(out []output) {
	for _, o := range out {
		fmt.Println(strings.Join(o.toStringSlice(), "|"))
	}
}
