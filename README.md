Golang Pointer Converter

```bash
go get -u github.com/catchv/ptrconv
```

```Go
import "github.com/catchv/ptrconv"
```

```Go
s := ptrconv.AutoPtr("Hello world!!!")
fmt.Println(*s)

i := ptrconv.AutoPtr(1234)
fmt.Println(*i)

ui := ptrconv.AutoPtr(uint(1234))
fmt.Println(*ui)
```




