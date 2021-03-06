package ipfs

import (
	"io"
	"io/ioutil"
	"os"

	"gx/ipfs/QmQvJiADDe7JR4m968MwXobTCCzUqQkP87aRHe29MEBGHV/go-logging"
	"gx/ipfs/QmcVVHfdyv15GVPk7NrxdWjh2hLVccXnoD8j2tyQShiXJb/go-log"
	logWriter "gx/ipfs/QmcVVHfdyv15GVPk7NrxdWjh2hLVccXnoD8j2tyQShiXJb/go-log/writer"

	ipfsconfig "github.com/ipfs/go-ipfs/repo/config"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

// Init creates an initialized .ipfs directory in the directory `path`.
// The generated RSA key will have `keySize` bits.
func Init(path string, keySize int) error {
	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}

	// init, but discard the log messages about generating a key.
	cfg, err := ipfsconfig.Init(ioutil.Discard, keySize)
	if err != nil {
		return err
	}

	// Init the actual data store.
	if err := fsrepo.Init(path, cfg); err != nil {
		return err
	}

	return nil
}

// ForwardLog routes all ipfs logs to a file provided by brig.
// Only messages >= INFO are logged.
func (nd *Node) ForwardLog(w io.Writer) {
	logWriter.Configure(logWriter.Output(w))
	log.SetAllLoggers(logging.INFO)
}
