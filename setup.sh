function gencode() {
   echo $RANDOM | md5sum | head -c 10
}

function RandomFlag() {
  echo "ADMIN_COOKIE_IS_`gencode`" > keys/admin_cookie.txt
  echo "DASCTF{wow,_`gencode`_YOU_KN0w_Th3_Cl13nt,S1d3_T3mpl4t3_1njection}" > keys/flag.txt
}

docker compose up -d

