## Step 1
- Tạo contract bằng solidity.
- Lưu vào thư mục contract.

## Step 2
- Clone go-ethereum source code vào ```$GOPATH/src```.
- Build abigen:
  - Truy cập Remix IDE.
  - Compile và copy ABI.
- Copy file abigen đã build vào ```$GOPATH/bin```

## Step 3
- Build go file:
  ```
  $GOPATH/bin/abigen --abi ./contract/deposit.abi --pkg deposit --type Deposit --out ./contract/deposit.go
  ```
## Step 4
- Chỉnh sửa ```read_event.go``` cho phù hợp với contract.

## Tham khảo
- https://geth.ethereum.org/docs/dapp/native-bindings#what-is-an-abi
- https://goethereumbook.org/event-read/