# web-crawler
A web crawler that crawls for all referrals in the same domain.

## Usage
Run the binary passing it the URL which needs to be crawled and the depth of iteration: 
1. `./main https://twitter.com 1`

Or run the main file passing the URL:

2. `go run main/web-crawler.go https://twitter.com 1`

## Current Implementation Limitations

1. Single threaded
2. Duplicate calls 
3. Untested