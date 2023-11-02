SET @TargetIP6 = '2001:0db8:85a3:0000:0000:8a2e:0370:7334';

SELECT network, netmask, country_code -- select the columns
FROM ip_networks -- from the table
WHERE INET6_ATON(@TargetIP6) >= INET6_ATON(network) -- where the target IP is greater than or equal to the network
  AND INET6_ATON(@TargetIP6) <= (INET6_ATON(network) + (1 << (128 - netmask)) - 1) -- and the target IP is less than or equal to the network + the number of IPs in the network
ORDER BY netmask ASC -- order by the netmask ascending
LIMIT 10; -- limit to 10

-- INET6_ATON() is a function that converts a ipv6 address to numbers