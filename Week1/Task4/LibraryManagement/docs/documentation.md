# Library Management System with Concurrency reservation

## Workflow

1. A user calls `ReserveBook(bookID, memberID)`.
2. The request is sent to `ReserveQueue` (a channel).
3. A background worker processes requests concurrently.
4. The book is marked “Reserved”.
5. If not borrowed in 5 seconds, it automatically becomes “Available” again.
