# Solution
This solution was written in Go so to install it do the usual:

```bash
$ go get -u git.enova.com/fun/the-coding-zone
$ cd ~/go/src/git.enova.com/fun/the-coding-zone/problem-2017-07/solutions/zamir_syed
$ go install ./...
```
Then you can run the app over test files:

```bash
$ mario ../../sample.txt
$ mario ../../medium.txt
$ mario ../../large.txt
```

Here are the console outputs for each of the three sample files:

![example](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-07/solutions/zamir_syed/sample.png)
![example](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-07/solutions/zamir_syed/medium.png)
![example](https://git.enova.com/raw/fun/the-coding-zone/master/problem-2017-07/solutions/zamir_syed/large.png)

# Generating a Random Board!
You can also generate a random board and solve it:

```bash
$ mario 40 20 0.25 0 board.txt
```

You can run the command with no arguments to see its usage.
