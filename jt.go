package main

import (
    "flag"
    "github.com/gorilla/handlers"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    gw "jwt_test/jwt"
    "log"
    "net"
    "net/http"
    "strings"
)

var (
    EndpointGrpc     = flag.String("GrpcEndpoint", ":9090", "endpoint of grpc service")
    EndpointGrpcGw   = flag.String("GwEndPoint", ":8080", "endpoint of gw service")
    EndpointGrpcFull string
    Version          string = "Dev"
)

func startHttpServer() error {

    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
        runtime.WithIncomingHeaderMatcher(func(h string) (string, bool) {
            if strings.EqualFold(h, "Cres-Client-IP") {
                return h, true
            }
            if strings.EqualFold(h, "X-Real-ip") {
                return h, true
            }
            if strings.EqualFold(h, "user-agent") {
                h = "grpcgateway-user-agent"
                return h, true
            }
            if strings.EqualFold(h, "l5d-ctx-trace") {
                return h, true
            }
            return "", false
        }))


    opts := []grpc.DialOption{grpc.WithInsecure()}

    err := gw.RegisterJwtServiceHandlerFromEndpoint(ctx, mux, EndpointGrpcFull, opts)
    if err != nil {
        log.Println(err.Error())
        return err
    }

    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Request-Method", "content-type"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
    muxMain := http.NewServeMux()
    muxMain.Handle("/", mux)

    return http.ListenAndServe(*EndpointGrpcGw, handlers.CORS(originsOk, headersOk, methodsOk)(muxMain))
}

func startGrpcServer() {
    lis, err := net.Listen("tcp", *EndpointGrpc)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcserver := grpc.NewServer()
    d := GrpcRecServer{}
    gw.RegisterJwtServiceServer(grpcserver, &d)

    if err := grpcserver.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %s", err)
    }
}

func main() {
    flag.Parse()
    flag.Usage()

    if strings.HasPrefix(*EndpointGrpc, ":") {
        EndpointGrpcFull = "127.0.0.1" + *EndpointGrpc
    } else {
        EndpointGrpcFull = *EndpointGrpc
    }

    go startGrpcServer()
    if err := startHttpServer(); err != nil {
        log.Fatal(err)
    }
}