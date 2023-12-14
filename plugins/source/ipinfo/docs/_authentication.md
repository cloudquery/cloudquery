CloudQuery requires only `IP` and `TOKEN`. Follow [this guide](https://ipinfo.io/faq/article/127-how-to-use-the-token-security-feature) for how to create an token for ipinfo CloudQuery.


On the free tier, users have access to geolocation information of IP addresses, and the response payload looks like this:
```json copy
{
  "ip": "207.31.10.72",
  "hostname": "a207-31-10-072.unionsd.k12.ca.us",
  "city": "San Jose",
  "region": "California",
  "country": "US",
  "loc": "37.2563,-121.9229",
  "org": "AS3734 Santa Clara County Office of Education",
  "postal": "95124",
  "timezone": "America/Los_Angeles",
}
```

If the user is signed up to a [paid tier](https://ipinfo.io/pricing), they can access several different IP data such as IP privacy/VPN detection, company insights, ASN data, carrier data etc. The API response payload depends on the pricing the user is on. The highest tier API response payload looks like this:
```json copy
{
  "ip": "43.241.123.0",
  "city": "Mahbūbnagar",
  "region": "Telangana",
  "country": "IN",
  "loc": "17.3250,78.5623",
  "postal": "509129",
  "timezone": "Asia/Kolkata",
  "asn": {
    "asn": "AS134033",
    "name": "MITHRIL TELECOMMUNICATIONS PVT. LTD.",
    "domain": "hireachbroadband.com",
    "route": "43.241.123.0/24",
    "type": "isp"
  },
  "company": {
    "name": "MITHRIL TELECOMMUNICATIONS PVT. LTD.",
    "domain": "mithriltele.com",
    "type": "isp"
  },
  "carrier": {
    "name": "AirTel",
    "mcc": "404",
    "mnc": "2"
  },
  "privacy": {
    "vpn": false,
    "proxy": false,
    "tor": false,
    "relay": false,
    "hosting": false,
    "service": ""
  },
  "abuse": {
    "address": "5-5-126/1/PLOT NO.5, KRANTHI HILLS COLONY,VANSATHALIPURAM, Hyderabad, Telangana, 500070",
    "country": "IN",
    "email": "support@mithriltelecom.com",
    "name": "RAJESH PONNA",
    "network": "43.241.120.0/22",
    "phone": "+91 9951512121"
  },
  "domains": {
    "page": 0,
    "total": 0,
    "domains": []
  }
}
```