## sareq patch

Send HTTP PATCH request to the specified URL

```bash
sareq patch URL --body BODY [flags]
```

### Examples

```bash
# basic PATCH request
sareq patch https://api.example.com/users/user123 --body '{"email": "john@example.com"}'

# PATCH request with header
sareq patch https://api.example.com/users/user123 --header "Authorization=abc123" --body '{"email": "john@example.com"}'
```

### Options

```bash
  -B, --body string             define the (JSON) body to send with HTTP request (e.g. '{"key1": 1, "key2": "abc"}') (required)
  -H, --header stringToString   add a header to include with HTTP request (e.g. "key=value") (default [])
      --timeout int             specify timeout to use with HTTP request (default 10)
      --no-color                disable coloring in HTTP response
      --no-prettify             disable prettification in HTTP response
  -h, --help                    help for patch
```

### SEE ALSO

* [sareq](sareq.md)	 - SAReq is a CLI-based HTTP client for modern developers

