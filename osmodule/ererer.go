#!/bin/bash

HOME_DIR="/root/scripts/cloudflare";

CLOUDFLARE_IPS_V4=$(/usr/bin/curl -s --max-time 10 https://www.cloudflare.com/ips-v4)
CLOUDFLARE_IPS_V6=$(/usr/bin/curl -s --max-time 10 https://www.cloudflare.com/ips-v6)

echo "$CLOUDFLARE_IPS_V4" > $HOME_DIR/CLOUDFLARE_IPS.txt
echo "$CLOUDFLARE_IPS_V6" >> $HOME_DIR/CLOUDFLARE_IPS.txt

if diff -B $HOME_DIR/CLOUDFLARE_IPS_TMP.txt $HOME_DIR/CLOUDFLARE_IPS.txt >/dev/null ; then
echo ""
else
if [ -n "$CLOUDFLARE_IPS_V4" ] && [ -n "$CLOUDFLARE_IPS_V6" ]; then
/usr/sbin/ufw --force reset
/usr/sbin/ufw default allow incoming
/usr/sbin/ufw default allow outgoing
for IP in $CLOUDFLARE_IPS_V4; do
/usr/sbin/ufw allow from $IP to any port 80
/usr/sbin/ufw allow from $IP to any port 443
done
for IP in $CLOUDFLARE_IPS_V6; do
/usr/sbin/ufw allow from $IP to any port 80
/usr/sbin/ufw allow from $IP to any port 443
done
###Open_Service
/usr/sbin/ufw allow from 44.28.149.158/32 to any port 80
/usr/sbin/ufw allow from 44.28.149.158/32 to any port 443

/usr/sbin/ufw deny 80/tcp
/usr/sbin/ufw deny 443/tcp
/usr/sbin/ufw --force enable
fi
fi
