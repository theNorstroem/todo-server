package main

import (
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/markbates/pkger"
	"github.com/spf13/viper"
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	"github.com/theNorstroem/todo-server/pkg/auth"
	"github.com/theNorstroem/todo-server/pkg/person"
	person_grpcHandler "github.com/theNorstroem/todo-server/pkg/person/infrastructure/grpcHandler"
	"github.com/theNorstroem/todo-server/pkg/task"
	task_grpcHandler "github.com/theNorstroem/todo-server/pkg/task/infrastructure/grpcHandler"
	"github.com/theNorstroem/todo-specs/pkg/grpc-gateway"
	middleware_auth "github.com/veith/fgs-lib/pkg/auth"
	"github.com/veith/fgs-lib/pkg/configloader"
	"github.com/veith/fgs-lib/pkg/microbus"
	"github.com/veith/fgs-lib/pkg/webserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {
	configloader.Load("todo")

	// Setup the infrastructure
	infrastructure.Infrastructure.Grpc = grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_auth.UnaryServerInterceptor(middleware_auth.JWTAuthFunc))))

	// rest Router ist für direkte REST APIs gedacht (maintenance und statusabfrage,...)
	//infrastructure.Infrastructure.AdminRest = httprouter.New()

	infrastructure.Infrastructure.InternalEventBus = microbus.NewMicrobus()
	//infrastructure.Infrastructure.DbSession = upper.ConnectMysql()
	//defer infrastructure.Infrastructure.DbSession.Close() // Closing the session is a good practice.

	// register modules
	// register the apis after the modules are ready
	auth.RegisterModule()
	task.RegisterApplicationModule()
	task_grpcHandler.RegisterGrpcApis()

	demo.RegisterApplicationModule()
	demo_grpcHandler.RegisterGrpcApis()

	person.RegisterApplicationModule()
	person_grpcHandler.RegisterGrpcApis()

	fmt.Println(time.Now())
	// start the Transcoder
	// transcoder is not in infrastructure because it starts as a standalone service
	if viper.GetBool("server.transcoder.start") {
		go func() {
			fmt.Println("TRANSCODER started on: " + viper.GetString("server.transcoder.addr"))
			if err := transcoder.Run(viper.GetString("server.grpc.addr"), viper.GetString("server.transcoder.addr")); err != nil {
				log.Fatal(err)
			}
		}()
	}

	fmt.Println("GRPC started on: " + viper.GetString("server.grpc.addr"))
	lis, err := net.Listen("tcp", viper.GetString("server.grpc.addr"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if viper.GetBool("server.webserver.start") {
		go func() {
			webserver.Start(pkger.Dir("/public"))
		}()
	}
	// Register reflection service on gRPC server.
	reflection.Register(infrastructure.Infrastructure.Grpc)
	if err := infrastructure.Infrastructure.Grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
