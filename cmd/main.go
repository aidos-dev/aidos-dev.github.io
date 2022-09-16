package main

func main(){
	HandleRequest()
}

func HandleRequest(){
	// Router
	router := http.NewServeMux()

	// files server
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// home page
	router.HandleFunc("/", homePage)

	// Server
	server := http.Server{
		Addr: ":3300532.github.io"
		Handler: router,
	}

	// Run Server
	log.Println("Listening on http://3300532.github.io\n")
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/favicon.ico" {	
		errorHandler(w, http.StatusNotFound, "Error: Page Not Found. Status: 404")	
		return
	}

	if r.Method != "GET" {
		errorHandler(w, http.StatusMethodNotAllowed, "Method Not Allowed. Status: 405")
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/header.html",
		"templates/footer.html",
		"templates/home.html",
	)
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("home handler template parsing error: %v", err)
		return
	}

	tmpl.ExecuteTemplate(w, "home")
}

func errorHandler(w http.ResponseWriter, status int, errMessage) {
	tmpl, err := template.ParseFiles(
		"templates/errorPages.html",
	)
	if err != nil {
		http.Error(w, "Internal server error. Status: 500", http.StatusInternalServerError)
		log.Printf("Error: handler template parsing error: %v", err)
		return
	}

	w.WriteHeader(status)
	err = tmpl.Execute(w, errMessage)

	if err != nil {
		log.Printf("Error: handler template execution error: %s", err.Error())
	}
	log.Println(errMessage)
}