package main

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	jsx := `
	import React from 'react';
	import ReactDOM from 'react-dom';
	import { Button } from '@bytedesign/web-react';
	import { Select, Switch } from '@bytedesign/web-react';
	import '@bytedesign/web-react/es/Button/style/index.css';

	class Page extends React.Component {
		render() {
			console.log(Select, Switch);
			return (
				<div>
					<span>adkfjakldfja</span>
					12312312
					<img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1598701029074&di=e1f5dbc7ca7647bf25cc8e1183f21600&imgtype=0&src=http%3A%2F%2Fpic4.zhimg.com%2Fv2-4bbc9550e65e890d333baad8fa971b0e_1200x500.jpg" />
					<Button>324234234</Button>
				</div>
			);
		}
	}
	
	ReactDOM.render(<Page />, document.getElementById('root'));
		`

	result := api.Transform(jsx, api.TransformOptions{
		Loader: api.LoaderTSX,
	})

	fmt.Printf("%d errors and %d warnings\n",
		len(result.Errors), len(result.Warnings))

	os.Stdout.Write(result.JS)
}
