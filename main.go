package main

import (
	"./myservice"
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/hooklift/gowsdl/soap"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	//withCustomCode()
	// nums_WithGeneratedCode()
	settings2_WithGeneratedCode()
}

func withCustomCode() {
	url := "https://www.dataaccess.com/webservicesserver/NumberConversion.wso"

	payload := []byte(strings.TrimSpace(`
    <?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<NumberToDollars xmlns="http://www.dataaccess.com/webservicesserver/">
		  <dNum>500.50</dNum>
		</NumberToDollars>
	  </soap:Body>
	</soap:Envelope>`,
	))

	httpMethod := "POST"

	// Create the request
	req, err := http.NewRequest(httpMethod, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal("Error on creating request object. ", err.Error())
		return
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	client := &http.Client{}

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on dispatching request. ", err.Error())
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var v Envelope
	if err := xml.Unmarshal(data, &v); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", v.Body.NumberToDollarsResponse.NumberToDollarsResult)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Body    struct {
		Text                    string `xml:",chardata"`
		NumberToDollarsResponse struct {
			Text                  string `xml:",chardata"`
			NumberToDollarsResult string `xml:"NumberToDollarsResult"`
		} `xml:"NumberToDollarsResponse"`
	} `xml:"Body"`
}

func nums_WithGeneratedCode() {

	c := soap.NewClient("https://www.dataaccess.com/webservicesserver/NumberConversion.wso", soap.WithHTTPHeaders(map[string]string{"Content-Type": "text/xml; charset=utf-8"}))

	numConv := myservice.NewNumberConversionSoapType(c)

	var n myservice.NumberToDollars
	n.DNum = 500.51

	res, err := numConv.NumberToDollars(&n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func settings2_WithGeneratedCode() {

	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
		"SOAPAction": "http://www.brinksoftware.com/webservices/settings/v2/ISettingsWebService2/GetEmployees",
		"AccessToken": "Wwe5fMJE90eBQ99MOqQIhQ==",
		"LocationToken": "9IIBlF4lmES4PYIqKtAbbw==",
	}
	c := soap.NewClient("https://api-apiint.brinkpos.net/Settings2.svc", soap.WithHTTPHeaders(headers))

	getEmployees := settings2.NewISettingsWebService2(c)

	var g settings2.GetEmployees

	res, err := getEmployees.GetEmployees(&g)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("%v", res.GetEmployeesResult.Collection.Employee[1].AlternateId)
	}
}