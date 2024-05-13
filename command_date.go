package main

import (
	"fmt"
	"time"
)

func commandDate(cfg *config) error{
	currentTime := time.Now()
	fmt.Println(currentTime.String())
	return nil
}