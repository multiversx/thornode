package multiversx

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cKeys "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"gitlab.com/thorchain/thornode/bifrost/metrics"
	"gitlab.com/thorchain/thornode/bifrost/pubkeymanager"
	"gitlab.com/thorchain/thornode/bifrost/thorclient"
	stypes "gitlab.com/thorchain/thornode/bifrost/thorclient/types"
	"gitlab.com/thorchain/thornode/cmd"
	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/config"
	openapi "gitlab.com/thorchain/thornode/openapi/gen"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
	. "gopkg.in/check.v1"
)

func TestMvXPackage(t *testing.T) { TestingT(t) }

// MultiversXSuite represent the struct for testing the MultiversX client
type MultiversXSuite struct {
	thordir  string
	thorKeys *thorclient.Keys
	bridge   thorclient.ThorchainBridge
	m        *metrics.Metrics
	server   *httptest.Server
}

var _ = Suite(&MultiversXSuite{})

var m *metrics.Metrics

func GetMetricForTest(c *C) *metrics.Metrics {
	if m == nil {
		var err error
		m, err = metrics.NewMetrics(config.BifrostMetricsConfiguration{
			Enabled:      false,
			ListenPort:   9000,
			ReadTimeout:  time.Second,
			WriteTimeout: time.Second,
			Chains:       common.Chains{common.MVXChain},
		})
		c.Assert(m, NotNil)
		c.Assert(err, IsNil)
	}
	return m
}

