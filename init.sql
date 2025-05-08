CREATE TABLE IF NOT EXISTS vehicle_locations (
    id SERIAL PRIMARY KEY,
    vehicle_id VARCHAR(20) NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    timestamp BIGINT NOT NULL
);

INSERT INTO vehicle_locations (vehicle_id, latitude, longitude, timestamp)
VALUES 
('B1234XYZ', -6.2088, 106.8456, 1715003456),
('C2345XYZ', -6.2089, 106.8457, 1715003556);