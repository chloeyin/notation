# Signing in cose envelope

This is a prototype of notation which supports signing in cose format.It is forked from the origin repository

Currenly,  we will sign the cose by using [notation plugin](https://github.com/notaryproject/notaryproject/blob/main/specs/plugin-extensibility.md)

Notice: Now I disable pushing/pulling signature by default for convenience.

## Install And Run


```sh

# install cose
mkdir -p ~/bin
curl -Lo ~/bin/notation https://github.com/chloeyin/notation/releases/download/cose-linux-amd64/notation && chmod 755 ~/bin/notation
# install plugin
cose_dir=~/.config/notation/plugins/cose
mkdir -p "$cose_dir"
curl -Lo ~/.config/notation/plugins/cose/notation-cose https://github.com/chloeyin/notation-cose/releases/download/plugin-linux-amd64/notation-cose && chmod 755 $cose_dir/notation-cose
notation plugin list

# this is a tool used for parsing cose envelope 
curl -Lo ~/bin/cq https://github.com/chloeyin/notation-cose/releases/download/plugin-linux-amd64/cq && chmod 755 ~/bin/cq
# this is a tool for generating generate self signed certificate
curl -Lo ~/bin/certificate-generator https://github.com/chloeyin/notation-cose/releases/download/plugin-linux-amd64/certificate-generator && chmod 755 ~/bin/certificate-generator

# create a self signed certificate
certificate-generator --host "some test host"

KEY_NAME="my-cose-key"
KEY_PATH=~/key.pem
CERT_PATH=~/certificate.pem
KEY_ID="$KEY_PATH:$CERT_PATH"
IMAGE=docker.io/library/nginx

notation key add --name ${KEY_NAME} --plugin cose --id ${KEY_ID}
notation key list

notation cert add --name ${KEY_NAME} ${CERT_PATH}
notation cert list

notation sign --key ${KEY_NAME} --signatureFormat cose $IMAGE

notation verify --cert ${KEY_NAME} $IMAGE

# view your signature
notation sign --key ${KEY_NAME} --signatureFormat cose --output /tmp/cose.sig $IMAGE
cat /tmp/cose.sig | cq
notation verify --cert ${KEY_NAME} --signature /tmp/cose.sig $IMAGE
```

