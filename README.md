<h1 align="center">BackendLoginGo</h1>
<p align="center">
A simple example backend web for account creation, login, logout, and get the account info using Go language.
</p>

<br>

## Tutorial
- Open your terminal/cmd and type `git clone https://github.com/rushkii/BackendLoginGo.git`
- Then `cd BackendLoginGo` and change your database DSN in `config.env`
- Once setup complete, run the program `go run server.go`

<br>

## Usage
<ul>
    <li>Go to postman and register with POST <code>http://127.0.0.1:5000/register</code></li>
    <ul>
        <li>Don't forget fill your _username_, email, name, password in the body field</li>
        <li>Once done click <b>Send</b> buton</li>
    </ul>
    <li>Then, login to store your cookie headers with POST <code>http://127.0.0.1:5000/login</code></li>
    <ul>
        <li>Don't forget fill with username or email and password in the body field
        <li>Once done click <b>Send</b> buton</li>
    </ul>
    <li>After login, you can see the account you just created with
    <br>GET <code>http://127.0.0.1:5000/me</code> without fill any body field, because it use cookie headers.</li>
</ul>

<br>
<br>

<h2 align="center">COPYRIGHT</h2>
<p align="center"><b><a href="github.com/rushkii">Kiizuha</a></b> 2022</p>