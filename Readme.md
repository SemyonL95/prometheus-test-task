## TASK
https://dev.azure.com/DataHow/DataHow%20Challenges/_wiki/wikis/Challenges.wiki/1/Backend-challenge

## Deploy
1. Pull project
2. Run command "make app-server"
## Testing
1. Run command "make run-tests"
2. After tests, run command "app-down" for stop docker

## Additional description
I made in memory storage (set type) to collect ip adresses, i decide to not use external dependecies, because it's little overhead, we can count memory usage per record 4byte + ~3byte(overhead for map is not constant but i found that it take ~3bytes over) = 7; 1073741824(1GB of RAM) / 7 = 153391689.143