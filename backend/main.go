package main

import (
	"C"
	// "strconv"
	"unsafe"
	// "github.com/jbarham/primegen.go"
	"math/big"
	// "fmt"
)

func Function(vals []uint64) unsafe.Pointer {
	result := make([]uint64, 100)
	// var result []uint64
	var numbers []*big.Int
	for _, e := range vals {
		tmp := new(big.Int)
		tmp.SetUint64(e)
		numbers = append(numbers, tmp)
	}
	rs := lmf(numbers)
	i := 0
    for _, r := range rs {
    	// result = append(result, r.number.Uint64())
        result[i] = r.number.Uint64()
        i++
        for _, val := range r.decomp {
        	// result = append(result, val.Uint64())
        	result[i] = val.Uint64()
        	i++
        }
    }
    // fmt.Println(result)
	return unsafe.Pointer(&result)
}

//export ExportedFunction
func ExportedFunction(vals []uint64) uintptr {
	p := Function(vals)
	s := *(*([]uint64))(p)
	return uintptr(unsafe.Pointer(&s[0]))
}


// adapted from https://rosettacode.org/wiki/Parallel_calculations#Go
type result struct {
    number *big.Int
    decomp []*big.Int
}
 
func lmf(numbers []*big.Int) []result {
    rCh := make(chan result)
    for _, n := range numbers {
        go decomp(n, rCh)
    }

    rs := []result{<-rCh}
    for i := 1; i < len(numbers); i++ {
    	r := <-rCh
    	rs = append(rs, r)
    }
    return rs
}
 
func decomp(n *big.Int, rCh chan result) {
    rCh <- result{n, Primes(new(big.Int).Set(n))}
}
 
var (
    ZERO = big.NewInt(0)
    ONE  = big.NewInt(1)
)
 
func Primes(n *big.Int) []*big.Int {
    res := []*big.Int{}
    mod, div := new(big.Int), new(big.Int)
    for i := big.NewInt(2); i.Cmp(n) != 1; {
        div.DivMod(n, i, mod)
        for mod.Cmp(ZERO) == 0 {
            res = append(res, new(big.Int).Set(i))
            n.Set(div)
            div.DivMod(n, i, mod)
        }
        i.Add(i, ONE)
    }
    return res
}

func main(){}
