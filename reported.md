## reported races

使用GRace在github的若干项目上进行了数据竞争检测，经过过滤后将部分数据竞争检测结果报告给了开发者，这些在下面列出。

### act

项目地址：https://github.com/nektos/act

commit id: 816a7d

检测结果：

```
Potential race condition:
 Access1:
        if err = fn(ctx); err != nil {
           ^
act/cmd/root.go:663:3 ->
act/cmd/root.go:651:7

 Access2:
        return err
               ^
act/cmd/root.go:666:9
===================================================
Potential race condition:
 Access1:
        if err = fn(ctx); err != nil {
           ^
act/cmd/root.go:663:3 ->
act/cmd/root.go:657:8

 Access2:
        return err
               ^
act/cmd/root.go:666:9
```

性质：已确认，已修复。

### AdGuard

项目地址：https://github.com/AdguardTeam/AdGuardHome

commit id: d6043e2

```
===================================================
Potential race condition:
 Access1:
        if clients.arpdb != nil {
                   ^
AdGuardHome/internal/home/home.go:134:3 ->
AdGuardHome/internal/home/home.go:126:30 ->
AdGuardHome/internal/home/clients.go:141:13
 
 Access2:
        clients.arpdb = arpdb
                ^
AdGuardHome/internal/home/home.go:143:5 ->
AdGuardHome/internal/home/home.go:461:19 ->
AdGuardHome/internal/home/home.go:341:22 ->
AdGuardHome/internal/home/clients.go:93:10
===================================================
Potential race condition:
 Access1:
        return Context.dnsServer != nil && Context.dnsServer.IsRunning()
                       ^
AdGuardHome/internal/home/home.go:134:3 ->
AdGuardHome/internal/home/home.go:129:12 ->
AdGuardHome/internal/home/home.go:726:22 ->
AdGuardHome/internal/home/dns.go:494:15 ->
AdGuardHome/internal/home/dns.go:206:17
 
 Access2:
        Context.dnsServer, err = dnsforward.NewServer(p)
                ^
AdGuardHome/internal/home/home.go:143:5 ->
AdGuardHome/internal/home/home.go:466:15 ->
AdGuardHome/internal/home/home.go:929:22 ->
AdGuardHome/internal/home/dns.go:145:10
===================================================
Potential race condition:
 Access1:
        tlsConf.PrivateKeyData = []byte(tlsConf.PrivateKey)
                ^
AdGuardHome/internal/home/home.go:134:3 ->
AdGuardHome/internal/home/home.go:127:23 ->
AdGuardHome/internal/home/tls.go:140:14 ->
AdGuardHome/internal/home/tls.go:68:19 ->
AdGuardHome/internal/home/tls.go:176:10
 
 Access2:
        tlsConf.PrivateKeyData = []byte(tlsConf.PrivateKey)
                ^
AdGuardHome/internal/home/home.go:143:5 ->
AdGuardHome/internal/home/home.go:507:34 ->
AdGuardHome/internal/home/tls.go:53:15 ->
AdGuardHome/internal/home/tls.go:68:19 ->
AdGuardHome/internal/home/tls.go:176:10
===================================================
Potential race condition:
 Access1:
        _, err = con6.WriteTo(data, msg, addr)
           ^
AdGuardHome/internal/dhcpd/routeradv.go:270:3 ->
AdGuardHome/internal/dhcpd/routeradv.go:263:7
 
 Access2:
        return nil
        ^
AdGuardHome/internal/dhcpd/routeradv.go:273:2 
```

性质：正在确认中

### croc

项目地址：https://github.com/schollz/croc

commit id: 3c1d60d

```
Potential race condition:
 Access1:
        if err == nil {
           ^
croc/src/croc/croc.go:635:4 ->
croc/src/croc/croc.go:573:8

 Access2:
        err = <-errchan
        ^
croc/src/croc/croc.go:638:2
```

性质：已确认，未修复

### gitea

项目地址：https://github.com/go-gitea/gitea

commit id: dad057b

```
===================================================
Potential race condition:
 Access1:
        comment.TreePath = util.CleanPath(comment.TreePath)
                ^
gitea/services/migrations/migrate.go:452:37 ->
gitea/services/migrations/gitea_uploader.go:869:12
 
 Access2:
        log.Warn("GetRepoRawDiffForFile failed when migrating [%s, %s, %s, %s]: %v", g.gitRepo.Path, pr.MergeBase, headCommitID, comment.TreePath, err)
                                                                                                                                         ^
gitea/services/migrations/migrate.go:445:38 ->
gitea/services/migrations/gitea_uploader.go:883:5 ->
gitea/services/migrations/gitea_uploader.go:880:135 
```

