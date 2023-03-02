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
	
	planet_pt, err := spec.GetPlacetypeByName("planet")

	if err != nil {
		log.Fatalf("Failed to load SFO Museum placetypes specification, %w", err)
	}

	/*

	These don't work as expected in an SFO Museum context yet

	ancestors_func := wasm.AncestorsFunc(spec)
	defer ancestors_func.Release()
	
	descendants_func := wasm.DescendantsFunc(spec)
	defer descendants_func.Release()


	js.Global().Set("sfomuseum_placetypes_descendants", descendants_func)	
	js.Global().Set("sfomuseum_placetypes_ancestors", ancestors_func)
	*/
	
	placetypes_func := wasm.PlacetypesFunc(spec, planet_pt)	
	defer placetypes_func.Release()
	
	js.Global().Set("sfomuseum_placetypes", placetypes_func)

	c := make(chan struct{}, 0)

	log.Println("SFO Museum placetypes WASM binary initialized")
	<-c
}
