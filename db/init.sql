CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS destinations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS launchpads (
    id VARCHAR(36) PRIMARY KEY,
    name TEXT NOT NULL,
    full_name TEXT NOT NULL,
    locality TEXT,
    region TEXT,
    status TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS weekly_schedule (
    launchpad_id TEXT,
    day_of_week INT,                   -- 0 for Monday, 1 for Tuesday
    destination_id UUID REFERENCES destinations(id) ON DELETE CASCADE,
    last_updated TIMESTAMPTZ,           
    PRIMARY KEY (launchpad_id, day_of_week)
);


CREATE TABLE IF NOT EXISTS bookings (
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gender VARCHAR(10),
    birthday DATE,
    launchpad_id VARCHAR(20) NOT NULL,
    destination_id VARCHAR(20) NOT NULL,
    launch_date TIMESTAMP NOT NULL
);

INSERT INTO destinations (name) VALUES
('Mars'),
('Moon'),
('Pluto'),
('Asteroid Belt'),
('Europa'),
('Titan'),
('Ganymede')
ON CONFLICT (name) DO NOTHING;