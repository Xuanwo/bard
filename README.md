# bard

bard is a pastebin service built by Go.

## Installation

## Config

```yaml
public_url: http://127.0.0.1:8080
listen: 127.0.0.1:8080

key: xxxxxxx
max_file_size: 104857600

database:
  type: sqlite3
  connection: "/tmp/bard/db"

storage: "fs:///?work_dir=/tmp/bard/data"
```

## Usage

### Create

```shell
> dmesg | curl -F "file=@-" http://127.0.0.1:8080
{"url":"http://127.0.0.1:8080/VtfLHv"}
```

### Get

```shell
> curl http://127.0.0.1:8080/VtfLHv
[content of this paste]
```
