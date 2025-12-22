# 4 in a Row – Real-Time Game

A real-time, backend-driven implementation of the classic **4 in a Row (Connect Four)** game.  
Players can play against another human in real time using WebSockets, with automatic game state synchronization.

Built using **Go (Golang)** for the backend.

---

## Features

- Real-time 1v1 gameplay using **WebSockets**
- Lobby system for player matchmaking
- Automatic game state updates for both players
- Reconnect handling using match ID
- Turn-based validation & win detection
- Live reload during development using **Air**

---

## Tech Stack

### Backend
- **Go**
- `net/http`
- `gorilla/websocket`

### Dev Tooling
- **Air** – live reload for Go