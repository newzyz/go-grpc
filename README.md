# go-grpc
<h1>#Set up</h1>
<b>#1 Install Protocol Buffer Compiler
<p>https://grpc.io/docs/protoc-installation/</p>
<br/>
<b>#2 Package Install protoc-gen-go and protoc-gen-go-grpc</b>
<p>https://grpc.io/docs/languages/go/quickstart/</p>
<br/>
<b>#3 go install these for grpc gateway</b>
<pre>
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
</pre>
<p>For more details: https://github.com/grpc-ecosystem/grpc-gateway</p>

<h1>#Run</h1>
<b>#1 Generate protocol buffer files cr</b>
<pre>
  make create
</pre>
<p>For more details: Makefile</p>

<b>#2 Run go</b>
<pre>
  make go-run
</pre>




