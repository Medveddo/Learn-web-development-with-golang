#!/bin/bash
cd "$GOPATH/src/learn-web-dev-with-go"

echo "==== Releasing lenslocked.com ===="
echo " Deleting the local binary if it exists (so it isn't uploaded)..."
rm learn-web-dev-with-go
echo " Done! "

echo " Deleting existing code..."
ssh root@138.68.132.2 "rm -rf /root/go/src/learn-web-dev-with-go"
echo " Code deleted successfully!"

echo " Uploading code..."
rsync -avr --exclude '.git/*' --exclude 'tmp/*' \
--exclude 'images/*' ./ \
root@138.68.132.2:/root/go/src/learn-web-dev-with-go/
echo " Code uploaded successfully!"

echo " Go getting deps..."
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get golang.org/x/crypto/bcrypt"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get github.com/gorilla/mux"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get github.com/gorilla/schema"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get github.com/lib/pq"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get github.com/jinzhu/gorm"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get github.com/gorilla/csrf"
ssh root@138.68.132.2 "export GOPATH=/root/go; \
/usr/local/go/bin/go get gopkg.in/mailgun/mailgun-go.v1"

echo " Building the code on remote server..."
ssh root@138.68.132.2 'export GOPATH=/root/go; \
cd /root/app; \
/usr/local/go/bin/go build -o ./server \
$GOPATH/src/learn-web-dev-with-go/*.go'
echo " Code built successfully!"

echo " Moving assets..."
ssh root@138.68.132.2 "cd /root/app; \
cp -R /root/go/src/learn-web-dev-with-go/assets ."
echo " Assets moved successfully!"

echo " Moving views..."
ssh root@138.68.132.2 "cd /root/app; \
cp -R /root/go/src/learn-web-dev-with-go/views ."
echo " Views moved successfully!"

echo " Moving Caddyfile..."
ssh root@138.68.132.2 "cp /root/go/src/learn-web-dev-with-go/Caddyfile /etc/caddy/Caddyfile"
echo " Caddyfile moved successfully!"

echo " Restarting the server..."
ssh root@138.68.132.2 "sudo service mywebapp restart"
echo " Server restarted successfully!"

echo " Restarting Caddy server..."
ssh root@138.68.132.2 "sudo service caddy restart"
echo " Caddy restarted successfully!"

echo "==== Done releasing lenslocked.com ===="
