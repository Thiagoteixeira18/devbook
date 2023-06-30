insert into usuarios (nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$x7vRLGiJUdH04tn984ya4uHhB2RuOARuFPgqYD84H2YguMV1f.i.W"), -- usuario1
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$x7vRLGiJUdH04tn984ya4uHhB2RuOARuFPgqYD84H2YguMV1f.i.W"), -- usuario2
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$x7vRLGiJUdH04tn984ya4uHhB2RuOARuFPgqYD84H2YguMV1f.i.W"); -- usuario3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);