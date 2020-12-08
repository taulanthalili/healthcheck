package healthcheck



/*

//func mySearch() {
func main() {
	//path := flag.String(NGINX_LOG_PATH, "", "file path to search in")
	search := flag.String("select", "", "search string to look for")
	flag.Parse()

	files, err := ioutil.ReadDir(NGINX_LOG_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Mode().IsDir() {
			//to do: traverse directories recursively
			continue
		}
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		//need to check for file type to detect binary content
		fileType := http.DetectContentType(data)
		for _, line := range strings.Split(string(data), "\n") {
			if strings.Index(line, *search) > -1 {
				if strings.Index(fileType, "text/plain") > -1 {
					fmt.Printf("%s: %s\n", file.Name(), line)
				} else {
					//best guess it is binary file
					//no need to go through all lines in binary file
					fmt.Printf("Binary file %s matches\n", file.Name())
					break
				}
			}
		}
	}
}

*/