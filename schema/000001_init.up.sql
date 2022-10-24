CREATE TABLE IF NOT EXISTS mail(
    sender_email varchar not null,
    sender_msg varchar not null,
    sender_password varchar not null,
    sender_hashed_password varchar not null, 
    recipients_email varchar[] not null,
    email_status varchar not null,
    sent_at timestamp default(now())
);