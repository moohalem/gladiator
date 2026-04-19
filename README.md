# ⚔️ Go Gladiator Battle Simulator

A lightweight, terminal-based RPG combat simulator built entirely in Go. This project demonstrates core Golang concepts like **Struct Embedding**, **Interfaces**, **Pointers**, and **Pass-by-Reference** state management through a fun, randomized battle system.

## ✨ Features

* **Dynamic Combat System:** An interface-driven battle loop (`model.Fighter`) that allows any entity to fight any other entity (Hero vs. Adversary, Hero vs. Hero, etc.).
* **Randomized Spawns:** Pulls from a global pool of 50 Hero archetypes and 50 Adversary types.
* **Loot & Equipment:** Heroes have a chance to spawn with random weapons from the armory, boosting their attack stats.
* **Special Attacks:** RNG-based special abilities (Mana spells for Heroes, Special Attacks for Adversaries) that trigger dynamically during combat.

## 📂 Project Structure

```text
gladiator/
├── core/
│   └── battle.go       # Contains the main combat loop and RNG turn logic
├── model/
│   ├── adversary.go    # Adversary struct and special attack logic
│   ├── base.go         # Base stats and shared methods for combatants
│   ├── hero.go         # Hero struct, equipment logic, and mana attacks
│   ├── interfaces.go   # Defines the Fighter interface used by the battle system
│   ├── pool.go         # Global slices holding the weapon, hero, and enemy databases
│   └── weapon.go       # Weapon structs and generation logic
├── .gitignore          # Ignored files for Git
├── go.mod              # Go module dependencies and configuration
├── LICENSE             # MIT License
├── main.go             # Entry point, initializes the combatants and starts the game
└── README.md           # Project documentation
