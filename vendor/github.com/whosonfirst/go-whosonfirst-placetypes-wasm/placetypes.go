package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst-placetypes"
)

func PlacetypesFunc(spec *placetypes.WOFPlacetypeSpecification) js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				pt, err := spec.Placetypes()

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
