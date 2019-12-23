package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func dirCreation(p string) {
	if _, err := os.Stat(p); err != nil {
		fmt.Printf("INFO: '%s' does not exist. Creating it now.\n", rawDataDirPath)
		os.MkdirAll(p, 0775)
		fmt.Printf("INFO: '%s' directory created successfully.\n", rawDataDirPath)
	}
	fmt.Printf("INFO: '%s' directory already exists. Moving on to next step.\n", rawDataDirPath)
}

func downloadFile(url string, dest string) {
	fmt.Println("INFO: Check if offer file exists.")
	if _, err := os.Stat(dest); err == nil {
		fmt.Printf("ERROR: Offer Index file '%s' already exists locally.\n", dest)
		return
	}
	fmt.Println("INFO: File does not exist. Proceeding to next step.")
	f, err := os.Create(dest)
	if err != nil {
		log.Fatal("Error: File creation failed for '%s'.", dest)
	}
	defer f.Close()

	fmt.Printf("INFO: Offer Index File '%s' does not exist locally. Proceeding to download it.\n", dest)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("ERROR: Download of Offer Index File failed. Remote Url: '%s', Local Path: '%s'\n", offerIndexURL, dest)
	}
	defer resp.Body.Close()
	fmt.Printf("INFO: Writing to local Offer Index File '%s'", dest)
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Printf("INFO: Offer Index File written to '%s' successfully.\n", dest)
	return
}
