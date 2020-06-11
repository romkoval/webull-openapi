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
// Auth struct for Auth
type Auth struct {
	AccessToken string `json:"access_token,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenExpireTime string `json:"token_expire_time,omitempty"`
	Username string `json:"username,omitempty"`
	Uuid string `json:"uuid,omitempty"`
}