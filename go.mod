module github.com/kukgini/marigo

go 1.15

require (
	github.com/hyperledger/aries-framework-go v0.1.6-0.20201230175227-2b6c7532c011
	github.com/hyperledger/aries-framework-go-ext v0.0.0-20210209170459-14c492334960 // indirect
	github.com/hyperledger/aries-framework-go-ext/component/vdr/indy v0.0.0-20210209170459-14c492334960 // indirect
	github.com/hyperledger/indy-vdr/wrappers/golang v0.0.0-20201031155907-5f437d26ed71 // indirect
	github.com/jaytaylor/mockery-example v0.0.0-20170323165341-bf04a9147d8e // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	goji.io v2.0.2+incompatible // indirect
)

replace (
	github.com/kilic/bls12-381 => github.com/trustbloc/bls12-381 v0.0.0-20201104214312-31de2a204df8
	github.com/piprate/json-gold => github.com/trustbloc/json-gold v0.3.1-0.20200414173446-30d742ee949e
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
)
