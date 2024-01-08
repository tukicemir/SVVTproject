package main

import (
	"errors"
	"github.com/tebeka/selenium"
)

func loadTest(wd selenium.WebDriver) error {
	if err := wd.Get("https://klix.ba"); err != nil {
		return err
	}

	html, err := wd.PageSource()
	if err != nil {
		return err
	}
	if len(html) < 10 {
		return errors.New("page loaded has size < 10")
	}
	return nil
}

func init() {
	addTest(test{
		name:  "Load test",
		tType: testTypeSmoke,
		fn:    loadTest,
	})
}
