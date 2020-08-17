package main
import (
	"context"
	"time"
	"log"
	"os"

	"google.golang.org/grpc"

	pb "github.com/vglvishnu/gosamples/helloworld/helloworldproto"
)
const(
	address = "localhost:50051"
	defaultName = "world"
)
func main(){
	con,err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatal("Did not connect :: %+v\n", err)
	}
	defer con.Close()
	c := pb.NewGreetServiceClient(con)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	r,err := c.SayHello(ctx,&pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet :: %+v", err)
	}
	log.Printf("Received greetings Reply::%+v\n", r.GetMessage())
}