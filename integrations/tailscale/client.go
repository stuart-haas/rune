package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"rune/db"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

var baseUrl = "https://api.tailscale.com/api/v2"

type TailscaleDeviceResponse struct {
	Devices []TailscaleDevice `json:"devices"`
}

type TailscaleDevice struct {
	ID string `json:"id"`
	Hostname string `json:"hostname"`
	Addresses []string `json:"addresses"`
	Permalink string `json:"permalink"`
}

type TailscaleClient struct {
	client *http.Client
	tailnet string
}

func NewTailscaleClient() *TailscaleClient {
	var oauthConfig = &clientcredentials.Config{
		ClientID:     os.Getenv("TAILSCALE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("TAILSCALE_OAUTH_CLIENT_SECRET"),
		TokenURL:     fmt.Sprintf("%s/oauth/token", baseUrl),
	}
	client := oauthConfig.Client(context.Background())

	return &TailscaleClient{
		client: client,
		tailnet: os.Getenv("TAILSCALE_TAILNET"),
	}
}

func (s *TailscaleClient) FetchDevices() ([]TailscaleDevice, error) {
	res, err := s.client.Get(fmt.Sprintf("%s/tailnet/%s/devices", baseUrl, s.tailnet))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	wrapper := TailscaleDeviceResponse{}
	json.Unmarshal(body, &wrapper)

	return wrapper.Devices, nil
}

func (s *TailscaleClient) FetchDevice(id string) (TailscaleDevice, error) {
	res, err := s.client.Get(fmt.Sprintf("%s/device/%s", baseUrl, id))
	if err != nil {
		return TailscaleDevice{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return TailscaleDevice{}, err
	}

	device := TailscaleDevice{}
	json.Unmarshal(body, &device)

	return device, nil
}

func (s *TailscaleClient) CheckDevices() ([]TailscaleDevice, error) {
	// Fetch devices from Tailscale
	devices, err := s.FetchDevices()
	if err != nil {
		return nil, err
	}

	// Get all nodes from local DB with Tailscale provider
	var nodes []db.Node
	result := db.Client.Where("external_provider = ?", "tailscale").Find(&nodes)
	if result.Error != nil {
		return nil, result.Error
	}

	// Create map of DB devices for easy lookup
	dbDevices := make(map[string]bool)
	for _, node := range nodes {
		dbDevices[node.ExternalID] = true
	}

	// Find Tailscale devices that don't exist in DB
	var missingDevices []TailscaleDevice
	for _, device := range devices {
		if !dbDevices[device.ID] {
			missingDevices = append(missingDevices, device)
		}
	}

	return missingDevices, nil
}

func (s *TailscaleClient) SyncDevices() error {
	devices, err := s.FetchDevices()
	if err != nil {
		return err
	}

	for _, device := range devices {
		var node db.Node
		result := db.Client.Where("external_id = ? AND external_provider = ?", device.ID, "tailscale").First(&node)
		
		if result.Error == nil {
			node.Hostname = device.Hostname
			node.SyncedAt = time.Now()
			db.Client.Save(&node)
		} else {
			db.Client.Create(&db.Node{
				Hostname:         device.Hostname,
				ExternalID:       device.ID,
				ExternalProvider: "tailscale",
				SyncedAt:         time.Now(),
			})
		}
	}

	return nil
}

func (s *TailscaleClient) SyncDevice(id string) error {
	device, err := s.FetchDevice(id)
	if err != nil {
		return err
	}

	var node db.Node
	result := db.Client.Where("external_id = ? AND external_provider = ?", device.ID, "tailscale").First(&node)

	if result.Error == nil {
		node.Hostname = device.Hostname
		node.SyncedAt = time.Now()
		db.Client.Save(&node)
	} else {
		db.Client.Create(&db.Node{
			Hostname:         device.Hostname,
			ExternalID:       device.ID,
			ExternalProvider: "tailscale",
			SyncedAt:         time.Now(),
		})
	}

	return nil
}
