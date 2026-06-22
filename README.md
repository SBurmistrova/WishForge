# WishForge

WishForge is a backend application designed to turn wishes into achievable outcomes.

Instead of maintaining a static wishlist, users choose a wish, transform it into a structured project, split it into smaller steps, and gradually move toward completion.

The project focuses on progress through action rather than collecting ideas.

---

## Features

* Create and manage wishes
* Convert wishes into structured projects
* Break goals into smaller actionable steps
* Track completion progress
* Update and complete project steps
* Remove wishes together with their related steps
* Preserve completed journeys for future reflection

---

## Tech Stack

* Go
* SQL Server
* Chi Router
* REST API

---

## Architecture

```text
handlers/
service/
storage/
model/
```

### Layers

* **Handlers** — HTTP request processing and response generation
* **Service** — business logic and validation
* **Storage** — database interaction
* **Model** — application data structures

---

## API Overview

### Wishes

```http
GET    /wishes
POST   /wishes
GET    /wishes/{wishID}
PATCH  /wishes/{wishID}
DELETE /wishes/{wishID}
```

### Steps

```http
GET    /wishes/{wishID}/steps
POST   /wishes/{wishID}/steps
PATCH  /wishes/{wishID}/steps/{stepID}
DELETE /wishes/{wishID}/steps/{stepID}
```

---

## Technical Goals

This project was built as a backend portfolio project to practice:

* REST API design
* Backend architecture
* Layer separation
* Business logic implementation
* SQL integration
* Transaction handling
* Working with nested resources

---

## Future Improvements

* Progress calculation for wishes
* Automatic wish completion
* User authentication
* Completion reports and journey history
* Tests and CI
