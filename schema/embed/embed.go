//+build ignore

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type schemaID struct {
	ID string `json:"id"`
}

func getSchemaIDFromReader(r io.Reader) (string, error) {
	dec := json.NewDecoder(r)

	var schID schemaID
	err := dec.Decode(&schID)
	if err != nil {
		return "", err
	}

	return schID.ID, nil
}

// Reads all .json files in the current folder
// and encodes them as strings literals in schemas.go
func main() {
	out, err := os.OpenFile("schemas.go", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	out.Write([]byte(`package schema

// Code generated by go generate; DO NOT EDIT.
//go:generate go run embed/embed.go


var SchemaByURL = map[string]string{
`))
	fs, _ := ioutil.ReadDir(".")
	first := true
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			//out.Write([]byte("\t" + capitalize(strings.TrimSuffix(f.Name(), ".json")) + " = `"))
			f, err := os.Open(f.Name())
			if err != nil {
				log.Fatalf("failed to read %s: %s", f.Name(), err.Error())
			}

			id, err := getSchemaIDFromReader(f)
			if err != nil {
				log.Fatalf("failed to extract id from %s: %s", f.Name(), err.Error())
			}

			f.Seek(0, 0)

			if !first {
				out.Write([]byte(","))
			} else {
				first = false
			}

			out.Write([]byte("\n"))

			out.Write([]byte("\"" + id + "\" : `"))

			io.Copy(out, f)
			out.Write([]byte("`"))
		}

	}
	out.Write([]byte("}\n"))
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	if len(s) == 1 {
		return strings.ToTitle(s[0:1])
	}

	return strings.ToTitle(s[0:1]) + s[1:]
}
