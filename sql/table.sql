-- Table User
CREATE TABLE ms_user (
  id VARCHAR(50) PRIMARY KEY,
  username VARCHAR(50),
  password VARCHAR(50),
  role VARCHAR(7),
  is_active BOOLEAN
);

-- Table ms_customer
CREATE TABLE ms_customer (
  id VARCHAR(50) PRIMARY KEY,
  id_user VARCHAR(50) REFERENCES ms_user (id),
  full_name VARCHAR(100) NOT NULL, 
  NIK VARCHAR(15) NOT NULL,
  noPhone VARCHAR(15) NOT NULL,
  email VARCHAR(100) NOT NULL,
  address VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL
);

-- Table ms_member
CREATE TABLE ms_member (
  id VARCHAR(50) PRIMARY KEY,
  id_customer VARCHAR(50) REFERENCES ms_customer (id),
  type VARCHAR(50) NOT NULL,
  expire DATE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL
);

-- Table ms_vehicle
CREATE TABLE ms_vehicle (
  id VARCHAR(50) PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  type VARCHAR(50) NOT NULL,
  identification_number INT NOT NULL,
  machine_number INT NOT NULL,
  release_date DATE NOT NULL,
  price DECIMAL NOT NULL,
  status VARCHAR(10) NOT NULL,
  is_available BOOLEAN NOT NULL,
  number_plate VARCHAR(10),
  stnk VARCHAR(10),
  no_bpkb VARCHAR(10)
);

-- Table tx_cash
CREATE TABLE tx_cash (
  id VARCHAR(50) PRIMARY KEY,
  id_vehicle VARCHAR(50) REFERENCES ms_vehicle (id),
  id_customer VARCHAR(50) REFERENCES ms_customer (id),
  price DECIMAL NOT NULL,
  date_payment DATE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL
);

-- Table tx_credit
CREATE TABLE tx_credit (
  id VARCHAR(50) PRIMARY KEY,
  id_vehicle VARCHAR(50) REFERENCES ms_vehicle (id),
  id_customer VARCHAR(50) REFERENCES ms_customer (id),
  price DECIMAL NOT NULL,
  interest DECIMAL NOT NULL,
  date_in DATE NOT NULL,
  date_out DATE NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NULL,
  updated_by VARCHAR(50) NULL
);

-- Table tx_installment_credit
CREATE TABLE tx_installment_credit (
  id VARCHAR(50) PRIMARY KEY,
  id_vehicle VARCHAR(50) REFERENCES ms_vehicle (id),
  id_credit VARCHAR(50) REFERENCES tx_credit (id),
  price DECIMAL NOT NULL,
  total_payment_now DECIMAL NOT NULL,
  date_payment DATE NOT NULL,
  date_finish DATE NOT NULL,
  due_date INT NOT NULL,
  status BOOLEAN NOT NULL,
  suspend BOOLEAN NOT NULL
);

-- Table tx_rent
CREATE TABLE tx_rent (
  id VARCHAR(50) PRIMARY KEY,
  id_vehicle VARCHAR(50) REFERENCES ms_vehicle (id),
  id_customer VARCHAR(50) REFERENCES ms_customer (id),
  price DECIMAL NOT NULL,
  date_in DATE NOT NULL,
  date_out DATE NOT NULL,
  status VARCHAR(50) NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_by VARCHAR(50) NULL
);
