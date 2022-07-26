# musizticle
Part 1 of a 2-piece music streaming app, designed to be self-hosted locally with media streamed through a related client app 
found here - https://github.com/Admiral-Piett/musizticle-gui.

test

## Environment Variables
Set these in the environment before start up or be prepared to deal with the defaults.
- MUSIZTICLE_PORT
- MUSIZTICLE_SQLITE_DB
- MUSIZTICLE_TOKEN_EXPIRATION
- MUSIZTICLE_TOKEN_KEY_LENGTH

#### Optional ENV Vars
- LOG_LEVEL - eg. `DEBUG`|`INFO`|`ERROR`

## Installation Instructions
These steps will pull and run the latest version of the docker image for this app.  The `runner.sh` script 
(included in here) is an example of where to start and has some niceties built in, such as container death and 
clean up on exit.  Feel free to change at your own peril.

#### Pre-Req
- Install docker on the machine you intend to be your "server"
  - I used a RaspberryPi 4 with Centos on it to start, but you could do whatever.  Even cloud if you wanted.  Benefits
of containers.

#### Deployment
- Copy `.env` to server and set all your environment variables (listed above) into it
- Copy `./runner.sh` to server
  - Update to a different bash shell if you need to
- OPTIONAL: copy `./update.sh`
- Mount the directory of your music files to `./music` in the same directory as `runner.sh`
- Run `./update.sh` and then `./runner.sh`

```shell
mkdir musizticle
cd musizticle
cp /path/to/runner.sh /destination/musizticle/runner.sh
cp /path/to/env /destination/musizticle/env
./runner.sh
```
