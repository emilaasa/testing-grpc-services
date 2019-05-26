package greeter

import (
	"context"
	fmt "fmt"
	"testing"
)

type service struct{}

func (s service) SayHello(ctx context.Context, request *HelloRequest) HelloReply {
	return HelloReply{Message: fmt.Sprintf("Hello %s!", request.Name)}
}

func TestHello(t *testing.T) {
	s := service{}
	tests := []struct {
		name  string
		input *HelloRequest
		want  string
	}{
		{"simple string", &HelloRequest{Name: "world"}, "Hello world!"},
		{"numbers", &HelloRequest{Name: "123"}, "Hello 123!"},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testcase: '%s' input: '%s'", tc.name, tc.input.Name), func(t *testing.T) {
			got := s.SayHello(context.Background(), tc.input)
			if got.Message != tc.want {
				t.Errorf("got %s; want %s", got.Message, tc.want)
			}

		})

	}
}
