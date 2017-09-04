# terminator
terminator is a simple application that lets you run multiple services in different terminals in `mac` as well asmonitor the logs of the terminals to watch out for `filtering keywords`. It writes all the logs that matches the filtering keywords in the console or any filtering file if mentioned.Moreover it also gives a http endpoint with live reload. Browse to the webpage at anytime to see any logs that match your filter criteria.

## Run

```shell
$ go get github.com/shreyaganguly/terminator
$ terminator -f commands.txt -words alertkeyword1 and alertkeyword2,alertkeyword3
```
This command will search for all logs matching `alertkeyword1` and `alertkeyword2` or `alertkeyword3`. Mention more `or` filter keywords separated by `,` and the `and` keywords separated by `and`.


### Start running your everyday services locally in your machine with ease!!!
