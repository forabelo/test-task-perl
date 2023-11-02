# SET @TargetIP = '2001:0db8:85a3:0000:0000:8a2e:0370:7334'; -- example input ipv6
SET @TargetIP = '192.168.1.5'; -- example input inv4

# We use the IF() function to determine if the input is IPv4 or IPv6 and do the same as in the previous examples

SELECT network, netmask, country_code -- select the columns
FROM ip_networks -- from the table
WHERE IF(CHAR_LENGTH(@TargetIP) <= 15, INET_ATON(@TargetIP) >= INET_ATON(network) -- if IPv4
    AND INET_ATON(@TargetIP) <= INET_ATON(network) + (1 << (32 - netmask)) - 1, -- then
         INET6_ATON(@TargetIP) >= INET6_ATON(network) -- else IPv6
             AND INET6_ATON(@TargetIP) <= (INET6_ATON(network) + (1 << (128 - netmask)) - 1)) -- then
ORDER BY netmask ASC
LIMIT 10;