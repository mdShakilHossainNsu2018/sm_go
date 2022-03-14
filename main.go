package main

import (
	"github.com/mdShakilHossainNsu2018/sm_go/accounts"
	"github.com/mdShakilHossainNsu2018/sm_go/interceptors"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/auth_pb"
	"github.com/mdShakilHossainNsu2018/sm_go/protos/user_pb"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"time"
)

const (
	port = ":50051"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	log.Println("Main Program started")

	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// Migrate the schema
	//dbMigrationErr := db.AutoMigrate(&Product{})
	//if dbMigrationErr != nil {
	//	log.Fatal(dbMigrationErr)
	//}
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userStore := accounts.NewInMemoryUserStore()
	jwtManager := accounts.NewJWTManager(secretKey, tokenDuration)
	authService := accounts.NewAuthServer(userStore, jwtManager)
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryInterceptor),
		grpc.StreamInterceptor(interceptors.StreamInterceptor))

	user_pb.RegisterUserServiceServer(s, &accounts.AccountServer{})
	auth_pb.RegisterAuthServiceServer(s, authService)
	log.Printf("Starting gRPC listener on port" + port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
