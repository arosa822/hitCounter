#! /bin/bash
echo "Initiating call"
echo "Response:"
curl -X GET \
  http://localhost:8080/getMetrics \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: add8d8d9-425c-4964-bd78-045b3563c116,b11de7f6-d752-4fb8-bef3-0bc16eb7a6d4' \
  -H 'User-Agent: PostmanRuntime/7.20.1' \
  -H 'cache-control: no-cache'

