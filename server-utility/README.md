# Server Utility

You should have OpenWRT router connected to your server, enable Wake-on-LAN in server BIOS and configure your Tailscale network to use this script.

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
server connect [-s server_name] [-p port]
```

## Stop server

```bash
server stop [-s server_name] [-p port]
```
