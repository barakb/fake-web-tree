package main

import (
	"fmt"
	"flag"
	"net/http"
	"regexp"
	"strconv"
	"html/template"
)

var re *regexp.Regexp
var treeTemplate *template.Template
var graphTemplate *template.Template
var graph  *bool;

func servTree(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico"{
		http.Error(w, `StatusNotFound`, http.StatusNotFound)
		return
	}
	var n int = 0
	n, _ = getRequestedNode(r.RequestURI)
	childrens := []int { 2 * n, (2 * n) + 1}
	if *graph {
		graphTemplate.Execute(w, struct {
			Childes []int
			Current int
		}{childrens, n / 2})
	}else{
		treeTemplate.Execute(w, struct {
			Childes []int
			Current int
		}{childrens, n / 2})
	}
}

func getRequestedNode(url string) (int, error) {
	if url == "/index.html" {
		return 1, nil
	} else if match := re.FindStringSubmatch(url); 1 < len(match) {
		return strconv.Atoi(match[1])
	} else {
		return 1, nil
	}
}

func main() {
	graph = flag.Bool("graph", false, "create graph instead of tree")
	flag.Parse()
	fmt.Printf("graph is %v\n", *graph)
	re = regexp.MustCompile("/([0-9]+)/index.html")
	treeTemplate = template.New("tree template")
	treeTemplate.Parse(`
<html>
<head></head>
<body>
    <ul>
    {{range $index, $element := .Childes}}
    <li>
       <a href="/{{$element}}/index.html">{{$element}}</a>
    </li>
    {{end}}
    </ul>
</body>
</html>`)

	graphTemplate = template.New("graph template")
	graphTemplate.Parse(`
<html>
<head></head>
<body>
    <ul>
    {{range $index, $element := .Childes}}
    <li>
       <a href="/{{$element}}/index.html">{{$element}}</a>
    </li>
    {{end}}
    </ul>
     <a href="/{{.Current}}/index.html">back</a>
</body>
</html>`)
	fmt.Println("starting web server")
	mux := http.NewServeMux()
	mux.HandleFunc("/", servTree)
	err := http.ListenAndServe(":8080", mux)
	if err != nil{
		fmt.Printf("Error: %v\n", err)
		return
	}

}
