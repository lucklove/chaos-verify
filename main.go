package main

import (
	"fmt"
	"path"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	bc, err := ioutil.ReadFile(path.Join("test", "metric.csv"))
	if err != nil {
		panic(err)
	}
	sc := strings.Split(string(bc), ",")
	vs := []float64{}
	p := NewPredictor("data")
	for _, sv := range sc {
		v, e := strconv.ParseFloat(strings.Trim(sv, "\n "), 64)
		if e != nil {
			panic(e)
		}
		vs = append(vs, v)
		if len(vs) == 4 {
			t := make([]float64, 4)
			copy(t, vs)
			if p.Predict(t) {
				fmt.Println(vs)
			}
			vs = vs[1:]
		}
	}
}
