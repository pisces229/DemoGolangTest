# DemoGolangTest 

## Unit Test 3A 原則
* Arrange – 初始化
* Act – 行為，測試對象的執行過程
* Assert – 驗證結果

[Module]

`go mod init demo.golang.test`

[test](https://github.com/stretchr/testify/assert)

`go get github.com/stretchr/testify/assert`

`go test <path> -v`

`go test <path> -run <method> -v -cover `

`go test ./home -v`

`go test ./home -coverprofile=cover.out`

`go tool cover -html=cover.out`

[mock](https://github.com/golang/mock)

`go get github.com/golang/mock/gomock`

`go install github.com/golang/mock/mockgen@latest`

`mockgen -destination <path>/mock_<name>.go -source <path>/<name>.go -package <package>`

`mockgen -destination mock/person.go -source person/person.go -package mock`

[gorm mock](https://github.com/DATA-DOG/go-sqlmock)

`go get github.com/DATA-DOG/go-sqlmock`

## [常用 mock 方法](https://pkg.go.dev/github.com/golang/mock@v1.6.0/gomock#section-documentation)
### 調用方法
* Call.Do()：聲明在匹配時要運行的操作
* Call.DoAndReturn()：聲明在匹配調用時要運行的操作，並且模擬返回該函數的返回值
* Call.MaxTimes()：設置最大的調用次數爲 n 次
* Call.MinTimes()：設置最小的調用次數爲 n 次
* Call.AnyTimes()：允許調用次數爲 0 次或更多次
* Call.Times()：設置調用次數爲 n 次
### 參數匹配
* gomock.Any()：匹配任意值
* gomock.Eq()：通過反射匹配到指定的類型值，而不需要手動設置
* gomock.Nil()：返回 nil

### 生成多個 mock 文件

`go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]`

> add person.go `//go:generate mockgen -destination ../mock/person.go -package mock demo.golang.test/person Person`

`go generate ./...`