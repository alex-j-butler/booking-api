package client

import "alex-j-butler.com/booking-api/rpc"

type ServerResource struct {
	rpc.V1Server

	// client represents the booking API client this server resource was retrieved from.
	client *Client
}

func (s ServerResource) Start(client *Client) error {
	_, err := client.dispatcherClient.Call("Start", &rpc.V1StartServerReq{Key: client.APIKey, UUID: s.UUID})
	if err != nil {
		return err
	}

	return nil
}

func (s ServerResource) Stop(client *Client) error {
	_, err := client.dispatcherClient.Call("Stop", &rpc.V1StopServerReq{Key: client.APIKey, UUID: s.UUID})
	if err != nil {
		return err
	}

	return nil
}

func (s ServerResource) SetPassword(client *Client, rconPassword, srvPassword string) error {
	_, err := client.dispatcherClient.Call("SetPassword", &rpc.V1SetPasswordReq{Key: client.APIKey, UUID: s.UUID, RCONPassword: rconPassword, ServerPassword: srvPassword})
	if err != nil {
		return err
	}

	return nil
}

func (s ServerResource) SendCommand(client *Client, command string) error {
	_, err := client.dispatcherClient.Call("SendCommand", &rpc.V1SendCommandReq{Key: client.APIKey, UUID: s.UUID, Command: command})
	if err != nil {
		return err
	}

	return nil
}

func (s ServerResource) Update(client *Client) (bool, error) {
	resp, err := client.dispatcherClient.Call("Update", &rpc.V1UpdateServerReq{Key: client.APIKey, UUID: s.UUID})
	if err != nil {
		return false, err
	}

	return resp.(bool), nil
}

func (s ServerResource) UploadDemos(client *Client, uploader string) ([]string, error) {
	resp, err := client.dispatcherClient.Call("UploadDemos", &rpc.V1UploadDemosReq{Key: client.APIKey, UUID: s.UUID, Uploader: uploader})
	if err != nil {
		return nil, err
	}

	uploadDemosResp := resp.(*rpc.V1UploadDemosResp)
	return uploadDemosResp.Demos, nil
}

func (s ServerResource) Console(client *Client, lines int) ([]string, error) {
	resp, err := client.dispatcherClient.Call("Console", &rpc.V1ConsoleServerReq{Key: client.APIKey, UUID: s.UUID, Lines: lines})
	if err != nil {
		return nil, err
	}

	consoleResp := resp.(*rpc.V1ConsoleServerResp)
	return consoleResp.ConsoleLines, nil
}

func (s ServerResource) AddTag(client *Client, tag string) (bool, error) {
	resp, err := client.dispatcherClient.Call("AddServerTag", &rpc.V1TagReq{Key: client.APIKey, UUID: s.UUID, Tag: tag})
	if err != nil {
		return false, err
	}

	return resp.(*rpc.V1TagResp).Success, nil
}

func (s ServerResource) RemoveTag(client *Client, tag string) (bool, error) {
	resp, err := client.dispatcherClient.Call("RemoveServerTag", &rpc.V1TagReq{Key: client.APIKey, UUID: s.UUID, Tag: tag})
	if err != nil {
		return false, err
	}

	return resp.(*rpc.V1TagResp).Success, nil
}
