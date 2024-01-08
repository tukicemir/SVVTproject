package main

import (
	"github.com/tebeka/selenium"
	"strings"
)

func findAndClick(wd selenium.WebDriver, by string, name string) error {
	element, err := wd.FindElement(by, name)
	if err != nil {
		return err
	}
	return element.Click()
}

func fillFormEntry(wd selenium.WebDriver, key string, value string) error {
	k, err := wd.FindElement(selenium.ByName, key)
	if err != nil {
		return err
	}
	return k.SendKeys(value)
}

func loginUser(wd selenium.WebDriver, username, password string) error {
	if err := findAndClick(wd, selenium.ByID, "user"); err != nil {
		return err
	}

	if err := fillFormEntry(wd, "username", username); err != nil {
		return err
	}

	if err := fillFormEntry(wd, "password", password); err != nil {
		return err
	}

	if err := findAndClick(wd, selenium.ByCSSSelector, ".shadow-md"); err != nil {
		return err
	}

	return nil
}

func logoutUser(wd selenium.WebDriver) error {
	wd.Refresh()

	if err := findAndClick(wd, selenium.ByCSSSelector, "#user"); err != nil {
		return err
	}
	return findAndClick(wd, selenium.ByLinkText, "Odjavi se?")
}

func openRandomArticle(wd selenium.WebDriver) error {
	links, err := wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return err
	}

	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		if strings.Contains(href, "vijesti/bih/") {
			if err := wd.Get("https://klix.ba/" + href); err != nil {
				return err
			}
			break
		}
	}

	links, err = wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return err
	}

	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		if (strings.Contains(href, "vijesti/bih/") ||
			strings.Contains(href, "vijesti/svijet/")) &&
			!strings.Contains(href, "komentari") {
			fullUri := ""
			if strings.Contains(href, "klix.ba/") {
				fullUri = href
			} else {
				fullUri = "https://klix.ba/" + href
			}
			if err := wd.Get(fullUri); err != nil {
				return err
			}
			break
		}
	}
	return nil
}
