-- Tabel Customer
CREATE TABLE Customer (
  id SERIAL PRIMARY KEY,
  full_name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  NIK VARCHAR(255) NOT NULL,
  noPhone VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  create_at TIMESTAMP NOT NULL,
  create_by VARCHAR(255) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(255) NOT NULL
);

-- Tabel Member
CREATE TABLE Member (
  id SERIAL PRIMARY KEY,
  id_user INT REFERENCES Customer (id),
  type VARCHAR(255) NOT NULL,
  expire DATE NOT NULL,
  create_at TIMESTAMP NOT NULL,
  create_by VARCHAR(255) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(255) NOT NULL
);

-- Tabel Vehicle
CREATE TABLE Vehicle (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(255) NOT NULL,
  release DATE NOT NULL,
  price DECIMAL NOT NULL,
  status VARCHAR(255) NOT NULL
);

-- Tabel Detail_Vehicle
CREATE TABLE Detail_Vehicle (
  id SERIAL PRIMARY KEY,
  id_vehicle INT REFERENCES Vehicle (id),
  stock INT NOT NULL,
  number_plate VARCHAR(255) NOT NULL,
  stnk VARCHAR(255) NOT NULL,
  no_bpkb VARCHAR(255) NOT NULL
);

-- Tabel Cash
CREATE TABLE Cash (
  id SERIAL PRIMARY KEY,
  id_vehicle INT REFERENCES Vehicle (id),
  id_customer INT REFERENCES Customer (id),
  price DECIMAL NOT NULL,
  date_payment DATE NOT NULL,
  create_at TIMESTAMP NOT NULL,
  create_by VARCHAR(255) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(255) NOT NULL
);

-- Tabel Credit
CREATE TABLE Credit (
  id SERIAL PRIMARY KEY,
  id_vehicle INT REFERENCES Vehicle (id),
  id_customer INT REFERENCES Customer (id),
  price DECIMAL NOT NULL,
  interest DECIMAL NOT NULL,
  date_In DATE NOT NULL,
  date_Out DATE NOT NULL,
  create_at TIMESTAMP NOT NULL,
  create_by VARCHAR(255) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(255) NOT NULL
);

-- Tabel Installment_Credit
CREATE TABLE Installment_Credit (
  id SERIAL PRIMARY KEY,
  id_vehicle INT REFERENCES Vehicle (id),
  id_credit INT REFERENCES Credit (id),
  price DECIMAL NOT NULL,
  total_payment_now DECIMAL NOT NULL,
  date_payment DATE NOT NULL,
  date_finish DATE NOT NULL,
  due_date INT NOT NULL,
  status BOOLEAN NOT NULL,
  suspend BOOLEAN NOT NULL
);

-- Tabel Rent
CREATE TABLE Rent (
  id SERIAL PRIMARY KEY,
  id_vehicle INT REFERENCES Vehicle (id),
  id_customer INT REFERENCES Customer (id),
  price DECIMAL NOT NULL,
  date_In DATE NOT NULL,
  date_Out DATE NOT NULL,
  status VARCHAR(255) NOT NULL,
  create_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255) NOT NULL
);
