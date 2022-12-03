INSERT INTO events (
        id,
        name,
        description,
        location,
        meeting_time,
        created_at
    )
VALUES (
        'event-1',
        'Car Meetup 1',
        'A meetup for car enthusiasts',
        'Central Park',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        NOW()
    ) ON CONFLICT (id) DO NOTHING;
-- Insert a row into the phone_numbers table
INSERT INTO phone_numbers (id, number, created_at)
VALUES ('2125551212', '2125551212', NOW()) ON CONFLICT (id) DO NOTHING;
-- Insert a row into the relationships table
INSERT INTO relationships (id, event_id, phone_number_id)
VALUES ('relationship-1', 'event-1', '2125551212') ON CONFLICT (id) DO NOTHING;
-- Insert a row into the notifications table
INSERT INTO notifications (id, message, sent_at, event_id, created_at)
VALUES (
        'notification-1',
        'Car Meetup 1 is coming up soon!',
        TO_TIMESTAMP('2022-12-01 10:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        'event-1',
        NOW()
    ) ON CONFLICT (id) DO NOTHING;