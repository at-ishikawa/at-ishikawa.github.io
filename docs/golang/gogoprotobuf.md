# Protocol Buffers for Go with Gadgets
Written in November, 2019.

gogo/protobuf is the library to store some extensions from [golang/protobuf](https://github.com/golang/protobuf) in [this repository](https://github.com/gogo/protobuf).
There are some useful packages that golang/protobuf does not provide.
For instance,

1. Custom tags

## Getting Started
In this example, protoc-gen-gogo is used but there are other binaries.
To install it, you have to run:
```
go get -u github.com/gogo/protobuf/protoc-gen-gogo
```

## Use custom tags for protobuf

As of Nov. 2019, adding custom tags for generated proto messages are under discussion in golang/protobuf.
[This comment](https://github.com/golang/protobuf/issues/52#issuecomment-372462620) is the formal proposal for it.
Until it's released, gogo/protobuf can be used for such a purpose.
The example is like below, which is stored [here](https://github.com/gogo/protobuf/blob/3f2ed6d/test/tags/tags.proto)

```protobuf
syntax = "proto2";

package tags;

import "github.com/gogo/protobuf/gogoproto/gogo.proto";

option (gogoproto.populate_all) = true;

message Outside {
	optional Inside Inside = 1 [(gogoproto.embed) = true, (gogoproto.jsontag) = ""];
	optional string Field2 = 2 [(gogoproto.jsontag) = "MyField2", (gogoproto.moretags) = "xml:\",comment\""];
	oneof filed {
		string Field3 = 3 [(gogoproto.jsontag) = "MyField3", (gogoproto.moretags) = "xml:\",comment\""];
	}
}

message Inside {
	optional string Field1 = 1 [(gogoproto.jsontag) = "MyField1", (gogoproto.moretags) = "xml:\",chardata\""];
}
```

You can generate golang files for the above proto file from `proto` to `gen` directory by:

```
> protoc --gogo_out gen -I ./proto -I $GOPATH/src proto/*.proto
```
