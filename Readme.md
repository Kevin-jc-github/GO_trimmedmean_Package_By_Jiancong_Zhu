# GO_trimmedmean_Package_By_Jiancong_Zhu

## Overview
This repository contains a Go package that provides a function to calculate the trimmed mean of a given set of numeric values. The trimmed mean is a statistical measure that removes a specified proportion of the lowest and highest values from the data set and calculates the mean of the remaining values.

The package includes the following files:

- **`trimmedmean.go`**: Implements the `TrimmedMean` function, which calculates the trimmed mean of a slice of floating-point numbers.
- **`trimmedmean_test.go`**: Contains unit tests for the `TrimmedMean` function to verify its correctness.

## Features
- **Flexible trimming**: The function allows specifying different trimming proportions for the lower and upper ends of the data.
- **Error handling**: The function handles various edge cases, including invalid trimming proportions and empty input slices.

## File Descriptions

### `trimmedmean.go`
This file contains the implementation of the `TrimmedMean` function. It takes a slice of `float64` values and optional trimming proportions to calculate the trimmed mean.

#### Function Signature
```go
func TrimmedMean(data []float64, trimProportion ...float64) (float64, error)
```
- **`data`**: A slice of `float64` values for which the trimmed mean will be calculated.
- **`trimProportion`**: Optional arguments to specify trimming proportions for the lower and upper ends of the data. If only one value is provided, symmetric trimming is assumed.

### `trimmedmean_test.go`
This file contains unit tests for the `TrimmedMean` function. The tests verify that the function behaves correctly for various input data sets and trimming proportions.

#### Example Test Case
A simple test case is provided to calculate the trimmed mean for a data set and compare it against an expected value.

## Installation
To use the `trimmedmean` package, follow these steps:

### Clone the repository:
```bash
git clone https://github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu.git
```

### Navigate to the package directory:
```bash
cd GO_trimmedmean_Package_By_Jiancong_Zhu
```

### Initialize the module:
```bash
go mod init github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu
```

### Use `go get` to install the package in your project:
```bash
go get github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu
```

## Usage
Here's an example of how to use the `TrimmedMean` function in your Go program.

```go
package main

import (
    "fmt"
    "github.com/Kevin-jc-github/GO_trimmedmean_Package_By_Jiancong_Zhu"
)

func main() {
    data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Calculate the trimmed mean with a 0.1 trimming proportion
    mean, err := trimmedmean.TrimmedMean(data, 0.1)
    if err != nil {
        fmt.Printf("Error calculating trimmed mean: %v\n", err)
        return
    }

    fmt.Printf("Trimmed mean: %.2f\n", mean)
}
```

## Running Tests
To verify that the `TrimmedMean` function works as expected, run the included tests using the `go test` command:

```bash
go test
```

This command will execute the unit tests in `trimmedmean_test.go` and provide feedback on whether all tests pass successfully.

