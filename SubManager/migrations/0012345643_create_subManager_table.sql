-- +goose Up
create table subscriptions(
    sub_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(60) NOT NULL,
    user_id UUID NOT NULL,
    price INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
drop table subscriptions;