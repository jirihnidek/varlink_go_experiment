package main

import (
	"context"
	"github.com/jirihnidek/varlink_go_experiment/interface/orgexamplethis"
	"github.com/varlink/go/varlink"
)

type Data struct {
	orgexamplethis.VarlinkInterface
	data string
}

func (d *Data) Ping(ctx context.Context, call orgexamplethis.VarlinkCall, ping string) error {
	return call.ReplyPing(ctx, ping)
}

func main() {
	data := Data{data: "test"}

	service, _ := varlink.NewService(
		"Example",
		"This",
		"1",
		"https://example.org/this",
	)

	ctx := context.TODO()

	err := service.RegisterInterface(orgexamplethis.VarlinkNew(&data))
	if err != nil {
		panic(err)
	}
	err = service.Listen(ctx, "unix:/run/org.example.this", 0)

	if err != nil {
		panic(err)
	}
}
