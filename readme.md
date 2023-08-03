
#  Bank - API for Merchant & Bank

This repository contains the API documentation for  Bank merchant and customer banking services.

## Setup Project
steps to set up the project:

1. **Download or Clone the Project:**
   You can download the project as a zip file from the following link: [Merchant and Bank Project](https://github.com/fajarsyaa/merchant-and-bank/archive/refs/heads/main.zip). Alternatively, if you have Git installed, you can clone the project using the following command in your terminal or command prompt:
   ```
   git clone https://github.com/fajarsyaa/merchant-and-bank.git
   ```

2. **Navigate to the Project Directory:**
   Open your terminal or command prompt and navigate to the root directory of the cloned or downloaded project:
   ```
   cd merchant-and-bank
   ```

3. **Install Dependencies:**
   To ensure all the required dependencies are installed, run the following command to use Go modules and tidy the dependencies:
   ```
   go mod tidy
   ```

4. **Run the Project:**
   Finally, start the application by running the following command:
   ```
   go run main.go
   ```

   This command will start the server, and you should see the application running locally. You can now interact with the API using the specified endpoints.

Make sure you have Go installed on your system before running these commands. If you encounter any issues during the setup, please check the project documentation or consult the project's repository for further instructions.

## Table of Contents

- [Introduction](#introduction)
- [API Endpoints](#api-endpoints)
  - [POST /login](#post-login)
  - [POST /logout](#post-logout)
  - [POST /register](#post-register)
  - [POST /transaction/create](#post-transaction-create)
  - [GET /transactions](#get-transactions)
  - [GET /merchants](#get-merchants)
  - [POST /topup](#post-topup)

## Introduction

This API provides endpoints for customers and merchants to access various banking functionalities offered by Bank. The API supports actions such as customer login, logout, registration, creating transactions, viewing transaction history, and more.

The API requires authentication using a Bearer Token, which is obtained during the login process.

## API Endpoints

### POST /login

Endpoint for customer login.

**URL:** `localhost:8000/login`

**Request Method:** POST

**Request Body:**
```json
{
    "username": "user1",
    "password": "password"
}
```

## 
### POST /logout

Endpoint to log out the customer.

**URL:** `localhost:8000/logout`

**Request Method:** POST

## 
### POST /register

Endpoint for customer registration if the user is not registered.

**URL:** `localhost:8000/register`

**Request Method:** POST

**Request Body:**
```json
{
    "username": "wahyu",
    "password": "password",
    "NoRek": "098709870987",
    "balance": 2000000
}
```

## 
### POST /transaction/create

Endpoint to create a new transaction. Only authenticated customers can perform transactions.

**URL:** `localhost:8000/transaction/create`

**Request Method:** POST

**Authorization:** Bearer Token

**Request Body:**
```json
{
    "merchant_rek": "123456789",
    "amount": 10000
}
```

## 
### GET /transactions

Endpoint to view all transaction history.

**URL:** `localhost:8000/transactions`

**Request Method:** GET

**Authorization:** Bearer Token

## 
### GET /merchants

Endpoint to view all available merchants.

**URL:** `localhost:8000/merchants`

**Request Method:** GET

## 
### POST /topup

Endpoint for customer top-up (add funds to the account).

**URL:** `localhost:8000/topup`

**Request Method:** POST

**Authorization:** Bearer Token

**Request Body:**
```json
{
    "balance": 1000000
}
```

---
