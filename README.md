# streamsplit
Simple Go package to split data written to an io.Writer.

## MIME's Base64
Could for example be used with a base64 stream encoder to chunk the output to a limited line length.
```go
encoder := base64.NewEncoder(base64.StdEncoding, streamsplit.New(76, []byte("\r\n"), os.Stdout))
```
