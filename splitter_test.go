package streamsplit_test

import (
	"encoding/base64"
	"os"

	"github.com/ferdypruis/streamsplit"
)

func ExampleNew_base64ContentTransferEncoding() {
	w := streamsplit.New(76, []byte("\n"), os.Stdout)
	mime := base64.NewEncoder(base64.StdEncoding, w)
	mime.Write([]byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque dui lorem, venenatis sed convallis in, rhoncus ut enim.`))
	mime.Close()

	// Output:
	// TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4g
	// UXVpc3F1ZSBkdWkgbG9yZW0sIHZlbmVuYXRpcyBzZWQgY29udmFsbGlzIGluLCByaG9uY3VzIHV0
	// IGVuaW0u
}

func ExampleNew_spell() {
	w := streamsplit.New(1, []byte("-"), os.Stdout)
	w.Write([]byte(`abcdefg`))

	// Output:
	// a-b-c-d-e-f-g
}
