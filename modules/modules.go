package modules

import (
	"html/template"
	"io/ioutil"
	"log"
)

var m = make(map[string]*template.Template)

func Load(name, path string) {

	// Read in template file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	t, err := template.New(name).Parse(string(data))
	if err != nil {
		log.Fatal(err.Error())
	}

	if _, exists := m[name]; exists {
		log.Fatalf("Module %s already loaded", name)
	}

	m[name] = t
}
