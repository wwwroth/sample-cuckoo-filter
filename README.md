# sample-cuckoo-filter

```shell
# Generate 5,000 sample keys to load into filter
export TMP_KEY_FILE="temp/.keys"
go run . -action=generateKeys -numKeys=50000
```

```shell
# Build cuckoo filter and start web server
export HTTP_PORT=8000
export TMP_KEY_FILE="temp/.keys"
go run . -action=buildFilter
```

### Check if key exists in cuckoo filter
```shell
export HTTP_PORT=8000
curl --location "http://127.0.0.1:${HTTP_PORT}/key/53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239"

{
  "success": true,
  "data": {
    "key": "53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239",
    "message": "Key found in filter."
  }
}

```

### Delete key from cuckoo filter
```shell
export HTTP_PORT=8000
curl --location "http://127.0.0.1:${HTTP_PORT}/key/53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239/remove"

{
  "success": true,
  "data": {
    "key": "53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239",
    "message": "Key deleted from filter."
  }
}
```

### Insert key into cuckoo filter
```shell
export HTTP_PORT=8000
curl --location "http://127.0.0.1:${HTTP_PORT}/key/53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239/insert"

{
  "success": true,
  "data": {
    "key": "53df161fa3a852c4a5300b5960982f3c2ab4bc18d40cf31e941bf994e1917239",
    "message": "Key inserted into filter."
  }
}
```