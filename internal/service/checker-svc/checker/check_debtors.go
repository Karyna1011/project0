package checker

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3"
)

func Checker(Contract *bind.BoundContract, log *logan.Entry, addressPerson common.Address) bool {
	result := make([]interface{}, 0)

	log.Info(addressPerson.String())

	err := Contract.Call(&bind.CallOpts{}, &result, "getAvailableLiquidity", addressPerson)
	if err != nil {
		log.WithError(err).Error("error during calling contract")
		return false
	}

	log.Info("result=", result, " result[0]=", result[0], " result[1]=", result[1])

	if fmt.Sprintf("%v", result[1]) != "0" {
		log.Info("NOT DEBTOR")
		log.Info("myResult=", result[1])
		return false
	}
	log.Info("DEBTOR")
	return true
}
