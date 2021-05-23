package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))

	if err != nil {
		log.Panic(err)
	}

	return []byte(decode)
}

//project/
//	css/
//		old_css/
//		new_css/
//	indexes/
//		template1/
//			oldtemplate1.html
//			newtemplate1.html
//		template2/
//			oldtemplate2.html
//			newtemplate2.html
//.......