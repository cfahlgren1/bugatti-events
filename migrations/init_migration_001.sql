-- Add seed data for events
INSERT INTO events (
        id,
        name,
        description,
        location,
        created_at,
        updated_at
    )
VALUES (
        '1',
        'Ferrari Meet',
        'Monthly meetup for Ferrari owners and enthusiasts',
        'Ferrari dealership',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS')
    ),
    (
        '2',
        'Exotic Car Show',
        'Annual car show featuring exotic and luxury vehicles',
        'Convention center',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS')
    ),
    (
        '3',
        'Import Meet',
        'Weekly meetup for owners of imported performance vehicles',
        'Local park',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS')
    ) ON CONFLICT (id) DO NOTHING;
-- Add seed data for notifications
INSERT INTO notifications (id, message, sent_at, event_id)
VALUES (
        '1',
        'Ferrari Meet: Due to inclement weather, the meet has been moved to next weekend. Sorry for any inconvenience.',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        '1'
    ),
    (
        '2',
        'Exotic Car Show: The date of the show has been changed to next month. Mark your calendars!',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        '2'
    ),
    (
        '3',
        'Import Meet: The meet location has been changed to a new parking lot. See you there!',
        TO_TIMESTAMP('2022-12-01 12:00:00', 'YYYY-MM-DD HH24:MI:SS'),
        '3'
    ) ON CONFLICT (id) DO NOTHING;