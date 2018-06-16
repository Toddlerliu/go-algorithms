package graph

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func LoadGraphFromFile(graphtype, filename string) interface{} {
	// ioutil.ReadFile()  O_RDONLY
	path := "G:\\Code\\goAlgorithms\\src\\algorithms\\graph\\" + filename
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var line []byte
	reader := bufio.NewReader(file)
	line, err = reader.ReadBytes('\n')
	s := strings.Split(string(line), " ")
	vertexs, _ := strconv.Atoi(s[0])
	edges, _ := strconv.Atoi(strings.Replace(s[1], "\r\n", "", -1))
	var graph1 *AdjacencyMatrix
	var graph2 *AdjacencyList
	if graphtype == "DenseGraph" {
		graph1 = NewDenseGraph(vertexs, false)
		for i := 0; i < edges; i++ {
			line, err = reader.ReadBytes('\n')
			s := strings.Split(string(line), " ")
			v1, _ := strconv.Atoi(s[0])
			v2, _ := strconv.Atoi(strings.Replace(s[1], "\r\n", "", -1))
			graph1.AddEdge(v1, v2)
		}
	} else if graphtype == "SparseGraph" {
		graph2 = NewSparseGraph(vertexs, false)
		for i := 0; i < edges; i++ {
			line, err = reader.ReadBytes('\n')
			s := strings.Split(string(line), " ")
			v1, _ := strconv.Atoi(s[0])
			v2, _ := strconv.Atoi(strings.Replace(s[1], "\r\n", "", -1))
			graph2.AddEdge(v1, v2)
		}
	} else {
		return errors.New("wrong graph type!")
	}

	if graphtype == "DenseGraph" {
		return graph1
	}
	return graph2
}
