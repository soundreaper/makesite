package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// Referenced Dani's Solution!

// Page structure contains all the information needed to generate HTML page from text file.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func createPageFromTextFile(filePath string) Page {
	// Makes sure we can read the content of the file.
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// Get the name of the file without the .txt
	fileNameWithoutExtension := strings.Split(filePath, ".txt")[0]

	// Instantiate new Page and populate each of the fields and return.
	return Page{
		TextFilePath: filePath,
		TextFileName: fileNameWithoutExtension,
		HTMLPagePath: fileNameWithoutExtension + ".html",
		Content:      string(fileContents),
	}
}

func renderTemplateFromPage(templateFilePath string, page Page) {
	// Create a new template file in memory
	t := template.Must(template.New(templateFilePath).ParseFiles(templateFilePath))

	// Create new, blank HTML file
	newFile, err := os.Create(page.HTMLPagePath)
	if err != nil {
		panic(err)
	}

	// Executing injects the Page instance's data into the template file
	// we create in memory earlier. This allows us to see the text file's
	// content in the rendered template.
	t.Execute(newFile, page)
	fmt.Println("Generate following file in local directory: ", page.HTMLPagePath)
}

func main() {
	// Adds a console flag `--file=` to reference a particular text file
	var textFilePath string
	flag.StringVar(&textFilePath, "file", "", "path to a text file")
	flag.Parse()

	// Checks to see if `--file=` flag is empty
	if textFilePath == "" {
		panic("Missing the --file flag! Please supply one.")
	}

	// Read in specified text file and instantiate Page with it's information
	newPage := createPageFromTextFile(textFilePath)

	// Use the instantiated Page to generate a new HTML page based on the provided template
	renderTemplateFromPage("template.tmpl", newPage)
}
