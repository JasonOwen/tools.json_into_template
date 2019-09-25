# tools.json_into_template
Small tool to merge a json object into a template file.   

Uses the golang templating engine, for details about templating syntax please refer to https://golang.org/pkg/text/template/

# Build
go build -o json_into_template json_into_template.go

# Usage
**input.json**
```
{
  "Shop": "MiniMart",
  "Details": {
    "Address": "Corner Store",
    "Owners": "Mum and Pa"
  },
  "Fruits": [
    {
      "Fruit": "Banana",
      "Colour": "Yellow"
    },
    {
      "Fruit": "Apple",
      "Colour": "Red"
    },
    {
      "Fruit": "Orange",
      "Colour": "Orange"
    }
  ]
}
```

**template.txt**
```
Shop Title: {{ .Shop }}

============
Located at {{ .Details.Address }} and owned by {{ .Details.Owners }}

Sales Catalogue:
{{ range .Fruits }}
{{ .Fruit }}	-	{{ .Colour }}
{{ end }}
```
./json_to_template -i input.json -t template.txt -o output.conf

**output.conf** Output file created from above.
```
Shop Title: MiniMart

============
Located at Corner Store and owned by Mum and Pa

Sales Catalogue:

Banana	-	Yellow

Apple	-	Red

Orange	-	Orange
```

# Other Info
To maintain template syntax through a process you can escape the encapsulating the {{ in ". E.g.
```
RESULT="{{ ."{{ .value1 }}" }}"
```
This would make the output able to be a template file again, you could use one json file to create the template file for a second pass.