func (s *MultiversXSuite) SetUpTest(c *C) {
	s.m = GetMetricForTest(c)
	c.Assert(s.m, NotNil)
	types.SetupConfigForTest()
	c.Assert(os.Setenv("NET", "mocknet"), IsNil)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.RequestURI {
		case thorclient.ChainVersionEndpoint:
			_, err := rw.Write([]byte(`{"current":"` + types.GetCurrentVersion().String() + `"}`))
			c.Assert(err, IsNil)
		case thorclient.PubKeysEndpoint:
			priKey, _ := s.thorKeys.GetPrivateKey()
			tm, _ := codec.ToTmPubKeyInterface(priKey.PubKey())
			pk, err := common.NewPubKeyFromCrypto(tm)
			c.Assert(err, IsNil)
			content, err := os.ReadFile("../../../../test/fixtures/endpoints/vaults/pubKeys.json")
			c.Assert(err, IsNil)
			var pubKeysVault openapi.VaultPubkeysResponse
			c.Assert(json.Unmarshal(content, &pubKeysVault), IsNil)
			chain := common.ETHChain.String()
			router := "0xE65e9d372F8cAcc7b6dfcd4af6507851Ed31bb44"
			pubKeysVault.Asgard = append(pubKeysVault.Asgard, openapi.VaultInfo{
				PubKey: pk.String(),
				Routers: []openapi.VaultRouter{
					{
						Chain:  &chain,
						Router: &router,
					},
				},
			})
			buf, err := json.MarshalIndent(pubKeysVault, "", "	")
			c.Assert(err, IsNil)
			_, err = rw.Write(buf)
			c.Assert(err, IsNil)
		case thorclient.InboundAddressesEndpoint:
			httpTestHandler(c, rw, "../../../../test/fixtures/endpoints/inbound_addresses/inbound_addresses.json")
		case thorclient.AsgardVault:
			httpTestHandler(c, rw, "../../../../test/fixtures/endpoints/vaults/asgard.json")
		case thorclient.LastBlockEndpoint:
			httpTestHandler(c, rw, "../../../../test/fixtures/endpoints/lastblock/root.json")
		case thorclient.NodeAccountEndpoint:
			httpTestHandler(c, rw, "../../../../test/fixtures/endpoints/nodeaccount/template.json")
		case "/thorchain/mimir/key/MaxUTXOsToSpend":
			_, err := rw.Write([]byte(`-1`))
			c.Assert(err, IsNil)
		default:
			body, err := io.ReadAll(req.Body)
			c.Assert(err, IsNil)
			type RPCRequest struct {
				JSONRPC string          `json:"jsonrpc"`
				ID      interface{}     `json:"id"`
				Method  string          `json:"method"`
				Params  json.RawMessage `json:"params"`
			}
			var rpcRequest RPCRequest
			err = json.Unmarshal(body, &rpcRequest)
			c.Assert(err, IsNil)
			if rpcRequest.Method == "eth_getBalance" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x3b9aca00"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_getTransactionCount" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x0"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_chainId" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xf"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_gasPrice" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_estimateGas" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x493df"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_getBlockByNumber" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x2","extraData":"0xd88301090e846765746888676f312e31342e32856c696e757800000000000000ef855333e6b03b825c2f1381f111e278232688e21ba8c36aa35689505d9470704420825b302cd70cc6610f1334a3d7c801ac4b8871bd9f0692c1c96f0f60ee0f01","gasLimit":"0x7a1200","gasUsed":"0xfbc9","hash":"0x45f139a64f563e12f61824a4b44edc2c955818d176b160538ae24f566a006c00","logsBloom":"0x00000000000000000002000000000000000000100000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400000000000800000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000004000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0x7","parentHash":"0x2f202f8aa7355e77bfbdcd63c08f7c4e43e0bcca61b45fe6a2bdb950d777fa38","receiptsRoot":"0xe1cf0352843e29447633b9f1710e443f2582691e4278febf322c0bb7f86202cc","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x38c","stateRoot":"0x303f9a24ba76fa8f350d36f4cef139e6be023f95646e2602cf9e6f939f91beea","timestamp":"0x5fde861b","totalDifficulty":"0xf","transactions":[{"blockHash":"0x45f139a64f563e12f61824a4b44edc2c955818d176b160538ae24f566a006c00","blockNumber":"0x7","from":"0xfb337706200a55009e6bbd41e4dc164d59bc9aa2","gas":"0x17cdc","gasPrice":"0x1","hash":"0x042602a2dff77111f3e711ab7c81adbcbc9a2d87973f4afb8dc0f2856021ec74","input":"0x31a053cf000000000000000000000000fd5111db462a68cfd6df19fb110dc8e9116a90e9000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000444f55543a3841313034343144354241424535443444434138443531324646363236313039394135343741393739394536334337323238384530453742303534313444433200000000000000000000000000000000000000000000000000000000","nonce":"0x0","to":"0xe65e9d372f8cacc7b6dfcd4af6507851ed31bb44","transactionIndex":"0x0","value":"0xd6d8","v":"0x41","r":"0xbce697be8572d1543cd8c191c409cee2b4999a538e707286b5e14f7e8ff442b8","s":"0x4b8f8e8a14fb60dbe981f6ddbb31300bbc2ce8753ad6b82bdce8147280cd8e43"}],"transactionsRoot":"0xd42e9b932bffb89da313a7f9370d83c2fb4082a2d8ff162b70dcb36330a476db","uncles":[]}}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_sendRawTransaction" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_getTransactionReceipt" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{
				"transactionHash":"0x88df016429689c079f3b2f6ad39fa052532c56795b733da78a91ebe6a713944b",
				"transactionIndex":"0x0",
				"blockNumber":"0x1",
				"blockHash":"0x78bfef68fccd4507f9f4804ba5c65eb2f928ea45b3383ade88aaa720f1209cba",
				"cumulativeGasUsed":"0xc350",
				"contractAddress":"0x2a65aca4d5fc5b5c859090a6c34d164135398226",
				"gasUsed":"0x4dc",
				"logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
				"logs":[],
				"status":"0x1"
			}}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_blockNumber" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x7"}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_getBlockByNumber" {
				_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{
				"difficulty": "0x31962a3fc82b",
				"extraData": "0x4477617266506f6f6c",
				"gasLimit": "0x47c3d8",
				"gasUsed": "0x0",
				"hash": "0x78bfef68fccd4507f9f4804ba5c65eb2f928ea45b3383ade88aaa720f1209cba",
				"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
				"miner": "0x2a65aca4d5fc5b5c859090a6c34d164135398226",
				"nonce": "0xa5e8fb780cc2cd5e",
				"number": "0x1",
				"parentHash": "0x8b535592eb3192017a527bbf8e3596da86b3abea51d6257898b2ced9d3a83826",
				"receiptsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
				"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
				"size": "0x20e",
				"stateRoot": "0xdc6ed0a382e50edfedb6bd296892690eb97eb3fc88fd55088d5ea753c48253dc",
				"timestamp": "0x579f4981",
				"totalDifficulty": "0x25cff06a0d96f4bee",
				"transactions": [],
				"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
				"uncles": [
		]}}`))
				c.Assert(err, IsNil)
			}
			if rpcRequest.Method == "eth_call" {
				if string(rpcRequest.Params) == `[{"from":"0x9f4aab49a9cd8fc54dcb3701846f608a6f2c44da","input":"0x03b6a6730000000000000000000000009f4aab49a9cd8fc54dcb3701846f608a6f2c44da0000000000000000000000003b7fa4dd21c6f9ba3ca375217ead7cab9d6bf483","to":"0xe65e9d372f8cacc7b6dfcd4af6507851ed31bb44"},"latest"]` {
					_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":5,"result":"0x0000000000000000000000000000000000000000000000000000000000000012"}`))
					c.Assert(err, IsNil)
				} else if string(rpcRequest.Params) == `[{"data":"0x95d89b41","from":"0x0000000000000000000000000000000000000000","to":"0x3b7fa4dd21c6f9ba3ca375217ead7cab9d6bf483"},"latest"]` {
					_, err = rw.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000003544b4e0000000000000000000000000000000000000000000000000000000000"}`))
					c.Assert(err, IsNil)
				}
			}
		}
	}))
	s.server = server
	cfg := config.BifrostClientConfiguration{
		ChainID:         "thorchain",
		ChainHost:       server.Listener.Addr().String(),
		SignerName:      "bob",
		SignerPasswd:    "password",
		ChainHomeFolder: s.thordir,
	}

	kb := cKeys.NewInMemory()
	_, _, err := kb.NewMnemonic(cfg.SignerName, cKeys.English, cmd.THORChainHDPath, cfg.SignerPasswd, hd.Secp256k1)
	c.Assert(err, IsNil)
	s.thorKeys = thorclient.NewKeysWithKeybase(kb, cfg.SignerName, cfg.SignerPasswd)
	s.bridge, err = thorclient.NewThorchainBridge(cfg, s.m, s.thorKeys)
	c.Assert(err, IsNil)
}

