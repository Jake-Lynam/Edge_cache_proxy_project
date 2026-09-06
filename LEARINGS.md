25/08/2026
Today I setup the origin server to test my proxy with.

What is a proxy?
essentially traffic control, aproxy's job is to direct requests and responses to and from either the client or server. It sits between the client and server as a middle man. 

why is a proxy needed?
Many reasons; content filtering, privacy and saving bandwith

26/08/2026
Today I created the first stage of my proxy, essenstially just enabling it to send a request
and recieve a response from the server. I used functions such as "Get", "Close" and "ReadAll" 

I've learned the http concepts of go such as goroutine which allows for quicker communication
between a proxy and client/server by snycing multiple functions together within the proxy to run
at the same time using avaliable CPU cores. I will potentially learn more about this concept later
down the line.

06/09/2026
Today I have implemented my finished cache into this project. For now the system is ready to save 
different requests and responses but first I need to implement these different items into client
and server. 

What's left? I need to implement the different request and response types, include mutexing 
for concurrency, create ttl within the cache, write tests and gather metrics for a benchmark