package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	depth := 1
	maxDepth := 10
	dir, _ := filepath.Abs(".")
	for depth < maxDepth {
		envFile := filepath.Join(dir, ".env")
		if fi, err := os.Stat(envFile); err == nil && !fi.IsDir() {
			err = godotenv.Load(envFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[WARN] failed to load %s: %v\n", envFile, err)
			}
		}
		modFile := filepath.Join(dir, "go.mod")
		if fi, err := os.Stat(modFile); err == nil && !fi.IsDir() {
			break
		}
		d := filepath.Dir(dir)
		if d == dir {
			break
		}
		dir = d
		depth++
	}
}
