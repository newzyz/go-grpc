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
<b>(Optional) Generate protocol buffer files</b>
<pre>
  make create
</pre>
<p>For more details: Makefile</p>
    
<b>#1 Run Docker</b>
<pre>
  docker build .
</pre>
<pre>
  docker compose up
</pre>
 <p> #Use dbbackup.sql in dir /restored_db/dbbackup.sql to restored DB </p>
 
<b>#1 Run go</b>
<p>Run go locally </p>
<pre>
  make go-run 
</pre>

<h1>#Request Spec</h1>
<pre>
    POST http://localhost:3001/book #Getbooks require raw json {"page": int64,"per_page": int64}</br>
    GET http://localhost:3001/book/{id} #Getbook</br>
    POST http://localhost:3001/createbook #CreateBook require raw json {"title": string, "genre":string, "author":string}</br>
    PUT http://localhost:3001/book #UpdateBook require raw json {"id":int64 ,"title": string, "genre":string, "author":string}</br>
    DELETE http://localhost:3001/book/{id} #DeleteBook
</pre>

<pre>
    POST http://localhost:3001/customer #Getcustomers require raw json {"page": int64,"per_page": int64}</br>
    GET http://localhost:3001/customer/{id} #Getcustomer</br>
    POST http://localhost:3001/createcustomer #Createcustomer require raw json {"firstname":string,"lastname":string,"age": int64}</br>
    PUT http://localhost:3001/customer #Updatecustomer require raw json {"id":int64 ,"firstname":string,"lastname":string,"age": int64}</br>
    DELETE http://localhost:3001/customer/{id} #Deletecustomer
</pre>

