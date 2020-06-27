# my_fin

# Dep update
``
go mod vendor
``

# Nginx config

    limit_req_zone $binary_remote_addr zone=antibot:16m rate=6r/m;
    limit_req_log_level warn;
    limit_req_status 403;
    

    #server path
    location = /api/auth/login {
        limit_req zone=antibot burst=2 nodelay;
    }
    
    if ($request_method !~ ^(GET|HEAD|POST)$ ) {
        return 444;
    }
