# extchanger

## How to build

```shell
$ go build -o extchanger
```

## Usage

```shell
$ ./extchanger
Usage:
        All of the args are required.

  -d string
        target directory
  -from string
        from extension
  -r    change extension recursively
  -to string
        new extension

$ ls ./sample/sample2/sample3
dojo4.jpeg      dojo4.jpg

./exchanger -d=./sample/sample2/sample3 -from=jpeg -to=png
Converted:       sample/sample2/sample3/dojo4.jpeg      ->       sample/sample2/sample3/dojo4.png
Successfully converted!

$ ls ./sample/sample2/sample3
dojo4.jpeg      dojo4.jpg       dojo4.png

# Recursive
./exchanger -d=./sample/sample2 -from=png -to=jpeg -r=true
Converted:       sample/sample2/dojo3.png       ->       sample/sample2/dojo3.jpeg
Successfully converted!
Converted:       sample/sample2/sample3/dojo4.png       ->       sample/sample2/sample3/dojo4.jpeg
Successfully converted!
```

