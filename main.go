package main

import (
	"strings"

	"github.com/MalikovSoft/coverted_ctx_links_validator/database"
	opencms "github.com/MalikovSoft/coverted_ctx_links_validator/xml"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := database.InitDatabase(`root@/ncfu?charset=utf8&parseTime=true&loc=Local`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	XMLs := opencms.GetAllOpenCMSNews("./files-to-convert/")
	_, linksFromDatabase := database.GetAllLinksToResources(db)
	outputXMLs := make(map[string]*opencms.OpenCMSNewsBlocks, 0)

	for fullFilename, currentXML := range XMLs {
		currentXML.XMLSchemaLocation = `opencms://system/modules/ru.soft.malikov.web/schemas/NewsBlock.xsd`
		currentXML.XMLAttr = `http://www.w3.org/2001/XMLSchema-instance`
		content := currentXML.NewsBlock.FullDescription.Content.Value

		htmlCtx, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
		htmlCtx.Find("a").Each(func(loop int, a *goquery.Selection) {
			tmpLink, tmpLinkExists := a.Attr("href")
			if tmpLinkExists {
				href := strings.TrimSpace(tmpLink)
				if href == "" || linksFromDatabase[href] == "1" {
					a.ReplaceWithSelection(a.Contents())
				} else {
					a.SetAttr("href", linksFromDatabase[href])
				}

			}

		})
	}
}
