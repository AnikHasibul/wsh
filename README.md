# wsh
A shell for ignoring Android linker warning on termux!

### Installation

```
apt install golang git
git clone https://github.com/anikhasibul/wsh
cd wsh
go build
mv wsh $PREFIX/bin/
```

### Usage

```sh
wsh
```

Then you will be prompted to choose your working shell! Enter your prefered shell name! in my case it's `bash`

Done!

Now you can run `wsh` and it will cut off all of your `Warning: linker: *` messages!

### Using on startup!

For doing this trick you need to install another login shell. ex: `fish` or `zsh`

```
apt install zsh
```

In my case I've used `zsh` as a parent shell!

Then add this line to your `~/.zshrc` 

```sh
exec wsh
```

Now run 

```sh
chsh -s zsh
```

voila!

you've been bypassed `WARNING: linker: Unsupported flags DT_FLAGS_1=0x8`
