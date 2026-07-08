# Determinant Calculation Optimization using Laplace Expansion

Comparison between a baseline implementation and an optimized implementation of the Laplace expansion algorithm for determinant calculation of square matrices.

## Overview

This project was developed as part of the **Introduction to Programming** course at the **Federal University of Amazonas (UFAM)**.

The objective is to compare experimentally two versions of the determinant calculation algorithm:

- **Baseline:** classical recursive Laplace expansion.
- **Optimized:** expansion over the row or column containing the largest number of zero elements, avoiding unnecessary recursive calls when the selected cofactor is zero.

The implementation was developed entirely in **Go**.

---

## Project Structure

```text
.
├── main.go                  # Executes the experiments
├── laPlace.go               # Baseline implementation
├── laPlaceOtimizado.go      # Optimized implementation
├── go.mod
├── Relatorio.pdf
```

---

## Experimental Setup

The algorithms were evaluated using randomly generated square matrices of orders:

- 3 × 3
- 5 × 5
- 7 × 7
- 9 × 9
- 11 × 11

Each experiment was executed three times, and the average execution time (in nanoseconds) was used for comparison.

---

## Sample Results

| Matrix Order | Baseline (ns) | Optimized (ns) |
|--------------:|--------------:|---------------:|
| 3 | 0 | 0 |
| 5 | 0 | 0 |
| 7 | 193,033 | 254,533 |
| 9 | 25,124,800 | 7,840,166 |
| 11 | 3,941,542,133 | 1,626,135,166 |

The optimized implementation significantly reduced execution time for larger matrices by selecting the most suitable expansion row or column and eliminating unnecessary recursive calls.

---

## Technologies

- Go
- GoLand

---

## How to Run

Clone the repository:

```bash
git clone https://github.com/C-Cdev/laplace-determinant-optimization.git
```

Enter the project directory:

```bash
cd laplace-determinant-optimization
```

Run:

```bash
go run .
```

---

## Authors

- Caio Henrique Soares Melo
- André Luís de Souza Justino
- Lucas Chacon Vásquez Vale

---

## Academic Information

This repository contains the implementation and the experimental report produced for the Introduction to Programming course at UFAM.
