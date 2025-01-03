# easy-novel

# 当前命令行版本已经不在维护，请到 <https://github.com/767829413/fy-novel> 提供了最新的linux和windows 客户端版本方便使用！！！

## 介绍

用户欲免费阅新书且求最佳阅读体验者之所需。

今国内网络小说多为完本，若欲观新书，或需于正版平台付费，或借笔趣阁等网站，或用“阅读”等安卓应用。

此二法虽足常人，然众口难调，或嫌界面不美，或讥功能不足，或受平台所限。此时，阅读器之长处显现：可自定义。此工具之大用在于能将连载新书下载为epub等格式，便可导入所好之阅读器。

欲求下载小说之工具，得一Java版本，然Java深奥难测，故以Golang权作替代。

至于完本小说，亦可搜索下载。若有错字排版之误，宜自行搜寻精校版。

## 使用说明

***抄的 https://github.com/freeok/so-novel 使用方式请君对其***

```bash
./easynovel -c $HOME/.easy-novel.yaml 
按 Tab 键选择功能: 1
开始下载小说...
==> 请输入书名或作者（宁少字别错字）: fg
<== 正在搜索...
<== 搜索到 0 条记录，耗时 0.992025 s
按 Tab 键选择功能: gg
无效的选项，请重新选择
按 Tab 键选择功能: 1
开始下载小说...
==> 请输入书名或作者（宁少字别错字）: gg
<== 正在搜索...
<== 搜索到 5 条记录，耗时 0.516676 s
|------|----------------------|------------|--------------------------------|--------------|
| 序号 |         书名         |    作者    |            最新章节            | 最后更新时间 |
|------|----------------------|------------|--------------------------------|--------------|
|  1   |    猎户家的俏农媳    |   景辰gg   |       第50章.疯狂的月儿        |    12-09     |
|------|----------------------|------------|--------------------------------|--------------|
|  2   |   火影：我是大反派   |  墨墨子GG  |       第三百零九章 获救        |    12-31     |
|------|----------------------|------------|--------------------------------|--------------|
|  3   | 超级无敌天帝养成系统 | 母猪上树GG |        第190章 双喜设计        |    04-15     |
|------|----------------------|------------|--------------------------------|--------------|
|  4   |      序列零封神      |   零神GG   |     第二卷：新生 第十三章      |    07-24     |
|      |                      |            |      地底变故，初遇海妖！      |              |
|------|----------------------|------------|--------------------------------|--------------|
|  5   |  都市怪谈：前世今生  |   一休GG   |       第二十二章  帆布鞋       |    12-17     |
|------|----------------------|------------|--------------------------------|--------------|
==> 请输入下载序号（首列的数字，或输入 0 返回）： 
```

## 免责声明

此程序乃作者研习Go语言之练习项目，倘使用中有何问题，皆与作者无关！！！