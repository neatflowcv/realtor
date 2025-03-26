#!/bin/bash

curl \
'https://apis.zigbang.com/apt/locals/11680106/item-catalogs?tranTypeIn%5B0%5D=trade&tranTypeIn%5B1%5D=charter&tranTypeIn%5B2%5D=rental&includeOfferItem=true&offset=0&limit=10' --compressed \
-H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:136.0) Gecko/20100101 Firefox/136.0' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Accept-Language: en-US,en;q=0.5' \
-H 'Accept-Encoding: gzip, deflate, br, zstd' \
-H 'sdk-version: 0.87.0' \
-H 'X-Zigbang-Platform: www' \
-H 'Origin: https://www.zigbang.com' \
-H 'DNT: 1' \
-H 'Connection: keep-alive' \
-H 'Referer: https://www.zigbang.com/' \
-H 'Sec-Fetch-Dest: empty' \
-H 'Sec-Fetch-Mode: cors' \
-H 'Sec-Fetch-Site: same-site' \
-H 'Pragma: no-cache' \
-H 'Cache-Control: no-cache' \
-H 'TE: trailers'