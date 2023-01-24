package model

import "time"

type InputDeposit struct {
	Id     int32   `json:"id"`
	Amount float64 `json:"amount"`
}

type WalletDetails struct {
	WalletId       int     `json:"wallet_id"`
	Balance        float64 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}

type ThresholdCheck struct {
	Amounts        float64
	StartTime      time.Time
	AboveThreshold bool
}

var ThresholdList = make(map[int32]*ThresholdCheck)

var WalletsBalance = make(map[int32]float64)
