// Package svix this file is @generated DO NOT EDIT
package models

type TemplateIn struct {
	Description      *string                     `json:"description,omitempty"`
	FeatureFlag      *string                     `json:"featureFlag,omitempty"`
	FilterTypes      []string                    `json:"filterTypes,omitempty"`
	Instructions     *string                     `json:"instructions,omitempty"`
	InstructionsLink *string                     `json:"instructionsLink,omitempty"`
	Kind             *TransformationTemplateKind `json:"kind,omitempty"`
	Logo             string                      `json:"logo"`
	Name             string                      `json:"name"`
	Transformation   string                      `json:"transformation"`
}
