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
├── main.go             # Entry point, initializes the combatants and starts the game
├── battle/
│   └── battle.go       # Contains the battle loop and RNG turn logic
└── model/
    ├── base.go         # Defines the Fighter interface and base stats
    ├── hero.go         # Hero struct, equipment logic, and Special Attack
    ├── adversary.go    # Adversary struct and Special Attack
    └── pool.go         # Global slices holding the weapon, hero, and enemy databases
