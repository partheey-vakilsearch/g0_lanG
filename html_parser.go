package main
import(
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
)
func main() {
	input, _ := ioutil.ReadFile("aaa_2.html")
	doc, err := gokogiri.ParseHtml([]byte(input))
	if err != nil {
		fmt.Println("Parsing has error:", err)
		return
	}
	hrs, err := doc.Search("//hr")

	//data :=	make(map[string]string)
	type data struct {
	    application_number string
	    class_number string
	    conflicting_number string
	    journal_number string
	    proprietor_name string
	    proprietor_address string
	    status string
	    image string
	    application_date string
	    user_date string
	    goods_services string
	    vienna_code string 
	}


	hrs_count := len(hrs)
	if hrs_count > 2 {
		for i := 0; i < hrs_count; i++ {
			if hrs[i].PreviousSibling().PreviousSibling().Name() == "table" {
				fmt.Println(hrs[i].PreviousSibling().PreviousSibling().Search(".//tr"))
			}
			//trs, err := hrs[i].PreviousSibling().Search(".//tr")
			//	if err != nil {fmt.Println("search has error:", err)
			//		return}
			//zx, err := trs[0].Search(".//td")
			//	if err != nil {fmt.Println("search has error:", err)
			//		return}
			//fmt.Println(zx[0])				
		}
	}
	doc.Free()
}

/*func get_data(trs,data) {
	trs_count := len(trs)
	for i := 0; i < trs_count; i++ {
		
	}
	
}*/
