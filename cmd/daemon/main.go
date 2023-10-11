package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/0xdeadbad/simple-raft/sraft"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Bind    string   `short:"b" long:"bind" description:"RPC bind address" required:"true" default:"0.0.0.0:9091"`
	Cluster []string `short:"c" long:"cluster" description:"Cluster addresses list" required:"true" default:"127.0.0.1:9091"`
	Id      int      `short:"i" long:"id" description:"This node's ID" required:"true"`
	Debug   bool     `short:"d" long:"debug" description:"Show debugging messages" default:"false"`
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		panic(err)
	}

	bind := opts.Bind
	clusters := opts.Cluster
	id := opts.Id

	ns := make(map[int]*sraft.Node)
	for k, v := range clusters {
		node, err := sraft.NewNode(v)
		if err != nil {
			panic(err)
		}
		ns[k] = node
	}

	laddr, err := net.ResolveTCPAddr("tcp4", bind)
	if err != nil {
		panic(err)
	}

	raft, err := sraft.NewRaftNode(ctx, id, laddr, ns)
	if err != nil {
		panic(err)
	}

	raft.Start(ctx)

	select {}

}
