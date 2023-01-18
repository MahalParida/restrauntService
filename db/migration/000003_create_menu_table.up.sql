CREATE TABLE IF NOT EXISTS menu(
    Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255) NOT NULL,
    Type ENUM('veg', 'non-veg') NOT NULL,
    Spice_level ENUM('low','medium','high') NOT NULL,
    Available_on SET('sunday','monday','tuesday','wednesday','thursday','friday','saturday') NOT NULL,
    Is_vegan BOOLEAN,
    Is_available BOOLEAN,
    Cost INT,
    Preparation_time INT,
    Created_at TIMESTAMP DEFAULT NOW(),
    Updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW()
);