func httpTestHandler(c *C, rw http.ResponseWriter, fixture string) {
	var content []byte
	var err error

	switch fixture {
	case "500":
		rw.WriteHeader(http.StatusInternalServerError)
	default:
		content, err = os.ReadFile(fixture)
		if err != nil {
			c.Fatal(err)
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	if _, err = rw.Write(content); err != nil {
		c.Fatal(err)
	}
}

func (s *MultiversXSuite) TearDownTest(c *C) {
	c.Assert(os.Unsetenv("NET"), IsNil)

	if err := os.RemoveAll(s.thordir); err != nil {
		c.Error(err)
	}
}

func (s *MultiversXSuite) TestClient(c *C) {
	pubkeyMgr, err := pubkeymanager.NewPubKeyManager(s.bridge, s.m)
	c.Assert(err, IsNil)
	// TODO fix
	//poolMgr := thorclient.NewPoolMgr(s.bridge)
	e2, err2 := NewClient()
	c.Assert(err2, IsNil)
	c.Assert(e2, NotNil)
	c.Assert(pubkeyMgr.Start(), IsNil)
	defer func() { c.Assert(pubkeyMgr.Stop(), IsNil) }()
	c.Check(e2.GetChain(), Equals, common.MVXChain)
	height, err := e2.GetHeight()
	c.Assert(err, IsNil)
	c.Check(height, Equals, int64(7))
	//TODO fix
	//gasPrice := e2.GetGasPrice()
	//c.Check(gasPrice.Uint64(), Equals, uint64(initialGasPrice))

	acct, err := e2.GetAccount(types.GetRandomPubKey(), nil)
	c.Assert(err, IsNil)
	c.Check(acct.Sequence, Equals, int64(0))
	c.Check(acct.Coins[0].Amount.Uint64(), Equals, uint64(0))
	pk := types.GetRandomPubKey()
	addr := e2.GetAddress(pk)
	c.Check(len(addr), Equals, 42)
	_, err = e2.BroadcastTx(stypes.TxOutItem{}, []byte(`{
		"from":"0xa7d9ddbe1f17865597fbd27ec712455208b6b76d",
		"gas":"0xc350",
		"gasPrice":"0x4a817c800",
		"input":"0x68656c6c6f21",
		"nonce":"0x15",
		"to":"0xf02c1c8e6114b1dbe8937a39260b5b0a374432bb",
		"transactionIndex":"0x41",
		"value":"0xf3dbb76162000",
		"v":"0x25",
		"r":"0x1b5e176d927f8e9ab405058b2d2457392da3e20f328b16ddabcebc33eaac5fea",
		"s":"0x4ba69724e8f69de52f0125ad8b3c5c2cef33019bac3249e2c0a2192766d1721c"
	}`))
	c.Assert(err, IsNil)
	input := []byte(`{
    "height": 1,
    "tx_array": [
        {
            "vault_pub_key": "tthorpub1addwnpepq2mza4j4vplyjw295pkq8j2dan627lz6vufeu22pjx5vnnyjted5vwq3e3d",
            "chain": "ETH",
			"from_address":"0xa7d9ddbe1f17865597fbd27ec712455208b6b76d",
            "to_address": "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
            "coin": {
                "asset": "ETH.ETH",
                "amount": "194765912"
            },
            "max_gas": [
                {
                    "asset": "ETH.ETH",
                    "amount": "600000"
                }
            ],
			"gas_rate":1
        }
    ]
}`)
	var txOut stypes.TxOut
	err = json.Unmarshal(input, &txOut)
	c.Assert(err, IsNil)

	// TODO fix
	//txOut.TxArray[0].VaultPubKey = e2.kw.GetPubKey()
	//c.Logf(txOut.TxArray[0].VaultPubKey.String())
	//c.Logf(e2.kw.GetPubKey().String())
	out := txOut.TxArray[0].TxOutItem(txOut.Height)
	out.Chain = common.ETHChain
	out.Memo = "OUT:B6BD1A69831B9CCC0A1E9939E9AFBFCA144C427B3F61E176EBDCB14E57981C1B"
	r, _, obs, err := e2.SignTx(out, 1)
	c.Assert(err, IsNil)
	c.Assert(r, NotNil)
	c.Assert(obs, NotNil)
	fromAddr, err := out.VaultPubKey.GetAddress(common.ETHChain)
	c.Assert(err, IsNil)
	c.Assert(obs.Sender, Equals, fromAddr.String())

	_, err = e2.BroadcastTx(out, r)
	c.Assert(err, IsNil)
}
