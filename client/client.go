package client

import (
	"fmt"

	"alex-j-butler.com/booking-api/rpc"
	"github.com/valyala/gorpc"
)

// Client is an instance of the Booking API client.
type Client struct {
	rpcClient        *gorpc.Client
	dispatcherClient *gorpc.DispatcherClient
	APIKey           string
}

// New creates a new instance of the Booking API Client
// using the specified address and port.
func New(address string, port int, apiKey string) *Client {
	client := &gorpc.Client{
		Addr: fmt.Sprintf("%s:%d", address, port),
	}
	client.Start()

	dispatcher := gorpc.NewDispatcher()
	dispatcher.AddService("v1servers", &rpc.Handlers{})
	serversDispatcher := dispatcher.NewServiceClient("v1servers", client)

	return &Client{
		rpcClient:        client,
		dispatcherClient: serversDispatcher,
		APIKey:           apiKey,
	}
}

// GetServer retrieves a server from the Booking API from a UUID.
func (c Client) GetServer(uuid string) (ServerResource, error) {
	resp, err := c.dispatcherClient.Call("ListServer", &rpc.V1ListServerReq{Key: c.APIKey, UUID: uuid})
	if err != nil {
		return ServerResource{}, err
	}

	listServerResp := resp.(*rpc.V1ListServerResp)
	return ServerResource{V1Server: listServerResp.V1Server}, nil
}

// GetServers retrieves all servers from the Booking API.
func (c Client) GetServers() ([]ServerResource, error) {
	resp, err := c.dispatcherClient.Call("ListServers", &rpc.V1ListServersReq{Key: c.APIKey, Tag: ""})
	if err != nil {
		return nil, err
	}

	listServersResp := resp.(*rpc.V1ListServersResp)
	servers := make([]ServerResource, 0, len(listServersResp.Servers))
	for _, server := range listServersResp.Servers {
		servers = append(servers, ServerResource{V1Server: server})
	}
	return servers, nil
}

// GetServersByTag retrieves all servers with the specified tag from the Booking API.
func (c Client) GetServersByTag(tag string) ([]ServerResource, error) {
	resp, err := c.dispatcherClient.Call("ListServers", &rpc.V1ListServersReq{Key: c.APIKey, Tag: tag})
	if err != nil {
		return nil, err
	}

	listServersResp := resp.(*rpc.V1ListServersResp)
	servers := make([]ServerResource, 0, len(listServersResp.Servers))
	for _, server := range listServersResp.Servers {
		servers = append(servers, ServerResource{V1Server: server})
	}
	return servers, nil
}
