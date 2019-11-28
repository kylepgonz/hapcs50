package main

import (
	"html/template"
)

var tmpl = make(map[string]*template.Template)

func init() {
	m := template.Must
	p := template.ParseFiles
	tmpl["index"] = m(p("templates/index.gohtml", "templates/layout.gohtml"))
	tmpl["report1"] = m(p("templates/SoftwareDevelopment_Sprint1.gohtml", "templates/layout.gohtml"))
	tmpl["report2"] = m(p("templates/SoftwareDevelopment_Sprint2.gohtml", "templates/layout.gohtml"))
	tmpl["report3"] = m(p("templates/SoftwareDevelopment_Sprint3.gohtml", "templates/layout.gohtml"))
	tmpl["event-detail"] = m(p("templates/event-detail.gohtml", "templates/layout.gohtml"))
	tmpl["about"] = m(p("templates/aboutpage.gohtml", "templates/layout.gohtml"))
	tmpl["event-new"] = m(p("templates/event-new.gohtml", "templates/layout.gohtml"))
	tmpl["event-search"] = m(p("templates/search.gohtml", "templates/layout.gohtml"))
	tmpl["event-category"] = m(p("templates/searchcategory.gohtml", "templates/layout.gohtml"))
}
