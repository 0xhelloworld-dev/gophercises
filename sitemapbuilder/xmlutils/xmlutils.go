package xmlutils

import "encoding/xml"

var Sitemap *URLSet

// defines a single <url></url> record
type URL struct {
	Loc string `xml:"loc"`
}

// defines the set holding the urls record
type URLSet struct {
	//no need to explicitly define XMLName - the struct definition is parsed when xml.Marshal is called
	//xml.Name fields define the root element of the struct
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}
