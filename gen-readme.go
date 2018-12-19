package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

const templateFile = "_README.template"

// Generate table in README from profiles in .md and a templaet
// Usage:
//  go run gen-readme.go *.md > README.md

type person struct {
	Name   string
	File   string
	Period string
}

var (
	reName   = regexp.MustCompile("name:([^\n]+)")
	rePeriod = regexp.MustCompile("period:([^\n]+)")
)

func main() {
	var people []*person
	fmt.Fprintf(os.Stderr, "processing %d files\n", len(os.Args)-1)
	for _, file := range os.Args[1:] {
		fmt.Fprintf(os.Stderr, "\tfile %s\n", file)
		person, err := readFrom(file)
		check(err, "failed to read person file "+file)
		people = append(people, person)
	}

	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Period < people[j].Period
	})

	t, err := template.ParseFiles(templateFile)
	check(err, "failed to read template file")

	t.Execute(os.Stdout, people)
}

// readFrom reads .period and .name from a given file.
func readFrom(file string) (*person, error) {
	p := &person{File: file}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	p.Name = strings.TrimSpace(string(reName.FindSubmatch(data)[1]))
	p.Period = strings.TrimSpace(string(rePeriod.FindSubmatch(data)[1]))

	return p, nil
}

func check(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg, err)
		os.Exit(1)
	}
}
