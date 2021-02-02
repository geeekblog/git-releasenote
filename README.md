# git-releasenote
一个用于生成change log和releasenote的工具，可以通过修改响应的模板来生成不同的结果

当前版本 `0.2.0`

### 目前已实现功能：

* 可以生成releasenote
* 可以生成changelog 
* 可以使用参数选择目标项目的文件夹
* 可以生成指定最新tag数的changelog
* 可以生成最后一个tag之后提交的changelog
* 可以生成指定时间点之后的changelog
* 可以将生成结果输出到指定文件

### 安装：

#### 从release页面下载对应平台的压缩包，解压后使用(推荐)

[release页面](https://github.com/geeekblog/git-releasenote/releases)

#### 通过源码进行安装

1. 通过 `git clone https://github.com/geeekblog/git-releasenote` 下载源码到本地

2. 确保当前在 `release` 分支

3. 运行 `make release` ，编译后的文件和模板文件会在 `bin` 文件夹中

4. 可以手动copy`bin`文件夹中的所有内容到你需要安装的地方

```
PS:不建议直接使用go build直接编译，因为编译的过程中因为参数的不同，有可能会编译出与你需求不同的结果
```

### 使用：

目前只支持一个 `changelog` 子命令

命令中支持的参数可以通过

```
git-releasenote changelog --help
```

来查看

例如：

生成当前目录当前项目分支中的所有changelog到 `CHANGELOG.md` 文件

```
./git-releasenote changelog -o CHANGELOG.md
```

可以使用 `git-releasenote help` 来获得帮助信息

### 未来规划：

* 可以生成releasenote并将changelog对应的功能也实现

### ISSUE：

欢迎在ISSUE中提出你在实际使用中需要的功能，经讨论确认后，会排期进行开发。请优先使用中文，多谢。