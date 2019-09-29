package main

import (
	"os"
)

func main() {
	setEnvVars()
	removeStateFiles()
}

func setEnvVars() {
	os.Setenv("TF_ACC", "1")
	os.Setenv("TF_LOG", "DEBUG")
	os.Setenv("TF_LOG_PATH", "trace.log")
	os.Setenv("API_TOKEN", "<Nagios API Token>")
	os.Setenv("NAGIOS_URL", "http://localhost/nagiosxi")
}

func removeStateFiles() {
	filesToRmove := []string{"terraform.tfstate", "terraform.tfstate.backup", "trace.log", "the_plan", "nagios/trace.log"}

	for _, file := range filesToRmove {
		if _, err := os.Stat(file); err == nil {
			os.Remove(file)
		}
	}
}
