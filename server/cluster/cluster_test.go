package cluster

import (
	"github.com/smancke/guble/protocol"

	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"

	"errors"
	"net"
	"testing"
	"time"
)

const basePort = 10000

var (
	index = 1
)

func testConfig() (config Config) {
	remoteAddr := net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: basePort + index}
	var remotes []*net.TCPAddr
	remotes = append(remotes, &remoteAddr)
	config = Config{ID: index, Host: "127.0.0.1", Port: basePort + index, Remotes: remotes}
	index++
	return
}

func testConfigAnother() (config Config) {
	remoteAddr := net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: basePort + index - 1}
	var remotes []*net.TCPAddr
	remotes = append(remotes, &remoteAddr)
	config = Config{ID: index, Host: "127.0.0.1", Port: basePort + index, Remotes: remotes}
	index++
	return
}

func TestCluster_StartCheckStop(t *testing.T) {
	a := assert.New(t)

	conf := testConfig()
	node, err := New(&conf)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node.MessageHandler = DummyMessageHandler{}

	err = node.Start()
	a.NoError(err, "No error should be raised when Starting the Cluster")

	err = node.Check()
	a.NoError(err, "Health-check score of a Cluster with a single node should be OK")

	err = node.Stop()
	a.NoError(err, "No error should be raised when Stopping the Cluster")
}

func TestCluster_BroadcastStringAndMessageAndCheck(t *testing.T) {
	// defer testutil.EnableDebugForMethod()
	a := assert.New(t)

	config1 := testConfig()
	node1, err := New(&config1)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node1.MessageHandler = DummyMessageHandler{}

	//start the cluster node 1
	defer node1.Stop()
	err = node1.Start()
	a.NoError(err, "No error should be raised when starting node 1 of the Cluster")

	config2 := testConfigAnother()
	node2, err := New(&config2)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node2.MessageHandler = DummyMessageHandler{}

	//start the cluster node 2
	defer node2.Stop()
	err = node2.Start()
	a.NoError(err, "No error should be raised when starting node 2 of the Cluster")

	// Send a String Message
	str := "TEST"
	err = node1.BroadcastString(&str)
	a.NoError(err, "No error should be raised when sending a string to Cluster")

	// and a protocol message
	pmsg := protocol.Message{
		ID:            1,
		Path:          "/stuff",
		UserID:        "id",
		ApplicationID: "appId",
		Time:          time.Now().Unix(),
		HeaderJSON:    "{}",
		Body:          []byte("test"),
		NodeID:        1}
	err = node1.BroadcastMessage(&pmsg)
	a.NoError(err, "No error should be raised when sending a protocol message to Cluster")

	err = node1.Check()
	a.NoError(err, "Health-check score of a Cluster with 2 nodes should be OK for node 1")

	err = node2.Check()
	a.NoError(err, "Health-check score of a Cluster with 2 nodes should be OK for node 2")
}

func TestCluster_NewShouldReturnErrorWhenPortIsInvalid(t *testing.T) {
	a := assert.New(t)

	remoteAddr := net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: basePort + index - 1}
	var remotes []*net.TCPAddr
	remotes = append(remotes, &remoteAddr)
	index++

	config := Config{ID: 1, Host: "localhost", Port: -1, Remotes: remotes}
	_, err := New(&config)
	if a.Error(err, "An error was expected when Creating the Cluster") {
		a.Equal(err, errors.New("Failed to start TCP listener. Err: listen tcp :-1: bind: invalid argument"),
			"Error should be precisely defined")
	}
}

func TestCluster_StartShouldReturnErrorWhenNoRemotes(t *testing.T) {
	a := assert.New(t)

	var remotes []*net.TCPAddr
	index++

	config := Config{ID: 1, Host: "localhost", Port: basePort + index - 1, Remotes: remotes}
	node, err := New(&config)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node.MessageHandler = DummyMessageHandler{}

	defer node.Stop()
	err = node.Start()
	if a.Error(err, "An error is expected when Starting the Cluster") {
		a.Equal(err, errors.New("No remote hosts were successfully contacted when this node wanted to join the cluster"),
			"Error should be precisely defined")
	}
}

func TestCluster_StartShouldReturnErrorWhenInvalidRemotes(t *testing.T) {
	a := assert.New(t)

	remoteAddr := net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: 0}
	var remotes []*net.TCPAddr
	remotes = append(remotes, &remoteAddr)
	index++

	config := Config{ID: 1, Host: "localhost", Port: basePort + index - 1, Remotes: remotes}
	node, err := New(&config)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node.MessageHandler = DummyMessageHandler{}

	defer node.Stop()
	err = node.Start()
	if a.Error(err, "An error is expected when Starting the Cluster") {
		expected := multierror.Append(errors.New("Failed to join 127.0.0.1: dial tcp 127.0.0.1:0: getsockopt: connection refused"))
		a.Equal(err, expected, "Error should be precisely defined")
	}
}

func TestCluster_StartShouldReturnErrorWhenNoMessageHandler(t *testing.T) {
	a := assert.New(t)

	config := testConfig()
	node, err := New(&config)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	defer node.Stop()
	err = node.Start()
	if a.Error(err, "An error is expected when Starting the Cluster") {
		expected := errors.New("There should be a valid MessageHandler already set-up")
		a.Equal(err, expected, "Error should be precisely defined")
	}
}

func TestCluster_NotifyMsgShouldSimplyReturnWhenDecodingInvalidMessage(t *testing.T) {
	a := assert.New(t)

	config := testConfig()
	node, err := New(&config)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node.MessageHandler = DummyMessageHandler{}

	defer node.Stop()
	err = node.Start()
	a.NoError(err, "No error should be raised when Starting the Cluster")

	node.NotifyMsg([]byte{})

	//TODO Cosmin check that HandleMessage is not invoked (i.e. invalid message is not dispatched)
}

func TestCluster_broadcastClusterMessage(t *testing.T) {
	a := assert.New(t)

	config := testConfig()
	node, err := New(&config)
	a.NoError(err, "No error should be raised when Creating the Cluster")

	node.MessageHandler = DummyMessageHandler{}

	defer node.Stop()
	err = node.Start()
	a.NoError(err, "No error should be raised when Starting the Cluster")

	err = node.broadcastClusterMessage(nil)
	if a.Error(err, "An error is expected from broadcastClusterMessage") {
		expected := errors.New("Could not broadcast a nil cluster-message")
		a.Equal(err, expected, "Error should be precisely defined")
	}
}

type DummyMessageHandler struct {
}

func (dmh DummyMessageHandler) HandleMessage(pmsg *protocol.Message) error {
	return nil
}
