---
marp: true
---

# <!--fit-->PWD SECURITY

## by Codemunha

## 23 Feb, 2019  00:27AM

### github: [pwd-security](https://github.com/go-codemunha/pwd-security)

---
## Run go test

```go
go test -v ./...
```

## resule test

```go
?       github.com/chalermporn/pwd-security     [no test files]
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
ok      github.com/chalermporn/pwd-security/validator   0.687s

```