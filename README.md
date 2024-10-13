# Como iniciar

Configure os .env que quiser ou se preferir apenas retire o "_sample" do .env que vai funcionar também.

Necessita de ter o docker e o docker-compose baixado em sua máquina.

## Comando para iniciar e desativar

- docker-compose up --build
- docker-compose down
Faça uma requisição para o http://localhost:8080/ping se retornar:
```
{
    "message": "pong",
    "status": 200
}
```

Está tudo certo e já pode utilizar a API

# APIs

## SISTEMA DE BUSCAS COMANDOS ADICIONAIS

q = Você consegue escrever uma condição para a query que será feita por exemplo:
name = 'Saullo' 
limit = Você consegue botar um limite em sua consulta por exemplo 10 registros
offset = Você consegue fazer com que a busca pule algumas linhas

### USER / AUTH

Para desenvolvimento você consegue utilizar o login admin e senha admin

<strong>/login</strong> (POST) com o body:
```
{
    "username": "admin",
    "password": "admin"
}
```
Vai retornar um token que é utilizado para fazer permissões de administrador.

<strong>/register</strong> (POST) com o body:
```
{
    "username": "saulloreis",
    "password": "12345678"
}
```
Você consegue registrar um usuário no sistema, lembrando apenas permissões de usuário apenas o usuário admin tem permissões de adm.

<strong>/user/delete/:idUser</strong> (DELETE):
Você consegue deletar um usuário do sistema.

<strong>/user/update/:idUser</strong> (PATCH) com o body:
```
{   
    "username": "ADS",
    "password": "1234" 
}
```
Você consegue atualizar um usuário.

<strong>/user/get</strong> (GET):
Você consegue fazer buscas de todos os usuários podendo utilizar q, limit e offset 
campos:
    id_user,
    username

<strong>/user/get/:idUser</strong> (GET):
Busca individual de um usuário.

## PERSON

<strong>/persons/create</strong> (POST):
Cria uma pessoa body:
```
{
    "name": "Saullo Reis",
    "birth_date": "2001-01-01",
    "rg": "12.345.678-9",
    "cpf": "18872800",
    "cad_unico": "1234567890",
    "nis": "9876543210",
    "school": "Example School",
    "address": "123 Example Street",
    "address_number": "456",
    "blood_type": "O+",
    "neighborhood": "Example Neighborhood",
    "city": "Example City",
    "cep": "12345-678",
    "home_phone": "1234-5678",
    "cell_phone": "91234-5678",
    "contact_phone": "1",
    "email": "saulloreis01@hotmail.com",
    "current_age": 34,
    "responsible_person": {
        "name": "Jane Doe",
        "relationship": "Mother",
        "rg": "23.456.789-0",
        "cpf": "123.456.789-09",
        "cell_phone": "98765-4321"
    }
}
```

<strong>/persons/delete/:idPerson</strong> (DELETE):
Deleta uma pessoa pelo seu ID ( Ela não é excluida do banco apenas fica inativa )

<strong>/persons/update</strong> (PATCH): 
Com o mesmo body de criação OBS: CPF é utilizado para identificação ou seja obrigatório no body.

<strong>/persons/get/:idPerson</strong> (GET):
Você busca individualmente uma pessoa pelo seu ID.

<strong>/persons/get</strong> (GET): 
Você busca várias pessoas. É possivel utilizar o q, limit e offset.
Os nomes dos campos são iguais ao body e se for fazer uma condição para uma pessoa utilize por exemplo: person.name e para uma pessoa responsável: rperson.name

<strong>/persons/active/:idPerson</strong> (POST):
Você ativa uma pessoa que estava desativada.

## Activity

<strong>/activityList/create</strong> (POST):
Você cria uma linha na lista de atividades.

<strong>/activityList/delete/idActivityList</strong> (DELETE):
Você deleta uma atividade.

<strong>/activityList/get</strong> (GET):
Você busca a lista de atividades. É possivel utilizar o q, limit e offset.

<strong>/activities/action/link</strong> (POST)
Você consegue linkar uma atividade a uma pessoa body:
```
{
    "id_person": 1,
    "id_period": 1,
    "id_activity": 3,
    "hour_start": "18:00",
    "hour_end": "20:00"
}
```

<strong>/activities/action/link/delete/:idActivities</strong> (DELETE)
Você consegue retirar uma atividade de uma pessoa.

<strong>/activities/getAll/:idPerson</strong> (GET)
Você consegue listar as atividade de uma pessoa. Não é possivel utilizar o q, limit e offset.

