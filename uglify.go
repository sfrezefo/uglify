package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	inputfile  = flag.String("inputfile", "", "JSON file to be processed")
	outputfile = flag.String("outputfile", "", "JSON file to be processed")
	jsonarray  = flag.Bool("jsonarray", false, "suround file with array bracket")
)

func parsej(outfile io.Writer, f interface{}) {
	switch vv := f.(type) {
	case string:
		fmt.Print("\"", vv, "\"")
	case float64:
		fmt.Print(vv)
	case []interface{}:
		fmt.Print("[")
		for i, u := range vv {
			if i != 0 {
				fmt.Print(",")
			}
			parsej(outfile, u)
		}
		fmt.Print("]")
	case map[string]interface{}:
		fmt.Print("{")
		l := 0
		for k, v := range vv {
			if l != 0 {
				fmt.Print(",")
			}
			fmt.Print("\"", k, "\":")
			parsej(outfile, v)
			l++
		}
		fmt.Print("}")
	case interface{}:
		fmt.Fprintln(outfile, vv)
		fmt.Fprintln(outfile, "is an interface:")
	default:
		fmt.Fprintln(outfile, vv)
		fmt.Fprintln(outfile, "is of a type I don't know how to handle")
	}
}

func openStdinOrFile(filename string) io.Reader {
	var err error
	r := os.Stdin
	if filename != "" {
		r, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
	}
	return r
}

func openStdoutOrFile(filename string) io.Writer {
	var err error
	w := os.Stdout
	if filename != "" {
		w, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
	return w
}

func main() {
	var g interface{}

	flag.Parse()
	infile := openStdinOrFile(*inputfile)
	outfile := openStdoutOrFile(*outputfile)

	fmt.Fprintln(outfile, outfile, "xxxxxxxxxx")

	fmt.Fprintln(outfile, "inputfile: ", *inputfile)
	fmt.Fprintln(outfile, "outputfile: ", *outputfile)
	fmt.Fprintln(outfile, "jsonarray: ", *jsonarray)
	//	if *inputfile == "" {
	//		in = os.Stdin
	//	}
	//	file, err := ioutil.Open
	//	File("test1.json")
	//if *outputfile == "" {
	//}

	//	file, err := ioutil.ReadFile("test1.json")
	file, err := ioutil.ReadAll(infile)
	//file, err := infile.Read()
	s := string(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(outfile, s)
	err1 := json.Unmarshal(file, &g)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Fprintln(outfile, "unmarshalled : ", g)
	fmt.Fprintln(outfile, " ")
	if *jsonarray == true {
		fmt.Print("[")
	}
	parsej(outfile, g)
	if *jsonarray == true {
		fmt.Print("]")
	}
	fmt.Fprintln(outfile, " ")

}
