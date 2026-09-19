# EEPROM Password Manager TUI

A terminal-based password manager written in **Go**.

The application uses **Cobra** for command handling and communicates with an **Arduino** through Serial. The Arduino manages an external **I²C EEPROM** for persistent storage.

## Architecture

```text
┌──────────────────────────┐
│       Go TUI / CLI       │
│         Cobra            │
└────────────┬─────────────┘
             │
          Serial
             │
             ▼
      ┌─────────────┐
      │   Arduino   │
      └──────┬──────┘
             │
            I²C
             │
             ▼
┌─────────────────────────┐
│   AT24C256 / 24LC512    │
│         EEPROM          │
└─────────────────────────┘
```

## Requirements

* Go
* Arduino
* AT24C256 / 24LC512 EEPROM
* USB serial connection
* Arduino firmware

## Usage

The application is run directly with:

```bash
go run .
```

Commands are provided as arguments to the application.

### Add

Add a new password record:

```bash
go run . add "platform" "login" "password"
```

Example:

```bash
go run . add "GitHub" "mochafik" "mypassword"
```

### Read

Read a password record:

```bash
go run . read "platform"
```

Example:

```bash
go run . read "GitHub"
```

### Delete

Delete a password record:

```bash
go run . delete "platform"
```

Example:

```bash
go run . delete "GitHub"
```

### Reset

Reset the password manager:

```bash
go run . reset
```

## Commands

| Command  | Usage                                        | Description        |
| -------- | -------------------------------------------- | ------------------ |
| `add`    | `go run . add "platform" "login" "password"` | Add a record       |
| `read`   | `go run . read "platform"`                   | Read a record      |
| `delete` | `go run . delete "platform"`                 | Delete a record    |
| `reset`  | `go run . reset`                             | Reset the database |

## Communication

The Go application communicates with the Arduino through **Serial**.

```text
Go application
      │
      │ Serial
      ▼
   Arduino
      │
      │ I²C
      ▼
   EEPROM
```

The Go application sends commands to the Arduino. The Arduino performs the EEPROM operations and sends the result back through Serial.

## EEPROM

The Arduino uses an external EEPROM for persistent storage.

Supported EEPROMs:

* AT24C256
* 24LC512

I²C address:

```text
0x50
```

Current conceptual memory layout:

```text
EEPROM
│
├── Configuration Byte
├── Header
├── Index
└── Records
```

The Go application does not directly access the EEPROM. EEPROM addressing and storage management are handled by the Arduino.

## Technologies

* **Go**
* **Cobra**
* **Arduino**
* **C++**
* **Serial communication**
* **I²C**
* **AT24C256 / 24LC512 EEPROM**

## Project Status

**In development**

Current commands:

```bash
go run . add "platform" "login" "password"
go run . read "platform"
go run . delete "platform"
go run . reset
```
