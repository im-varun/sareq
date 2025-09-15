## sareq post

Send HTTP POST request to the specified URL

```bash
sareq post URL --body BODY [flags]
```

### Examples

```bash
# basic POST request
sareq post https://api.example.com/users --body '{"name": "John Doe"}'

# POST request with timeout of 5 seconds
sareq post https://api.example.com/users --body '{"name": "John Doe"}' --timeout 5

# POST request with headers
sareq post https://api.example.com/users --header "Authorization=abc123" --header "Content-Type=application/json" --body '{"name": "John Doe"}'
```

### Options

```bash
  -B, --body string             set the body to send with HTTP request (e.g '{"key1": int, "key2": "string"}') (required)
  -H, --header stringToString   set a header to send with HTTP request (e.g "key=value") (default [])
      --timeout int             set timeout for HTTP request (default 10)
      --no-color                disable coloring for HTTP response
      --no-prettify             disable prettification for HTTP response
  -h, --help                    help for post
```

### SEE ALSO

* [sareq](sareq.md)	 - SAReq is a CLI-based HTTP client for modern developers

