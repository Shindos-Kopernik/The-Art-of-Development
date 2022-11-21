package Playlist2_Lesson8_PostgreSQL

CREATE TABLE public.author
(
id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR (100) NOT NULL,
author_id UUID NOT NULL ,
CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id)
);

INSERT INTO author (name) VALUES ('Народ'); --02a3d7a0-d03a-4908-9f80-67e0495b9e2b
INSERT INTO author (name) VALUES ('Джофн Роулинг'); --9d98ee85-dbee-4f69-b2fa-4f78131d0646
INSERT INTO author (name) VALUES ('Джек Лондон'); --f1f3fc2f-c497-488b-add2-0a11094f3993

INSERT INTO  book (name, author_id) VALUES ('Колобок', '02a3d7a0-d03a-4908-9f80-67e0495b9e2b')
INSERT INTO book (name, author_id) VALUES ('Гарри Поттер', '9d98ee85-dbee-4f69-b2fa-4f78131d0646')
INSERT INTO book (name, author_id) VALUES ('Бриллиант', 'f1f3fc2f-c497-488b-add2-0a11094f3993')