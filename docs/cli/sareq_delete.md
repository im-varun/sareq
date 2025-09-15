## sareq delete

Send HTTP DELETE request to the specified URL

```bash
sareq delete URL [flags]
```

### Examples

```bash
# basic DELETE request
sareq delete https://api.example.com/users/user123

# DELETE request with header
sareq delete https://api.example.com/users/user123 --header "Authorization=abc123"
```

### Options

```bash
  -B, --body string             set the body to send with HTTP request (e.g '{"key1": int, "key2": "string"}')
  -H, --header stringToString   set a header to send with HTTP request (e.g "key=value") (default [])
      --timeout int             set timeout for HTTP request (default 10)
      --no-color                disable coloring for HTTP response
      --no-prettify             disable prettification for HTTP response
  -h, --help                    help for delete
```

### SEE ALSO

* [sareq](sareq.md)	 - SAReq is a CLI-based HTTP client for modern developers

