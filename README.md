# multiset

![https://github.com/p1ck0/multiset/actions](https://github.com/p1ck0/multiset/actions/workflows/go.yml/badge.svg)

Tool for doing multiple request on multiple servers easier.

## Installation

### Go
```bash
$ go install github.com/p1ck0/multiset@latest
```

### Build
```bash
$ git clone https://github.com/p1ck0/multiset
$ cd multiset
$ make build
```

## Usage

```
Usage of multiset:
  -a    async mode
  -d    debug mode
  -f string
        file to read
```

## Example 

### Run
```bash
multiset -f file.json -a
```

### JSON file

```json
{
    "auto_complete": { // field that auto completes request
        "body": {
            "enable": true,
            "type": "auto"
        },
        "headers": {
            "Content-Type": "application/json"
        },
        "method": "POST"
    },
    "requests": [ // requests which will be sent
        {
            "url": "http://localhost:8080/api/test",
            "headers": {
                "Authorization": "Token 124"
            },
            "body": {
                "name": "test1"
            }
        },
        {
            "method": "PUT",
            "url": "http://localhost:8081/api/test",
            "headers": {
                "Authorization": "Token 123"
            },
            "body": {
                "name": "test2",
                "enable": false
            }
        }
    ]
}
```

### Example dir

See more in [example](https://github.com/p1ck0/multiset/tree/main/example)

## License

[MIT License](https://github.com/p1ck0/multiset/blob/main/LICENSE)
