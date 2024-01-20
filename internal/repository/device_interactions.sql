CREATE TABLE IF NOT EXISTS device_interactions (
   interaction_id varchar(36) PRIMARY KEY,
   latitude numeric(11,8) NOT NULL,
   longitude numeric(11,8) NOT NULL,
   device_id varchar(36) NOT NULL,
   device_name varchar(15) NOT NULL,
   timestamp TIMESTAMP NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);
