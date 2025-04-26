package models

type BlogList struct {
	Contents []Blog `json:"contents"`
	TotalCount int    `json:"totalCount"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}