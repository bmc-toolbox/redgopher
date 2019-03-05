#### Method compatibility between BMCs

| Dell iDRAC9         |  HP iLO5            | Method          |
|  :-----------:      |  :-----:            | :--------------:| 
|  404                |  :heavy_check_mark: | RedfishV1SessionServiceSessionsGet |
|  :heavy_check_mark: |  :heavy_check_mark: | RedfishV1ManagersManagerIdNetworkProtocolGet |

- Dell iDRAC9 version 3.21.26.22
- HP iLO5     version 1.40

#### Generate go client from redfish spec

1. Retrieve redfish schema.

`wget https://www.dmtf.org/sites/default/files/standards/documents/DSP8010_2018.3.zip  -O ./spec/DSP8010_2018.3.zip`

- Build openapi-generator 

We need to build openapi-generator and apply patches for https://github.com/OpenAPITools/openapi-generator/issues/535
which are found at https://github.com/OpenAPITools/openapi-generator/pull/542/commits/5a6d4fc1f844181fa113c18c3cf8d2b720f811e8

```
git clone https://github.com/OpenAPITools/openapi-generator.git
cd openapi-generator/
git checkout v3.3.4
git fetch origin pull/542/head:go-fix
./mvnw clean install
cp ./modules/openapi-generator-cli/target/openapi-generator-cli.jar /opt/openapi-generator-cli-3.3.4-patched.jar
```

- Run `./scripts/fix_generate.sh`
