package main

import "os"

func callBackExit(c *config) error {
	os.Exit(0)
	return nil
}
