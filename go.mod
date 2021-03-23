module github.com/kukgini/marigo

go 1.15

require (
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/golang/snappy v0.0.3 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/hyperledger/aries-framework-go v0.1.6
	github.com/hyperledger/aries-framework-go/component/storage/leveldb v0.0.0-20210318194411-d7768a4e9ef5
	github.com/hyperledger/aries-framework-go/component/storageutil v0.0.0-20210310001230-bc1bd8ea889c
	github.com/hyperledger/aries-framework-go/spi v0.0.0-20210310160016-d5eea2ecdd50
	github.com/mitchellh/mapstructure v1.3.2 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/kilic/bls12-381 => github.com/trustbloc/bls12-381 v0.0.0-20201104214312-31de2a204df8
	github.com/piprate/json-gold => github.com/trustbloc/json-gold v0.3.1-0.20200414173446-30d742ee949e
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
)
