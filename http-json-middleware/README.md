### Correct requests
```
curl -i -H "Content-Type: application/json" -X POST  http://localhost:8000/city -d '{"name":"New York", "area":304}'

HTTP/1.1 200 OK
Date: Mon, 20 Dec 2021 21:10:16 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8

201 - Created
```

```
curl -i -H "Content-Type: application/json" -X POST  http://localhost:8000/city -d '{"name":"Boston", "area":89}'

HTTP/1.1 200 OK
Date: Mon, 20 Dec 2021 21:07:41 GMT
Content-Length: 13
Content-Type: text/plain; charset=utf-8

201 - Created
```

### Bad requests
```
curl -i -X POST  http://localhost:8000/city -d '{"name":"New York", "area":304}'

HTTP/1.1 415 Unsupported Media Type
Date: Mon, 20 Dec 2021 11:12:05 GMT
Content-Length: 46
Content-Type: text/plain; charset=utf-8

415 - Unsupported Media Type. Please send JSON
```