# Server Utility

To use this script, you should have:
- OpenWRT router connected to your server via LAN
- Wake-on-LAN enabled in server BIOS 
- Configured Tailscale network
- SSH connection to your server

Also you can add this script to your PATH in your .zshrc or .bashrc:
```sh
export PATH=$PATH:<path_to_this_script>
```

## Create .env file

Copy .env example:
```bash
cp .env.example .env
```
Then change values in .env file to your own data.

## Run Server

```bash
server start
```

## Connect to server

```bash
server connect [-s server_name] [-u user] [-p port]
```

## Stop server

```bash
server stop [-s server_name] [-u user] [-p port]
```
