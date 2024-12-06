package main

import (
	"context"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/bonavadeur/katyusha/pkg/katyusha"
)

func main() {
	bonalib.Log("Konnichiwa, Katyusha-sama desu . Watashi wa Yukino Yukinoshita <3 <3 <3")
	ctx := context.Background()
	// do something here ...

	<-ctx.Done() // hangout forever
}
