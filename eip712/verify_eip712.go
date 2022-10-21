package eip712

import (
	"encoding/json"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/storyicon/sigverify"
)

const ExampleTypedData = `
{
    "types": {
        "EIP712Domain": [
            {
                "name": "name",
                "type": "string"
            },
            {
                "name": "chainId",
                "type": "uint256"
            }
        ],
        "RandomAmbireTypeStruct": [
            {
                "name": "identity",
                "type": "address"
            },
            {
                "name": "rewards",
                "type": "uint256"
            }
        ]
    },
    "domain": {
        "name": "Ambire Typed test message",
        "chainId": "1"
    },
    "primaryType": "RandomAmbireTypeStruct",
    "message": {
        "identity": "0x0000000000000000000000000000000000000000",
        "rewards": 0
    }
}
`

const BeeExampleTypedData = `
{
    "types": {
        "EIP712Domain": [
            {
                "name": "name",
                "type": "string"
            },
            {
                "name": "version",
                "type": "string"
            }
        ],
        "MyType": [
			{
				"name": "test",
				"type": "string"
			}
		]
    },
    "domain": {
        "name": "test",
        "version": "1.0"
    },
    "primaryType": "MyType",
    "message": {
        "test": "abc"
    }
}
`

func CheckVerification() {
	var typedData apitypes.TypedData
	if err := json.Unmarshal([]byte(ExampleTypedData), &typedData); err != nil {
		panic(err)
	}
	valid, err := sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress("0xaC39b311DCEb2A4b2f5d8461c1cdaF756F4F7Ae9"),
		typedData,
		"0xee0d9f9e63fa7183bea2ca2e614cf539464a4c120c8dfc1d5ccc367f242a2c5939d7f59ec2ab413b8a9047de5de2f1e5e97da4eba2ef0d6a89136464f992dae11c",
	)
	fmt.Println(valid, err) // true <nil>

	var beeTypedData apitypes.TypedData
	if err := json.Unmarshal([]byte(BeeExampleTypedData), &beeTypedData); err != nil {
		panic(err)
	}
	valid, err = sigverify.VerifyTypedDataHexSignatureEx(
		ethcommon.HexToAddress("0x8D3766440f0d7B949a5E32995D09619A7F86e632"),
		beeTypedData,
		"0x60f054c45d37a0359d4935da0454bc19f02a8c01ceee8a112cfe48c8e2357b842e897f76389fb96947c6d2c80cbfe081052204e7b0c3cc1194a973a09b1614f71c",
	)
	fmt.Println(valid, err) // true <nil>
}
