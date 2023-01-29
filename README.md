<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="./img/jlog_logo.png" alt="jLog logo"></a>
</p>

<h3 align="center">jLog</h3>

<p align="center"> jLog description.
    <br> 
</p>

## 📝 Table of Contents

- ~~[About](#about)~~
- [Getting Started](#getting-started)
  - [Installing](#installing)
  - ~~[Running](#deployment)~~
  - [Usage](#usage)
- ~~[Built Using](#built_using)~~
- [Developer convention](#developer-convention)


## Getting Started
### Installing
```bash
go get github.com/Sales-Analysis/jLog
```

### ~~Running~~

### Usage

#### import
```go
import (
  jLog "github.com/Sales-Analysis/jLog"
)
```

#### init

```go
j := jLog.Init(envFile string) *jlog 
```
Create new Logger.
* [envFile[Optional]](https://github.com/Sales-Analysis/jLog/blob/main/.env_examples): The path to .exe file with parameters.


#### Example

```go
package SomePackageName

import (
  jLog "github.com/Sales-Analysis/jLog"
)

func SomeFuncName() {
  j := jLog.Init("")
  
  j.Info("This is info")
  j.Warning("This is warning")
  j.Error("This is error")
  j.Dummy("This is dummy")
}
```

```bash
[2023-01-29 19:52:16][SomePackageName][SomeFuncName][INFO]: This is info
[2023-01-29 19:52:16][SomePackageName][SomeFuncName][WARNING]: This is warning
[2023-01-29 19:52:16][SomePackageName][SomeFuncName][ERROR]: This is error
[2023-01-29 19:52:16][SomePackageName][SomeFuncName][DUMMY]: This is dummy.
```


## Developer convention
- [Branch](./docs/Branch.md) The branch name convention
- [Commit](./docs/Commit.md) The commit message convention
