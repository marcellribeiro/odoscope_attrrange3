package main

import (
	"encoding/json"
	"fmt"

	"github.com/marcellribeiro/odoscope_attrrange3/repository"
)

func main() {
	sampleData := repository.MakeSampleData()

	b, err := json.Marshal(sampleData)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	// delivery.NewCalculatorHandler(r, config)
}
