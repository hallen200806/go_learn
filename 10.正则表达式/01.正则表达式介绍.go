package main

import (
	"regexp"
	"fmt"
)

func main() {
	/*
	步骤：
		1.导入包：regexp
			import "regexp"
		2.解析规则
			reg1 := regexp.MustCompile(`a.c`)		这里必须使用tab键上面的那个引号，意思是原生的
		3.根据规则提取关键信息
			ret1 := reg1.FindAllStringSubmatch(context,-1)   // -1是指提取所有匹配到的
	 */
	context1 := "12.fsd.53.asc.666.nnn.c"

	// 解析规则
	reg1 := regexp.MustCompile(`a.c`)   // 这里必须使用tab键上面的那个引号，意思是原生的
	if reg1 == nil{
		fmt.Println("regexp error")
		return
	}

	// 根据规则提取关键信息
	ret1 := reg1.FindAllStringSubmatch(context1,-1)   // -1是指提取所有匹配到的
	fmt.Println(ret1)


	// 匹配：数字.数字
	context2 := "123.456.789"
	reg2 := regexp.MustCompile(`\d+\.\d+`)		// +匹配多个， .需要转移
	if reg2 == nil{
		return
	}
	ret2 :=reg2.FindAllStringSubmatch(context2,-1)
	fmt.Println(ret2)

	// 匹配：网页数据
	context3 := `
			<!DOCTYPE html>
			<html lang="zh-cn">
			<head>
				<meta charset="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<meta name="referrer" content="origin" />
				<meta property="og:description" content="一、背景 在最近的项目中的一个需求是消息实时推送消息以及通知功能，项目使用django写的所以决定采用django-channels来实现websocket进行实时通讯。目前官方已经更新到2.1版本，" />
				<meta http-equiv="Cache-Control" content="no-transform" />
				<meta http-equiv="Cache-Control" content="no-siteapp" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<title>django实时通讯--channels2.x使用 - W-D - 博客园</title>
				
				<link rel="stylesheet" href="/css/blog-common.min.css?v=PyyhYDRKBG1sYtpoHme_xHO5AMd5iN57I45iBKF8FVY" />
				<link id="MainCss" rel="stylesheet" href="/skins/darkgreentrip/bundle-darkgreentrip.min.css?v=4KE41eS1YQSSwl64fGzzTUj6ijs-YQFat4AaN-g_jxc" />
				<link type="text/css" rel="stylesheet" href="https://www.cnblogs.com/wdliu/custom.css?v=GOPYYv35uWYu&#x2B;L/HUooSUSoD5Ng=" />
				<link id="mobile-style" media="only screen and (max-width: 767px)" type="text/css" rel="stylesheet" href="/skins/darkgreentrip/bundle-darkgreentrip-mobile.min.css?v=14y836ONdWu0_jxwS0i0t1S3KjlMl55rzUI5k1qhIdc" />
				
				<link type="application/rss+xml" rel="alternate" href="https://www.cnblogs.com/wdliu/rss" />
				<link type="application/rsd+xml" rel="EditURI" href="https://www.cnblogs.com/wdliu/rsd.xml" />
				<link type="application/wlwmanifest+xml" rel="wlwmanifest" href="https://www.cnblogs.com/wdliu/wlwmanifest.xml" />
				<script src="https://common.cnblogs.com/scripts/jquery-2.2.0.min.js"></script>
				<script src="/js/blog-common.min.js?v=9AUvV7CGRXeBxRXvxWjIJxJZemektvGlJfH8yBiYzRQ"></script>
				<script>
					var currentBlogId = 326904;
					var currentBlogApp = 'wdliu';
					var cb_enable_mathjax = false;
					var isLogined = false;
				</script>
			</head>
			<body>
				<a name="top"></a>
				
			<!--done-->
			<div id="home">
			<div id="header">
				<div id="blogTitle">
					<a id="lnkBlogLogo" href="https://www.cnblogs.com/wdliu/"><img id="blogLogo" src="/skins/custom/images/logo.gif" alt="返回主页" /></a>		
					
			<!--done-->
			<h1><a id="Header1_HeaderTitle" class="headermaintitle HeaderMainTitle" href="https://www.cnblogs.com/wdliu/">W-D</a>
			</h1>
			<h2>
			起风的日子
			</h2>

			<div>测试1</div>
			<div>测试2</div>
			<div>测试3</div>
			<div>测试4</div>
			
			</body>
			</html>
	`

	reg3 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)  // ?s 表示\n也处理
	if reg3 == nil{
		return
	}
	ret3 := reg3.FindAllStringSubmatch(context3,-1)
	fmt.Println(ret3)  // [[<div>测试1</div> 测试1] [<div>测试2</div> 测试2] [<div>测试3</div> 测试3] [<div>测试4</div> 测试4]]

	for _,data := range ret3{
		//fmt.Println(data[0])  // 有div标签的
		fmt.Println(data[1])	// 没有div标签的
	}


}
