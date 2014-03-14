package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type RankNode struct {
	Name    string
	Rank    float32
	Related []RankNode
}

func main() {
	body, err := ioutil.ReadFile("test.gv")
	if err != nil {
		fmt.Println(err.Error())
	}
	relationRegEx := regexp.MustCompile("([A-z][A-z]*) *-> *([A-z][A-z]*)")
	for _, line := range strings.Split(string(body), "\n") {

	}
	fmt.Printf("%v\n", string(body))
}
