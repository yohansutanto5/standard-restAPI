package model

import "time"



type CustomLog struct {
	TransactionID int    `json:"transactionID"`
	Code          string `json:"code"`
	Status        int    `json:"status"`
	Message       string `json:"message"`
	Method        string `json:"method"`
	Path          string
	Duration      time.Duration
	ClientIp      string
	Agent         string

	Data interface{} `json:"data"`
}
