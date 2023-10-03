package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/pkcs12"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	templatePath := flag.String("template", "template.configprofile", "path to the template to use")
	csvPath := flag.String("csv", "vpns.csv", "path to the CSV of VPN configs")
	p12Path := flag.String("certificate", "engineering.p12", "path to certificate (in PKCS12 format)")
	p12Password := flag.String("password", "", "password for PKCS12 certificate")
	flag.Parse()

	if *p12Password == "" {
		fmt.Print("Please enter the PKCS12 password: ")
		bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Println("Error reading password:", err)
			os.Exit(1)
		}
		*p12Password = string(bytePassword)
		fmt.Println() // Print a newline after password input
	}

	tmpl := LoadTemplate(*templatePath)

	p12File, err := os.Open(*p12Path)
	if err != nil {
		log.Fatal(err)
	}
	defer p12File.Close()

	p12Bytes, err := io.ReadAll(p12File)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, certificate, err := pkcs12.Decode(p12Bytes, *p12Password)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	csvfile, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	// TODO MAYBE make this tolerate columns in different orderd
	// discard the header line
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// process csv
	for {
		data, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		profile, err := GenerateProfile(tmpl, data[0], data[2], data[1], data[3])
		if err != nil {
			log.Fatal(err)
		}

		signedProfile, err := Sign(privateKey, certificate, profile)
		if err != nil {
			log.Fatal(err)
		}

		outPath := data[0] + ".mobileconfig"
		outFile, err := os.Create(outPath)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		// write it
		outFile.Write(signedProfile)
		log.Printf("created %s", outPath)
	}

}
