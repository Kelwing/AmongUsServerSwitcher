package main

type RegionInfo struct {
	CurrentRegionIdx int      `json:"CurrentRegionIdx"`
	Regions          []Region `json:"Regions"`
}
type Server struct {
	Name               string `json:"Name"`
	IP                 string `json:"Ip"`
	Port               int    `json:"Port"`
	Players            int    `json:"Players"`
	ConnectionFailures int    `json:"ConnectionFailures"`
}
type Region struct {
	Type          string   `json:"$type"`
	Fqdn          string   `json:"Fqdn,omitempty"`
	DefaultIP     string   `json:"DefaultIp,omitempty"`
	Port          int      `json:"Port,omitempty"`
	Name          string   `json:"Name"`
	TranslateName int      `json:"TranslateName"`
	PingServer    string   `json:"PingServer,omitempty"`
	Servers       []Server `json:"Servers,omitempty"`
}
