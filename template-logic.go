package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

func LoadTemplate(path string) *template.Template {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Fatal(err)
	}

	return tmpl
}

func GenerateProfile(tmpl *template.Template, displayName string, groupName string, remoteAddress string, sharedSecret string) ([]byte, error) {
	vpn := NewVpnConfigurarion(displayName, groupName, remoteAddress, sharedSecret)

	var buf bytes.Buffer

	// fill in the template
	err := tmpl.Execute(&buf, vpn)
	if err != nil {
		return nil, fmt.Errorf("couldn't create profile: %w", err)
	}

	return buf.Bytes(), nil
}
