# yekGo-ProtoBuffer
yekGo-ProtoBuffer Protocol buffer example which is supported with gorm(golang Object Relation Mapping) library and also this is proto

# What are protocol buffers?
Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

For detail information https://developers.google.com/protocol-buffers/ 

For detail usage https://grpc.io/docs/quickstart/go.html 

## Quick Start

```bash
git clone https://github.com/ysBayram/yekGo-ProtoBuffer.git
go run ./yekGo-ProtoBuffer/server/server.go
go run ./yekGo-ProtoBuffer/client/client.go

```

## To Do
- [ ] Add Dynamic Proto File Creater
- [ ] Add Different Type Database Configuration
- [ ] Add unit tests
- [ ] Changeable Port Number 
- [ ] Struct Create By Existiance Table
- [ ] API endpoint create depend on structs