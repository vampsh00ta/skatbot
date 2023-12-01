begin;

insert  INTO subject ( name ) VALUES ('матан'),('линал'),('писюн') on CONFLICT(name) do nothing;
insert  INTO subject_type ( name ) VALUES ('Курсовая работа'),('Контрольная работа'),('Экзаменационные вопросы'),('Домашняя работа') on CONFLICT(name) do nothing;
insert  INTO variant_type ( name ) VALUES ('Вопросы'),('Готовое решение') on CONFLICT(name) do nothing;
insert into semester(number) values(1),
                           (2),
                           (3),
                           (6),
                           (7),
                           (5),
                           (8),
                           (9),
                           (10),
                           (11)on CONFLICT(number) do nothing;
insert into instistute (number ) values(1),
                                      (2),
                                      (3),
                                      (6),
                                      (7),
                                      (5),
                                      (8),
                                      (9),
                                      (10),
                                      (11)on CONFLICT(number) do nothing;

commit;