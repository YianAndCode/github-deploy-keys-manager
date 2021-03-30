# GitHub Deploy Keys Manager

## What is

When you use deploy key to deploy multiple repositories on one server, you will need to generate a dedicated key pair for each one (GitHub does not allow you to reuse a deploy key for multiple repositories), `GitHub Deploy Keys Manager` (I call it `gdkm` below) can help you manage your deploy keys easily.

[GitHub Docs about Deploy keys](https://docs.github.com/en/developers/overview/managing-deploy-keys#using-multiple-repositories-on-one-server)

## Usage

### Install

Build from source

```bash
git clone https://github.com/YianAndCode/github-deploy-keys-manager.git
cd github-deploy-keys-manager
make
```

### Basic use

```bash
./bin/gdkm -repo={YOUR_REPO_URL}
# e.g. ./bin/gdkm -repo=git@github.com:YianAndCode/github-deploy-keys-manager.git
```

After you excuting the command above, the key pair will be generated and stored in `$HOME/.ssh/deploy/` by default.

If you want to specify key path, you can pass the path through `-key-path=`:
```bash
./bin/gdkm -repo={YOUR_REPO_URL} -key-path=/path/to/save/key
```

## TODO:

 - Modify ssh_config file automatically
