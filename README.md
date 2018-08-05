# web-crawler
A web crawler that crawls for all referrals in the same domain.

## Usage
Run the binary passing it the URL which needs to be crawled: 
1. ./main https://twitter.com

Or run the main file passing the URL:

2. go run web-crawler.go https://twitter.com

## Current Implementation Limitations

1. Single threaded
2. Duplicate calls 
3. Untested