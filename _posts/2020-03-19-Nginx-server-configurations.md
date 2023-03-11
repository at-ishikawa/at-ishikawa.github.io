---
date: 2020-03-19
title: Nginx server configurations
---

Performance
===

proxy_cache
---
The below configuration is the example to cache images.
```
proxy_cache_path /path/to/cache levels=1:2 keys_zone=my_cache:10m max_size=10g
                 inactive=60m use_temp_path=off;

server {
    # ...
    location / {
        proxy_cache my_cache;
        proxy_pass http://my_upstream;
    }

    location ~ .*\.(jpg|JPG|gif|GIF|png|PNG|swf|SWF|css|CSS|js|JS|inc|INC|ico|ICO) {
        # Server proxy cache
        proxy_cache my_cache;
        proxy_ignore_headers Cache-Control;
        proxy_cache_valid any 30m;

        # Client cache
        expires 30d;
        ...
    }
}
```

- `proxy_ignore_headers` ignores Cache-Control header from a client. In this case, `proxy_cache_valid` is required.
- `expires` sends `Cache-Control: max-age=\d+` and `Expires` header in order to cache on browsers.

References
- [Server cache](https://www.nginx.com/blog/nginx-caching-guide/)
- [Client cache](https://www.howtoforge.com/make-browsers-cache-static-files-on-nginx)

gzip compression
---
The example of gzip compression is followings.
```
http {
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_proxied any;
    gzip_types text/css text/javascript application/javascript application/json application/font-woff application/font-tff image/jpeg image/gif image/png application/octet-stream;
    gzip_disable "MSIE [1-6]\.";

    # gzip_static
    gzip_static on;
}
```
- `gzip_proxied`: gzip works even if nginx works as a reverse proxy server
- `gzip_static`: use a precompressed file if it exists.


References
- [gzip configuration](http://www.techrepublic.com/article/how-to-configure-gzip-compression-with-nginx/)
- [gzip_static module document](http://nginx.org/en/docs/http/ngx_http_gzip_static_module.html)


Connections
---
The example for configurations related with connections.
```
events {
    use epoll;
    multi_accept on;
}
http {
    # copies data between one FD and other from within the kernel
    # faster then read() + write()
    sendfile on;
    # send headers in one peace, its better then sending them one by one
    tcp_nopush on;
    # don't buffer data sent, good for small data bursts in real time
    tcp_nodelay on;

    # server will close connection after this time -- default 75
    keepalive_timeout 30;

    # number of requests client can make over keep-alive -- for testing environment
    keepalive_requests 100000;
}
```

- `use epoll`: use multiplex I/O. better than select/poll.

References
- [gist for NGINX Tuning For Best Performance](https://gist.github.com/denji/8359866)


File cache
---
The example for file cache.
```
http {
    # cache informations about FDs, frequently accessed files
    # can boost performance, but you need to test those values
    open_file_cache max=200000 inactive=20s;
    open_file_cache_valid 30s;
    open_file_cache_min_uses 2;
    open_file_cache_errors on;
}
```

References
- [gist for NGINX Tuning For Best Performance](https://gist.github.com/denji/8359866)
