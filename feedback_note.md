# Feedback

## Language
 Score : 4/5 ==> Exceeds Expectations

1. I'm not sure what is meant here https://github.com/gideon-maina/rss-processor/blob/master/fetchrss/fetch.go#L95
but this is necessary to ensure we capture and are using that current loop value in the go routines, else we would be 
using stale values since in the for loop, we cannot gurantee when each go routine would be executed.

2. Commenting of functions doesn't follow standard Go comments e.g This should start with the function name
https://github.com/gideon-maina/rss-processor/blob/master/fetchrss/fetch.go#L64 . Also linters could be setup for such
cases.

3. Uses a basic http handler/router which is just perfect for this project but I believe he's aware of other handler/router 
setups where multiple routes with different HTTP methods are expected so we don't end up check things like HTTP 
Method over and over again. https://github.com/gideon-maina/rss-processor/blob/master/serverss/server.go#L38

4. I do not see any reason why we need to handle http Requests in new Go routines. Each request to the server is 
already handled in a new Go routine in Go's http lib. 
https://github.com/gideon-maina/rss-processor/blob/master/serverss/server.go#L44

5.  Uses Go Modules

## Architecture and Design 
Score : 3/5 ==> Meets Expectations

1. Follows more of a CLI approach to building a Web serivce.

## Testing & Security
Score: 4/5 ==> Exceeds Expectations

 1. Demonstrates use of Tokens to ensure endpoints are protected.
 2. We could implement RateLimiting to prevent potential DDOS attacks. (I know this is an overkill, but would be nice if it was written as TODO in Readme ;))
 
 # Summary
 ## Overall Average Rating = 3.7/5 ==> Meets Expectations
