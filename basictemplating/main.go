package main

import(
    "bytes"
    "fmt"
    "os"
    "text/template"
    "log"
)

type Warrior struct {
  Name string
  Power int
}

func main() {
    // The template we are filling
    const myTemplate = "{{.Name}} has a power level of: {{.Power}}"

    // Creating the new templating
    tmpl, err := template.New("WarriorTemplate").Parse(myTemplate)
    if err != nil {
        log.Fatal(err)
    }

    // Execute the template and direct it to stdout
    randy := &Warrior{Name: "Randolf the Elder", Power: 250}
    err = tmpl.Execute(os.Stdout, randy)
    if err != nil {
        log.Fatal(err) 
    }

    // Similar, but execute the template and direct it to our Buffer
    var placeholder bytes.Buffer
    barclay := &Warrior{Name: "Barclay the Bold", Power: 300}
    err = tmpl.Execute(&placeholder, barclay)
    if err != nil {
        log.Fatal(err) 
    }

    // Print the Buffer
    fmt.Printf("\n%s\n", placeholder.String())
}


