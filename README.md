# commoninitialisms

For always getting the latest　`lint.commonInitialisms`. (Use go/ast)


## Why need it

I want to get commonInitialisms on lint package.  

https://github.com/golang/lint/blob/8f45f776aaf18cebc8d65861cc70c33c60471952/lint.go#L771-L810

But this package is unexport field. So Many gophers write duplicate code.

https://github.com/search?l=Go&p=1&q=commonInitialisms+API+ID&type=Code


## Example 

```go
package main

import (
    "fmt"
    "github.com/yudppp/commoninitialisms"
)

func main() {
    commonInitialisms, _ := commoninitialisms.GetCommonInitialisms()
    fmt.Println(commonInitialisms)
    // or 
    fmt.Println(commoninitialisms.Must(commoninitialisms.GetCommonInitialisms()))
}
```


## Related post

- https://blog.yudppp.com/posts/commoninitialisms/