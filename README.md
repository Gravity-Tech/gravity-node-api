
# Gravity Node API

### Overview

Current http server implementation provides two distinct routes for mockup data of **Gravity Nebula** and **Gravity Node** models.

### Params

| Parameter | Default | Usage 
|---------|-------|-----|
| `--port` | `8090` | `Port for server to run on`

```
> go build -o mockup-server
> ./mockup-server --port 8094
```

### Routes


	GetCommonStats        = "/"
	GetNodeRewards        = ""
	GetNodeActionsHistory = ""

| Route | Method | Response Description
|-------|-------|-----|
| `/nebulas/all` | `GET` | `Nebulas list`
| `/nodes/all` | `GET` | `Nodes list`
| `/common/stats` | `GET` | `Common gravity node stats`
| `/nodes/rewards/all` | `GET` | `Node rewards`
| `/nodes/actions/history` | `GET` | `Nodes actions history`


