---
title: sareq delete
---

Send HTTP DELETE request to the specified URL

```bash
sareq delete URL [flags]
```

### Options

```bash
  -B, --body string             define the body to send with HTTP request (e.g. '{"key1": 1, "key2": "abc"}')
  -H, --header stringToString   add a header to include with HTTP request (e.g. "key=value") (default [])
      --timeout int             specify timeout to use with HTTP request (default 10)
      --no-color                disable coloring in HTTP response
      --no-prettify             disable prettification in HTTP response
  -h, --help                    help for delete
```

### Examples

```bash
# basic DELETE request
sareq delete https://api.example.com/users/user123

# DELETE request with header
sareq delete https://api.example.com/users/user123 --header "Authorization=abc123"
```

### SEE ALSO

* [sareq](sareq.md)	 - A CLI-based HTTP client for modern developers

