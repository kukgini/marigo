module github.com/kukgini/marigo

go 1.15

require (
	github.com/gorilla/mux v1.7.3
	github.com/hyperledger/aries-framework-go v0.1.6-0.20201230175227-2b6c7532c011
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
)

replace (
	github.com/kilic/bls12-381 => github.com/trustbloc/bls12-381 v0.0.0-20201104214312-31de2a204df8
	github.com/piprate/json-gold => github.com/trustbloc/json-gold v0.3.1-0.20200414173446-30d742ee949e
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
)
