package main

import (
	"time"
)

type LinkType struct {
	URL        string
	Title      string
	Tag        string
	CreateTime time.Time
}

type LinksType struct {
	Links []LinkType
}

type CreateRequest struct {
	URL        string
	Title      string
	Tag        string
	CreateTime time.Time
}

type CreateResponse struct {
	Message bool
}

type FetchLinkRequest struct {
	SearchTerm string
}

type FetchLinkResponse struct {
	Message bool
	Links   []LinkType
}

type FetchTagRequest struct {
	Tag string
}

type FetchTagResponse struct {
	Message bool
	Size    int
	Links   []LinkType
}
