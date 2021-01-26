package main

import (
	"log"

	"github.com/go-openapi/analysis"
	"github.com/go-openapi/loads"
)

func main() {
	doc, err := loads.Spec("./openapi.yaml")
	if err != nil {
		log.Fatalf("Load Failed: %v", err)
	}
	an := analysis.New(doc.Spec()) // Analyze spec
	err = analysis.Flatten(analysis.FlattenOpts{
		Spec:   an,
		Expand: true,
	})
	if err != nil {
		log.Fatalf("Flatten Failed: %v", err)
	}
}
