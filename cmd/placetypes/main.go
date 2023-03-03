package main

import (
	"log"
	"syscall/js"

	sfom_placetypes "github.com/sfomuseum/go-sfomuseum-placetypes"
	"github.com/sfomuseum/go-sfomuseum-placetypes-wasm"
)

func main() {

	spec, err := sfom_placetypes.SFOMuseumPlacetypeSpecification()

	if err != nil {
		log.Fatalf("Failed to load SFO Museum placetypes specification, %w", err)
	}

	ancestors_func := wasm.AncestorsFunc(spec)
	defer ancestors_func.Release()

	descendants_func := wasm.DescendantsFunc(spec)
	defer descendants_func.Release()

	placetypes_func := wasm.PlacetypesFunc(spec)
	defer placetypes_func.Release()

	js.Global().Set("sfomuseum_placetypes_descendants", descendants_func)
	js.Global().Set("sfomuseum_placetypes_ancestors", ancestors_func)
	js.Global().Set("sfomuseum_placetypes", placetypes_func)

	c := make(chan struct{}, 0)

	log.Println("SFO Museum placetypes WASM binary initialized")
	<-c
}
