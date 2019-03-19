package xml

// CDATA имплементированный тип поля, обрамляющий содержимое служебным тэгом "CDATA"
type CDATA struct {
	Value string `xml:",cdata"`
}

// LinkToResource имплементированный тип поля, отражающий ссылку
type LinkToResource struct {
	Name     string `xml:"name,attr"`
	Internal bool   `xml:"internal,attr"`
	Type     string `xml:"type,attr"`
	Target   CDATA  `xml:"target"`
	//Query    CDATA  `xml:"query"`
	//UUID     string `xml:"uuid"`
}

type linkToPreview struct {
	Type   string `xml:"type,attr"`
	Target CDATA  `xml:"target"`
}
