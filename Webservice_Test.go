package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/webconnex/xmlutil"
	//"net/http"
)

func TestMain(t *testing.T) {
	buf := new(bytes.Buffer)

	buf.WriteByte('\n')
	buf.WriteString(generateRequestContent())
	
	// Print request
	bs := buf.Bytes()
	bs = bytes.Replace(bs, []byte{'>', '<'}, []byte{'>', '\n', '<'}, -1)
	fmt.Printf("bs : %s\n\n", bs)
	fmt.Printf("buf : %s\n\n", buf)
	url := "https://wcc.sc.egov.usda.gov/awdbWebService/services HTTP/1.1"
	r, err := http.Post(url, "text/xml; charset=utf-8; ", buf)
	if err != nil {
		// return
	}
	fmt.Printf("Test 1 :%v\n", r)
	fmt.Printf("Test 2 :%v\n", err)

	dec := x.NewDecoder(r.Body)
	find := []xml.Name{
		xml.Name{"", "areYouThereResponse"},
		xml.Name{"http://www.w3.org/2003/05/soap-envelope", "Fault"},
	}
	var start *xml.StartElement
	if start, err = dec.Find(find); err != nil {
		log.Fatal(err)
		fmt.Printf("Resp : %#v\n\n", err)
	}
	if start.Name.Local == "Fault" {
		log.Fatal("Fault!") // Here you can decode a Soap Fault
	}
	var resp areYouThereResponse
	if err := dec.DecodeElement(&resp, start); err != nil {
		log.Fatal(err)
		fmt.Printf("Resp : %#v\n\n", err)
	}
	fmt.Printf("Resp : %#v\n\n", err)
}
func generateRequestContent() string {
	const getTemplate = `
	<soapenv:Envelope xmlns:awd="http://www.wcc.nrcs.usda.gov/ns/awdbWebService" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
	<soapenv:Body>
	  <awd:areYouThere>
	  </awd:areYouThere>
	  </soapenv:Body>
	</soapenv:Envelope>`
	return getTemplate
}
