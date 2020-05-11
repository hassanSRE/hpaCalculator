# hpaCalculator

This repo is to calculate the average metric value divided by the number of replicas. It is taking from the Kubernetes  source code [here](https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/podautoscaler/horizontal.go#L485) to validate and to understand the resulting values.

## How to run it
You can run using the binary uploaded like so:

``` Go
go run ./main.go math.go amount.go 13 56
```

The result should be something like this

``` Go
Calculating value for replicas: 13 with AverageValue 55
Decimal Value:  4231
```

Or you can build and run if you have go installed

``` Go
go run ./main.go math.go amount.go 13 55
```
