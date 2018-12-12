# vdothu

VdoThu - Video cache & Thumb service and utilities

vdothu:
  - -a string
     * endpoint of the gRPC service (default ":0")
  - -h, --help
     * show help
  - -n string
     * a valid network type which is consistent to -a (default "tcp")
  - -v value
     * log level for V logs
  - -z
     * compress with gzip

client:
  - -a string
     * endpoint of the gRPC service
  - -c string
     * specify transport credentials (TLS) from file
  - -h, --help
     * show help
  - -H string
     * host name or IP address (default "localhost")
  - -v value
     * log level for V logs
  - -z
     * compress with gzip

proxy:
  - -e string
     * endpoint of the gRPC service to proxy (default ":0")
  - -h, --help
     * show help
  - -p string
     * port of the proxy service (default ":0")
  - -v value
     * log level for V logs

direct:
  - $HOST
     * host of the web service (default "")
  - $PORT
     * port of the web service (default "")

### Download and Install

#### Use Binary Distributions

##### common

1.Install docker on your operating system

2.Execute command
  - `docker pull dlot/$R:latest`

3.Organize your commands & options to replace the following "$@"

4.Execute command
  - `docker run --rm dlot/$R:latest /app/vdolot/$EXE "$@"`

5.(Optional) Save the above script as an executable file EXE for executing `./$EXE $CMDs $OPTs` when needed

##### vdothu

1.Follow #common

2.If treating above the above variables as environment variables, then
  - R=vdolot_rpc
  - EXE=vdothu

##### client

1.Follow #common

2.If treating above the above variables as environment variables, then
  - R=vdolot_cli
  - EXE=client

##### proxy

1.Follow #common

2.If treating above the above variables as environment variables, then
  - R=vdolot_wgw
  - EXE=proxy

##### direct

1.Follow #common

2.If treating above the above variables as environment variables, then
  - R=vdolot_web
  - EXE=direct

#### Install From Source

1.Select your working directory (as environment variable D)

2.Execute command
  - `git clone $GIT_HOST/$USR_NAME/$PRO_NAME.git "$D/$PRO_NAME"`

3.Follow your Go knowledge & your wants...

### Advanced usage

#### Serve RPC service

Execute command `docker run --rm dlot/vdolot_rpc:latest /app/vdolot/vdothu "$@"`

#### Serve RESTful service (by RPC proxy)

Execute command `docker run --rm dlot/vdolot_wgw:latest /app/vdolot/proxy "$@"`

#### Serve RESTful service (by web server)

Execute command `docker run --rm dlot/vdolot_web:latest /app/vdolot/direct "$@"`

#### Serve RESTful service (by ngnix proxy)

Configure your nginx.conf

### Contributing

To contribute, please pull requests of your contributing code to above.

## Let your hot heart go!