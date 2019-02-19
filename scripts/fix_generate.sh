#!/bin/bash
set -xe

DIR=spec/DSP8010_2018.3/openapi/

## NOTES
## Ensure openapi-generator-cli jar is avaiable under /opt
## https://github.com/OpenAPITools/openapi-generator


#https://github.com/OpenAPITools/openapi-generator/issues/455
#for f in $(grep 'redfish.dmtf.org' $DIR -l );
# do sed -e 's@http://redfish.dmtf.org/schemas/v1/@@g' -i $f; 
#done

#fix Could not find definitions/GenerateCSRResponse in contents of ./CertificateService.v1_0_0.yaml
sed -e 's@definitions/GenerateCSRResponse@components/schemas/GenerateCSRResponse@g' -i $DIR/openapi.yaml

java -jar /opt/openapi-generator-cli-3.3.4-patched.jar generate -i $DIR/openapi.yaml -g go -o client 
