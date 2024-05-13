package main

import (
	"errors"
	"fmt"
	"log"
)

func commandMap(cfg *config) error{
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Location areas: ")
	for _, area:= range resp.Results{
		fmt.Printf(" - %v\n",area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cfg *config) error{
	if cfg.prevLocationAreaURL == nil{
		return errors.New("You are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Location areas: ")
	for _, area:= range resp.Results{
		fmt.Printf(" - %v\n",area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}