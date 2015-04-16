DONE - 5501 - The port accepts traffic but never sends back data

DONE - 5502 - The port sends back an empty string immediately upon connection

DONE - 5503 - The port sends back an empty string after the client sends data

DONE - 5504 - The port sends back a malformed response ("foo bar") immediately upon connection

DONE - 5505 - The port sends back a malformed response ("foo bar") after the client sends data

DONE - 5506 - The client accepts the request, and sends back one byte every 5 seconds

DONE - 5507 - The client accepts the request, and sends back one byte every 30 seconds

DONE - 5508 - Send a request to localhost:5508/sleep/<float> to sleep for float number of seconds. If no value is provided, sleep for 5 seconds.

DONE - 5008 - Send a request to localhost:5509/status/<int> to return a response with HTTP status code status. If no value is provided, return status code 200.

5510 - The server will send a response with a Content-Length: 3 header, however the response is actually 1 MB in size. This can break clients that reuse a socket.

5511 - Send a request to localhost:5511/size/<int> to return a Cookie header that is n bytes long. By default, return a 63KB header. 1KB larger will break many popular clients (curl, requests, for example)

5512 - Use this port to test retry logic in your client - to ensure that it retries on failure.

5513 - Send a request to localhost:5513/failrate/<float>. The server will drop requests with a frequency of failrate.

5514 - The server will try as hard as it can to return a content type that is not parseable by the Accept header provided by the request. Specify a Accept: application/json header in your request and the server will return data with the text/morse content type. 

5515 - The server will return a response with a content-type that matches the request, but it will be incomplete. The server will advertise an incorrect, too long Content-Length, and the response body will not be complete. The practical effect is that the server will hang halfway through the response download.

5516 - Same semantics as port 5515, but the server will close the connection partway through, instead of hanging indefinitely.
