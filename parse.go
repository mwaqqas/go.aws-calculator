package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func parseOfferIndex(file string) (out offerIndex, err error) {
	j, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
		return out, err
	}
	plan, _ := ioutil.ReadAll(j)
	err = json.Unmarshal(plan, &out)
	if err != nil {
		fmt.Println(err)
		return out, err
	}
	return out, err
}

func parseRegionIndex(file string) (out regionIndex, err error) {
	j, err := os.Open(file)
	if err != nil {
		fmt.Print(err)
		return out, err
	}
	plan, _ := ioutil.ReadAll(j)
	err = json.Unmarshal(plan, &out)
	if err != nil {
		fmt.Println(err)
		return out, err
	}
	return out, err
}
