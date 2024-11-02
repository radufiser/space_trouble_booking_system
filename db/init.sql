CREATE TABLE IF NOT EXISTS destinations (
    id VARCHAR(36) PRIMARY KEY,
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
    launchpad_id VARCHAR(36) REFERENCES launchpads(id),
    day_of_week INT,                   -- 0 for Sunday, 1 for Monday
    destination_id VARCHAR(36) REFERENCES destinations(id),
    last_updated TIMESTAMPTZ DEFAULT NOW(),               
    PRIMARY KEY (launchpad_id, day_of_week, destination_id)
);

CREATE TABLE IF NOT EXISTS bookings (
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    gender VARCHAR(10),
    birthday DATE,
    launchpad_id VARCHAR(36) NOT NULL,
    destination_id VARCHAR(36) NOT NULL,
    launch_date TIMESTAMP NOT NULL
);

INSERT INTO destinations (id, name) VALUES
('29a8f36a-14eb-47a3-beeb-bf48b5d2fefe', 'Mars'),
('1443a911-b39c-4404-bdab-dcffe4a6c019', 'Moon'),
('0d49ae5c-9efb-45f7-9e82-7f5569f7324e', 'Pluto'),
('cf134939-a20a-4c17-a011-981c964574fc', 'Asteroid Belt'),
('ebc98b61-ba27-414d-a75d-db910ddf2ea2', 'Europa'),
('69aea949-8fba-4059-a9bc-a5de0d9f9b59', 'Titan'),
('954b6450-fd41-41d5-877a-af665309d25d', 'Ganymede');

INSERT INTO launchpads (id, name, full_name, locality, region, status) VALUES
('5e9e4501f509094ba4566f84', 'CCSFS SLC 40', 'Cape Canaveral Space Force Station Space Launch Complex 40', 'Cape Canaveral', 'Florida', 'active'),
('5e9e4502f5090927f8566f85', 'STLS', 'SpaceX South Texas Launch Site', 'Boca Chica Village', 'Texas', 'under construction'),
('5e9e4502f509092b78566f87', 'VAFB SLC 4E', 'Vandenberg Space Force Base Space Launch Complex 4E', 'Vandenberg Space Force Base', 'California', 'active'),
('5e9e4502f509094188566f88', 'KSC LC 39A', 'Kennedy Space Center Historic Launch Complex 39A', 'Cape Canaveral', 'Florida', 'active');

INSERT INTO weekly_schedule (launchpad_id, day_of_week, destination_id, last_updated) VALUES
('5e9e4502f509094188566f88', 0, 'cf134939-a20a-4c17-a011-981c964574fc', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 1, 'ebc98b61-ba27-414d-a75d-db910ddf2ea2', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 2, '69aea949-8fba-4059-a9bc-a5de0d9f9b59', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 3, '954b6450-fd41-41d5-877a-af665309d25d', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 4, '29a8f36a-14eb-47a3-beeb-bf48b5d2fefe', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 5, '1443a911-b39c-4404-bdab-dcffe4a6c019', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509094188566f88', 6, '0d49ae5c-9efb-45f7-9e82-7f5569f7324e', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 0, '29a8f36a-14eb-47a3-beeb-bf48b5d2fefe', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 1, '1443a911-b39c-4404-bdab-dcffe4a6c019', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 2, '0d49ae5c-9efb-45f7-9e82-7f5569f7324e', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 3, 'cf134939-a20a-4c17-a011-981c964574fc', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 4, 'ebc98b61-ba27-414d-a75d-db910ddf2ea2', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 5, '69aea949-8fba-4059-a9bc-a5de0d9f9b59', '2024-11-02 11:45:50.003287+00'),
('5e9e4501f509094ba4566f84', 6, '954b6450-fd41-41d5-877a-af665309d25d', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 0, '1443a911-b39c-4404-bdab-dcffe4a6c019', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 1, '0d49ae5c-9efb-45f7-9e82-7f5569f7324e', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 2, 'cf134939-a20a-4c17-a011-981c964574fc', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 3, 'ebc98b61-ba27-414d-a75d-db910ddf2ea2', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 4, '69aea949-8fba-4059-a9bc-a5de0d9f9b59', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 5, '954b6450-fd41-41d5-877a-af665309d25d', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f5090927f8566f85', 6, '29a8f36a-14eb-47a3-beeb-bf48b5d2fefe', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 0, '0d49ae5c-9efb-45f7-9e82-7f5569f7324e', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 1, 'cf134939-a20a-4c17-a011-981c964574fc', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 2, 'ebc98b61-ba27-414d-a75d-db910ddf2ea2', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 3, '69aea949-8fba-4059-a9bc-a5de0d9f9b59', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 4, '954b6450-fd41-41d5-877a-af665309d25d', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 5, '29a8f36a-14eb-47a3-beeb-bf48b5d2fefe', '2024-11-02 11:45:50.003287+00'),
('5e9e4502f509092b78566f87', 6, '1443a911-b39c-4404-bdab-dcffe4a6c019', '2024-11-02 11:45:50.003287+00');