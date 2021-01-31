SERVER_ADDR="localhost:50052"
CLIENT_PORT=8082
GOPATH=~/go
HTTPS=false

# check https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy

if [ "$HTTPS" = true ] ; then
  #https & http
  $GOPATH/bin/grpcwebproxy \
    --server_tls_cert_file=./misc/localhost.crt \
    --server_tls_key_file=./misc/localhost.key \
    --backend_addr=$SERVER_ADDR \
    --use_websockets \
    --allow_all_origins
else
  #only http
  $GOPATH/bin/grpcwebproxy \
    --backend_addr=$SERVER_ADDR \
    --server_http_debug_port=$CLIENT_PORT \
    --use_websockets \
    --run_tls_server=false \
    --allow_all_origins
fi
