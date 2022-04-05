CREATE TABLE public.delivery
(
    id_delivery serial PRIMARY KEY UNIQUE,
    name        VARCHAR(255) NOT NULL,
    phone       VARCHAR(255) NOT NULL,
    zip         VARCHAR(255) NOT NULL,
    city        VARCHAR(255) NOT NULL,
    address     VARCHAR(255) NOT NULL,
    region      VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL
);

CREATE TABLE public.payment
(
    id_payment    serial PRIMARY KEY UNIQUE,
    transaction   VARCHAR(255) NOT NULL,
    request_id    VARCHAR(255) NOT NULL,
    currency      VARCHAR(255) NOT NULL,
    provider      VARCHAR(255) NOT NULL,
    amount        INT,
    payment_dt    INT,
    bank          VARCHAR(255) NOT NULL,
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT
);

CREATE TABLE public.item
(
    id_item      serial PRIMARY KEY UNIQUE,
    chrt_id      INT,
    track_number VARCHAR(255) NOT NULL,
    price        INT,
    rid          VARCHAR(255) NOT NULL,
    name         VARCHAR(255) NOT NULL,
    sale         INT,
    size         VARCHAR(255) NOT NULL,
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR(255) NOT NULL,
    status       INT
);

CREATE TABLE public.order
(
    id_order           serial UNIQUE PRIMARY KEY,
    order_uid          VARCHAR(255) NOT NULL UNIQUE,
    track_number       VARCHAR(255) NOT NULL,
    entry              VARCHAR(255) NOT NULL,
    delivery_id        INT,
    payment_id         INT,
    locale             VARCHAR(255) NOT NULL,
    internal_signature VARCHAR(255) NOT NULL,
    customer_id        VARCHAR(255) NOT NULL,
    delivery_service   VARCHAR(255) NOT NULL,
    shardkey           VARCHAR(255) NOT NULL,
    sm_id              VARCHAR(255) NOT NULL,
    date_created       VARCHAR(255) NOT NULL,
    oof_shard          VARCHAR(255) NOT NULL,

    FOREIGN KEY (delivery_id) REFERENCES Delivery (id_delivery)
        ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (payment_id) REFERENCES Payment (id_payment)
        ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE public.order_items
(
    id       serial UNIQUE PRIMARY KEY,
    order_id INT,
    item_id  INT,
    FOREIGN KEY (order_id) REFERENCES public.order (id_order)
        ON DELETE CASCADE ON UPDATE CASCADE,

    FOREIGN KEY (item_id) REFERENCES public.item (id_item)
        ON DELETE CASCADE ON UPDATE CASCADE
);
