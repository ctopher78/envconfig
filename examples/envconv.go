package main

import (
	"fmt"
	"log"
	"time"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Debug      bool
	Port       int `default:"8080"`
	User       string `default:"Chris"`
	Host       Host `default:"foo%s"`
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
}

type Host struct {
	IP	string
}

func (h *Host) Decode(value string) error {
	if _, ok := os.LookupEnv("MYAPP_HOST"); ok {
		return nil
	}
	h.IP = fmt.Sprintf(value, "bar")
	return nil
}

func main() {
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s\nHost: %s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout, s.Host.IP)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
}