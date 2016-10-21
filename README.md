## Overview

httpdig - allows to make DNS queries using Google's HTTPS DNS service.

## Install

``` bash
go get -u github.com/firstrow/httpdig
```

## Usage

``` go
import (
       "fmt"
       "github.com/firstrow/httpdig"
)

resp, _ := httpdig.Query("google.com", "NS")
fmt.Print(resp.Answer)
```

## Links
[RR Types](https://en.wikipedia.org/wiki/List_of_DNS_record_types)
