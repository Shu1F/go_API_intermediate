insert into articles (title, contents, username, nice, created_at) values
('first Post', 'This is my first blog', 'Shuichi', 2, now());
insert into articles (title, contents, username, nice) values
('2nd', 'Second post', 'Shuichi', 4);

insert into comments (article_id, message, created_at) values
(1, 'Great post!', now());
insert into comments (article_id, message) values
(1, 'Thanks for sharing');