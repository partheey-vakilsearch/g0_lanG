package main

import(
	"fmt"
	"github.com/moovweb/gokogiri"
	"io/ioutil"
	"strings"
	"encoding/json"
	"testing"
	"strconv"
)

type data struct {
    Application_number string 
    Class_number string 
    Conflicting_number string 
    Journal_number string 
    Proprietor_name string 
    Proprietor_address string 
    Status string 
    Application_date string 
    User_date string 
    Goods_services string 
    Vienna_code string  
}

var alpha = []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

func main() {
	time := testing.Benchmark(Benchmarkparser)
	fmt.Println(time)
}

func Benchmarkparser(b *testing.B) {
	caller()
}

func caller() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			for k := 0; k < 25; k++ {
				for l := 0; l < 45; l++ {
					path := "/home/partheey/data/trademark/results/"+alpha[i]+"/"+alpha[j]+"/"+alpha[i]+alpha[j]+alpha[k]+"_"+strconv.Itoa(l)+".html"
					parser(path)
				}
				
			}
			
		}
		
	}
}

func parser(path string) {

	input, _ := ioutil.ReadFile(path)
	doc, err := gokogiri.ParseHtml([]byte(input))
	if err != nil {
		fmt.Println("Parsing has error:", err)
		return
	}
	hrs, err := doc.Search("//hr")

	hrs_count := len(hrs)
	if hrs_count > 2 {
		for i := 0; i < hrs_count; i++ {
		var dota data
			as, err := (hrs[i].PreviousSibling().Search(".//tr"))
				if err != nil {fmt.Println("search has error:", err)
				return}
			if len(as) > 2 {
				for i := 0; i < len(as); i++ {
					record, err := as[i].Search(".//td")
					if i == 0 {
							if err != nil {fmt.Println("search has error:", err)
							return}
						dota.Application_number = record[0].InnerHtml()
						dota.Class_number = record[1].InnerHtml()
						dota.Conflicting_number = record[2].InnerHtml()
						dota.Journal_number = record[3].InnerHtml()
						dota.Proprietor_name = record[4].InnerHtml()
						dota.Proprietor_address = record[5].InnerHtml()
						dota.Status = record[6].InnerHtml()
					} else if i == 1{
						app_date := strings.Split(record[0].InnerHtml(), ":") 
						dota.Application_date = app_date[1]
						usr_date := strings.Split(record[1].InnerHtml(), ":") 
						dota.User_date = usr_date[1]
					} else{
						gs := strings.Split(record[0].InnerHtml(), ":")
						dota.Goods_services = gs[1]
						/*fmt.Println(dota)*/
					}
				}
			} else if len(as) == 1 {
				record, err := as[0].Search(".//td")
					if err != nil {fmt.Println("search has error:", err)
					return}
				vc := strings.Split(record[1].InnerHtml(), ":")
				dota.Vienna_code = vc[1]
			}

			if hrs[i].PreviousSibling().PreviousSibling().Name() == "table" {
				asa, err := hrs[i].PreviousSibling().PreviousSibling().Search(".//tr")
					if err != nil {fmt.Println("search has error:", err)
					return}
				if len(asa) > 2 {
					for i := 0; i < len(asa); i++ {
						record, err := asa[i].Search(".//td")
						if i == 0 {
								if err != nil {fmt.Println("search has error:", err)
								return}
							dota.Application_number = record[0].InnerHtml()
							dota.Class_number = record[1].InnerHtml()
							dota.Conflicting_number = record[2].InnerHtml()
							dota.Journal_number = record[3].InnerHtml()
							dota.Proprietor_name = record[4].InnerHtml()
							dota.Proprietor_address = record[5].InnerHtml()
							dota.Status = record[6].InnerHtml()
						} else if i == 1{
							app_date := strings.Split(record[0].InnerHtml(), ":") 
							dota.Application_date = app_date[1]
							usr_date := strings.Split(record[1].InnerHtml(), ":") 
							dota.User_date = usr_date[1]
						} else{
							gs := strings.Split(record[0].InnerHtml(), ":")
							dota.Goods_services = gs[1]
						}
					}
				} 
			}

		res1D := data{ dota.Application_number, dota.Class_number, dota.Conflicting_number, dota.Journal_number, dota.Proprietor_name, dota.Proprietor_address, dota.Status, dota.Application_date, dota.User_date, dota.Goods_services, dota.Vienna_code }

		final_data, _ := json.Marshal(res1D)
		/*fmt.Println(string(final_data))*/
		path := "/home/partheey/gotry/results/" + dota.Class_number + "_" + dota.Application_number +".json" 	
		ioutil.WriteFile(path,final_data,0644)
		}
	}
	doc.Free()
}
