package models

import "marcode.io/url-parser/pkg/parser"

type LinkDetails struct {
	PageURL      string
	HTMLVersion  string
	Title        string
	Headings     parser.HeadingCount
	Links        parser.LinkCount
	HasLoginForm bool
}
