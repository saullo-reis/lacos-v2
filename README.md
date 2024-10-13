
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>API Documentation</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            background-color: #f4f4f4;
            margin: 0;
            padding: 20px;
        }
        h1, h2, h3 {
            color: #333;
        }
        .container {
            max-width: 900px;
            margin: auto;
            background: white;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
        pre {
            background: #333;
            color: white;
            padding: 10px;
            border-radius: 5px;
            overflow-x: auto;
        }
        code {
            color: #c7254e;
            background-color: #f9f2f4;
            padding: 2px 4px;
            border-radius: 4px;
            font-size: 90%;
        }
        .section {
            margin-bottom: 30px;
        }
        ul {
            margin: 10px 0;
            padding-left: 20px;
        }
        li {
            margin-bottom: 10px;
        }
        .endpoint {
            font-weight: bold;
            color: #007BFF;
        }
        .note {
            color: #555;
            font-style: italic;
        }
    </style>
</head>
<body>

<div class="container">

    <h1>API Documentation</h1>

    <div class="section">
        <h2>Guia de Início</h2>

        <h3>Configuração</h3>
        <ul>
            <li>Configure o arquivo <code>.env</code> conforme necessário, ou renomeie o arquivo <code>.env_sample</code> para <code>.env</code> para usar as configurações padrão.</li>
            <li>Certifique-se de ter o Docker e o Docker Compose instalados na sua máquina.</li>
        </ul>

        <h3>Executando a Aplicação</h3>
        <p>Comandos para iniciar e parar o servidor:</p>
        <ul>
            <li><strong>Iniciar o servidor:</strong></li>
            <pre><code>docker-compose up --build</code></pre>

            <li><strong>Parar o servidor:</strong></li>
            <pre><code>docker-compose down</code></pre>
        </ul>

        <h3>Verificação</h3>
        <p>Faça uma requisição GET para <code>http://localhost:8080/ping</code>. A resposta esperada é:</p>
        <pre><code>{
    "message": "pong",
    "status": 200
}</code></pre>
        <p class="note">Se essa resposta for retornada, a API está pronta para uso.</p>
    </div>

    <div class="section">
        <h2>APIs Disponíveis</h2>

        <h3>Sistema de Busca</h3>
        <p>A API suporta os seguintes parâmetros para busca:</p>
        <ul>
            <li><strong>q:</strong> Condição para a consulta, por exemplo: <code>name='Saullo'</code>.</li>
            <li><strong>limit:</strong> Limita o número de registros retornados, por exemplo, <code>limit=10</code>.</li>
            <li><strong>offset:</strong> Define o número de registros a pular, por exemplo, <code>offset=5</code>.</li>
        </ul>
    </div>

    <div class="section">
        <h2>Autenticação e Gerenciamento de Usuários</h2>

        <h3><span class="endpoint">/login</span> (POST)</h3>
        <p>Realiza login com um usuário administrador.</p>
        <p><strong>Body:</strong></p>
        <pre><code>{
    "username": "admin",
    "password": "admin"
}</code></pre>
        <p><strong>Resposta:</strong> Retorna um token JWT que será utilizado para permissões de administrador.</p>

        <h3><span class="endpoint">/register</span> (POST)</h3>
        <p>Registra um novo usuário.</p>
        <p><strong>Body:</strong></p>
        <pre><code>{
    "username": "saulloreis",
    "password": "12345678"
}</code></pre>

        <h3><span class="endpoint">/user/delete/:idUser</span> (DELETE)</h3>
        <p>Deleta um usuário pelo ID.</p>
        <pre><code>DELETE /user/delete/1</code></pre>

        <h3><span class="endpoint">/user/update/:idUser</span> (PATCH)</h3>
        <p>Atualiza as informações de um usuário.</p>
        <p><strong>Body:</strong></p>
        <pre><code>{
    "username": "ADS",
    "password": "1234"
}</code></pre>

        <h3><span class="endpoint">/user/get</span> (GET)</h3>
        <p>Busca todos os usuários. Você pode usar os parâmetros <code>q</code>, <code>limit</code> e <code>offset</code>.</p>

        <h3><span class="endpoint">/user/get/:idUser</span> (GET)</h3>
        <p>Busca um usuário específico pelo ID.</p>
        <pre><code>GET /user/get/1</code></pre>
    </div>

    <div class="section">
        <h2>Gerenciamento de Pessoas</h2>

        <h3><span class="endpoint">/persons/create</span> (POST)</h3>
        <p>Cria um novo registro de pessoa.</p>
        <p><strong>Body:</strong></p>
        <pre><code>{
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
}</code></pre>

        <h3><span class="endpoint">/persons/delete/:idPerson</span> (DELETE)</h3>
        <p>Marca uma pessoa como inativa, sem removê-la do banco de dados.</p>
        <pre><code>DELETE /persons/delete/1</code></pre>

        <h3><span class="endpoint">/persons/update</span> (PATCH)</h3>
        <p>Atualiza o registro de uma pessoa. O CPF é obrigatório e usado para identificação.</p>
        <p><strong>Body:</strong> Mesmo formato do body de criação.</p>

        <h3><span class="endpoint">/persons/get/:idPerson</span> (GET)</h3>
        <p>Busca uma pessoa específica pelo ID.</p>
        <pre><code>GET /persons/get/1</code></pre>

        <h3><span class="endpoint">/persons/get</span> (GET)</h3>
        <p>Busca várias pessoas, suportando os parâmetros <code>q</code>, <code>limit</code>, e <code>offset</code>.</p>

        <h3><span class="endpoint">/persons/active/:idPerson</span> (POST)</h3>
        <p>Ativa o registro de uma pessoa que estava desativada.</p>
        <pre><code>POST /persons/active/1</code></pre>
    </div>

    <div class="section">
        <h2>Gerenciamento de Atividades</h2>

        <h3><span class="endpoint">/activityList/create</span> (POST)</h3>
        <p>Cria uma nova atividade.</p>

        <h3><span class="endpoint">/activityList/delete/:idActivityList</span> (DELETE)</h3>
        <p>Deleta uma atividade.</p>

        <h3><span class="endpoint">/activityList/update</span> (PATCH)</h3>
        <p>Atualiza uma atividade.</p>

        <h3><span class="endpoint">/activityList/get</span> (GET)</h3>
        <p>Busca todas as atividades.</p>

        <h3><span class="endpoint">/activityList/get/:idActivityList</span> (GET)</h3>
        <p>Busca uma atividade específica pelo ID.</p>
    </div>

</div>

