## ROY1SME
<p>
   <a href="https://github.com/rroy233/roy1sme">
      <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/rroy233/roy1sme?style=flat-square">
   </a>
   <a href="https://github.com/rroy233/roy1sme/releases">
      <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/rroy233/roy1sme?style=flat-square">
   </a>
   <a href="https://github.com/rroy233/roy1sme/blob/main/LICENSE">
      <img alt="GitHub license" src="https://img.shields.io/github/license/rroy233/roy1sme?style=flat-square">
   </a>
   <a href="https://github.com/rroy233/roy1sme/commits/main">
      <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/rroy233/roy1sme?style=flat-square">
   </a>
</p>

> An API client for a URL shortener service

[中文](README.md) | EN

### Getting the API Key

> The default language of the website is Simplified Chinese. You can use your browser's translator to translate it to English.

Go to http://roy1s.me, log in, and navigate to the "API Access" page to obtain your API key.

![img_1.png](docs%2Fimg_1.png)

### Installation

```shell
go get -u github.com/rroy233/roy1sme
```

### Usage

```go
package main

import (
	"log"
	"github.com/rroy233/roy1sme"
)

func main(){
	// Initialize a client instance with your API key
	client := roy1sme.NewClient("YOUR_API_KEY")

	// Create a short URL with a 1-day expiration
	myUrl, err := client.CreateUrl("https://github.com/rroy233/roy1sme", roy1sme.ExpireOneDay)
	if err != nil {
		panic(err)
	}
	log.Printf("My new URL: %v", myUrl)

	// Retrieve the creation history
	myHistory, err := client.GetHistory()
	if err != nil {
		panic(err)
	}
	log.Printf("My history: %v", myHistory)
}

```

### License

GPL-3.0 license