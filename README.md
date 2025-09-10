# parachute 🎯

[![Go Reference](https://pkg.go.dev/badge/github.com/daalfox/parachute.svg)](https://pkg.go.dev/github.com/daalfox/parachute)
[![Go Report Card](https://goreportcard.com/badge/github.com/daalfox/parachute)](https://goreportcard.com/report/github.com/daalfox/parachute)

A **type-safe wrapper** around Go’s [`singleflight`](https://pkg.go.dev/golang.org/x/sync/singleflight), powered by generics.  
No more `interface{}`, no more unsafe type assertions 🚀

---

## ✨ Why Parachute?

`golang.org/x/sync/singleflight` is awesome for deduplicating concurrent calls,  
but it returns `interface{}`, so you end up with unsafe type assertions:

```go
val, err, _ := g.Do("key", func() (interface{}, error) {
    return fetchData(), nil
})
data := val.(MyType) // panic if wrong type
```
With parachute, you get compile-time safety:
```go
p := parachute.Group[string]{} // parachute for string results

// string
val, err, shared := p.Do("key", func() (string, error) {
    return fetchData(), nil
})

fmt.Println(val)   // already string ✅
```
## 📦 Installation
```
go get github.com/daalfox/parachute
```
