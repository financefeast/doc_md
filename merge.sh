
# build
cd /var/tmp/source
wget https://github.com/gohugoio/hugo/releases/download/v0.80.0/hugo_0.80.0_Linux\-64bit.tar.gz
tar -zxvf hugo_0.80.0_Linux-64bit.tar.gz

# compile
echo compile
./hugo --minify --enableGitInfo

# copy public
cp -r ./public/* /usr/share/nginx/html/


