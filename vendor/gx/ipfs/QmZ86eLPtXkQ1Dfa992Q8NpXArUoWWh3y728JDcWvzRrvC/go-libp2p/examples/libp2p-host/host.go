package main

import (
	"context"
	"crypto/rand"
	"fmt"

	libp2p "gx/ipfs/QmZ86eLPtXkQ1Dfa992Q8NpXArUoWWh3y728JDcWvzRrvC/go-libp2p"
	crypto "gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
)

func main() {
	// The context governs the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// To construct a simple host with all the default settings, just use `New`
	h, err := libp2p.New(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my hosts ID is %s\n", h.ID())

	// If you want more control over the configuration, you can specify some
	// options to the constructor

	// Set your own keypair
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}

	h2, err := libp2p.New(ctx,
		// Use your own created keypair
		libp2p.Identity(priv),

		// Set your own listen address
		// The config takes an array of addresses, specify as many as you want.
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello World, my second hosts ID is %s\n", h2.ID())
}
