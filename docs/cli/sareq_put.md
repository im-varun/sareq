---
title: sareq put
---

Send HTTP PUT request to the specified URL

```bash
sareq put URL --body BODY [flags]
```

### Options

```bash
  -B, --body string             define the body to send with HTTP request (e.g. '{"key1": 1, "key2": "abc"}') (required)
  -H, --header stringToString   add a header to include with HTTP request (e.g. "key=value") (default [])
      --timeout int             specify timeout to use with HTTP request (default 10)
      --no-color                disable coloring in HTTP response
      --no-prettify             disable prettification in HTTP response
  -h, --help                    help for put
```

### Examples

```bash
# basic PUT request
sareq put https://api.example.com/users/user123 --body '{"name": "John Doe"}'

# PUT request with header
sareq put https://api.example.com/users/user123 --header "Authorization=abc123" --body '{"name": "John Doe"}'
```

### SEE ALSO

* [sareq](sareq.md)	 - SAReq is a CLI-based HTTP client for modern developers

