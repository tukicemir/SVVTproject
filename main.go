package main

import (
	"flag"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"os"
)

func main() {
	headless := flag.Bool("headless", false, "start chrome in headless mode")
	verbose := flag.Bool("verbose", false, "use verbose mode (print error messages)")
	flag.Parse()

	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	if *headless {
		caps.AddChrome(chrome.Capabilities{Args: []string{
			"--headless",
		}})
	}

	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}

	err = driver.Get("https://klix.ba")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ocurred: %v\n", err)
		os.Exit(1)
	}

	failedTests := 0

	for i, t := range tests {
		fmt.Printf("#%d %s (%s)", i+1, t.name, t.tType)
		if err := t.fn(driver); err != nil {
			fmt.Printf(" [FAILED]")
			if *verbose {
				fmt.Printf(" %s", err.Error())
			}
			fmt.Printf("\n")
			failedTests++
		} else {
			fmt.Printf(" [PASS]\n")
		}
		driver.Get("https://klix.ba")
	}

	fmt.Printf("Total tests: %d\n", len(tests))
	fmt.Printf("Passed tests: %.2f%%\n", ((float32(len(tests)) - float32(failedTests)) / float32(len(tests)) * 100))
	fmt.Printf("Failed tests: %.2f%%\n", (float32(failedTests) / float32(len(tests)) * 100))
}
