
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