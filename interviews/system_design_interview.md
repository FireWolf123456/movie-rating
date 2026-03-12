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

## Interviewer Reference — Answers to Common Clarifying Questions

> This section is for the interviewer only. These answers are ready to give when the candidate asks clarifying questions. Keep answers short and direct.

### Traffic & Scale

**How many users are there / what is the read-to-write ratio?**
> Approximately 1 million daily active users. Availability checks are much more frequent than orders — assume roughly 50 reads per write (users browse before buying).

**Are there traffic peaks?**
> Yes, evenings and weekends. Assume peak traffic is approximately 5 times the daily average. The system should handle bursts gracefully.

**Can multiple Distribution Centers serve the same customer location?**
> Yes. A customer location may be covered by 1 to 3 Distribution Centers at any given time. Availability is the union of all copies across those Distribution Centers.

---

### Data Model

**What is a "physical copy"? Is it an individual disc?**
> Yes. Each physical copy is a single disc (DVD or Blu-ray). There can be multiple copies of the same movie title across different Distribution Centers, or even within the same Distribution Center.

**How many copies does a typical Distribution Center hold?**
> On average, about 10 copies per Distribution Center (100,000 total copies across 10,000 Distribution Centers). Some Distribution Centers may hold more popular titles in higher quantities.

**How is customer location represented?**
> As geographic coordinates (latitude/longitude), provided by the client app.

---

### The "1 Hour Delivery" Rule

**How do you define "nearby" for a Distribution Center?**
> A fixed radius of 20 km from the customer's location. No need to account for traffic — keep it simple.

**Is the 1-hour radius checked at query time or enforced at order time?**
> Both. Availability queries filter by Distribution Centers within range. At order time the system should re-validate that the Distribution Center is still within range.

---

### Consistency & Ordering

**Can availability reads be slightly stale?**
> Yes, a few seconds of staleness is acceptable for availability queries. The important thing is that the order itself is strongly consistent — no two customers can claim the same physical copy.

**Is there a cart or reservation step?**
> That's a good question to leave open for the candidate to propose. A short-lived reservation (for example, 10 minutes) is a reasonable and expected design choice.