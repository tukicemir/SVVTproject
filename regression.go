package main

import (
	"errors"
	"fmt"
	"github.com/tebeka/selenium"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func registerTest(wd selenium.WebDriver) error {
	if err := findAndClick(wd, selenium.ByID, "user"); err != nil {
		return err
	}

	if err := findAndClick(wd, selenium.ByCSSSelector, "#registerbtn"); err != nil {
		return err
	}

	inputs, err := wd.FindElements(selenium.ByTagName, "input")
	if err != nil {
		return err
	}

	for i := range inputs {
		tp, _ := inputs[i].GetAttribute("type")
		switch tp {
		case "username":
			inputs[i].SendKeys("svvtklix123")
		case "email":
			inputs[i].SendKeys("svvtklix123")
		case "password":
			inputs[i].SendKeys("Password123")
		}
	}

	return nil
}

func loginTest(driver selenium.WebDriver) error {
	if err := loginUser(driver, "svvtklix@gmail.com", "svvtklix123"); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	driver.Refresh()

	if err := findAndClick(driver, selenium.ByCSSSelector, "#user"); err != nil {
		return err
	}

	loggedInDiv, err := driver.FindElement(selenium.ByCSSSelector, "#loggedindiv")
	if err != nil {
		return err
	}

	innerDiv, err := loggedInDiv.FindElement(selenium.ByTagName, "div")
	if err != nil {
		return err
	}

	name, err := innerDiv.FindElement(selenium.ByCSSSelector, ".font-semibold")
	if err != nil {
		return err
	}

	nameValue, err := name.Text()
	if err != nil {
		return err
	}

	if nameValue != "SVVTprojekat" {
		return errors.New("value mismatch")
	}

	return logoutUser(driver)
}

func searchTest(wd selenium.WebDriver) error {
	if err := findAndClick(wd, selenium.ByCSSSelector, "#search-open"); err != nil {
		return err
	}

	if err := fillFormEntry(wd, "q", "test"); err != nil {
		return err
	}

	wd.KeyDown(selenium.EnterKey)
	wd.KeyUp(selenium.EnterKey)

	nav, err := wd.FindElement(selenium.ByCSSSelector, "#breadcrumb")
	if err != nil {
		return err
	}

	ol, err := nav.FindElement(selenium.ByTagName, "ol")
	if err != nil {
		return err
	}

	lis, err := ol.FindElements(selenium.ByTagName, "li")
	if err != nil {
		return err
	}

	if len(lis) != 2 {
		return errors.New("did not get two elements")
	}

	txt, err := lis[1].Text()
	if err != nil {
		return err
	}

	if txt != "test" {
		return errors.New("value mismatch")
	}

	return nil
}

func articleNavigationTest(wd selenium.WebDriver) error {
	return openRandomArticle(wd)
}

func commentFunctionality(wd selenium.WebDriver) error {
	if err := openRandomArticle(wd); err != nil {
		return err
	}

	if err := loginUser(wd, "svvtklix@gmail.com", "svvtklix123"); err != nil {
		return err
	}

	wd.Refresh()

	form, err := wd.FindElement(selenium.ByCSSSelector, "#samokomentar")
	if err != nil {
		return err
	}

	k, err := wd.FindElement(selenium.ByCSSSelector, "#komentarinput")
	if err != nil {
		return err
	}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	num := r.Intn(100)
	k.SendKeys(fmt.Sprintf("test comment%d", num))

	div, err := form.FindElement(selenium.ByTagName, "div")
	if err != nil {
		return err
	}

	btn, err := div.FindElement(selenium.ByTagName, "button")
	if err != nil {
		return err
	}

	if err := btn.Click(); err != nil {
		return err
	}

	source, _ := wd.PageSource()

	if strings.Contains(source, "5 komentara") {
		return errors.New("comment limit")
	}

	return logoutUser(wd)
}

func categoriesTest(wd selenium.WebDriver) error {
	if err := findAndClick(wd, selenium.ByCSSSelector, "#sidebar-open"); err != nil {
		return err
	}

	links, err := wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return err
	}

	for _, link := range links {
		txt, err := link.Text()
		if err != nil {
			continue
		}
		if strings.Contains(txt, "Naslovnica") {
			return nil
		}
	}

	return errors.New("test failed; couldn't find \"Naslovnica\"")
}

