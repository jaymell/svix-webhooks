// Package svix this file is @generated DO NOT EDIT
package models

type BackgroundTaskOut struct {
	Data   map[string]any       `json:"data"`
	Id     string               `json:"id"`
	Status BackgroundTaskStatus `json:"status"`
	Task   BackgroundTaskType   `json:"task"`
}
