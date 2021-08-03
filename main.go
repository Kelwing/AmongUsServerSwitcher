package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

//go:generate goversioninfo -icon=among-us.ico -manifest=ServerSwitcher.exe.manifest

func main() {
	appDataLoc := os.Getenv("APPDATA")
	regionInfoLoc := filepath.Join(appDataLoc, "..\\LocalLow\\Innersloth\\Among Us\\regionInfo.json")

	defer func() {
		fmt.Println("You can now close this window.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}()

	f, err := os.Open(regionInfoLoc)
	if err != nil {
		fmt.Println("Failed to find Among Us Installation")
		return
	}
	decoder := json.NewDecoder(f)
	var regionInfo RegionInfo
	err = decoder.Decode(&regionInfo)
	if err != nil {
		fmt.Println("Failed to read in current server info")
		return
	}
	f.Close()

	// Check if region already exists
	for _, region := range regionInfo.Regions {
		if region.Name == Name && region.Fqdn == Fqdn {
			fmt.Printf("%s already installed, select it in your server selector\n", Name)
			return
		}
	}

	out, err := os.Create(regionInfoLoc)
	if err != nil {
		fmt.Println("Failed to create new regionInfo.json")
		return
	}
	defer out.Close()

	encoder := json.NewEncoder(out)

	regionInfo.Regions = append(regionInfo.Regions, Region{
		Type:          "DnsRegionInfo, Assembly-CSharp",
		Name:          Name,
		TranslateName: 1003,
		Fqdn:          Fqdn,
		DefaultIP:     DefaultIP,
		Port:          Port,
	})

	regionInfo.CurrentRegionIdx = len(regionInfo.Regions) - 1

	err = encoder.Encode(regionInfo)
	if err != nil {
		fmt.Println("Failed to write regionInfo.json")
		return
	}
	fmt.Printf("Successfully installed the %s server.\n", Name)
}