func tagsTest(wd selenium.WebDriver) error {
	links, err := wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return err
	}

	tagName := ""
	found := false

	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		if strings.Contains(href, "/tagovi/") {
			tagName = strings.Split(href, "/")[2]
			ok := true
			for _, r := range []rune(tagName) {
				if !unicode.IsLetter(r) {
					ok = false
					break
				}
			}
			if ok {
				found = true
				if err := wd.Get("https://klix.ba/" + href); err != nil {
					return err
				}
				break
			}
		}
	}

	if !found {
		return errors.New("couldn't find tag with only letters")
	}

	nav, err := wd.FindElement(selenium.ByCSSSelector, "#breadcrumb")
	if err != nil {
		return err
	}

	ol, err := nav.FindElement(selenium.ByTagName, "ol")
	if err != nil {
		return err
	}

	lis, err := ol.FindElements(selenium.ByTagName, "li")
	if err != nil {
		return err
	}

	if len(lis) != 2 {
		return errors.New("did not get two elements")
	}

	txt, err := lis[1].Text()
	if err != nil {
		return err
	}

	if strings.ToLower(txt) != strings.ToLower(tagName) {
		return errors.New("value mismatch")
	}

	return nil
}

func jobsListingTest(wd selenium.WebDriver) error {
	if err := wd.Get("https://posao.klix.ba"); err != nil {
		return err
	}

	input, err := wd.FindElement(selenium.ByCSSSelector, "#simple-search")
	if err != nil {
		return err
	}

	placeholder, err := input.GetAttribute("placeholder")
	if err != nil {
		return err
	}

	if placeholder != "pretraži oglase za posao" {
		return errors.New("could not find correct placeholder")
	}

	return nil
}

func jobsSearchTest(wd selenium.WebDriver) error {
	if err := wd.Get("https://posao.klix.ba"); err != nil {
		return err
	}

	searchField, err := wd.FindElement(selenium.ByCSSSelector, "#simple-search")
	if err != nil {
		return err
	}

	searchField.SendKeys("saradnik")

	elems, err := wd.FindElements(selenium.ByCSSSelector, ".text-gray-500")
	if err != nil {
		return err
	}

	found := false

	for _, elem := range elems {
		txt, err := elem.Text()
		if err != nil {
			continue
		}

		if strings.Contains(txt, "Detaljna pretraga") {
			found = true
			elem.Click()
		}
	}

	if !found {
		return errors.New("could not find button to search")
	}

	return nil
}

func jobsApplyTest(wd selenium.WebDriver) error {
	if err := wd.Get("https://posao.klix.ba"); err != nil {
		return err
	}

	searchField, err := wd.FindElement(selenium.ByCSSSelector, "#simple-search")
	if err != nil {
		return err
	}

	searchField.SendKeys("saradnik")

	elems, err := wd.FindElements(selenium.ByCSSSelector, ".text-gray-500")
	if err != nil {
		return err
	}

	found := false

	for _, elem := range elems {
		txt, err := elem.Text()
		if err != nil {
			continue
		}
		if strings.Contains(txt, "Detaljna pretraga") {
			found = true
			elem.Click()
		}
	}

	if !found {
		return errors.New("could not find button to search")
	}

	links, err := wd.FindElements(selenium.ByTagName, "a")
	if err != nil {
		return err
	}

	for _, link := range links {
		href, err := link.GetAttribute("href")
		if err != nil {
			continue
		}
		if strings.Contains(href, "/oglasi/") {
			if err := wd.Get(href); err != nil {
				return err
			}
			break
		}
	}

	name, err := wd.FindElement(selenium.ByCSSSelector, "#name")
	if err != nil {
		return err
	}

	name.SendKeys("SVVTprojekat")

	email, err := wd.FindElement(selenium.ByCSSSelector, "#email")
	if err != nil {
		return err
	}

	email.SendKeys("svvtklix@gmail.com")

	phone, err := wd.FindElement(selenium.ByCSSSelector, "#phone")
	if err != nil {
		return err
	}

	phone.SendKeys("+38761000000")

	return nil
}

func init() {
	addTest(test{
		name:  "Register",
		tType: testTypeRegression,
		fn:    registerTest,
	})
	addTest(test{
		name:  "Login",
		tType: testTypeRegression,
		fn:    loginTest,
	})
	addTest(test{
		name:  "Search",
		tType: testTypeRegression,
		fn:    searchTest,
	})
	addTest(test{
		name:  "Article Navigation",
		tType: testTypeRegression,
		fn:    articleNavigationTest,
	})
	addTest(test{
		name:  "Comment",
		tType: testTypeRegression,
		fn:    commentFunctionality,
	})
	addTest(test{
		name:  "Categories",
		tType: testTypeRegression,
		fn:    categoriesTest,
	})
	addTest(test{
		name:  "Tags",
		tType: testTypeRegression,
		fn:    tagsTest,
	})
	addTest(test{
		name:  "Jobs Listing",
		tType: testTypeRegression,
		fn:    jobsListingTest,
	})
	addTest(test{
		name:  "Jobs search",
		tType: testTypeRegression,
		fn:    jobsSearchTest,
	})
	addTest(test{
		name:  "Jobs Apply",
		tType: testTypeRegression,
		fn:    jobsApplyTest,
	})
}
