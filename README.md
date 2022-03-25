<h1 align="center">BackendLoginGo</h1>
<p align="center">
A simple example backend web for account register, login, logout, and get the account info using Go language.
</p>

<br>

### Tutorial
- Open your terminal/cmd and type `git clone https://github.com/rushkii/BackendLoginGo.git`
- Then `cd BackendLoginGo` and change your database DSN in `config.env`
- Once setup complete, run the program `go run server.go`

<br>

### Usage
<ul>
    <li>Go to postman and register with <b>POST</b> <code>http://127.0.0.1:5000/register</code></li>
    <ul>
        <li>Don't forget fill your <u>username</u>, <u>email</u>, <u>name</u>, <u>password</u> in the body field</li>
        <li>Once done click <b>Send</b> buton</li>
    </ul>
    <li>Then, login to store your cookie headers with <b>POST</b> <code>http://127.0.0.1:5000/login</code></li>
    <ul>
        <li>Don't forget fill with <u>username</u> or <u>email</u> and <u>password</u> in the body field
        <li>Once done click <b>Send</b> buton</li>
    </ul>
    <li>After login, you can see the account you just created with
    <br><b>GET</b> <code>http://127.0.0.1:5000/me</code> without fill any body field, because it use cookie headers.</li>
</ul>
<br>

### Inspirations
I was have an idea to create like an e-commerce web named **dijokiin.id** but,<br>
I kinda lazy to do it alone :( so I just creating this simple web backend for my portfolio.

<br>
<br>

<h2 align="center">COPYRIGHT</h2>
<p align="center"><b><a href="https://github.com/rushkii">Kiizuha</a></b> 2022 under the <a href="LICENSE">MIT License</a>.</p>
