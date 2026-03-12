# Movies Rating API — System Design Interview

## Position: Senior Software Engineer - Backend Go
## Time Limit: 45 minutes

---

## Summary

This is a working Go API for rating movies. Your task is to design a Local Movie Delivery Service

## Functional Requirements

### Core Requirements

- Customers should be able to query availability of movies, deliverable in 1 hour, by location (that is, the effective availability is the union of all inventory nearby Distribution Centers).
- Customers should be able to order multiple movies at the same time.

### Below the line (out of scope)
- Handling payments/purchases.
- Handling driver routing and deliveries.
- Search functionality and catalog APIs. (The system is strictly concerned with availability and ordering).
- Cancellations and returns.

## Non-Functional Requirements

### Core Requirements
- Availability requests should be fast (<100ms) because they will be called frequently as users browse by location.
- Ordering should be strongly consistent: two customers should not be able to purchase the same physical product.
- System should be able to support 10,000 Distribution Centers and 100,000 total physical copies distributed across all Distribution Centers.
- Order volume will be 10 million orders per day

### Below the line (out of scope)
- Privacy and security.
- Disaster recovery.
- Here's how these might be shorthanded in an interview. Note that out-of-scope requirements usually stem from questions like "do we need to handle privacy?". Interviewers are usually comfortable with you making assertions "I'm going to leave privacy out of scope for the start" and will correct you if needed.

---
