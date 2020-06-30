/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Transfer struct for Transfer
type Transfer struct {
	Amount float32 `json:"amount,omitempty"`
	Date string `json:"date,omitempty"`
	Direction TransferDirection `json:"direction,omitempty"`
	Legs []Leg `json:"legs,omitempty"`
	// Example: COMPLETED
	SubStatus string `json:"sub_status,omitempty"`
	TransferType string `json:"transfer_type,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
}
