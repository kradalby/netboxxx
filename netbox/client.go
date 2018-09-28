package netboxxx

import (
	// "github.com/digitalocean/go-netbox/netbox"
	//"context"
	"fmt"
	"github.com/digitalocean/go-netbox/netbox/client"
	runtimeclient "github.com/go-openapi/runtime/client"

	//"github.com/digitalocean/go-netbox/netbox/client/tenancy"
	cidr "github.com/apparentlymart/go-cidr/cidr"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"net"

	"github.com/fatih/structs"
	"github.com/flosch/pongo2"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

type Client struct {
	c *client.NetBox
}

type TokenAuth struct {
	token string
}

type NetBoxNetwork struct {
	Description string
	IPNet       *net.IPNet
}

type Network struct {
	Name    string
	Network string
	Netmask string
	Gateway string
	Start   string
	End     string
}

func (t TokenAuth) AuthenticateRequest(req runtime.ClientRequest, _ strfmt.Registry) error {
	req.SetHeaderParam("Authorization", fmt.Sprintf("Token %s", t.token))
	return nil
}

func NewClient(host string, apiKey string) *Client {
	//keyTransport := TokenAuth{token: apiKey}

	transport := runtimeclient.New(host, client.DefaultBasePath, []string{"https"})
	transport.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiKey))

	return &Client{c: client.New(transport, strfmt.Default)}
}

func (c *Client) getNetBoxNetworks() ([]*NetBoxNetwork, error) {

	rs, err := c.c.IPAM.IPAMPrefixesList(nil, nil)
	if err != nil {
		return []*NetBoxNetwork{}, err
	}

	networks := make([]*NetBoxNetwork, *rs.Payload.Count)

	for index, prefix := range rs.Payload.Results {
		_, ipNet, err := net.ParseCIDR(*prefix.Prefix)
		if err != nil {
			return []*NetBoxNetwork{}, err
		}

		networks[index] = &NetBoxNetwork{
			Description: prefix.Description,
			IPNet:       ipNet,
		}

	}

	return networks, nil

}

func (c *Client) PrintNetworks() {

	networks, err := c.getNetBoxNetworks()
	if err != nil {
		fmt.Printf("[ERROR] %#v\n", err)
	}

	for _, network := range networks {
		fmt.Printf("%#v\n", network)
	}

}

func (c *Client) PrintTemplate(templateFile string) {

	networks, err := c.createNetworks()
	if err != nil {
		fmt.Printf("[ERROR] %#v\n", err)
		panic("Exited")
	}

	for _, network := range networks {

		m := structs.Map(network)

		template := pongo2.Must(pongo2.FromFile(templateFile))

		out, err := template.Execute(m)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)

	}
}

func (c *Client) createNetworks() ([]*Network, error) {

	netBoxNetworks, err := c.getNetBoxNetworks()
	if err != nil {
		return []*Network{}, err
	}

	networks := make([]*Network, len(netBoxNetworks))

	for index, nbNet := range netBoxNetworks {

		networkAddress, broadcast := cidr.AddressRange(nbNet.IPNet)
		gateway := cidr.Inc(networkAddress)
		start := cidr.Inc(gateway)
		end := cidr.Dec(broadcast)

		networks[index] = &Network{
			Name:    nbNet.Description,
			Network: networkAddress.String(),
			Netmask: fmt.Sprintf("%d.%d.%d.%d",
				nbNet.IPNet.Mask[0],
				nbNet.IPNet.Mask[1],
				nbNet.IPNet.Mask[2],
				nbNet.IPNet.Mask[3]),
			Gateway: gateway.String(),
			Start:   start.String(),
			End:     end.String(),
		}
	}

	return networks, nil

}
