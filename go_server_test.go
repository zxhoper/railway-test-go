package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
	//	"time"
	//	"bufio"
	//	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestGetRoot(t *testing.T) {

	url := "http://localhost:1323/"
	fmt.Printf("TestGetRoot : %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)

	patt := `.*(Hello.*World!).*`
	if m, err := regexp.MatchString(patt, string(html)); err != nil {
		t.Fatalf("Expected Hello.*World!, Got %v", m)

	}

}

func TestGetUserID(t *testing.T) {

	url := "http://localhost:1323/users/bill "
	fmt.Printf("TestGetUserID : %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)

	patt := `bill`
	if m, err := regexp.MatchString(patt, string(html)); err != nil {
		t.Fatalf("Expected 'bill', Got %v", m)

	}

}

func TestGetQueryString(t *testing.T) {

	url := "http://localhost:1323/show?team=bill&member=dev"
	fmt.Printf("TestGetQueryString :  %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)

	patt := `.*team:bill, member:dev.*`
	//team:bill, member:dev
	if m, err := regexp.MatchString(patt, string(html)); err != nil {
		t.Fatalf("Expected `.*team:bill, member:dev.*`, Got %v", m)

	}

}
