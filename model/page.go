package model

import "io/ioutil"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("memory/" + filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("memory/" + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
