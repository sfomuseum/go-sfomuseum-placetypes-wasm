package main

import (
	"log"
	"syscall/js"

	sfom_placetypes "github.com/sfomuseum/go-sfomuseum-placetypes"
	wof_wasm "github.com/whosonfirst/go-whosonfirst-placetypes-wasm"
)

func main() {

	spec, err := sfom_placetypes.SFOMuseumPlacetypeSpecification()

	if err != nil {
		log.Fatalf("Failed to load SFO Museum placetypes specification, %w", err)
	}

	descendants_func := wof_wasm.DescendantsFunc(spec)
	defer descendants_func.Release()

	ancestors_func := wof_wasm.AncestorsFunc(spec)
	defer ancestors_func.Release()

	placetypes_func := wof_wasm.PlacetypesFunc(spec)
	defer placetypes_func.Release()

	parents_func := wof_wasm.ParentsFunc(spec)
	defer parents_func.Release()

	children_func := wof_wasm.ChildrenFunc(spec)
	defer children_func.Release()

	isvalid_func := wof_wasm.IsValidFunc(spec)
	defer isvalid_func.Release()

	js.Global().Set("sfomuseum_placetypes_descendants", descendants_func)
	js.Global().Set("sfomuseum_placetypes_ancestors", ancestors_func)
	js.Global().Set("sfomuseum_placetypes_children", children_func)
	js.Global().Set("sfomuseum_placetypes_parents", parents_func)

	js.Global().Set("sfomuseum_placetypes_is_valid", isvalid_func)
	js.Global().Set("sfomuseum_placetypes", placetypes_func)

	c := make(chan struct{}, 0)

	log.Println("SFO Museum placetypes WASM binary initialized")
	<-c
}
