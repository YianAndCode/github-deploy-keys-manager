# GitHub Deploy Keys Manager

[English](README.md) | 中文

## 关于

众所周知，GitHub 不允许将同一个部署密钥用在不同的代码仓库中，所以当你在一台服务器上用部署密钥来部署多个项目时，你需要为每个项目生成不同的密钥，`GitHub Deploy Keys Manager`（下称 `gdkm`）就可以帮你轻松管理部署密钥。

[关于部署密钥的 GitHub 官方文档](https://docs.github.com/cn/developers/overview/managing-deploy-keys#%E5%9C%A8%E4%B8%80%E5%8F%B0%E6%9C%8D%E5%8A%A1%E5%99%A8%E4%B8%8A%E4%BD%BF%E7%94%A8%E5%A4%9A%E4%B8%AA%E4%BB%93%E5%BA%93)

## 使用方法

### 安装

从源码编译安装

```bash
git clone https://github.com/YianAndCode/github-deploy-keys-manager.git
cd github-deploy-keys-manager
make
```

### 基本用法

```bash
./bin/gdkm -repo={YOUR_REPO_URL}
# 例如 ./bin/gdkm -repo=git@github.com:YianAndCode/github-deploy-keys-manager.git
```

当你执行完上面的命令，一个新生成的密钥对会被保存在 `$HOME/.ssh/deploy/`。

你也可以用 `-key-path=` 来指定密钥对的存储路径：
```bash
./bin/gdkm -repo={YOUR_REPO_URL} -key-path=/path/to/save/key
```

## TODO:

 - [x] 自动更新 ssh_config
