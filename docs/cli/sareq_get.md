## sareq get

Send HTTP GET request to the specified URL

```bash
sareq get URL [flags]
```

### Examples

```bash
# basic GET request
sareq get https://api.example.com/users/user123

# GET request with timeout of 5 seconds
sareq get https://api.example.com/users/user123 --timeout 5

# GET request with header
sareq get https://api.example.com/users/user123 --header "Authorization=abc123"
```

### Options

```bash
  -B, --body string             set the body to send with HTTP request (e.g '{"key1": int, "key2": "string"}')
  -H, --header stringToString   set a header to send with HTTP request (e.g "key=value") (default [])
      --timeout int             set timeout for HTTP request (default 10)
      --no-color                disable coloring for HTTP response
      --no-prettify             disable prettification for HTTP response
  -h, --help                    help for get
```

### SEE ALSO

* [sareq](sareq.md)	 - SAReq is a CLI-based HTTP client for modern developers

