package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/LassiHeikkila/flatmap/flatmap"
)

func main() {
	os.Exit(run())
}

func run() int {
	// read JSON from stdin
	// flatten it
	// and write JSON to stdout
	// any errors should go to stderr

	d := json.NewDecoder(os.Stdin)

	m := make(map[string]interface{})

	err := d.Decode(&m)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading JSON from stdin:", err)
		return 1
	}

	fm := flatmap.Flatten(m)

	e := json.NewEncoder(os.Stdout)
	err = e.Encode(&fm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing JSON to stdout:", err)
		return 1
	}
	return 0
}
