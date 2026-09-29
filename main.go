package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sb strings.Builder
	createHTMLBoilerplate(&sb, "<h1>Markdown to HTML</h1>")
	os.WriteFile("index.html", []byte(sb.String()), 0644)
}

func createHTMLBoilerplate(sb *strings.Builder, content string) *strings.Builder {
	docString := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    %s
</body>
</html>`, content)

	sb.WriteString(docString)

	return sb
}
