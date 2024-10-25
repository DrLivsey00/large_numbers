package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GetKeysNumber(length int) (*big.Int, error) {
	if length < 8 || length%8 != 0 || length > 4096 {
		return big.NewInt(0), fmt.Errorf("invalid length %d: must be >= 8, <= 4096, and a multiple of 8", length)
	}
	max := new(big.Int).Lsh(big.NewInt(1), uint(length))
	return max, nil
}

func GenKey(length int) (*big.Int, error) {
	if length < 8 || length%8 != 0 || length > 4096 {
		return nil, fmt.Errorf("invalid key length")
	}
	max := new(big.Int).Lsh(big.NewInt(1), uint(length))
	max.Sub(max, big.NewInt(1))

	key, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func Brutforce(targetKey *big.Int, length int) (int64, error) {
	start := time.Now()

	max := new(big.Int).Lsh(big.NewInt(1), uint(length))

	for i := big.NewInt(0); i.Cmp(max) < 0; i.Set(i) {

		if i.Cmp(targetKey) == 0 {
			break
		}
		i.Add(i, big.NewInt(1))
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}
func main() {
	fmt.Println("Starting...")
	bits := []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	for _, bitLength := range bits {
		num, err := GetKeysNumber(bitLength)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Number of keys with length %d is %d\n", bitLength, num) //get keys number

		key, err := GenKey(bitLength)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Generated key with length %d: %d\n", bitLength, key)
		time, err := Brutforce(key, bitLength)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Time to brutforce this key: %dms\n", time)

	}
}
