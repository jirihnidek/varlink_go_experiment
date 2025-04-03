package main

import ("github.com/jirihnidek/varlink_go_experiment/interface/orgexamplethis")

type Data struct {
	orgexamplethis.VarlinkInterface
	data string
}

func main() {
        data := Data{data: "test"}

        func (d *Data) Ping(call orgexamplethis.VarlinkCall, ping string) error {
                return call.ReplyPing(ping)
        }

        service, _ = varlink.NewService(
                "Example",
                "This",
                "1",
                 "https://example.org/this",
        )

        service.RegisterInterface(orgexamplethis.VarlinkNew(&data))
        err := service.Listen("unix:/run/org.example.this", 0)
}