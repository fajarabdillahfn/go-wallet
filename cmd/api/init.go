package api

import (
	"log"

	kafkaWallet "github.com/fajarabdillahfn/go-wallet/internal/broker/kafka"
)

func initialize()  {
	log.Println("init broker...")
	kafkaWallet.Init()
	
	go kafkaWallet.AboveThresholdProcessor()
	
	go kafkaWallet.BalanceProcessor()
}