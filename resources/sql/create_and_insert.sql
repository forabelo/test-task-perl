DROP table IF EXISTS ip_networks;

CREATE TABLE ip_networks (
    id INT AUTO_INCREMENT PRIMARY KEY, -- unique id
    network VARCHAR(45) NOT NULL, -- ipv4 network or ipv6 network. I assume that we keep the network in the string format.
    netmask INT NOT NULL, -- uint 0-32 for ipv4, 0-128 for ipv6
    country_code CHAR(2) NOT NULL -- ISO 3166-1 alpha-2
);

INSERT INTO ip_networks (network, netmask, country_code) VALUES
-- ipv4
('192.168.0.0', 32, 'DE'),
('192.168.0.1', 16, 'DE'),
('192.168.1.0', 32, 'DE'),
('192.168.0.0', 24, 'DE'),
('192.168.0.10', 16, 'DE'),
('192.0.0.0', 8, 'DE');

INSERT INTO ip_networks (network, netmask, country_code) VALUES
-- ipv6
 ('2001:0db8:85a3:0000:0000:8a2e:0370:0000', 128, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a2e:0000:0000', 112, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a2e:0370:7000', 120, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a00:0000:0000', 104, 'US'), -- +
('2001:0db8:85a3:0000:0000:0000:0000:0000', 96, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a2e:0300:0000', 110, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a2e:0370:7300', 124, 'US'), -- +
('2001:0db8:85a3:0000:0000:8a2e:0370:7330', 126, 'US'), -- +
('2002:0db8:85a3:0000:0000:8a2e:0370:7334', 128, 'DE'), -- -
('3001:0db8:85a3:0000:0000:8a2e:0370:7334', 128, 'DE'); -- -