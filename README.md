# HEN encyption service and cli

### Overview

web encryption service to provide a way to store encypted secrets in config files in the repo.

![](https://i.pinimg.com/originals/ce/fa/41/cefa416338834d90cc7e8f5b41a3ac78.png)

### Features

  - allows configs checked into repo kind to provide KISS approch to secrets

### Diagram 

TODO


## goals and future features

  - check in secure config values
  - have a ui html form to generate egg(s) from string
  - allow developers with token ability to check in a one way secure values
  - allow admin and gitlab ci decrypt aka "hatch eggs"
  - make it easy to update config values
  - make it easy to change configs
  - make code integrations
  - if integrated into key, enviroment configs could be built into container image at `/config/foo.(dev|stage|prod).json`
  - configs could be processed in ci and pushed as secrets to k8s and our container could read the secret
  - future allow cli admin edit vi mode with temp file and `egg(...)` replaced with `cracked-egg(string)`
  - settings in one location the repo vs gitlab var, terraform, parameter store
  - settings not dependent on infra setup 
  - future for auth / token
  - future would like to use jwt for tokens
  - proposed server url https://hen.gomonger.io
  - tokens database, use existing jwt token, maybe use token service?
  - maybe values files or direct config edits
  - md5 check sum, why to validate x sent = x encrypted
  - simplify config
  - Can still use if moving away from json
  - could build into kit and env vars
  - part of ci process could remove parameter store
  - could use k8s secrets to store json
  - can track settings CI changes
  - CON: doesn't allow for dynamic changes without CI
  - configs/settings files can put shipped with image kit 
  - could call hen service to decrypt for hen/egg('values')

## TODO
  - read up on jwt and how it works
  - read up on tokens, token a service
  - read up on mysql storage 
  - setup a service like go-gin

## server 

```bash
export ENCRYPTION_KEY="32-bit-string" 
./hen server
```

## client

```bash
./hen protect -m "hello world"
{"egg":"egg(97b2fe93cccdfbb794ca2a5131ee7afe9bac07758797705df2008234b7aa9966652a4f802e9d0c,5eb63bbbe01eeed093cb22bb8f5acdc3)"}
```

### Decrypt curl example with Token

```bash
curl -X POST http://localhost:3000/hatch \
     -H 'Content-Type: application/json' \
     -H 'Token: foobar' \
     -d '{"egg":"231d4adc3e9f6a62640e49b882747a376a5a5127401b1d1cdc05c7e2a39ece1ef63e8c58"}'                                      
# {"secret":"test abc"}%       
```



## Encryption Ref


https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
