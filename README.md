# Golang go-micro

# go-micro的框架結構

![image](./images/20210110110029.png)

# 下載安裝

```bash
go get -u -v github.com/micro/micro
go get -u -v github.com/micro/go-micro
go get -u -v github.com/micro/protoc-gen-micro
```

![image](./images/20210110112157.png)

# msg.proto

```bash
syntax = "proto3";

// 定義需要的結構體參數
message InfoRequest{
    string username = 1;
}

message InfoResponse{
    string msg = 2;
}

// 定義接口
service Hello{
    rpc Info(InfoRequest) returns (InfoResponse){
        
    }
}
```

# 生成.go文件(會有兩個檔案，另一個是micro)

```bash
protoc -I . --micro_out=. --go_out=. ./*.proto
```



![image](./images/20210110112455.png)

![image](./images/20210110112518.png)

# 生成的文件似乎有點問題

報了這個錯誤，但這些包下了對應的go get也捉不到

![image](./images/20210110131250.png)

![image](./images/20210110121217.png)

![image](./images/20210110121108.png)

後來靈機一動將v2拿掉

![image](./images/20210110121157.png)

就可以運行了

![image](./images/20210110121557.png)

未完待續