package xml

import "encoding/xml"

type imagePreview struct {
	Link linkToPreview `xml:"link"`
}

type photo struct {
	Link linkToPreview `xml:"link"`
}

type linksCollection struct {
	LinksList []*LinkToResource `xml:"link"`
}

type fullDescription struct {
	Name    *string          `xml:"name,attr"`
	Links   *linksCollection `xml:"links"`
	Content *CDATA           `xml:"content"`
}

type newsCategories struct {
	LinksList []linkToPreview `xml:"link"`
}

type newsContentImages struct {
	Photo []photo `xml:"Photo"`
}

type newsBlock struct {
	Lang             string            `xml:"language,attr"`
	Title            CDATA             `xml:"Title"`
	ShortDescription CDATA             `xml:"ShortDescription"`
	ImagePreview     imagePreview      `xml:"ImagePreview"`
	Date             string            `xml:"Date"`
	FullDescription  *fullDescription  `xml:"FullDescription"`
	Categories       newsCategories    `xml:"Category"`
	Images           newsContentImages `xml:"Images"`
	Counter          CDATA             `xml:"Counter"`
}

// OpenCMSNewsBlocks тип xml-структуры новости в OpenCMS
type OpenCMSNewsBlocks struct {
	XMLName           xml.Name   `xml:"NewsBlocks"`
	XMLAttr           string     `xml:"xmlns:xsi,attr"`
	XMLSchemaLocation string     `xml:"xsi:noNamespaceSchemaLocation,attr"`
	NewsBlock         *newsBlock `xml:"NewsBlock"`
}
