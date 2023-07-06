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

insert into publicacoes(titulo, conteudo, autor_id)
values
("publicacao do usuario 1", "essa é a publicaçao do usuario 1! oba", 1),
("publicacao do usuario 2", "essa é a publicaçao do usuario 2! oba", 2),
("publicacao do usuario 3", "essa é a publicaçao do usuario 3! oba", 3);