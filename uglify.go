package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	inputfile   = flag.String("inputfile", "", "JSON file to be processed")
	outputfile  = flag.String("outputfile", "", "JSON file to be processed")
	jsonarray   = flag.Bool("jsonarray", false, "suround file with array bracket")
	prettyprint = flag.Bool("prettyprint", false, "suround file with array bracket")
	debuglevel  = flag.Int("debuglevel", 0, "debug level")
)

//jconst debug nodebugging = true

const debug debugging = true

type debugging bool
type nodebugging bool

func (d debugging) Printf(format string, args ...interface{}) {
	if *debuglevel != 0 {
		log.Printf(format, args...)
	}
}

func (d nodebugging) Printf(format string, args ...interface{}) {
}

func prettyprintf(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func parsej(outfile io.Writer, f interface{}, str *string) {
	switch vv := f.(type) {
	case string:
		fmt.Fprint(outfile, "\"", vv, "\"")
		*str += fmt.Sprint("\"", vv, "\"")
	case float64:
		fmt.Fprint(outfile, vv)
		*str += fmt.Sprint(vv)
	case []interface{}:
		fmt.Fprint(outfile, "[")
		*str += fmt.Sprint("[")
		for i, u := range vv {
			if i != 0 {
				fmt.Fprint(outfile, ",")
				*str += fmt.Sprint(",")
			}
			parsej(outfile, u, str)
		}
		fmt.Fprint(outfile, "]")
		*str += fmt.Sprint("]")
	case map[string]interface{}:
		fmt.Fprint(outfile, "{")
		*str += fmt.Sprint("{")
		l := 0
		for k, v := range vv {
			if l != 0 {
				fmt.Fprint(outfile, ",")
				*str += fmt.Sprint(",")
			}
			fmt.Fprint(outfile, "\"", k, "\":")
			*str += fmt.Sprint("\"", k, "\":")
			parsej(outfile, v, str)
			l++
		}
		fmt.Fprint(outfile, "}")
		*str += fmt.Sprint("}")
	case interface{}:
		fmt.Fprintln(outfile, vv)
		fmt.Fprintln(outfile, "is an interface:")
	default:
		fmt.Fprintln(outfile, vv)
		fmt.Fprintln(outfile, "is of a type I don't know how to handle")
	}
	//	return str
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
		//w, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		w, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			panic(err)
		}
	}
	return w
}

func main() {
	var g interface{}
	var str string

	flag.Parse()
	infile := openStdinOrFile(*inputfile)
	outfile := openStdoutOrFile(*outputfile)

	debug.Printf("inputfile: ", *inputfile)
	debug.Printf("outputfile: ", *outputfile)
	debug.Printf("jsonarray: ", *jsonarray)
	debug.Printf("prettyprint: ", *prettyprint)
	debug.Printf("debuglevel: ", *debuglevel)
	file, err := ioutil.ReadAll(infile)
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)
	if *debuglevel != 0 {
		fmt.Fprintln(outfile, s)
	}

	if *prettyprint == true {
		res, _ := prettyprintf(file)
		ss := string(res)
		//	fmt.Fprintln(outfile, res)
		fmt.Fprintln(outfile, ss)
	} else {

		err1 := json.Unmarshal(file, &g)

		if err1 != nil {
			log.Fatal(err1)
		}

		debug.Printf("xxxxxxxxxxunmarshalled : %s ", g)
		debug.Printf("unmarshalled : ", g)
		debug.Printf(" \n")
		if *jsonarray == true {
			fmt.Fprint(outfile, "[")
			str += fmt.Sprint("[")
		}
		//pj := parsej(outfile, g, str)
		parsej(outfile, g, &str)

		if *jsonarray == true {
			fmt.Fprint(outfile, "]")
			str += fmt.Sprint("]")
		}
		fmt.Fprint(outfile, "\n")
		fmt.Fprint(outfile, str)
		fmt.Fprintln(outfile, " ")
	}

}
