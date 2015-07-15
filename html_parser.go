package main

import(
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
)

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

func main() {
	input, _ := ioutil.ReadFile("aaa_2.html")
	doc, err := gokogiri.ParseHtml([]byte(input))
	if err != nil {
		fmt.Println("Parsing has error:", err)
		return
	}
	hrs, err := doc.Search("//hr")

	//data :=	make(map[string]string)
	


	hrs_count := len(hrs)
	if hrs_count > 2 {
		for i := 0; i < hrs_count; i++ {
			//var dota data
			as, err := (hrs[i].PreviousSibling().Search(".//tr"))
				if err != nil {fmt.Println("search has error:", err)
				return}
			for i := 0; i < len(as); i++ {
				if i == 0 {
					application_number, err := as[0].Search(".//td")
						if err != nil {fmt.Println("search has error:", err)
						return}
					fmt.Println(application_number[1].InnerHtml())					
				}
			}
			if hrs[i].PreviousSibling().PreviousSibling().Name() == "table" {
				aa := hrs[i].PreviousSibling().PreviousSibling()
				fmt.Println(aa)
				asa ,err := aa.Search("./tr")
					if err != nil {fmt.Println("search has error:", err)
					return}
				if i == 1 {
				viennacode, err := asa[0].Search(".//td")
					if err != nil {fmt.Println("search has error:", err)
					return}
				fmt.Println(viennacode)	
				}
					if err != nil {fmt.Println("search has error:", err)
					return}
			}
		/*	trs, err := hrs[i].PreviousSibling().Search(".//tr")
				if err != nil {fmt.Println("search has error:", err)
					return}
			zx, err := trs[0].Search(".//td")
				if err != nil {fmt.Println("search has error:", err)
					return}
			fmt.Println(zx[0])*/				
		}
	}
	doc.Free()
}

/*func get_data(dota data, AS *Node , count int){
	for i := 0; i < count; i++ {
		if count >1 {
			fmt.Println("other infos")
			fmt.Println(trs)
		}else {
			fmt.Println("viennacode")
			fmt.Println(trs)
		}
	}
}*/
