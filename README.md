<h1 align="center"> Book Matcher </h1>
<h5 align="center"> Development Started - Summer 2023 </h5>

English | [日本語](https://github.com/nao-United92/book-matcher/blob/main/README.ja.md)</h2>

<!-- TABLE OF CONTENTS -->
<h2 id="table-of-contents"> :book: TABLE OF CONTENTS</h2>

<details open="open">
  <summary>TABLE OF CONTENTS</summary>
  <ol>
    <li><a href="#overview"> ➤ Overview</a></li>
    <li><a href="#technologies"> ➤ Technologies</a></li>
    <li><a href="#project-files-description"> ➤ Project Files Description</a></li>
    <li><a href="#getting-started"> ➤ Getting Started</a></li>
    <li><a href="#ongoing"> ➤ Ongoing</a></li>
    <li><a href="#to-be-implemented"> ➤ To Be Implemented</a></li>
    <li><a href="#references"> ➤ References</a></li>
  </ol>
</details>


![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- Overview -->
<h2 id="overview"> :cloud: Overview</h2>

<p align="justify">
  This is a project repository for <b>Book Matcher</b>, a community service that connects people with their favorite books.<br>
  <b>※It's still unfinished. From now on, we will implement new main functions, expand and improve existing functions.</b>
</p>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- technologies -->
<h2 id="technologies"> :gear: Technologies</h2>

<p align="justify">
<img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=plastic">
<img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=plastic">
<img src="https://img.shields.io/badge/-Mysql-4479A1.svg?logo=mysql&style=plastic">
<img src="https://img.shields.io/badge/-Html5-E34F26.svg?logo=html5&style=plastic">
<img src="https://img.shields.io/badge/-Css3-1572B6.svg?logo=css3&style=plastic">
<img src="https://img.shields.io/badge/-Jquery-0769AD.svg?logo=jquery&style=plastic">
</p>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- PROJECT FILES DESCRIPTION -->
<h2 id="project-files-description"> :floppy_disk: Project Files Description</h2>

<h3>controller</h3>
<ul>
  <li><b>bookmark/bookmark.go</b> - Event handling related to bookmark functionality</li>
  <li><b>login/login.go</b> - Event handling related to login functionality</li>
  <li><b>password/password.go</b> - Event processing related to password reissue function</li>
  <li><b>profile/profile.go</b> - Event handling related to profile features</li>
  <li><b>search/search.go</b> - Event processing related to book search function</li>
  <li><b>signup/signup.go</b> - Event processing related to new user registration function</li>
</ul>

<h3>db/my.cnf</h3>
<ul>
  This is a DB (MySQL) server configuration file that is read when the server starts. Options related to DB operation are set.
</ul>

<h3>model</h3>
<ul>
  <li><b>bookmark.go</b> - Bookmark data table column definition, data CRUD processing</li>
  <li><b>user</b> - User data table column definition, data CRUD processing</li>
</ul>

<h3>static</h3>
<ul>
  <li><b>css/style.css</b> - Common CSS file</li>
  <li><b>images/</b> - Contains sample images for testing.</li>
  <li><b>js/script.js</b> - Common JS file</li>
</ul>

<h3>view</h3>
<ul>
  <li><b>base.html</b> - Common page design</li>
  <li><b>bookmark.html</b> - Bookmark list form design</li>
  <li><b>login.html</b> - Login list form design</li>
  <li><b>menu.html</b> - Menu (initial display) page design</li>
  <li><b>password_reset.html</b> - Password reissue form design</li>
  <li><b>profile.html</b> - Profile form design</li>
  <li><b>search_books.html</b> - Book search list form design</li>
  <li><b>search_books_detail.html</b> - Book details form design</li>
  <li><b>signup.html</b> - New user registration form design</li>
</ul>

<h3>.env</h3>
<ul>
  Contains database connection information. Read from main.go and used in DB connection.
</ul>

<h3>Dockerfile</h3>
<ul>
  Specify Docker image, install hot reload tool (air), and change WORKDIR
</ul>

<h3>compose.yaml</h3>
<ul>
  Container definitions for the app server (go), DB server (db), and DB client tool (phpmyadmin) are described.
</ul>

<h3>go.mod</h3>
<ul>
  Describes module dependencies used in Go programming.
</ul>

<h3>go.sum</h3>
<ul>
  Contains hashed information about module dependencies in the go.mod file.
</ul>

<h3>main.go</h3>
<ul>
  Entry point for Go programs. Immediately after the program starts, the "main function" belonging to the "main" package is executed first.
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- Getting Started -->
<h2 id="getting-started"> :arrow_forward: Getting Started</h2>

<p><b>1. When starting for the first time</b> - Execute image construction, container construction and startup all at once</p>
<pre><code>$ docker-compose up --build -d</code></pre>

<p><b>2. When starting after the second time</b> - Execute container construction and start</p>
<pre><code>$ docker-compose up -d</code></pre>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- Ongoing -->
<h2 id="ongoing"> :crossed_swords: Ongoing</h2>
<ul>
  <li>Password reset function</li>
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- To Be Implemented -->
<h2 id="to-be-implemented"> :memo: To Be Implemented</h2>
<ul>
  <li>Book review posting function</li>
  <li>View and add friends function</li>
  <li>Notification function</li>
</ul>

![-----------------------------------------------------](https://raw.githubusercontent.com/andreasbm/readme/master/assets/lines/rainbow.png)

<!-- References -->
<h2 id="references"> :scroll: References</h2>
<a href="https://developers.google.com/books/docs/v1/using?hl=ja">Google Books APIs</a><br>
<a href="https://qiita.com/s_emoto/items/975cc38a3e0de462966a">About the MVC model</a><br>
<a href="https://qiita.com/soicchi/items/2637a9195e64fdc73609">What are go.mod and go.sum files?</a><br>
<a href="https://the-simple.jp/what-is-o-r-mapper-explains-the-basic-concept-of-o-r-mapping-and-how-to-use-it-in-practice">What is O/R Mapper? Explaining the basic concept and practical usage of O/R mapping</a><br>
<a href="https://fonts.google.com/icons">Material Symbols and Icons - Google Fonts</a>
<br><br>

<a href="https://github.com/nao-United92" target="_blank"><img alt="Github" src="https://img.shields.io/badge/GitHub-%2312100E.svg?&style=for-the-badge&logo=Github&logoColor=white" /></a>
<a href="https://qiita.com/nao-United92" target="_blank"><img alt="Medium" src="https://img.shields.io/badge/qiita-55C500.svg?&style=for-the-badge&logo=qiita&logoColor=white" /></a>
