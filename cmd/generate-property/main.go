package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"strings"
)

var interfaceName string
var propertyName string
var typeName string

func parseTmpl(tmpl string, data map[string]interface{}) ([]byte, error) {
	parsedTmpl, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := parsedTmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}

func main() {
	flag.StringVar(&interfaceName, "i", "", "interface name")
	flag.StringVar(&propertyName, "p", "", "property name")
	flag.StringVar(&typeName, "t", "", "type name")
	flag.Parse()

	propertyNameUpper := strings.ToUpper(string(propertyName[0])) + propertyName[1:]
	typeNameUpper := strings.ToUpper(string(typeName[0])) + typeName[1:]
	if typeName == "float" {
		typeName = "float64"
	}
	t := `

{{ .property_upper }}() {{ .type }}
Set{{ .property_upper }}(s {{ .type }})

func (r *{{ .interface }}) {{ .property_upper }}() {{ .type }} {
	return r.v.Get("{{ .property }}").{{ .type_upper }}()
}

func (r *{{ .interface }}) Set{{ .property_upper }}(s {{ .type }}) {
	r.v.Set("{{ .property }}", s)
}

`
	v := map[string]interface{}{
		"interface":      interfaceName,
		"property":       propertyName,
		"property_upper": propertyNameUpper,
		"type":           typeName,
		"type_upper":     typeNameUpper,
	}
	bs, err := parseTmpl(t, v)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf(string(bs))
	}
}
