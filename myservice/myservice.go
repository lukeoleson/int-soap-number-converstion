// Code generated by gowsdl DO NOT EDIT.

package myservice

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type NumberToWords struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWords"`

	UbiNum uint64 `xml:"ubiNum,omitempty" json:"ubiNum,omitempty"`
}

type NumberToWordsResponse struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToWordsResponse"`

	NumberToWordsResult string `xml:"NumberToWordsResult,omitempty" json:"NumberToWordsResult,omitempty"`
}

type NumberToDollars struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToDollars"`

	DNum float64 `xml:"dNum,omitempty" json:"dNum,omitempty"`
}

type NumberToDollarsResponse struct {
	XMLName xml.Name `xml:"http://www.dataaccess.com/webservicesserver/ NumberToDollarsResponse"`

	NumberToDollarsResult string `xml:"NumberToDollarsResult,omitempty" json:"NumberToDollarsResult,omitempty"`
}

type NumberConversionSoapType interface {

	/* Returns the word corresponding to the positive number passed as parameter. Limited to quadrillions. */
	NumberToWords(request *NumberToWords) (*NumberToWordsResponse, error)

	NumberToWordsContext(ctx context.Context, request *NumberToWords) (*NumberToWordsResponse, error)

	/* Returns the non-zero dollar amount of the passed number. */
	NumberToDollars(request *NumberToDollars) (*NumberToDollarsResponse, error)

	NumberToDollarsContext(ctx context.Context, request *NumberToDollars) (*NumberToDollarsResponse, error)
}

type numberConversionSoapType struct {
	client *soap.Client
}

func NewNumberConversionSoapType(client *soap.Client) NumberConversionSoapType {
	return &numberConversionSoapType{
		client: client,
	}
}

func (service *numberConversionSoapType) NumberToWordsContext(ctx context.Context, request *NumberToWords) (*NumberToWordsResponse, error) {
	response := new(NumberToWordsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *numberConversionSoapType) NumberToWords(request *NumberToWords) (*NumberToWordsResponse, error) {
	return service.NumberToWordsContext(
		context.Background(),
		request,
	)
}

func (service *numberConversionSoapType) NumberToDollarsContext(ctx context.Context, request *NumberToDollars) (*NumberToDollarsResponse, error) {
	response := new(NumberToDollarsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *numberConversionSoapType) NumberToDollars(request *NumberToDollars) (*NumberToDollarsResponse, error) {
	return service.NumberToDollarsContext(
		context.Background(),
		request,
	)
}
