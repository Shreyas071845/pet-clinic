select * from owners;

SELECT * FROM pets;

SELECT * FROM appointments;


-- Drop existing tables first (if needed)
DROP TABLE IF EXISTS appointments CASCADE;
DROP TABLE IF EXISTS pets CASCADE;
DROP TABLE IF EXISTS owners CASCADE;

-- Owners table
CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    contact VARCHAR(20),
    email VARCHAR(100) UNIQUE
);

-- Pets table
CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    species VARCHAR(50),
    breed VARCHAR(50),
    owner_id INT NOT NULL REFERENCES owners(id) ON DELETE CASCADE,
    medical_history TEXT
);

-- Appointments table
CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    time TIME NOT NULL,
    pet_id INT NOT NULL REFERENCES pets(id) ON DELETE CASCADE,
    reason TEXT
);

INSERT INTO owners (name, contact, email) VALUES
('Ganesh', '9876543210', 'john.doe@gmail.com'),
('Riya', '9998877665', 'priya.sharma@gmail.com'),
('Ravi Kumar', '8899001122', 'ravi.kumar@yahoo.com');

INSERT INTO pets (name, species, breed, owner_id, medical_history) VALUES
('Bruno', 'Dog', 'Labrador', 1, 'Vaccinated and dewormed'),
('Misty', 'Cat', 'Persian', 2, 'Allergic to certain foods'),
('Rocky', 'Dog', 'German Shepherd', 3, 'Hip dysplasia treatment ongoing'),
('Simba', 'Cat', 'Maine Coon', 2, 'Neutered last month');

INSERT INTO appointments (date, time, pet_id, reason) VALUES
('2025-10-30', '10:00', 1, 'Routine check-up'),
('2025-10-31', '16:00', 2, 'Vaccination booster'),
('2025-11-01', '09:30', 3, 'Follow-up on hip treatment');

