# streamsplit
Simple Go package to split data written to an io.Writer.

## MIME-Base64
Could for example be used to implement base64 Content-Transfer-Encoding when sending e-mail.
```go
mime64 := base64.NewEncoder(base64.StdEncoding, streamsplit.New(76, []byte("\n"), os.Stdout))
mime64.Write([]byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Quisque dui lorem, venenatis sed convallis in, rhoncus ut enim.`))
mime64.Close()

// Output:
// TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4g
// UXVpc3F1ZSBkdWkgbG9yZW0sIHZlbmVuYXRpcyBzZWQgY29udmFsbGlzIGluLCByaG9uY3VzIHV0
// IGVuaW0u
```
