package model

import "time"

type (
	SwarmNode struct {
		ID           string    `json:"id"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		Role         string    `json:"role"`
		Availability string    `json:"availability"`
		HostName     string    `json:"hostname"`
		State        string    `json:"state"`
		IPaddr       string    `json:"ipaddr"`
	}
)
