package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Names []N

type N struct {
	Name    string
	Types   string
	Comment string
}

func main() {
	var class string
	var ns Names

	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		err := fmt.Errorf("引数足りないよ")
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	for i, n := range args {
		if i == 0 {
			class = n
		} else {
			f := strings.Index(n, ":")
			l := strings.LastIndex(n, ":")
			if f == -1 {
				err := fmt.Errorf(": フォーマットが間違ってるよ")
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}

			if f == l {
				n := N{Name: n[:f], Types: n[f+1:]}
				ns = append(ns, n)
			} else {
				n := N{Name: n[:f], Types: n[f+1 : l], Comment: n[l+1:]}
				ns = append(ns, n)
			}
		}
	}

	var params []string
	var setters []string
	var getters []string

	for _, n := range ns {
		params = append(params, makeParam(n.Name, n.Types, n.Comment))
		setters = append(setters, makeSetter(n.Name, n.Types))
		getters = append(getters, makeGetter(n.Name, n.Types))
	}

	php := "<?php\n\n"

	className := makeLeadClass(class)

	php = php + className

	for _, pa := range params {
		php = php + pa + "\n"
	}

	for _, se := range setters {
		php = php + se + "\n"
	}

	for _, ge := range getters {
		php = php + ge + "\n"
	}

	php = php + "}\n"

	write(class, php)
	os.Exit(0)
}

func write(fileName string, php string) {
	fileName = fileName + ".php"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, php)
}

func makeLeadClass(name string) string {
	m := strings.ToUpper(name[:1])
	m1 := name[1:]
	class := m + m1

	t := "class %s\n{\n"
	s := fmt.Sprintf(t, class)
	return s
}

func makeParam(name string, types string, comment string) string {

	t := "    /** @var %s %s */\n    private $%s;\n"
	s := fmt.Sprintf(t, types, comment, name)
	return s

}

// makeGetter is create Get Method
func makeGetter(name string, types string) string {
	m := strings.ToUpper(name[:1])
	m1 := name[1:]
	method := m + m1

	t := "    /**\n     * %s Getter\n     * \n     * @return %s %s\n     */\n    public function get%s()\n    {\n        return $this->%s;\n    }\n"
	s := fmt.Sprintf(t, method, types, name, method, name)
	return s
}

// makeSetter is create Set Method
func makeSetter(name string, types string) string {
	m := strings.ToUpper(name[:1])
	m1 := name[1:]
	method := m + m1

	t := "    /**\n     * %s Setter\n     * \n     * @param %s %s\n     */\n    public function set%s($%s)\n    {\n        $this->%s = $%s;\n    }\n"
	s := fmt.Sprintf(t, method, types, name, method, name, name, name)
	return s
}
