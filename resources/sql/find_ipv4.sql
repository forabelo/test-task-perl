SET @TargetIP = '192.168.1.5';

SELECT network, netmask, country_code -- select the columns
FROM ip_networks -- from the table
WHERE INET_ATON(@TargetIP) >= INET_ATON(network) -- where the target IP is greater than or equal to the network
  AND INET_ATON(@TargetIP) <= INET_ATON(network) + (1 << (32 - netmask)) - 1  -- and the target IP is less than or equal to the network + the number of IPs in the network
ORDER BY netmask ASC -- order by netmask ascending
LIMIT 10; -- limit to 10 results

-- INET_ATON() is a function that converts a ipv4 address to numbers, << is a bits shift operator

