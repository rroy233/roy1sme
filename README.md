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

> 一个短链接生成网站的api客户端

中文 | [EN](README_en.md)

### 获取Api-Key

前往 http://roy1s.me 登录，并前往“api访问”页面，获取到API KEY。
![img.png](docs%2Fimg.png)


### 安装

```shell
go get -u github.com/rroy233/roy1sme
```

### 使用

```go
package main

import (
	"log"
	"github.com/rroy233/roy1sme"
)

func main(){
	//初始化一个客户端实例，填入自己的api key
	client := roy1sme.NewClient("YOUR_API_KEY")

	//创建一个短链接，1day有效期
	myUrl, err := client.CreateUrl("https://github.com/rroy233/roy1sme", roy1sme.ExpireOneDay)
	if err != nil {
		panic(err)
	}
	log.Printf("my new url:%v", myUrl)

	//获取自己的创建记录
	myHistory, err := client.GetHistory()
	if err != nil {
		panic(err)
	}
	log.Printf("my history:%v", myHistory)
}
```

### 协议

GPL-3.0 license