性质：已确认，已修复

```
===================================================
Potential race condition:
 Access1:
        err = SanitizeReader(pr2, renderer.Name(), output)
        ^
gitea/modules/markup/renderer.go:239:4 ->
gitea/modules/markup/renderer.go:236:4

 Access2:
        err = PostProcess(ctx, pr, pw2)
        ^
gitea/modules/markup/renderer.go:254:3 ->
gitea/modules/markup/renderer.go:247:4
```

性质：已确认，未修复

### gost

项目地址：https://github.com/ginuerzh/gost

commit id: 1c62376

```
===================================================
Potential race condition:
 Access1:
        clientAddr = addr
        ^
gost/socks.go:1357:3 ->
gost/socks.go:1341:5

 Access2:
        if clientAddr == nil {
           ^
gost/socks.go:1389:3 ->
gost/socks.go:1369:7
===================================================
Potential race condition:
 Access1:
        clientAddr = addr
        ^
gost/ss.go:375:3 ->
gost/ss.go:368:5 ->
gost/ss.go:351:6
 
 Access2:
        if clientAddr == nil {
           ^
gost/ss.go:409:3 ->
gost/ss.go:402:5 ->
gost/ss.go:387:8
```

性质：未回复

### photoprism 

项目地址：https://github.com/photoprism/photoprism

commit id: 0d00dc

```
===================================================
Potential race condition:
 Access1:
        *c.r = append(*c.r, i)
           ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:99:23 ->
photoprism/pkg/clusters/dbscan.go:235:21 ->
photoprism/pkg/clusters/dbscan.go:253:8
 
 Access2:
        c.r = r
          ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:101:7 ->
photoprism/pkg/clusters/dbscan.go:163:12 ->
photoprism/pkg/clusters/dbscan.go:207:4 
===================================================
Potential race condition:
 Access1:
        *c.r = append(*c.r, i)
                         ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:99:23 ->
photoprism/pkg/clusters/dbscan.go:235:21 ->
photoprism/pkg/clusters/dbscan.go:253:22
 
 Access2:
        c.r = r
          ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:101:7 ->
photoprism/pkg/clusters/dbscan.go:163:12 ->
photoprism/pkg/clusters/dbscan.go:207:4 
===================================================
Potential race condition:
 Access1:
        c.p = c.d[p]
          ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:101:7 ->
photoprism/pkg/clusters/dbscan.go:163:12 ->
photoprism/pkg/clusters/dbscan.go:206:4
 
 Access2:
        if c.distance(c.p, c.d[i]) < c.eps {
                        ^
photoprism/internal/photoprism/faces_cluster.go:45:26 ->
photoprism/pkg/clusters/dbscan.go:99:23 ->
photoprism/pkg/clusters/dbscan.go:235:21 ->
photoprism/pkg/clusters/dbscan.go:251:20 
```

性质：未回复

### rclone

项目地址：https://github.com/rclone/rclone

commit id: 927e72

```
===================================================
Potential race condition:
 Access1:
        abort = true
        ^
rclone/vfs/read.go:240:3 ->
rclone/vfs/read.go:234:4
 
 Access2:
        for *poff != off && !abort {
                             ^
rclone/vfs/read.go:241:23
===================================================
Potential race condition:
 Access1:
        err = s.serveHTTP()
        ^
rclone/cmd/serve/dlna/dlna.go:310:3 ->
rclone/cmd/serve/dlna/dlna.go:306:3
 
 Access2:
        return nil
        ^
rclone/cmd/serve/dlna/dlna.go:312:2 
===================================================
Potential race condition:
 Access1:
        if err != nil {
           ^
rclone/cmd/serve/dlna/dlna.go:310:3 ->
rclone/cmd/serve/dlna/dlna.go:307:6
 
 Access2:
        return nil
        ^
rclone/cmd/serve/dlna/dlna.go:312:2 
===================================================
Potential race condition:
 Access1:
        fs.Logf(s.f, "Error on serving HTTP server: %v", err)
                                                         ^
rclone/cmd/serve/dlna/dlna.go:310:3 ->
rclone/cmd/serve/dlna/dlna.go:308:53
 
 Access2:
        return nil
        ^
rclone/cmd/serve/dlna/dlna.go:312:2 
```

性质：已确认，已修复
