
echo "Starting word-to-abbreviation backend"
/usr/word-to-abbreviation &

echo "Starting nginx webserver"
nginx -g "daemon off;"