package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func main() {
	m := yaml.MapSlice{}
	if err := yaml.Unmarshal(GetYaml(), &m); err != nil {
		log.Panic(err)
	}

	routes := GetYamlRoutes(m)

	PrintMapSlice(m)

	fmt.Println()
	for _, v := range routes {
		fmt.Println(v)
	}
}

func GetYamlRoutes(m yaml.MapSlice) [][]string {
	rst := [][]string{}
	var walk func(m yaml.MapSlice, depth int, parent []string)
	walk = func(m yaml.MapSlice, depth int, parent []string) {
		n := []yaml.MapItem(m)
		for i, v := range n {
			tmp := make([]string, depth+1)
			copy(tmp, append(parent, fmt.Sprint(i, ":", v.Key)))
			rst = append(rst, tmp)

			if slice, ok := n[i].Value.(yaml.MapSlice); ok == true {
				walk(slice, depth+1, tmp)
			}
		}
	}

	walk(m, 0, []string{})
	return rst
}

func PrintMapSlice(m yaml.MapSlice) {
	var walk func(m yaml.MapSlice, depth int)
	walk = func(m yaml.MapSlice, depth int) {
		for i, v := range m {
			for j := 0; j < depth; j++ {
				fmt.Print("  ")
			}
			fmt.Print(i, " ", v.Key, " : ")
			if _, ok := v.Value.(yaml.MapSlice); ok == false {
				fmt.Print(v.Value)
			}
			fmt.Println()

			if slice, ok := v.Value.(yaml.MapSlice); ok == true {
				walk(slice, depth+1)
			}
		}
	}
	walk(m, 0)
}

func GetYaml() []byte {
	b, err := ioutil.ReadFile("sample.yaml")
	if err != nil {
		log.Panic(err)
	}
	return b
}
