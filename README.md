
simple implementation of raft algorithm

---

### TODO
- [ ] Leader election
- [x] Log replication

### Branch
- master： Leader election & Log replication
- election：Leader election
- replication： Log replication

### build
```
go build -o simple-raft
```

### run
```
go get go get github.com/mattn/goreman
goreman start
```

