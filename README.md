# go-piratebay-link-fetcher

Piratby link featcher is a simple scrip to find piratby magnet links

# Usage 

Create en exacutable file 
```
go build 
```
Run the executable with parameters 
```
./pirateby param1 [param2 param3]
```
Parameters:
* param1 (obrigatory) - the searched srting ("e.g. "the big bang theory")
* param2 (optional) - url of prefered piratbyproxy (default "pirateproxy.tv")
* param3 (optional) - result limit (default 5) 

Example:
```
./pirateby "the big bang theory" "pirateproxy.pl" 3
```
