package main

import (
	"context"

	isuports "github.com/isucon/isucon12-qualify/webapp/go"
)

func main() {
	isuports.InstallTracerProvider(context.Background())
	isuports.Run()
}
