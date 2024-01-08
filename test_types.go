package main

import "github.com/tebeka/selenium"

type testFn func(wd selenium.WebDriver) error
type testType string

const (
	testTypeSmoke      testType = "Smoke"
	testTypeRegression testType = "Regression"
)

type test struct {
	name  string
	tType testType
	fn    testFn
}

var tests []test

func addTest(t test) {
	tests = append(tests, t)
}
