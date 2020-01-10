package main

import (
	"io/ioutil"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
	"math"
	"fmt"
)

type Predictor struct {
	Network
}

func NewPredictor(dataDir string) *Predictor {
	const slice = 4
	rand.Seed(time.Now().UTC().UnixNano())
	net := NewNetwork(slice, 1024, 1, 1)
	for epochs := 0; epochs < 1024; epochs++ {
		trainNetwork(&net, dataDir)
	}
	return &Predictor{net}
}

func (p *Predictor) Predict(inputs []float64) bool {
	normalize(inputs)
	outputs := p.Network.Predict(inputs)
	return outputs.At(0, 0) > 0.7
}

func normalize(x []float64) {
	var sum float64
	for i := range x {
		sum += x[i] * x[i]
	}

	mag := math.Sqrt(sum)

	for i := range x {
		if math.IsInf(x[i]/mag, 0) || math.IsNaN(x[i]/mag) {
			// fallback to zero when dividing by 0
			x[i] = 0
			continue
		}

		x[i] /= mag
	}
}

func trainNetwork(net *Network, dataDir string) {
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		bc, err := ioutil.ReadFile(path.Join(dataDir, file.Name()))
		if err != nil {
			panic(err)
		}
		sc := strings.Split(string(bc), "\n")
		for _, line := range sc {
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			vs := []float64{}
			ss := strings.Split(line, ",")
			for _, sv := range ss {
				v, e := strconv.ParseFloat(strings.Trim(sv, " "), 64)
				if e != nil {
					fmt.Println(line, sv)
					panic(e)
				}
				vs = append(vs, v)
			}
			normalize(vs[0:net.inputs])
			net.Train(vs[0:net.inputs], []float64{vs[net.inputs]})
		}
	}
}