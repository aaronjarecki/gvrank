package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"log"
	"math"
	"sort"
	"os"
)

type RankNode struct {
	Name          string
	RankLastRound float64
	CurrentRank   float64
	NumChildren   int
	Parents       []*RankNode
}
func (node RankNode) String() string {
	return fmt.Sprintf("%13s: %.2f%%", node.Name, node.CurrentRank * 100.0)
}
func (node RankNode) DotString() string {
	return fmt.Sprintf("%s [label = \"%s: %.2f%%\"]\n", node.Name, node.Name, node.CurrentRank * 100.0)
}

type ByRank []RankNode
func (a ByRank) Len() int           {return len(a)}
func (a ByRank) Swap(i, j int)      {a[i], a[j] = a[j], a[i]}
func (a ByRank) Less(i, j int) bool {return a[i].CurrentRank > a[j].CurrentRank}

var AllNodes map[string]*RankNode
var DampingFactor float64 = 0.85
var Threshold float64 = 0.001
var InputFile string = "test.gv"
var Verbose bool = false
var OutputDot bool = false
var FileContents string = ""

func main() {
	fmt.Printf("")
	AllNodes = make(map[string]*RankNode, 0)
	parseInput()
	getFileContents()
	processContents()
	postProcessNodes()
	discoverRank()
	outputRankings()
}

func parseInput() {
	if len(os.Args) > 1 && os.Args[1] != "" {
		InputFile = os.Args[1]
	}
	if len(os.Args) > 2 && os.Args[2] == "-dot" {
		OutputDot = true
	} else if len(os.Args) > 2 && os.Args[2] != "" {
		Verbose = true
	}
}

func getFileContents() string {
	body, err := ioutil.ReadFile(InputFile)
	if err != nil {
		log.Fatalf("getFileContents: %v\n", err.Error())
	}
	FileContents = string(body)
	return FileContents
}

func processContents() {
	relationRegEx := regexp.MustCompile("([A-z0-9][A-z0-9]*) *-> *([A-z0-9][A-z0-9]*)")
	for _, line := range strings.Split(FileContents, "\n") {
		matches := relationRegEx.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			processRelation(match[1], match[2])
		}
	}
}

func processRelation(firstName string, secondName string) {
	firstNode := findOrMakeNode(firstName)
	secondNode := findOrMakeNode(secondName)

	firstNode.NumChildren++
	secondNode.Parents = append(secondNode.Parents, firstNode)
}

func findOrMakeNode(name string) *RankNode {
	nodeRef, ok := AllNodes[name]
	if !ok {
		node := RankNode{Name: name}
		node.Parents = make([]*RankNode, 0)
		nodeRef = &node
		AllNodes[name] = nodeRef
	}
	return nodeRef
}

func postProcessNodes() {
	for _, node := range AllNodes {
		if node.NumChildren == 0 {
			// Nodes that don't link to anyone are
			// assumed to link to everyone
			for _, otherNode := range AllNodes {
				if otherNode.Name != node.Name {
					//except for themselves
					node.NumChildren++
					otherNode.Parents = append(otherNode.Parents, node)
				}
			}
		}
	}
}

func discoverRank() {
	percentageToTransfer := 1.0
	maxDelta := 99.0
	initializeRank()
	for maxDelta > Threshold {
		maxDelta = 0.0
		for _, node := range AllNodes {
			delta := getNewRank(node, percentageToTransfer)
			if delta > maxDelta {
				maxDelta = delta
			}
		}
		percentageToTransfer *= 0.8
		finalizeRank()
	}
}

func initializeRank() {
	numNodes := len(AllNodes)
	for _, node := range AllNodes {
		node.RankLastRound = 1.0 / float64(numNodes)
		node.CurrentRank = 0
		if Verbose {
			fmt.Printf("%13s", node.Name)
		}
	}
	if Verbose {
		fmt.Printf("%13s\n", "_Total_")
	}
}

func finalizeRank() {
	total := 0.0
	for _, node := range AllNodes {
		node.RankLastRound = node.CurrentRank
		if Verbose {
			fmt.Printf("%13.3f", node.CurrentRank)
			total += node.CurrentRank
		}
	}
	if Verbose {
		fmt.Printf("%13.3f\n", total)
	}
}

func getNewRank(node *RankNode, percent float64) float64 {
	//node.CurrentRank = node.RankLastRound * (1.0 - percent)
	node.CurrentRank = (1.0 - DampingFactor) / float64(len(AllNodes))
	for _, parent := range node.Parents {
		//node.CurrentRank += (parent.RankLastRound * percent) / float64(parent.NumChildren)
		node.CurrentRank += (parent.RankLastRound * DampingFactor) / float64(parent.NumChildren)
	}
	return math.Abs(node.CurrentRank - node.RankLastRound)
}

func outputRankings() {
	if OutputDot {
		writeNewDot()
	} else {
		printRankings()
	}
}

func writeNewDot() {
	fileName := InputFile + ".ranked"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for index, line := range strings.Split(FileContents, "\n") {
		if index == 1 {
			for _, node := range AllNodes {
				_, err := file.Write([]byte(node.DotString()))
				if err != nil {
					log.Fatalf("Writing Nodes: ", err.Error())
				}
			}
			_, err := file.Write([]byte(line+"\n"))
			if err != nil {
				log.Fatalf("Writing existing content: ", err.Error())
			}
		} else {
			_, err := file.Write([]byte(line+"\n"))
			if err != nil {
				log.Fatalf("Writing existing content: ", err.Error())
			}
		}
	}
}

func printRankings() {
	nodes := make([]RankNode, 0)
	for _, node := range AllNodes {
		nodes = append(nodes, *node)
	}
	sort.Sort(ByRank(nodes))
	for _, node := range nodes {
		fmt.Println(node)
	}
}













