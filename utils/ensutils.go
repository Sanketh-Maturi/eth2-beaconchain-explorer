package utils

import (
	"eth2-exporter/rpc"
	"github.com/ethereum/go-ethereum/common"
	ens "github.com/wealdtech/go-ens/v3"
)

func IsENSDomainPresent(ensDomain string) bool {
	//Resolve ENS address and return false if ENS address doesn't exist
	address, err := ResolveENSDomain(ensDomain)
	if err != nil {
		if address.Hex() != "0" {
			return true
		}
	}

	return false
}

func GetENSAddress(ensDomain string) string {
	address, err := ResolveENSDomain(ensDomain)
	if err != nil {
		return address.Hex()
	}

	return ""
}

func ResolveENSDomain(ensDomain string) (common.Address, error) {
	// Resolve the ENS domain to an Ethereum address
	return ens.Resolve(rpc.CurrentErigonClient.GetNativeClient(), ensDomain)
}
