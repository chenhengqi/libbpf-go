package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// all headers contain LIBAPI definitions
var apiHeaders = []string{
	"bpf.h",
	"btf.h",
	"libbpf.h",
	"xsk.h",
}

// regular expression to match LIBAPI definition
//     LIBBPF_API[ LIBBPF_DEPRECATED("msg")] name(type1 arg1,...);
var regex = regexp.MustCompile("LIBBPF_API[a-zA-Z0-9_,*()\\s\"]+;")

// -path path/to/libbpf/src
var libbpfHeadersPath = flag.String("path", "", "path to libbpf headers")

// -dest path/to/libbpf-go
var libbpfGoPath = flag.String("dest", "..", "directory for the generated files")

func main() {
	flag.Parse()
	if *libbpfHeadersPath == "" {
		flag.Usage()
		os.Exit(0)
	}

	for _, header := range apiHeaders {
		data, err := ioutil.ReadFile(filepath.Join(*libbpfHeadersPath, header))
		if err != nil {
			panic(err)
		}
		matches := regex.FindAll(data, -1)
		gen(matches, header, *libbpfGoPath)
	}
}
