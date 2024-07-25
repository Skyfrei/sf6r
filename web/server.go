package main

import(
	"fmt"
	"net/http"
	"os"
)

type Page struct{
	title string
	body string
}

func viewHandler(w http.ResponseWriter, r *http.Request){

	filePath := "./pages/login.html"
	mainPage := loadPage(filePath)
	
	fmt.Fprintf(w, mainPage.body)
}

func check(e error){
	if e != nil{
		panic(e)
	}
}

func readFile(filePath string) string{
	dat, err := os.ReadFile(filePath)
	check(err)
	return string(dat)
}

func loadPage(path string) Page{
	body := readFile(path)
	return Page{body: body}
}

func main(){


	http.HandleFunc("/", viewHandler)







	log.Fatal(http.ListenAndServe(":80", nil))
}


