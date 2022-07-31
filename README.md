# use_protobuf_define_multi_version_example
使用 protobuf 定义多版本的例子

## 依赖

```shell
cd example1
make install_protoc
```

## 生成 pb 文件

```shell
cd example1
make proto
```

## proto 文件定义

```protobuf
syntax = "proto3";
package examplepb;
option go_package = "examplepb/";
import "github.com/fananchong/versionpb/version.proto";

message Msg1 {
  option (versionpb.version_msg) = "3.0";
  enum Enum1 {
    option (versionpb.version_enum) = "3.0";
    E1 = 0;
    E2 = 1;
    E3 = 2 [ (versionpb.version_enum_value) = "3.3" ];
  }
  bytes f1 = 1;
  int64 f2 = 2;
  Enum1 f3 = 3;
  bool f4 = 4;
  int64 f5 = 5 [ (versionpb.version_field) = "3.1" ];
  Enum1 f6 = 6 [ (versionpb.version_field) = "3.2" ];
}
```

## 例子介绍

| 例子     | 协议版本 | 说明                                                           |
| :------- | :------- | :------------------------------------------------------------- |
| example1 | proto 3  | **注意， f5 字段赋值 0 ，被优化掉了。导致 3.1.0 版本没识别到** |
| example2 | proto 2  | **正确识别 f5 字段的 3.1.0 版本**                              |
