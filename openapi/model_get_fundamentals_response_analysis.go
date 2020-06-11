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
// GetFundamentalsResponseAnalysis struct for GetFundamentalsResponseAnalysis
type GetFundamentalsResponseAnalysis struct {
	Datas []GetFundamentalsResponseAnalysisDatas `json:"datas,omitempty"`
	IndustryName string `json:"industryName,omitempty"`
	TotalCount float32 `json:"totalCount,omitempty"`
}