package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// saves the default seccomp profile as a json file so people can use it as a
// base for their own custom profiles
func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f := filepath.Join(wd, "default.json")

	// write the default profile to the file
	b, err := json.MarshalIndent(UserProfile(), "", "\t")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(f, b, 0o644); err != nil {
		panic(err)
	}
}
