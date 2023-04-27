package main

import (
	"encoding/base64"

	"github.com/google/uuid"
)

type VpnConfiguration struct {
	DisplayName   string
	GroupName     string
	RemoteAddress string
	SharedSecret  string
	interiorUUID  string
	exteriorUUID  string
}

func NewVpnConfigurarion(displayName string, groupName string, remoteAddress string, sharedSecret string) *VpnConfiguration {
	return &VpnConfiguration{
		DisplayName:   displayName,
		GroupName:     groupName,
		RemoteAddress: remoteAddress,
		SharedSecret:  sharedSecret,
	}
}

func (vpn *VpnConfiguration) InteriorUUID() string {
	if vpn.interiorUUID == "" {
		vpn.interiorUUID = uuid.NewString()
	}
	return vpn.interiorUUID
}

func (vpn *VpnConfiguration) ExteriorUUID() string {
	if vpn.exteriorUUID == "" {
		vpn.exteriorUUID = uuid.NewString()
	}
	return vpn.exteriorUUID
}

func (vpn *VpnConfiguration) Base64Secret() string {
	return base64.StdEncoding.EncodeToString([]byte(vpn.SharedSecret))
}
