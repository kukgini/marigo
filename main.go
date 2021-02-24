package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hyperledger/aries-framework-go-ext/component/vdr/indy"
	"github.com/hyperledger/aries-framework-go/pkg/client/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries"
	"github.com/hyperledger/indy-vdr/wrappers/golang/vdr"
)

func main() {
	log.Println("load genesis tx from internet")
	genesisFile, err := http.Get("https://raw.githubusercontent.com/sovrin-foundation/sovrin/master/sovrin/pool_transactions_builder_genesis")
	if err != nil {
		log.Fatalln(err)
	}
	defer genesisFile.Body.Close()

	fmt.Println("create indy vdr client")
	indyClient, err := vdr.New(genesisFile.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(indyClient)

	fmt.Println("create indy vdr client")
	client, err := indy.New("sov", indy.WithIndyClient(indyClient))
	if err != nil {
		log.Fatalln(err)
	}
	log.Fatalln(indyClient)

	fmt.Println("create aries framework")
	framework, _ := aries.New()

	fmt.Println("get current aries context")
	ctx, _ := framework.Context()

	fmt.Println("create new didexchange client")
	client, _ := didexchange.New(ctx)

	fmt.Println("create invitation")
	invitation, _ := client.CreateInvitation("test")

	fmt.Println("invitation created:")
	fmt.Println(invitation.Invitation.ID)
	fmt.Println(invitation.Invitation.Label)
	fmt.Println(invitation.Invitation.DID)
	fmt.Println(invitation.Invitation.Type)
}
