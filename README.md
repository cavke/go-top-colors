# go-top-colors

**go-top-colors** is a program that takes a list of image urls from the input file 
and for every image it finds top 3 most prevalent colors 
in the RGB scheme in a hexadecimal format (#000000 - #FFFFFF).
Application writes results in a CSV file in form of: `url,color,color,color`

Gouroutines are used to parallelize image processing. Results are stored in the LRU cache.

#### Running the application

You should have Go 1.15 or later to build this application.

To start the application, execute following commands in the root directory of the project:

```
$ go build
$ ./go-top-colors -input=input.txt -output=result.csv
```

#### Runtime parameters

Application takes following command parameters (**bold** are the required ones):
- **-input** - file path to the input file that contains URLs in separate rows.
- **-output** - file path to the output CSV file where the results will be stored.
- -max-parallels - limits the maximum number of goroutines for the image processing. Default value is: *runtime.NumCPU()*.
- -cache-size - maximum number of elements in the LRU cache. Default value is 50.

With optional parameters you can tune the application, so that it scales better with the running environment:

Example: 
```
$ ./go-top-colors -input=input.txt -output=result.csv -max-parallels=4 -cache-size=100
```