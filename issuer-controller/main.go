package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	protocol "github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/issuecredential"
	client "github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"
	controller "github.com/hyperledger/aries-framework-go/pkg/controller/command/issuecredential"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	holderDid := "did:peer:1zQmZ5XyDuVejqba3nHwKqEVBCCEGaXPnkfK6qKCqTe32b3q"
	issuerDid := "did:peer:1zQmPpxTcgeN9XW4j36DAD7F7766KHzw4rVxhWVgusZJki78"

	//baseUrl := "https://localhost:8082"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	//httpGet(baseUrl + "/connections")

	req := &controller.SendOfferArgs{
		MyDID:    issuerDid,
		TheirDID: holderDid,
		OfferCredential: &client.OfferCredential{
			Type:    "https://didcomm.org/issue-credential/2.0/offer-credential",
			Comment: "This is comment",
			Formats: &protocol.Format
		},
	}
	log.Printf("%+v", req)
	//httpPost(baseUrl + "/issuecredential/send-offer")

	//jsonDataPostExample()
}

func httpGet(url string) {
	log.Printf("http get %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	str := string(respBody)
	log.Printf(str)
}

func httpPost(url, data string) {
	log.Println("json data post example")
	person := Person{Name: "Alex", Age: 10}
	pbytes, _ := json.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post(url, "application/json", buff)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	str := string(respBody)
	log.Println(str)
}
