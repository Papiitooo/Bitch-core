## BitchCoin-core

Official Golang implementation of the bitch protocol.


## Building BitcCoin-core from the source

```shell
 git clone https://github.com/Papiitoo/BitchCoin-core.git
```

Building `bitch` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make bitch
```

or, to build the full suite of utilities:

```shell
make all
```

## Executables

The go-bitch project comes with several wrappers/executables found in the `cmd`
directory.

|    Command    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| :-----------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  **`bitch`**   | Our main bitch CLI client. It is the entry point into the bitch network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the bitch network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `bitch --help` and the [CLI page](https://bitch.bitch.org/docs/interface/command-line-options) for command line options.          |
|   `clef`    | Stand-alone signing tool, which can be used as a backend signer for `bitch`.  |


## Running `bitch`



### Full node on the main Bitch network

By far the most common scenario is people wanting to simply interact with the bitch
network: create accounts; transfer funds; deploy and interact with contracts. For this
particular use-case the user doesn't care about years-old historical data, so we can
fast-sync quickly to the current state of the network. To do so:

```shell
$ bitch console
```

This command will:
 * Start `bitch` in fast sync mode (default, can be changed with the `--syncmode` flag),
   causing it to download more data in exchange for avoiding processing the entire history
   of the bitch network, which is very CPU intensive.
 * Start up `bitch`'s built-in interactive [JavaScript console],
   (via the trailing `console` subcommand) through which you can interact using [`web3` methods](https://web3js.readthedocs.io/en/) 
   (note: the `web3` version bundled within `bitch` is very old, and not up to date with official docs),
   as well as `bitch`'s own [management APIs].
   This tool is optional and if you leave it out you can always attach to an already running
   `bitch` instance with `bitch attach`.



### Configuration

As an alternative to passing the numerous flags to the `bitch` binary, you can also pass a
configuration file via:

```shell
$ bitch --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ bitch --your-favourite-flags dumpconfig
```

*Note: This works only with `bitch` v1.6.0 and above.*

#### Docker quick start

One of the quickest ways to get bitch up and running on your machine is by using
Docker:

```shell
docker run -d --name bitch-node -v /Users/alice/bitch:/root \
           -p 8545:8545 -p 30303:30303 \
           bitch/client-go
```

This will start `bitch` in fast-sync mode with a DB memory allowance of 1GB just as the
above command does.  It will also create a persistent volume in your home directory for
saving your blockchain as well as map the default ports. There is also an `alpine` tag
available for a slim version of the image.

Do not forget `--http.addr 0.0.0.0`, if you want to access RPC from other containers
and/or hosts. By default, `bitch` binds to the local interface and RPC endpoints is not
accessible from the outside.

### Programmatically interfacing `bitch` nodes


The IPC interface is enabled by default and exposes all the APIs supported by `bitch`,
whereas the HTTP and WS interfaces need to manually be enabled and only expose a
subset of APIs due to security reasons. These can be turned on/off and configured as
you'd expect.

HTTP based JSON-RPC API options:

  * `--http` Enable the HTTP-RPC server
  * `--http.addr` HTTP-RPC server listening interface (default: `localhost`)
  * `--http.port` HTTP-RPC server listening port (default: `8545`)
  * `--http.api` API's offered over the HTTP-RPC interface (default: `eth,net,web3`)
  * `--http.corsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--ws.addr` WS-RPC server listening interface (default: `localhost`)
  * `--ws.port` WS-RPC server listening port (default: `8546`)
  * `--ws.api` API's offered over the WS-RPC interface (default: `eth,net,web3`)
  * `--ws.origins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: `admin,debug,eth,miner,net,personal,shh,txpool,web3`)
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to
connect via HTTP, WS or IPC to a `bitch` node configured with the above flags and you'll
need to speak [JSON-RPC](https://www.jsonrpc.org/specification) on all transports. You
can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based
transport before doing so! Hackers on the internet are actively trying to subvert
bitch nodes with exposed APIs! Further, all browser tabs can access locally
running web servers, so malicious web pages could try to subvert locally available
APIs!**



## License

The go-bitch library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-bitch binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
