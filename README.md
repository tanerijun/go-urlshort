# URLShort

An URL shortener written with Go.

## Introduction

URLShort maps short paths to a real URL based on JSON/YAML file.

```json
[
  {
    "path": "/goo",
    "url": "https://google.com"
  },
  {
    "path": "/fb",
    "url": "https://facebook.com"
  }
]
```

```yaml
- path: /goo
  url: https://google.com
- path: /fb
  url: https://facebook.com
```

Example url:

```
http://<link-to-server>/fb -> https://facebook.com
```

## Usage

1. Clone repo
2. Build the app or run the app directly providing the path to the JSON/YAML map as an argument.
   ```
   urlshort sample/sample.json
   go run main.go sample/sample.yaml
   ```
3. Server is live on port 8080.
