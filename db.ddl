create table ars.items
(
    id       integer not null
        unique,
    fname    text,
    fgroup   integer,
    fdesc    text,
    price    numeric(6, 2),
    quantity integer,
    show     boolean default true,
    new      boolean default true
);

alter table ars.items
    owner to postgres;

create table ars.tgroups
(
    id     integer not null
        unique,
    fname  text,
    fgroup integer,
    fdesc  text,
    show   boolean default true,
    new    boolean default true
);

alter table ars.tgroups
    owner to postgres;
