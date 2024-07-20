package main

import "os"

func callBackExit(c *config, args ...string) error {
	os.Exit(0)
	return nil
}
