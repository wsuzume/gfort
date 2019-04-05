package gfort

import (
	//"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/pkg/errors"
)

var suppress = false

func SuppressWarning() {
	suppress = true
}

func ActivateWarning() {
	suppress = false
}

func Ignore(err error) {
	if !suppress {
		var e error
		if err == nil {
			e = errors.New("<gfort dummy error>")
		} else {
			e = err
		}
		fmt.Fprintf(os.Stderr, "\nWarning: Ignoring error %+v\n", e)
	}
}

func IgnoreF(data interface{}, err error) {
	Ignore(err)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}

	if !suppress {
		var e error
		if err == nil {
			e = errors.New("<gfort dummy error>")
		} else {
			e = err
		}
		fmt.Fprintf(os.Stderr, "\nWarning: Code contains panic() %+v\n", e)
	}
}

func PanicF(data interface{}, err error) {
	Panic(err)
}

func M(data interface{}, err error) interface{} {
	return Must(data, err)
}

func Must(data interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}
	return data
}

func Filter(f interface{}, g interface{}, data ...interface{}) interface{} {
	fv, gv := reflect.ValueOf(f), reflect.ValueOf(g)

	var xs []reflect.Value

	for _, x := range data {
		xs = append(xs, reflect.ValueOf(x))
	}

	ys := gv.Call(xs)
	fv.Call(ys)
	return ys[0].Interface()
}
