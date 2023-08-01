CREATE TABLE ms_customer (
    id VARCHAR(100),    
    email VARCHAR(100),
    username VARCHAR(200),
    password VARCHAR(255),
    no_rek BIGINT
);

INSERT INTO ms_customer (id, email, username, password, no_rek) VALUES
('MSC001', 'agus@com.com', 'agus', '$2a$12$l2CpaqOPl1HzVKXE6TY0A.3VZKn5WlHgjHqwm0BlyKIm88RqS8MyG', 1234567890123456),
('MSC002', 'bagus@com.com', 'bagus', '$2a$12$l2CpaqOPl1HzVKXE6TY0A.3VZKn5WlHgjHqwm0BlyKIm88RqS8MyG', 9876543210987654);