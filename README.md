---
marp: true
---

# <!--fit-->PWD SECURITY

## by Codemunha

## 23 Feb, 2019  00:27AM

### github: [pwd-security](https://github.com/go-codemunha/pwd-security)

![Go](https://github.com/go-codemunha/pwd-security/workflows/Go/badge.svg)

---

## Source Code : validator.go

```go

package validator

import (
    "unicode"
)

func Password(s string) bool {
    if s == "" || len(s) == 0 {
        return false
    }
    if len(s) < 8 {
        return false
    }
    var hasUppercase = false
    var hasLowercase = false
    var isSpace = true
    var hasSymbol = false
    var hasNumber = false
    for _, r := range s {
        switch {
        case unicode.IsUpper(r):
            hasUppercase = true
        case unicode.IsLower(r):
            hasLowercase = true
        case unicode.IsNumber(r):
            hasNumber = true
        case unicode.IsSpace(r):
            isSpace = false
        case !unicode.IsSymbol(r):
            hasSymbol = true
        default:
        }

    }
    return hasLowercase && hasUppercase && hasNumber && isSpace && hasSymbol
}


```

---

## Run go test

```go
go test -v ./...
```

## result test

```go
?       github.com/go-codemunha/pwd-security     [no test files]
=== RUN   TestPasswordOriginal
=== RUN   TestPasswordOriginal/FooBar123!
=== RUN   TestPasswordOriginal/foobar
=== RUN   TestPasswordOriginal/foobar123
=== RUN   TestPasswordOriginal/@
--- PASS: TestPasswordOriginal (0.00s)
    --- PASS: TestPasswordOriginal/FooBar123! (0.00s)
    --- PASS: TestPasswordOriginal/foobar (0.00s)
    --- PASS: TestPasswordOriginal/foobar123 (0.00s)
    --- PASS: TestPasswordOriginal/@ (0.00s)
PASS
ok      github.com/go-codemunha/pwd-security/validator   0.687s

```
