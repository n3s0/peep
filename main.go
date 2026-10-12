package main

import (
	"fmt"
	"github.com/n3s0/peep/internal/httpclients"
)

func main() {
	basicData, err := httpclients.GetIpInfoData()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%s\n", basicData.IP)
}
