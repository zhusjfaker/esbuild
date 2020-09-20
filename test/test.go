package main

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	jsx := `
	import { Button } from '@bytedesign/web-react';
	import { Select, Switch } from '@bytedesign/web-react';
		`

	result := api.Transform(jsx, api.TransformOptions{
		Loader: api.LoaderTSX,
	})

	fmt.Printf("%d errors and %d warnings\n",
		len(result.Errors), len(result.Warnings))

	os.Stdout.Write(result.JS)
}
