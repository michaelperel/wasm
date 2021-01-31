# Demo
This app shows you how to use client side Vue with Go WASM

# Run
In one terminal tab: 
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
GOARCH=wasm GOOS=js go build -o lib.wasm
```

In another terminal tab:
```
go run server/server.go
```

Navigate to localhost:8080