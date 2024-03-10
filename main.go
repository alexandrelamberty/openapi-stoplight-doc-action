package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

// Data structure to hold data for the template
type PageData struct {
	Title string
	File  string
}

func main() {
	// Parse command line arguments
	var file string
	var title string
	var directory string
	flag.StringVar(&file, "file", "api.yaml", "Path to the input file")
	flag.StringVar(&title, "title", "API Documentation", "Custom title for the HTML doc")
	flag.StringVar(&directory, "directory", "", "Path to the output directory")
	flag.Parse()

	// Check if the input file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist\n", file)
		os.Exit(1)
	}

	// HTML template
	htmlTemplate := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, shrink-to-fit=no"
    />
    <title>{{.Title}}</title>
    <!-- Embed elements Elements via Web Component -->
    <script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
    <link
      rel="stylesheet"
      href="https://unpkg.com/@stoplight/elements/styles.min.css"
    />
  </head>
  <style>
    body {
      font-family: ui-sans-serif, sans-serif;
      font-size: 12px;
      height: 100vh;
    }
  </style>
  <body>
    <elements-api
      apiDescriptionUrl="{{.File}}"
      router="hash"
    />
  </body>
</html>
	`

	// Create a new template with the provided or default
	tmpl, err := template.New("index.html").Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	// Data to be used in the template
	data := PageData{
		Title: title,
		File:  file,
	}

	// Check if the directory parameter is provided
	if directory != "" {
		fmt.Printf("Output directory: %s\n", directory)
		// Create the directory if it doesn't exist
		err = os.Mkdir(directory, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
	}

	// Copy the api file to the build directory, only if passed.
	if directory != "" {
		err = os.Link(file, directory+"/api.yaml")
		if err != nil {
			panic(err)
		}
	}

	// Create the HTML file in the doc directory
	var path string = "index.html"
	if directory != "" {
		path = directory + "/index.html"
	}
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template with the provided data and write to the file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}

	if directory == "" {
		fmt.Println("HTML documentation created in root directory")
	} else {
		fmt.Println("HTML documentation created in " + directory)
	}
}
