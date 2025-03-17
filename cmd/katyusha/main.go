package main

import (
	"context"

	"github.com/bonavadeur/katyusha/pkg/bonalib"
	_ "github.com/yukino1710193/katyusha/pkg/katyusha"
)

func main() {
	bonalib.Log("Katyusha ON")
	ctx := context.Background()
	// do something here ...

	<-ctx.Done() // hangout forever
}
