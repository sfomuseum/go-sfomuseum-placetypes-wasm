package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	wof_placetypes "github.com/whosonfirst/go-whosonfirst-placetypes"
)

func PlacetypesFunc(spec *wof_placetypes.WOFPlacetypeSpecification, planet_pt *wof_placetypes.WOFPlacetype) js.Func {

	roles_custom := []string{
		wof_placetypes.CUSTOM_ROLE,
	}

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				pt := spec.DescendantsForRoles(planet_pt, roles_custom)

				enc_pt, err := json.Marshal(pt)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal placetypes, %v", err))
					return
				}

				resolve.Invoke(string(enc_pt))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
