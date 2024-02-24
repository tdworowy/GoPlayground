package main

import (
	"bufio"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalln(err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalln(err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalln(err)
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	staff, err := page.Locator("div[aria-label='Działanie dla tego posta']").All()
	if err != nil {
		log.Fatalln(err)
	}

	for _, entry := range staff {
		entry.Click()
		if err := page.GetByText("Zgłoś reklamę"); err != nil {
			continue
		}
	}
}
