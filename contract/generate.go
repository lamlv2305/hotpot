package contract

//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=erc20.json --pkg contract --type ERC20 --out erc20.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=multicall.json --pkg contract --type MultiCall --out multicall.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=proxy_factory.json --pkg contract --type ProxyFactory --out proxy_factory.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=gnosis_safe.json --pkg contract --type GnosisSafe --out gnosis_safe.go
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --abi=multisend.json --pkg contract --type MultiSend --out multisend.go
